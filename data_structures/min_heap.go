package main

import "container/heap"

type MinHeap []int

func NewMinHeap() *MinHeap {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	return minHeap
}

func (mh *MinHeap) Len() int {
	return len(*mh)
}

func (mh *MinHeap) Empty() bool {
	return len(*mh) == 0
}

func (mh *MinHeap) Less(i, j int) bool {
	return (*mh)[i] < (*mh)[j]
}

func (mh *MinHeap) Swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}

func (mh *MinHeap) Top() int {
	return (*mh)[0]
}

func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}

func (mh *MinHeap) Pop() interface{} {
	v := (*mh)[len(*mh)-1]
	*mh = (*mh)[0 : len(*mh)-1]
	return v
}
