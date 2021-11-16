// WorkerPool is a contract for Worker Pool implementation
package workerpool

import (
	"log"
)

type T = interface{}

type WorkerPool interface {
	Run()
	AddTask(task func())
}

type workerPool struct {
	maxWorker   int
	queuedTaskC chan func()
}

func NewWorkerPool(maxWorker int) *workerPool {
	wp := &workerPool{
		maxWorker:   maxWorker,
		queuedTaskC: make(chan func()),
	}
	return wp
}

func (wp *workerPool) GetTotalQueuedTask() int {
	return len(wp.queuedTaskC)
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		wID := i + 1
		log.Printf("[WorkerPool] Worker %d has been spawned", wID)

		go func(workerID int) {
			for task := range wp.queuedTaskC {
				log.Printf("[WorkerPool] Worker %d start processing task", wID)
				task()
				log.Printf("[WorkerPool] Worker %d finish processing task", wID)
			}
		}(wID)
	}
}

func (wp *workerPool) AddTask(task func()) {
	wp.queuedTaskC <- task
}
