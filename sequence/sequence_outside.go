package sequence

import (
	"fmt"
)

// List 数组初始化信息
type List struct {
	E        []int
	Capacity int
}

// NewList 初始化数据
func (l *List) NewList(e []int) (*List, error) {
	if len(e) == 0 {
		return &List{},
			fmt.Errorf("initialization must define attributes")
	}
	return &List{
		E:        e,
		Capacity: len(e),
	}, nil
}

// Insert 插入 移动n-i+1个元素元素  时间复杂度O(n)
func (l *List) Insert(index int, element int) (List, error) {
	if index > len(l.E)-1 {
		return List{}, fmt.Errorf(" index out of range")
	}
	copyList := make([]int, len(l.E))
	copy(copyList, l.E)
	copyList[index] = element
	data := append(copyList[:index+1], l.E[index:]...)
	l.E = data
	l.Capacity = len(data)
	return List{E: data, Capacity: len(data)}, nil
}

// Delete 删除  时间复杂度O(n)
func (l *List) Delete(element int) (List, error) {
	if l.Capacity < 0 {
		return List{}, fmt.Errorf("length less than 0 ")
	}
	for i := 0; i <= l.Capacity-1; i++ {
		if l.E[i] == element {
			copyList := make([]int, len(l.E))
			copy(copyList, l.E)
			data := append(copyList[:i], copyList[i+1:]...)
			return List{
				E:        data,
				Capacity: len(data),
			}, nil
		}
	}
	return List{
		E:        l.E,
		Capacity: l.Capacity,
	}, fmt.Errorf("not found %d", element)
}

// GettingElements 判断这个元素是否存在及获取下标  时间复杂度O(n)
func (l *List) GettingElements(element int) (int, bool) {
	set := make(map[any]int, len(l.E))
	for key, value := range l.E {
		set[value] = key
	}
	if _, ok := set[element]; ok {
		return set[element], true
	}
	return element, false
}

// Sort 希尔排序 时间复杂度O(n log n)
func (l *List) Sort() bool {
	// 获取排序得间隔
	n := len(l.E) / 2
	if n == 0 {
		return false
	}
	//间隔大于0则开始
	for n > 0 {
		//从间隔n开始遍历
		for i := n; i < len(l.E); i++ {
			temp := l.E[i]
			// 当当前位置间隔大于n，且当前元素小于前一个间隔得元素则开始交换
			for i >= n && temp < l.E[i-n] {
				// 前一个元素后移
				l.E[i] = l.E[i-n]
				// 更新索引
				i -= n
			}
			// 保存得元素插入正确得位置
			l.E[i] = temp
		}
		// 间隔缩小
		n /= 2
	}
	return true
}

// Output 输出所有元素
func (l *List) Output() {
	for i := 0; i < l.Capacity; i++ {
		fmt.Printf("index:%d price:%d \n", i, l.E[i])
	}
}
