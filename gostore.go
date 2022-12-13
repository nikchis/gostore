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

func newStore[T any](ttl time.Duration) *store[T] {
	s := store[T]{ttl: ttl}
	s.data = make(map[string]item[T])
	s.stop = make(chan struct{})
	if s.ttl > 0 {
		go func() {
			for {
				select {
				case <-s.stop:
					return
				case <-time.After(s.ttl):
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
	}
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

func (s *store[T]) setWithTtl(key string, value T, ttl time.Duration) {
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
	/*
		// Too expensive for performance
		if ok && time.Now().After(v.expire) {
			var r T
			s.Lock()
			delete(s.data, key)
			s.Unlock()
			return r, false
		}
	*/
	return v.value, ok
}

func (s *store[T]) delete(key string) {
	s.Lock()
	delete(s.data, key)
	s.Unlock()
}

// -----------------------------------------------------------------------------

type Store[T any] struct {
	sync.RWMutex
	arr    [16]*store[T]
	ttl    time.Duration
	closed bool
}

// Create a new Store with TTL
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

func (s *Store[T]) Set(uuidKey string, value T) {
	k := uuidKey[len(uuidKey)-1]
	s.RLock()
	if s.closed {
		s.RUnlock()
		return
	}
	switch k {
	case '0':
		s.arr[0].set(uuidKey, value)
	case '1':
		s.arr[1].set(uuidKey, value)
	case '2':
		s.arr[2].set(uuidKey, value)
	case '3':
		s.arr[3].set(uuidKey, value)
	case '4':
		s.arr[4].set(uuidKey, value)
	case '5':
		s.arr[5].set(uuidKey, value)
	case '6':
		s.arr[6].set(uuidKey, value)
	case '7':
		s.arr[7].set(uuidKey, value)
	case '8':
		s.arr[8].set(uuidKey, value)
	case '9':
		s.arr[9].set(uuidKey, value)
	case 'a':
		s.arr[10].set(uuidKey, value)
	case 'b':
		s.arr[11].set(uuidKey, value)
	case 'c':
		s.arr[12].set(uuidKey, value)
	case 'd':
		s.arr[13].set(uuidKey, value)
	case 'e':
		s.arr[14].set(uuidKey, value)
	case 'f':
		s.arr[15].set(uuidKey, value)
	}
	s.RUnlock()
}

func (s *Store[T]) SetWithTtl(uuidKey string, value T, ttl time.Duration) {
	k := uuidKey[len(uuidKey)-1]
	s.RLock()
	if s.closed {
		s.RUnlock()
		return
	}
	switch k {
	case '0':
		s.arr[0].setWithTtl(uuidKey, value, ttl)
	case '1':
		s.arr[1].setWithTtl(uuidKey, value, ttl)
	case '2':
		s.arr[2].setWithTtl(uuidKey, value, ttl)
	case '3':
		s.arr[3].setWithTtl(uuidKey, value, ttl)
	case '4':
		s.arr[4].setWithTtl(uuidKey, value, ttl)
	case '5':
		s.arr[5].setWithTtl(uuidKey, value, ttl)
	case '6':
		s.arr[6].setWithTtl(uuidKey, value, ttl)
	case '7':
		s.arr[7].setWithTtl(uuidKey, value, ttl)
	case '8':
		s.arr[8].setWithTtl(uuidKey, value, ttl)
	case '9':
		s.arr[9].setWithTtl(uuidKey, value, ttl)
	case 'a':
		s.arr[10].setWithTtl(uuidKey, value, ttl)
	case 'b':
		s.arr[11].setWithTtl(uuidKey, value, ttl)
	case 'c':
		s.arr[12].setWithTtl(uuidKey, value, ttl)
	case 'd':
		s.arr[13].setWithTtl(uuidKey, value, ttl)
	case 'e':
		s.arr[14].setWithTtl(uuidKey, value, ttl)
	case 'f':
		s.arr[15].setWithTtl(uuidKey, value, ttl)
	}
	s.RUnlock()
}

func (s *Store[T]) Get(uuidKey string) (T, bool) {
	var v T
	var ok bool
	k := uuidKey[len(uuidKey)-1]
	s.RLock()
	if s.closed {
		s.RUnlock()
		return v, ok
	}
	switch k {
	case '0':
		v, ok = s.arr[0].get(uuidKey)
	case '1':
		v, ok = s.arr[1].get(uuidKey)
	case '2':
		v, ok = s.arr[2].get(uuidKey)
	case '3':
		v, ok = s.arr[3].get(uuidKey)
	case '4':
		v, ok = s.arr[4].get(uuidKey)
	case '5':
		v, ok = s.arr[5].get(uuidKey)
	case '6':
		v, ok = s.arr[6].get(uuidKey)
	case '7':
		v, ok = s.arr[7].get(uuidKey)
	case '8':
		v, ok = s.arr[8].get(uuidKey)
	case '9':
		v, ok = s.arr[9].get(uuidKey)
	case 'a':
		v, ok = s.arr[10].get(uuidKey)
	case 'b':
		v, ok = s.arr[11].get(uuidKey)
	case 'c':
		v, ok = s.arr[12].get(uuidKey)
	case 'd':
		v, ok = s.arr[13].get(uuidKey)
	case 'e':
		v, ok = s.arr[14].get(uuidKey)
	case 'f':
		v, ok = s.arr[15].get(uuidKey)
	}
	s.RUnlock()
	return v, ok
}

func (s *Store[T]) Delete(uuidKey string) {
	k := uuidKey[len(uuidKey)-1]
	s.RLock()
	if s.closed {
		s.RUnlock()
		return
	}
	switch k {
	case '0':
		s.arr[0].delete(uuidKey)
	case '1':
		s.arr[1].delete(uuidKey)
	case '2':
		s.arr[2].delete(uuidKey)
	case '3':
		s.arr[3].delete(uuidKey)
	case '4':
		s.arr[4].delete(uuidKey)
	case '5':
		s.arr[5].delete(uuidKey)
	case '6':
		s.arr[6].delete(uuidKey)
	case '7':
		s.arr[7].delete(uuidKey)
	case '8':
		s.arr[8].delete(uuidKey)
	case '9':
		s.arr[9].delete(uuidKey)
	case 'a':
		s.arr[10].delete(uuidKey)
	case 'b':
		s.arr[11].delete(uuidKey)
	case 'c':
		s.arr[12].delete(uuidKey)
	case 'd':
		s.arr[13].delete(uuidKey)
	case 'e':
		s.arr[14].delete(uuidKey)
	case 'f':
		s.arr[15].delete(uuidKey)
	}
	s.RUnlock()
}
