package lc0295

type MedianFinder struct {
	arr []int
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) binarySearch(value int) int {
	low := 0
	high := len(mf.arr) - 1
	mid := 0
	for low <= high {
		mid = (low + high) >> 1
		midVal := mf.arr[mid]
		if midVal == value {
			return mid
		}
		if midVal > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func (mf *MedianFinder) AddNum(num int) {
	if len(mf.arr) == 0 {
		mf.arr = append(mf.arr, num)
		return
	}
	pos := mf.binarySearch(num)

	rightPart := mf.arr[pos:]
	mf.arr = append(mf.arr[:pos], num)
	mf.arr = append(mf.arr, rightPart...)
}

func (mf *MedianFinder) FindMedian() float64 {
	if len(mf.arr) == 0 {
		return 0
	}
	mid := len(mf.arr) >> 1
	if len(mf.arr)%2 == 1 {
		return float64(mf.arr[mid])
	}
	return (float64(mf.arr[mid]) + float64(mf.arr[mid-1])) * 0.5
}
