package compression

import (
	"binvault/pkg/clients/filesystem"
	"binvault/pkg/tasks"
)

var tempFolder = "temp"

func Init() {
	tasks.RegisterInitializer(TaskInitializer)
	tasks.RegisterHandler("compress", handleFile)
	go filesystem.WatchFolder(tempFolder, OnFileCreate)
}

func TaskInitializer() []tasks.Task {
	paths := filesystem.GetFiles(tempFolder)
	var taskList []tasks.Task
	for _, path := range paths {
		taskList = append(taskList, tasks.Task{
			Name: "compress",
			Data: path,
		})
	}
	return taskList
}

func OnFileCreate(path string, action string) {
	tasks.PushTask(tasks.Task{
		Name: "compress",
		Data: path,
	})
}
