package fhandler

import (
	"gofup"
)

// FileHandler -
type FileHandler struct {
	gofup.FileInfo
}

// NewFUpload factory
func NewFUpload(address string) (gofup.Gofup, error) {
	return &FileHandler{
		FileInfo: gofup.FileInfo{
			Address: address,
		},
	}, nil
}
