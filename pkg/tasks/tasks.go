package tasks

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
	for range workers {
		go taskExecutor()
	}
	for _, initializer := range initializers {
		for _, task := range initializer() {
			taskQueue <- task
		}
	}

}

func taskExecutor() {
	for {
		task := <-taskQueue
		if handler, ok := handlers[task.Name]; ok {
			handler(task.Data)
		}
	}
}
