package tests

import (
	"fmt"
	b "mamazo/sorting/bubblesort"

	m "mamazo/sorting/mergesort"

	s "mamazo/sorting/selectionsort"

	i "mamazo/sorting/insertionsort"

	q "mamazo/sorting/quicksort"

	"reflect"

	"mamazo/sorting/sort"

	"testing"
)

var algos = []sort.Sort{
	b.BubbleSort{},
	s.SelectionSort{},
	m.MergeSort{},
	i.InsertionSort{},
	q.Quicksort{},
}

func cases() []map[string][]int {
	return []map[string][]int{
		map[string][]int{
			"input":    []int{9, 1, 8, 6, 2, 4},
			"expected": []int{1, 2, 4, 6, 8, 9},
		},
		map[string][]int{
			"input":    []int{9, 8, 6, 4, 2, 1},
			"expected": []int{1, 2, 4, 6, 8, 9},
		},
		map[string][]int{
			"input":    []int{1, 2, 4, 6, 8, 9},
			"expected": []int{1, 2, 4, 6, 8, 9},
		},
		map[string][]int{
			"input":    []int{},
			"expected": []int{},
		},
		map[string][]int{
			"input":    []int{0},
			"expected": []int{0},
		},
		map[string][]int{
			"input":    []int{-1, -2, -4, -6, -8, -9},
			"expected": []int{-9, -8, -6, -4, -2, -1},
		},
	}
}

func TestSort(t *testing.T) {
	for _, algo := range algos {
		fmt.Printf("%+v\n", algo)

		fmt.Printf("cases %+v\n", cases())
		for _, testcase := range cases() {
			input := testcase["input"]
			expected := testcase["expected"]
			fmt.Printf("input: %+v\n", input)
			fmt.Printf("expected: %+v\n", expected)
			res := algo.Sort(input)
			fmt.Printf("result: %+v\n", input)

			fmt.Println("============================")
			if !reflect.DeepEqual(input, expected) {
				t.Errorf("Incorrect result %v,  should be %v", res, expected)
			}
		}
	}
}
