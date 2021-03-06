package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// return the number of elements
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word := x / 64
	s.words[word] &= ^(1 << uint(x))
}

// remove all elements from the set
func (s *IntSet) Clear() {
	for i, _ := range s.words {
		s.words[i] = 0
	}
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var t IntSet
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				t.Add(j)
			}
		}
	}
	return &t
}

func (s *IntSet) AddAll(args ...int) {
	for _, a := range args {
		s.Add(a)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) []uint64 {
	var ret []uint64
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if t.Has(j) {
					ret = append(ret, uint64(j))
				}
			}
		}
	}
	return ret
}

func (s *IntSet) DifferenceWith(t *IntSet) []uint64 {
	var ret []uint64
	return ret
}

func (s *IntSet) SymmetricDifference(t *IntSet) []uint64 {
	var ret []uint64
	return ret
}
