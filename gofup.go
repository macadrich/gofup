package gofup

// Gofup -
type Gofup interface {
	SetFileFolder(string)
	Send() error
	Receive() error
}

// FileInfo -
type FileInfo struct {
	Address    string
	FileFolder string
}

// SetFileFolder -
func (f *FileInfo) SetFileFolder(dir string) {
	f.FileFolder = dir
}
