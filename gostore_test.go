// Copyright (c) 2022 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license
package gostore

import (
	"testing"
	"time"
)

const (
	tokenJWT = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5R`
)

func TestSetGet(t *testing.T) {
	s, closer := New[string](time.Second)
	defer closer()

	s.Set("1F586B6c-68e9-4b59-8e2c-7e0dacf2409b", "first")
	s.Set("e30d341f-be81-487f-b368-9d045d263044", "second")
	s.Set("e30d341f-be81-487f-b368-9d045d26304B", "third")
	s.Set(tokenJWT, "jwt")

	if _, ok := s.Get("1f586b6c-68e9-4b59-8e2c-7e0dacf2409b"); ok {
		t.Fatalf("Failed on case 0")
	}
	if _, ok := s.Get("1F586B6c-68e9-4b59-8e2c-7e0dacf2409a"); ok {
		t.Fatalf("Failed on case 1")
	}
	if _, ok := s.Get("1F586B6c-68e9-4b59-8e2c-7e0dacf2409a"); ok {
		t.Fatalf("Failed on case 2")
	}
	if v, ok := s.Get("1F586B6c-68e9-4b59-8e2c-7e0dacf2409b"); ok && v != "first" {
		t.Fatalf("Failed on case 3")
	}
	if v, ok := s.Get("e30d341f-be81-487f-b368-9d045d263044"); ok && v != "second" {
		t.Fatalf("Failed on case 4")
	}
	if v, ok := s.Get("e30d341f-be81-487f-b368-9d045d26304B"); !ok || v != "third" {
		t.Fatalf("Failed on case 5")
	}
	if v, ok := s.Get(tokenJWT); !ok || v != "jwt" {
		t.Fatalf("Failed on case 6")
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
