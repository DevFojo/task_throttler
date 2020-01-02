package main

import (
	"fmt"
)

type task struct {
	ID int
}

func (t *task) run() {
	fmt.Printf("running task %v\n", t.ID)
}

const totalTaskCount = 30

func main() {
	tasks := make([]*task, 0)
	for i := 0; i < totalTaskCount; i++ {
		tasks = append(tasks, &task{ID: i})
	}
	throttle(tasks, 3)
}

func throttle(tasks []*task, count int) {
	runningCount := 0
	taskQueue := make([]*task, 0)
	for i := 0; i < len(tasks); i++ {
		if runningCount != count {
			tasks[i].run()
			runningCount++
			continue
		}
		taskQueue = append(taskQueue, tasks[i])
	}
}
