package gofup

// FUpload -
type FUpload interface {
	Send() error
	Receive() error
}
