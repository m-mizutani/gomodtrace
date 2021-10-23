package main

import "github.com/m-mizutani/goerr"

var (
	ErrInvalidFormat = goerr.New("invalid input")
	ErrInvalidOutput = goerr.New("invalid output format")
)
