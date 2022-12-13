// Copyright (c) 2022 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license
package gostore

import (
	"testing"
	"time"
)

func TestSetGet(t *testing.T) {
	s, closer := New[string](time.Second)
	defer closer()

	s.Set("1f586b6c-68e9-4b59-8e2c-7e0dacf2409b", "first")
	s.Set("e30d341f-be81-487f-b368-9d045d263044", "second")

	if _, ok := s.Get("1f586b6c-68e9-4b59-8e2c-7e0dacf2409b"); !ok {
		t.Fatalf("Failed on case 1")
	}
	if _, ok := s.Get("1f586b6c-68e9-4b59-8e2c-7e0dacf2409a"); ok {
		t.Fatalf("Failed on case 2")
	}
	if v, ok := s.Get("1f586b6c-68e9-4b59-8e2c-7e0dacf2409b"); ok && v != "first" {
		t.Fatalf("Failed on case 3")
	}
	if v, ok := s.Get("e30d341f-be81-487f-b368-9d045d263044"); ok && v != "second" {
		t.Fatalf("Failed on case 4")
	}
}

func BenchmarkGet(b *testing.B) {
	s, closer := New[string](10 * time.Second)
	defer closer()

	s.Set("1f586b6c-68e9-4b59-8e2c-7e0dacf2409b", "first")
	for i := 0; i < b.N; i++ {
		s.Get("1f586b6c-68e9-4b59-8e2c-7e0dacf2409b")
	}
}

func BenchmarkGetInternal(b *testing.B) {
	s := newStore[string](10 * time.Second)
	defer s.close()

	s.set("1f586b6c-68e9-4b59-8e2c-7e0dacf2409b", "first")
	for i := 0; i < b.N; i++ {
		s.get("1f586b6c-68e9-4b59-8e2c-7e0dacf2409b")
	}
}
