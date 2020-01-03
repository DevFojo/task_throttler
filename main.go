package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	ID   int
	Done bool
}

func (t *task) run() {
	fmt.Printf("running task %v \n", t.ID)
	t.Done = true
}

func throttle(tasks []*task, wg *sync.WaitGroup, count int) {
	runningTasksChannel := make(chan int, count)
	for i := 0; i < len(tasks); i++ {
		task := tasks[i]
		runningTasksChannel <- task.ID
		wg.Add(1)
		go runTask(task, wg, runningTasksChannel)
	}
}

func runTask(task *task, wg *sync.WaitGroup, runningTasksChannel chan int) {
	defer wg.Done()
	time.Sleep(time.Second)
	task.run()
	<-runningTasksChannel
}
