package main

import (
	"bufio"
	"io"
	"strings"

	"github.com/m-mizutani/goerr"
)

func read(rd io.Reader) ([]*Dependency, error) {
	buf := bufio.NewScanner(rd)
	var deps []*Dependency

	for buf.Scan() {
		dep, err := parseLine(buf.Text())
		if err != nil {
			return nil, err
		}
		deps = append(deps, dep)
	}

	return deps, nil
}

func parseLine(line string) (*Dependency, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return nil, goerr.Wrap(ErrInvalidFormat).With("input", line)
	}

	return &Dependency{
		Src: parts[0],
		Dst: parts[1],
	}, nil
}
