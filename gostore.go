// Copyright (c) 2022 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license
package gostore

import (
	"sync"
	"time"
)

type Store[T any] struct {
	sync.RWMutex
	arr    [storeArrLastIndex + 1]*store[T]
	ttl    time.Duration
	closed bool
}

// Create a new Store with TTL
// Use TTL = 0 to disable TTL invalidation
func New[T any](ttl time.Duration) (*Store[T], func()) {
	s := Store[T]{ttl: ttl}
	for i := range s.arr {
		s.arr[i] = newStore[T](ttl)
	}
	return &s, func() { s.Close() }
}

// Close Store and free resources
func (s *Store[T]) Close() {
	s.Lock()
	defer s.Unlock()
	if s.closed {
		return
	}
	s.closed = true
	for i := range s.arr {
		if s.arr[i] != nil {
			s.arr[i].close()
		}
		s.arr[i] = nil
	}
}

// Set is optimized for UUID keys, but other formats are allowed too
func (s *Store[T]) Set(uuidKey string, value T) {
	if uuidKey == "" {
		return
	}

	k := uuidKey[len(uuidKey)-1]

	s.RLock()
	if s.closed {
		s.RUnlock()
		return
	}
	s.arr[keyChar(k).storeArrIndex()].set(uuidKey, value)
	s.RUnlock()
}

// SetWithCustomTTL is optimized for UUID keys, but other formats are allowed too
func (s *Store[T]) SetWithCustomTTL(uuidKey string, value T, ttl time.Duration) {
	if uuidKey == "" {
		return
	}

	k := uuidKey[len(uuidKey)-1]

	s.RLock()
	if s.closed {
		s.RUnlock()
		return
	}
	s.arr[keyChar(k).storeArrIndex()].setWithTTL(uuidKey, value, ttl)
	s.RUnlock()
}

func (s *Store[T]) Get(uuidKey string) (T, bool) {
	var v T
	var ok bool

	if uuidKey == "" {
		return v, false
	}

	k := uuidKey[len(uuidKey)-1]

	s.RLock()
	if s.closed {
		s.RUnlock()
		return v, false
	}
	v, ok = s.arr[keyChar(k).storeArrIndex()].get(uuidKey)
	s.RUnlock()

	return v, ok
}

func (s *Store[T]) Delete(uuidKey string) {
	if uuidKey == "" {
		return
	}

	k := uuidKey[len(uuidKey)-1]

	s.RLock()
	if s.closed {
		s.RUnlock()
		return
	}
	s.arr[keyChar(k).storeArrIndex()].delete(uuidKey)
	s.RUnlock()
}
