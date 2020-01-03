package main

import (
	"sync"
	"testing"
)

func Test_throttle_for_30_tasks(t *testing.T) {
	tasks := make([]*task, 0)
	for i := 0; i < 30; i++ {
		tasks = append(tasks, &task{ID: i})
	}
	var wg sync.WaitGroup
	throttle(tasks, &wg, 3)
	wg.Wait()

	for _, task := range tasks {
		if !task.Done {
			t.Errorf("task %v not completed", task.ID)
		}
	}
}
