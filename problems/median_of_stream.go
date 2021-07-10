package problems

import (
	"container/heap"
)

type MedianFinder struct {
	bottom *MaxIntHeap
	top    *MinIntHeap
}

func NewMedianFinder() MedianFinder {
	return MedianFinder{&MaxIntHeap{}, &MinIntHeap{}}
}

func (this *MedianFinder) AddNum(num int) {
	if len(*this.bottom) == 0 {
		heap.Push(this.bottom, num)
		return
	}

	if num > this.bottom.Peek() {
		heap.Push(this.top, num)
	} else {
		heap.Push(this.bottom, num)
	}

	if this.top.Len() > this.bottom.Len() {
		oldTopLowest := heap.Pop(this.top).(int)
		heap.Push(this.bottom, oldTopLowest)
	}
	if this.bottom.Len() > this.top.Len()+1 {
		oldBottomHighest := heap.Pop(this.bottom).(int)
		heap.Push(this.top, oldBottomHighest)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.bottom.Len() == 0 {
		return 0
	}

	tot := this.top.Len() + this.bottom.Len()
	if tot%2 == 0 {
		return (float64(this.top.Peek()) + float64(this.bottom.Peek())) / float64(2)
	}
	return float64(this.bottom.Peek())
}
