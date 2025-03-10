package models

type Visibility string

const (
	Visibility_Public  Visibility = "public"
	Visibility_Private Visibility = "private"
)

type FileType string

const (
	FileType_Image     FileType = "image"
	FileType_Text      FileType = "text"
	FileType_Undefined FileType = "undefined"
)
