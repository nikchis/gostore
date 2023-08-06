// Copyright (c) 2022 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license
package gostore

import (
	"sync"
	"sync/atomic"
	"time"
)

type Store[T any] struct {
	arr    [arrLastIndex + 1]*store[T]
	ttl    time.Duration
	closed atomic.Bool
}

// Create a new Store with TTL optimized for alphanumeric keys.
// Use TTL = 0 to disable TTL invalidation.
func New[T any](ttl time.Duration) (*Store[T], func()) {
	if ttl == 0 {
		ttl = durationMax
	}
	s := Store[T]{ttl: ttl}
	for i := range s.arr {
		invalidationShift := time.Duration(i+1) * shiftDefault
		s.arr[i] = newStore[T](ttl, invalidationShift)
	}
	return &s, func() { s.Close() }
}

// Create a new Store with TTL optimized for UUID keys.
// Use TTL = 0 to disable TTL invalidation.
func NewUUID[T any](ttl time.Duration) (*Store[T], func()) {
	if ttl == 0 {
		ttl = durationMax
	}
	s := Store[T]{ttl: ttl}
	var lastIndexStore *store[T]
	for i := range s.arr {
		invalidationShift := time.Duration(i+1) * shiftDefault
		if i <= arrFIndex {
			s.arr[i] = newStore[T](ttl, invalidationShift)
			continue
		}
		if lastIndexStore == nil {
			lastIndexStore = newStore[T](ttl, invalidationShift)
		}
		s.arr[i] = lastIndexStore
	}
	return &s, func() { s.Close() }
}

// Create a new Store with TTL optimized for numeric keys.
// Use TTL = 0 to disable TTL invalidation.
func NewNumeric[T any](ttl time.Duration) (*Store[T], func()) {
	if ttl == 0 {
		ttl = durationMax
	}
	s := Store[T]{ttl: ttl}
	var lastIndexStore *store[T]
	for i := range s.arr {
		invalidationShift := time.Duration(i+1) * shiftDefault
		if i <= arr9Index {
			s.arr[i] = newStore[T](ttl, invalidationShift)
			continue
		}
		if lastIndexStore == nil {
			lastIndexStore = newStore[T](ttl, invalidationShift)
		}
		s.arr[i] = lastIndexStore
	}
	return &s, func() { s.Close() }
}

// Close Store and free resources.
func (s *Store[T]) Close() {
	if s == nil || !s.closed.CompareAndSwap(false, true) {
		return
	}
	for i := range s.arr {
		if s.arr[i] != nil {
			s.arr[i].close()
			s.arr[i] = nil
		}
	}
}

// Set is optimized for alphanumeric, UUID or numeric keys.
// But other formats are allowable too.
func (s *Store[T]) Set(key string, value T) {
	if s == nil || s.closed.Load() || key == "" {
		return
	}
	k := key[len(key)-1]
	s.arr[keyChar(k).arrIndex()].set(key, value)
}

// SetWithCustomTTL is optimized for alphanumeric, UUID or numeric keys.
// But other formats are allowable too.
func (s *Store[T]) SetWithCustomTTL(key string, value T, ttl time.Duration) {
	if s == nil || s.closed.Load() || key == "" {
		return
	}
	k := key[len(key)-1]
	s.arr[keyChar(k).arrIndex()].setWithTTL(key, value, ttl)
}

func (s *Store[T]) Get(key string) (T, bool) {
	if s == nil || s.closed.Load() || key == "" {
		var v T
		return v, false
	}
	k := key[len(key)-1]
	v, ok := s.arr[keyChar(k).arrIndex()].get(key)
	return v, ok
}

func (s *Store[T]) Delete(key string) {
	if s == nil || s.closed.Load() || key == "" {
		return
	}
	k := key[len(key)-1]
	s.arr[keyChar(k).arrIndex()].delete(key)
}

func (s *Store[T]) Size() int {
	if s == nil || s.closed.Load() {
		return 0
	}
	var size int = s.arr[0].size()
	for i := 1; i <= arrLastIndex; i++ {
		if s.arr[i] == s.arr[i-1] {
			break
		}
		size += s.arr[i].size()
	}
	return size
}

func (s *Store[T]) Keys() []string {
	if s == nil || s.closed.Load() {
		return []string{}
	}
	keys := make([]string, 0, s.Size())
	keys = append(keys, s.arr[0].keys()...)
	for i := 1; i <= arrLastIndex; i++ {
		if s.arr[i] == s.arr[i-1] {
			break
		}
		keys = append(keys, s.arr[i].keys()...)
	}
	return keys
}

// Deletes all expired by TTL values.
// Returns count of deleted values.
func (s *Store[T]) Invalidate() int64 {
	if s == nil || s.closed.Load() {
		return 0
	}
	var count atomic.Int64
	var wg sync.WaitGroup
	for i := 1; i <= arrLastIndex; i++ {
		if s.arr[i] == s.arr[i-1] {
			break
		}
		if s.arr[i].size() >= 1000 {
			wg.Add(1)
			go func(i int) {
				count.Add(s.arr[i].invalidate())
				wg.Done()
			}(i)
		} else {
			count.Add(s.arr[i].invalidate())
		}
	}
	count.Add(s.arr[0].invalidate())
	wg.Wait()

	return count.Load()
}
