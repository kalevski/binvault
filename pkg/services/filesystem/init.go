package filesystem

import "binvault/pkg/env"

func Init() {
	dataPath := env.GetPath("DATA_PATH")
	CreateFolder(dataPath, FOLDER_BUCKETS)
	CreateFolder(dataPath, FOLDER_PUBLIC)
	CreateFolder(dataPath, FOLDER_TEMP)
}
