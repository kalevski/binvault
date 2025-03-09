package compression

import (
	"binvault/pkg/cfg"
	"binvault/pkg/clients/filesystem"
	"binvault/pkg/tasks"
)

func Init() {
	tasks.RegisterInitializer(TaskInitializer)
	tasks.RegisterHandler("compress", handleFile)
	go filesystem.WatchFolder(cfg.GetVars().TEMP_DIR_NAME, OnFileCreate)
}

func TaskInitializer() []tasks.Task {
	paths := filesystem.GetFiles(cfg.GetVars().TEMP_DIR_NAME)
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
