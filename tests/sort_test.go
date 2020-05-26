package tests

import (
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
	i.InsertionSort{},
	q.Quicksort{},
	m.MergeSort{},
}

func cases() []map[string][]int {
	return []map[string][]int{
		{
			"input":    []int{9, 1, 8, 6, 2, 4},
			"expected": []int{1, 2, 4, 6, 8, 9},
		},
		{
			"input":    []int{9, 8, 6, 4, 2, 1},
			"expected": []int{1, 2, 4, 6, 8, 9},
		},
		{
			"input":    []int{1, 2, 4, 6, 8, 9},
			"expected": []int{1, 2, 4, 6, 8, 9},
		},
		{
			"input":    []int{},
			"expected": []int{},
		},
		{
			"input":    []int{0},
			"expected": []int{0},
		},
		{
			"input":    []int{-1, -2, -4, -6, -8, -9},
			"expected": []int{-9, -8, -6, -4, -2, -1},
		},
		{
			"input":    longRandNums(),
			"expected": []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 9},
		},
	}
}

func longRandNums() []int {
	return []int{7, 9, 4, 3, 7, 4, 8, 9, 0, 7, 5, 8, 8, 1, 4, 5, 8, 6, 8, 2, 8, 4, 9, 8, 8, 1, 5, 8, 9, 0, 6, 7, 7, 3, 1, 5, 0, 2, 5, 3, 1, 2, 0, 3, 4, 2, 1, 4, 8, 7, 0, 0, 6, 8, 6, 6, 8, 2, 7, 8, 3, 4, 6, 1, 4, 8, 4, 7, 7, 6, 2, 4, 5, 9, 6, 2, 3, 8, 7, 8, 7, 6, 8, 9, 9, 5, 4, 2, 4, 7, 5, 3, 1, 0, 2, 6, 9, 7, 2, 8}
}

// doTheTest is the actual test function
func doTheTest(algo sort.Sort, testcase map[string][]int, t *testing.T) {
	input := testcase["input"]
	expected := testcase["expected"]
	res := algo.Sort(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Incorrect result %v,  should be %v", res, expected)
	}
}

func TestSort(t *testing.T) {
	for _, algo := range algos {
		// some of the algo are in place
		// so we need to produce new cases each time
		for _, testcase := range cases() {
			doTheTest(algo, testcase, t)
		}
	}
}

func BenchmarkBubbleSort(testB *testing.B) {
	nums := longRandNums()
	for i := 0; i < testB.N; i++ {
		b.BubbleSort{}.Sort(nums)
	}
}

func BenchmarkSelectionSortSort(testB *testing.B) {
	nums := longRandNums()
	for i := 0; i < testB.N; i++ {
		s.SelectionSort{}.Sort(nums)
	}
}

func BenchmarkSortQuicksort(testB *testing.B) {
	nums := longRandNums()
	for i := 0; i < testB.N; i++ {
		q.Quicksort{}.Sort(nums)
	}
}

func BenchmarkMergeSort(testB *testing.B) {
	nums := longRandNums()
	mg := m.MergeSort{}
	for i := 0; i < testB.N; i++ {
		mg.Sort(nums)
	}
}
