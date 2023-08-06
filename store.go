// Copyright (c) 2022 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license
package gostore

import (
	"sync"
	"sync/atomic"
	"time"
)

type item[T any] struct {
	value  T
	expire time.Time
}

type store[T any] struct {
	sync.RWMutex
	data    map[string]item[T]
	ttl     time.Duration
	shift   time.Duration
	stop    chan struct{}
	started atomic.Bool // invalidator started
	closed  atomic.Bool
}

func newStore[T any](ttl, shift time.Duration) *store[T] {
	if ttl <= 0 {
		ttl = durationMax
	}
	if shift <= 0 {
		shift = shiftDefault
	}
	s := store[T]{ttl: ttl, shift: shift}
	s.data = make(map[string]item[T])
	s.stop = make(chan struct{})
	if s.ttl == durationMax {
		return &s
	}
	s.startInvalidator(ttl / ttlToPeriodDiv)
	return &s
}

func (s *store[T]) close() {
	if s == nil || !s.closed.CompareAndSwap(false, true) {
		return
	}
	close(s.stop)
	s.data = map[string]item[T]{}
}

func (s *store[T]) set(key string, value T) {
	if s == nil || s.closed.Load() {
		return
	}
	item := item[T]{
		value:  value,
		expire: time.Now().Add(s.ttl),
	}
	s.Lock()
	s.data[key] = item
	s.Unlock()
}

func (s *store[T]) setWithTTL(key string, value T, ttl time.Duration) {
	if s == nil || s.closed.Load() {
		return
	}
	item := item[T]{
		value:  value,
		expire: time.Now().Add(ttl),
	}
	s.Lock()
	s.data[key] = item
	s.Unlock()
	if !s.started.Load() {
		s.startInvalidator(ttl / ttlToPeriodDiv)
	}
}

func (s *store[T]) get(key string) (T, bool) {
	if s == nil || s.closed.Load() {
		var v T
		return v, false
	}
	s.RLock()
	item, found := s.data[key]
	s.RUnlock()
	return item.value, found
}

func (s *store[T]) delete(key string) {
	if s == nil || s.closed.Load() {
		return
	}
	s.Lock()
	delete(s.data, key)
	s.Unlock()
}

func (s *store[T]) size() int {
	if s == nil || s.closed.Load() {
		return 0
	}
	s.RLock()
	size := len(s.data)
	s.RUnlock()
	return size
}

func (s *store[T]) keys() []string {
	if s == nil || s.closed.Load() {
		return []string{}
	}
	s.RLock()
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	s.RUnlock()
	return keys
}

func (s *store[T]) invalidate() int64 {
	if s == nil || s.closed.Load() {
		return 0
	}
	var counter int64
	t := time.Now()
	s.Lock()
	for k, v := range s.data {
		if !t.After(v.expire) {
			continue
		}
		delete(s.data, k)
		counter++
	}
	s.Unlock()
	return counter
}

func (s *store[T]) startInvalidator(period time.Duration) {
	if s == nil || s.closed.Load() ||
		!s.started.CompareAndSwap(false, true) {
		return
	}
	go func() {
		<-time.After(s.shift)
		if period > periodMax {
			period = periodMax
		} else if period < periodMin {
			period = periodMin
		}
		ticker := time.NewTicker(period)
		defer ticker.Stop()
		for {
			select {
			case <-s.stop:
				return
			case <-ticker.C:
				if s.closed.Load() {
					return
				}
				s.invalidate()
			}
		}
	}()
}
