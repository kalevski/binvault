package tasks

import (
	"log"
	"sync"
	"time"
)

type Task struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

var taskQueue = make(chan Task, 10)
var initializers []TaskInitializer
var handlers = make(map[string]TaskHandler)

type TaskInitializer func() []Task
type TaskHandler func(data string)

func RegisterInitializer(initializer TaskInitializer) {
	initializers = append(initializers, initializer)
}

func RegisterHandler(name string, handler TaskHandler) {
	handlers[name] = handler
}

func PushTask(task Task) {
	taskQueue <- task
}

func Run(workers int) {
	time.Sleep(5 * time.Second)
	for _, initializer := range initializers {
		for _, task := range initializer() {
			taskQueue <- task
		}
	}

	log.Printf("Starting task workers [%d]", workers)
	var group sync.WaitGroup
	for range workers {
		group.Add(1)
		go func() {
			defer group.Done()
			for task := range taskQueue {
				handler, ok := handlers[task.Name]
				if !ok {
					log.Printf("No handler found for task [%s]", task.Name)
					continue
				}
				handler(task.Data)
			}
		}()
	}

}
