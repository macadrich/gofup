package upload

// File represents file details
type File struct {
	Name string // filename
	Size int64  // file size
	Hash string // hash file
}

const (
	// BUFFERSIZE size
	BUFFERSIZE = 1024
)
