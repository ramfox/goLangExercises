// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
	"gopl.io/ch2/popcount"
)

//!+intset

const platformBitSize = int(32 << (^uint(0) >> 63))

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/platformBitSize, uint(x%platformBitSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/platformBitSize, uint(x%platformBitSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/platformBitSize, uint(x%platformBitSize)
		s.words[word] &^= 1 << bit
	}
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

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// DifferenceWith sets s to s - t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

// SymetricDifference sets s to the symetric difference of s and t.
// what is in either s or t, but not in both
func (s *IntSet) SymetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Len returns the number of elements
func (s *IntSet) Len() int {
	var length int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		length += popcount.PopCount(word)
	}
	return length
}

// Clear removes all the elements of a set
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var n IntSet
	n.words = make([]uint64, len(s.words))
	copy(n.words, s.words)
	return &n
}

// Variadic method that allows a list of values to be added
func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

// Elems returns a slice containing the elements of the set
func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < platformBitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, platformBitSize*i+j)
			}
		}
	}
	return elems
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < platformBitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", platformBitSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string
