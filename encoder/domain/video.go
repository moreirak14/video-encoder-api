package domain

import "time"

type Video struct {
	ID         string
	ResourceID string
	FilePath   string
	CreateadAt time.Time
}
