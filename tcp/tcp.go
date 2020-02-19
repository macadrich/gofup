package tcp

// Protocol -
type Protocol interface {
	Start() error
	End() error
}
