// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func TestLength(t *testing.T) {
	var tests = []struct {
		elements []int
		want     int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{42, 100, 1}, 3},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("Len(): %v", test.elements)
		var s IntSet
		for _, element := range test.elements {
			s.Add(element)
		}
		got := s.Len()
		if got != test.want {
			t.Errorf("%s = %d, want %d", descr, got, test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		elements []int
		remove   int
		want     string
	}{
		{[]int{}, 10, "{}"},
		{[]int{1}, 1, "{}"},
		{[]int{1}, 2, "{1}"},
		{[]int{42, 100, 1}, 100, "{1 42}"},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("%v, Remove(%d):", test.elements, test.remove)
		var s IntSet
		for _, element := range test.elements {
			s.Add(element)
		}
		s.Remove(test.remove)
		got := s.String()
		if got != test.want {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		elements []int
		want     string
	}{
		{[]int{}, "{}"},
		{[]int{1}, "{}"},
		{[]int{1, 2, 4, 5, 6, 7}, "{}"},
		{[]int{42, 100, 1}, "{}"},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("%v, Clear():", test.elements)
		var s IntSet
		for _, element := range test.elements {
			s.Add(element)
		}
		s.Clear()
		got := s.String()
		if got != test.want {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}

func TestAddAll(t *testing.T) {
	var tests = []struct {
		elements []int
		add      []int
		want     string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{}, []int{1}, "{1}"},
		{[]int{}, []int{1, 2}, "{1 2}"},
		{[]int{1, 2}, []int{3, 4}, "{1 2 3 4}"},
		{[]int{100, 14, 59}, []int{25, 43, 67}, "{14 25 43 59 67 100}"},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("%v, Clear():", test.elements)
		var s IntSet
		for _, element := range test.elements {
			s.Add(element)
		}
		s.AddAll(test.add...)
		got := s.String()
		if got != test.want {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}

func TestElems(t *testing.T) {
	var tests = []struct {
		want []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{1, 2}},
		{[]int{1, 2, 3, 4}},
		{[]int{14, 25, 43, 59, 67, 100}},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("%v, Clear():", test.want)
		var s IntSet
		for _, element := range test.want {
			s.Add(element)
		}
		got := s.Elems()
		for i, _ := range got {
			if got[i] != test.want[i] {
				t.Errorf("%s = %s, want %s", descr, got, test.want)
				break
			}
		}
	}
}

func TestUnion(t *testing.T) {
	var tests = []struct {
		s    []int
		u    []int
		want string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{}, []int{1}, "{1}"},
		{[]int{1}, []int{2}, "{1 2}"},
		{[]int{1, 2}, []int{3, 4}, "{1 2 3 4}"},
		{[]int{1, 3}, []int{3, 4}, "{1 3 4}"},
		{[]int{100, 14, 59}, []int{25, 43, 67}, "{14 25 43 59 67 100}"},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("s.UnionWith(t):\ns: %v\nt: %v", test.s, test.u)
		var s, u IntSet
		for _, element := range test.s {
			s.Add(element)
		}
		for _, element := range test.u {
			u.Add(element)
		}
		s.UnionWith(&u)
		got := s.String()
		if got != test.want {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}

func TestIntersect(t *testing.T) {
	var tests = []struct {
		s    []int
		u    []int
		want string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{}, []int{1}, "{}"},
		{[]int{1}, []int{2}, "{}"},
		{[]int{1, 2}, []int{3, 4}, "{}"},
		{[]int{1, 3}, []int{3, 4}, "{3}"},
		{[]int{100, 14, 59}, []int{100, 14, 59}, "{14 59 100}"},
		{[]int{100, 14, 59}, []int{25, 43, 67}, "{}"},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("s.IntersectWith(t):\ns: %v\nt: %v", test.s, test.u)
		var s, u IntSet
		for _, element := range test.s {
			s.Add(element)
		}
		for _, element := range test.u {
			u.Add(element)
		}
		s.IntersectWith(&u)
		got := s.String()
		if got != test.want {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}

func TestDifference(t *testing.T) {
	var tests = []struct {
		s    []int
		u    []int
		want string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{}, []int{1}, "{}"},
		{[]int{1, 2}, []int{2}, "{1}"},
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 5}, "{1 4}"},
		{[]int{100, 14, 59}, []int{100, 14, 59}, "{}"},
		{[]int{100, 14, 59}, []int{25, 43, 67}, "{14 59 100}"},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("s.DifferenceWith(t):\ns: %v\nt: %v", test.s, test.u)
		var s, u IntSet
		for _, element := range test.s {
			s.Add(element)
		}
		for _, element := range test.u {
			u.Add(element)
		}
		s.DifferenceWith(&u)
		got := s.String()
		if got != test.want {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}

func TestSymetricDifference(t *testing.T) {
	var tests = []struct {
		s    []int
		u    []int
		want string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{}, []int{1}, "{1}"},
		{[]int{1, 2}, []int{2}, "{1}"},
		{[]int{1, 2, 4, 5}, []int{2, 3, 5}, "{1 3 4}"},
		{[]int{100, 14, 59}, []int{100, 14, 59}, "{}"},
		{[]int{100, 14, 59}, []int{25, 43, 67}, "{14 25 43 59 67 100}"},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("s.SymetricDifference(t):\ns: %v\nt: %v", test.s, test.u)
		var s, u IntSet
		for _, element := range test.s {
			s.Add(element)
		}
		for _, element := range test.u {
			u.Add(element)
		}
		s.SymetricDifference(&u)
		got := s.String()
		if got != test.want {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}
