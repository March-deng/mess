package timer

import (
	"container/heap"
	"log"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {
	s := NewSchedule(5 * time.Second)
	defer s.Close()
	id := s.AddScheduleTask(time.Now().Add(19*time.Second), func() {
		log.Println("执行任务, 时间是:", time.Now().Unix())
	})
	log.Println("任务ID 1 是:", id)
	id = s.AddScheduleTask(time.Now().Add(30*time.Second), func() {
		log.Println("执行任务, 时间是:", time.Now().Format(time.RFC3339))
	})
	log.Println("任务ID 2 是:", id)
	s.CancelTask(id)
	log.Println("取消任务:", id)
	id = s.AddScheduleTask(time.Now().Add(15*time.Second), func() {
		log.Println("执行任务, 时间是:", time.Now().Format(time.RFC3339))
	})
	log.Println("任务ID 3 是:", id)
	time.Sleep(5 * time.Minute)
}

func TestHeap(t *testing.T) {
	h := &taskHeap{2, 1, 5}
	heap.Init(h)
	ok, value := h.Peek()
	if ok {
		log.Println("peek first:", value)
	}
	h.Insert(3)
	for h.Len() > 0 {
		log.Println("value is:", h.PopTask())
	}
	log.Println(*h)
}
