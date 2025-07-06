package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return nil
	}
	if m <= 0 {
		// ignore errors mode
		m = len(tasks) + 1
	}
	var (
		wg         sync.WaitGroup
		errCounter int32
		tasksCh    = make(chan Task, len(tasks))
	)
	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for {
				task, ok := <-tasksCh
				if !ok || atomic.LoadInt32(&errCounter) >= int32(m) {
					return
				}
				err := task()
				if err != nil {
					atomic.AddInt32(&errCounter, 1)
				}
			}
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&errCounter) >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
