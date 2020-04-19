package selectionsort

// SelectionSort is
type SelectionSort struct{}

// Sort is
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
