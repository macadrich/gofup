package upload

import "gofup"

// Upload -
type Upload struct {
	Dir     string
	Address string
}

// NewFUpload factory
func NewFUpload(dir, address string) (gofup.FUpload, error) {
	return &Upload{
		Dir:     dir,
		Address: address,
	}, nil
}
