package array

import (
	"github.com/chentaihan/container/common"
)

func NewArraySortInt(cap int) *ArraySortInt {
	if cap < 0 {
		cap = 0
	}
	return &ArraySortInt{
		buf:  make([]int, cap),
		size: 0,
	}
}

type ArraySortInt struct {
	buf  []int
	size int
}

func (as *ArraySortInt) Add(val int) {
	as.checkCap()
	index, _ := common.IntBinarySearchPos(as.buf[:as.size], val)
	for i := as.size - 1; i >= index; i-- {
		as.buf[i+1] = as.buf[i]
	}
	as.buf[index] = val
	as.size++
}

func (as *ArraySortInt) Get(index int) (int, bool) {
	if index >= as.size || index < 0 {
		return -1, false
	}
	return as.buf[index], true
}

func (as *ArraySortInt) Index(val int) int {
	return common.IntBinarySearch(as.GetArray(), val)
}

func (as *ArraySortInt) RemoveIndex(index int) (int, bool) {
	if index >= as.size || index < 0 {
		return -1, false
	}
	val := as.buf[index]
	for i := index; i < as.size-1; i++ {
		as.buf[index] = as.buf[index+1]
	}
	as.size--
	return val, true
}

func (as *ArraySortInt) Remove(value int) int {
	index := common.IntBinarySearch(as.GetArray(), value)
	if index < 0 {
		return 0
	}
	start, end := index, index
	for start >= 0 && as.buf[start] == value {
		start--
	}
	for end < as.size && as.buf[end] == value {
		end++
	}
	as.buf = append(as.buf[:start+1], as.buf[end:]...)
	count := end - start - 1
	as.size -= count
	return count
}

func (as *ArraySortInt) Len() int {
	return as.size
}

func (as *ArraySortInt) checkCap() {
	if as.size < cap(as.buf) {
		return
	}
	newCap := 0
	if cap(as.buf) > 1024*1024 { //大于1M，每次扩大1M
		newCap = as.size + 1024*1024
	} else if cap(as.buf) < 4 { //小于4，就是4
		newCap = 4
	} else { //2倍扩容
		newCap = cap(as.buf) << 1
	}
	buf := make([]int, newCap)
	copy(buf, as.GetArray())
	as.buf = buf
}

func (as *ArraySortInt) Clear() {
	as.buf = []int{}
	as.size = 0
}

func (as *ArraySortInt) GetArray() []int {
	return as.buf[:as.size]
}

func (as *ArraySortInt) Copy() []int {
	list := make([]int, as.size)
	copy(list, as.GetArray())
	return list
}
