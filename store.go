// Copyright (c) 2022 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license
package gostore

import (
	"sync"
	"time"
)

type item[T any] struct {
	value  T
	expire time.Time
}

type store[T any] struct {
	sync.RWMutex
	data   map[string]item[T]
	ttl    time.Duration
	stop   chan struct{}
	closed bool
}

func newStore[T any](ttl, shift time.Duration) *store[T] {
	s := store[T]{ttl: ttl}
	s.data = make(map[string]item[T])
	s.stop = make(chan struct{})

	if s.ttl <= 0 {
		return &s
	}

	go func() {
		<-time.After(shift)

		period := ttl / 2
		ticker := time.NewTicker(period)
		defer ticker.Stop()

		for {
			select {
			case <-s.stop:
				return
			case <-ticker.C:
				t := time.Now()
				s.Lock()
				for k, v := range s.data {
					if t.After(v.expire) {
						delete(s.data, k)
					}
				}
				s.Unlock()
			}
		}
	}()

	return &s
}

func (s *store[T]) close() {
	s.Lock()
	defer s.Unlock()
	if s.closed {
		return
	}
	close(s.stop)
	s.closed = true
}

func (s *store[T]) set(key string, value T) {
	item := item[T]{
		value:  value,
		expire: time.Now().Add(s.ttl),
	}
	s.Lock()
	s.data[key] = item
	s.Unlock()
}

func (s *store[T]) setWithTTL(key string, value T, ttl time.Duration) {
	item := item[T]{
		value:  value,
		expire: time.Now().Add(ttl),
	}
	s.Lock()
	s.data[key] = item
	s.Unlock()
}

func (s *store[T]) get(key string) (T, bool) {
	s.RLock()
	v, ok := s.data[key]
	s.RUnlock()
	return v.value, ok
}

func (s *store[T]) delete(key string) {
	s.Lock()
	delete(s.data, key)
	s.Unlock()
}
