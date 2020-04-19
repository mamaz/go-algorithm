package bubblesort

// BubbleSort is
type BubbleSort struct{}

// Sort is
func (b BubbleSort) Sort(nums []int) []int {
	for o := 0; o < len(nums); o++ {
		for i := 0; i < len(nums); i++ {
			if nums[i] > nums[o] {
				temp := nums[o]
				nums[o] = nums[i]
				nums[i] = temp
			}
		}
	}
	return nums
}
