package filesystem

import (
	"binvault/pkg/models"
	"path/filepath"
)

func GetFileType(filename string) models.FileType {
	extension := GetFileExtension(filename)
	switch extension {
	case "jpg", "jpeg", "png":
		return models.FileType_Image
	case "txt":
		return models.FileType_Text
	}
	return models.FileType_Undefined
}

func GetFileExtension(filename string) string {
	extesion := filepath.Ext(filename)
	if len(extesion) > 0 {
		extesion = extesion[1:]
	}
	return extesion
}
