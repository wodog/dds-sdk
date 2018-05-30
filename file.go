package dds

import "time"

type File struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Bucket    string    `json:"bucket"`
	Size      int64     `json:"size"`
	Path      string    `json:"path"`
	Type      string    `json:"type"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
