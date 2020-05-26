// Package mergesort is a package implementing merge sort algorithm for int
package mergesort

// MergeSort is struct for Merge Sort
type MergeSort struct{}

// Sort executes merge sort
func (m MergeSort) Sort(nums []int) []int {
	return m.mSort(nums, 0, len(nums)-1)
}

func (m MergeSort) mSort(nums []int, left, right int) []int {
	if left < right {
		mid := (left + (right - 1)) / 2

		m.mSort(nums, left, mid)
		m.mSort(nums, mid+1, right)
		m.merge(nums, left, mid, right)
	}
	return nums
}

func (m MergeSort) merge(nums []int, left, mid, right int) {
	n1 := mid - left + 1
	n2 := right - mid

	leftSide := []int{}
	rightSide := []int{}

	for i := 0; i < n1; i++ {
		leftSide = append(leftSide, nums[left+i])
	}
	for i := 0; i < n2; i++ {
		rightSide = append(rightSide, nums[mid+1+i])
	}

	i := 0
	j := 0
	k := left

	for i < len(leftSide) && j < len(rightSide) {
		if leftSide[i] <= rightSide[j] {
			nums[k] = leftSide[i]
			i++
		} else {
			nums[k] = rightSide[j]
			j++
		}
		k++
	}

	for i < len(leftSide) {
		nums[k] = leftSide[i]
		i++
		k++
	}

	for j < len(rightSide) {
		nums[k] = rightSide[j]
		j++
		k++
	}
}
