package cli

import (
	"io"
	"os"
)

type Streams struct {
	In io.Reader
	Out io.Writer
	Err io.Writer
}

func NewStreams() *Streams {
	return &Streams{
		In: os.Stdin,
		Out: os.Stdout,
		Err: os.Stderr,
	}
}

