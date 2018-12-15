package main

import (
	"bytes"
	"fmt"
)

// IntSet is a set of small non-negative integers
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint64(x%64)
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

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("}") {
					buf.WriteByte(',')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns words length
func (s *IntSet) Len() int {
	return len(s.words) + 1
}

// Remove x from set
func (s *IntSet) Remove(x int) {
	words, bit := x/64, uint64(x%64)
	s.words[words] &^= (1 << bit)
}

// Clear clean all of the words
func (s *IntSet) Clear() {
	s.words = *new([]uint64)
}

// Copy return a copy of the set
func (s IntSet) Copy() *IntSet {
	return &s
}

func main() {
	var x, y IntSet
	x.Add(2)
	x.Add(144)
	x.Add(9)
	// fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	// fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
	// fmt.Println(x.Has(9), x.Has(123))

	fmt.Println(x.Len())
	x.Remove(42)
	fmt.Println(x.String())
	z := x.Copy()
	x.Clear()
	fmt.Println(x.String())
	fmt.Println(z.String())
}
