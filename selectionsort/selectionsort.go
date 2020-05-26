// Package selectionsort implements Selection Sort for int
package selectionsort

// SelectionSort is struct implementing Selection Sort
type SelectionSort struct{}

// Sort implements Selection sort
func (s SelectionSort) Sort(nums []int) []int {
	for step := range nums {
		minIndex := step
		for i := step + 1; i < len(nums); i++ {
			if nums[i] < nums[minIndex] {
				minIndex = i
			}
		}
		nums[step], nums[minIndex] = nums[minIndex], nums[step]
	}
	return nums
}
