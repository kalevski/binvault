package models

import "time"

type Bucket struct {
	Name       string     `json:"name"`
	CreatedBy  string     `json:"createdBy,omitempty"`
	CreatedAt  time.Time  `json:"createdAt"`
	Visibility Visibility `json:"visibility"`
}
