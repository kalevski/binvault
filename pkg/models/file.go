package models

import "time"

type File struct {
	Name       string     `json:"name"`
	Bucket     string     `json:"bucket"`
	Size       int64      `json:"size"`
	Extension  string     `json:"extension"`
	Type       FileType   `json:"type"`
	CreatedAt  time.Time  `json:"createdAt"`
	Visibility Visibility `json:"visibility"`
}
