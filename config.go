package main

import (
	"io"
	"os"

	"github.com/m-mizutani/goerr"
)

const (
	StdinMark = "-"
)

type config struct {
	Input  string
	Output string

	out    output
	reader io.ReadCloser
}

func (x *config) Setup() error {
	if x.Input == StdinMark {
		x.reader = os.Stdin
	} else {
		r, err := os.Open(x.Input)
		if err != nil {
			return goerr.Wrap(err, "open file").With("filename", x.Input)
		}
		x.reader = r
	}

	switch x.Output {
	case "tree":
		x.out = OutputTree
	case "json":
		x.out = OutputJson
	default:
		return goerr.Wrap(ErrInvalidOutput).With("output", x.Output)
	}

	return nil
}

func (x *config) Teardown() error {
	if err := x.reader.Close(); err != nil {
		return goerr.Wrap(err, "close reader")
	}

	return nil
}
