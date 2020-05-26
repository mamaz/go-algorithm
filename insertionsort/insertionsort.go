// Package insertionsort implements Insertion Sort
package insertionsort

// InsertionSort is struct for implementing Insertion Sort
type InsertionSort struct{}

// Sort implements insertion sort
func (i InsertionSort) Sort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		value := nums[i]
		counter := i - 1

		for counter >= 0 && value < nums[counter] {
			nums[counter+1] = nums[counter]
			counter--
		}
		nums[counter+1] = value
	}
	return nums
}
