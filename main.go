package main

import (
	"go_workerpool/workerpool"
	"log"
	"time"
)

func main() {
	// Start Worker Pool.
	totalWorker := 3
	var wp = workerpool.NewWorkerPool(totalWorker)
	wp.Run()

	type result struct {
		id    int
		value int
	}

	totalTask := 10
	resultC := make(chan result, totalTask)

	for i := 0; i < totalTask; i++ {
		id := i // this is important, anomony function can access outside variables
		wp.AddTask(func() {
			log.Printf("[main] Starting task %d", id)
			time.Sleep(5 * time.Second)
			resultC <- result{id, id * 2}
		})

	}

	for i := 0; i < totalTask; i++ {
		res := <-resultC
		log.Printf("[main] Task %d has been finished with result %d:", res.id, res.value)
	}

}
