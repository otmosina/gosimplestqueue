package queue

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Task interface {
	Exec() error
}

type Queue struct {
	counter int32
}

func New() *Queue {
	q := &Queue{}
	return q
}

func (q *Queue) Add(task Task, t time.Time) {
	go func() {
		scheduledTime := t
		id := atomic.AddInt32(&q.counter, 1)
		<-time.After(t.Sub(time.Now()))
		fmt.Printf("task queued num %d start execute at         %v\n", id, time.Now())
		fmt.Printf("task queued num %d start planned at         %v\n", id, scheduledTime)
		task.Exec()
		atomic.StoreInt32(&q.counter, id-1)
	}()
}
