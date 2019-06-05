package timer

import "container/heap"

type taskBucket interface {
	Peek() (bool, int64)
	PopTask() int64
	Insert(value int64)
}

//最小堆
type taskHeap []int64

func newTaskHeap() *taskHeap {
	h := &taskHeap{}
	heap.Init(h)
	return h
}

func (h taskHeap) Len() int {
	return len(h)
}

func (h taskHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h taskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *taskHeap) Push(value interface{}) {
	*h = append(*h, value.(int64))
}

func (h *taskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//查看最小堆最小的元素
func (h *taskHeap) Peek() (bool, int64) {
	if h.Len() > 0 {
		return true, (*h)[0]
	}
	return false, 0
}

//同时也删除了这个元素
func (h *taskHeap) PopTask() int64 {
	return heap.Pop(h).(int64)
}

func (h *taskHeap) Insert(value int64) {
	heap.Push(h, value)
}
