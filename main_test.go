package main

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		in   []int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 2, []int{2}},    // removing the head
		{[]int{1, 2}, 1, []int{1}},    // removing the tail
		{[]int{1, 2}, 5, []int{1, 2}}, // n longer than list: unchanged
	}
	for _, c := range cases {
		got := toSlice(RemoveNthFromEnd(fromSlice(c.in), c.n))
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("RemoveNthFromEnd(%v,%d) = %v, want %v", c.in, c.n, got, c.want)
		}
	}
}
