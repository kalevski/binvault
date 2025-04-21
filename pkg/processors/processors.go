package processors

import (
	"binvault/pkg/services/filesystem"
	"binvault/pkg/tasks"
)

func Init() {
	initializeTemplates()
	tasks.RegisterInitializer(TaskInitializer)
	tasks.RegisterHandler("compress", handleFile)
	go filesystem.WatchFolder(filesystem.GetFolderPath(filesystem.FOLDER_TEMP), OnFileCreate)
}

func TaskInitializer() []tasks.Task {
	paths := filesystem.GetFiles(filesystem.GetFolderPath(filesystem.FOLDER_TEMP))
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
