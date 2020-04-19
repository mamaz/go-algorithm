package quicksort

type Quicksort struct{}

func (q Quicksort) Sort(nums []int) []int {
	return q.qSort(nums, 0, len(nums)-1)
}

func (q Quicksort) qSort(nums []int, low, high int) []int {
	if low < high {
		partition := q.separate(nums, low, high)

		q.qSort(nums, 0, partition-1)
		q.qSort(nums, partition+1, high)
	}
	return nums
}

func (q Quicksort) separate(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1

	for index := low; index < high; index++ {
		if nums[index] <= pivot {
			i++
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]

	return i + 1
}
