package dds

import "time"

type File struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Bucket    string    `json:"bucket"`
	FileID    string    `json:"fileID"`
	Size      int64     `json:"size"`
	Type      string    `json:"type"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
