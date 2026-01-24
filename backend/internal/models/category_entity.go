package models

import "time"

type Category struct {
	Id          int32
	Name        string
	Slug        string
	Description string
	CreatedAt   time.Time
}
