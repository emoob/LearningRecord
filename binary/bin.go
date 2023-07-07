package binary

// Binary /*
// @Nums 原始数据
// @NumsSize 长度
// @target 搜索的数字
type Binary struct {
	Nums     []int
	NumsSize int
	Target   int
}

func (B *Binary) Search() int {
	return B.BinnarySearch(B.Nums, B.Target, 0, B.NumsSize-1)
}

func (B *Binary) BinnarySearch(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}
	/*
		nums是一个有序数组，如果nums中间的数大于target肯定在左边那么他们起始的位置就在target和mid-1间
		反之在mid+1和长度right之间
	*/
	if nums[mid] > target {
		return B.BinnarySearch(nums, target, left, mid-1)
	} else {
		return B.BinnarySearch(nums, target, mid+1, right)
	}
}
