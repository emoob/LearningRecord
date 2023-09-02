package leetcode

func twoSum(nums []int, target int) []int {
	prevNums := map[int]int{}
	for i, num := range nums {
		// 获取差值
		targetNum := target - num
		targetNumIndex, ok := prevNums[targetNum]
		// 看这个字典是否有差值
		if ok {
			return []int{targetNumIndex, i}
		}
		// 不存在则存入
		prevNums[num] = i
	}
	return []int{}
}
