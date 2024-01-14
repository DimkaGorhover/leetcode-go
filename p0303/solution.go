package p0303

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	arr := make([]int, len(nums)+1)
	arr[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		arr[i+1] = arr[i] + nums[i]
	}
	return NumArray{arr}
}

func (a NumArray) SumRange(left int, right int) int {
	return a.arr[right+1] - a.arr[left]
}
