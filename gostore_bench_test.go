// Copyright (c) 2022 Nikita Chisnikov <chisnikov@gmail.com>
// Distributed under the MIT/X11 software license
package gostore

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkGet(b *testing.B) {
	s, closer := New[string](0)
	defer closer()

	s.Set("1f586b6c-68e9-4b59-8e2c-7e0dacf2409b", "first")
	for i := 0; i < b.N; i++ {
		s.Get("1f586b6c-68e9-4b59-8e2c-7e0dacf2409b")
	}
}

func BenchmarkInvalidateNumeric(b *testing.B) {
	for _, v := range []int{1000, 10_000, 100_000, 1000_000} {
		b.Run(fmt.Sprintf("Count-%d", v), func(b *testing.B) {
			s, closeStore := NewNumeric[string](0)
			for i := 0; i < v; i++ {
				s.Set(strconv.Itoa(i), strconv.Itoa(i))
			}
			for i := 0; i < b.N; i++ {
				s.Invalidate()
			}
			closeStore()
		})
	}
}

func BenchmarkGetInternal(b *testing.B) {
	for _, v := range []int{1000, 10_000, 100_000, 1000_000} {
		b.Run(fmt.Sprintf("Count-%d", v), func(b *testing.B) {
			s := newStore[string](0, 0)
			for i := 0; i < v; i++ {
				s.set(strconv.Itoa(i), strconv.Itoa(i))
			}
			for i := 0; i < b.N; i++ {
				s.get(strconv.Itoa(i))
			}
			s.close()
		})
	}
}

func BenchmarkGetNumeric(b *testing.B) {
	for _, v := range []int{1000, 10_000, 100_000, 1000_000} {
		b.Run(fmt.Sprintf("Count-%d", v), func(b *testing.B) {
			s, closeStore := NewNumeric[string](0)
			for i := 0; i < v; i++ {
				s.Set(strconv.Itoa(i), strconv.Itoa(i))
			}
			for i := 0; i < b.N; i++ {
				s.Get(strconv.Itoa(i))
			}
			closeStore()
		})
	}
}

func BenchmarkSizeNumeric(b *testing.B) {
	for _, v := range []int{1000, 10_000, 100_000, 1000_000} {
		b.Run(fmt.Sprintf("Count-%d", v), func(b *testing.B) {
			s, closeStore := NewNumeric[string](0)
			for i := 0; i < v; i++ {
				s.Set(strconv.Itoa(i), strconv.Itoa(i))
			}
			for i := 0; i < b.N; i++ {
				s.Size()
			}
			closeStore()
		})
	}
}
