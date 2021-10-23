package main

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/m-mizutani/goerr"
)

type output func(w io.Writer, node *Node) error

func OutputJson(w io.Writer, node *Node) error {
	raw, err := json.Marshal(node)
	if err != nil {
		return goerr.Wrap(err, "Failed to marshal node")
	}

	if _, err := w.Write(raw); err != nil {
		return goerr.Wrap(err)
	}
	if _, err := w.Write([]byte("\n")); err != nil {
		return goerr.Wrap(err)
	}

	return nil
}

func OutputTree(w io.Writer, node *Node) error {
	if node == nil {
		w.Write([]byte("not found\n"))
	}
	return outputTree(w, node, 0)
}

func outputTree(w io.Writer, node *Node, margin int) error {
	var indent string
	if margin > 0 {
		indent = strings.Repeat(" ", margin) + "+- "
	}

	if _, err := w.Write([]byte(indent + node.Name + "\n")); err != nil {
		return goerr.Wrap(err)
	}

	for _, edge := range node.DependedBy {
		if err := outputTree(w, edge, margin+2); err != nil {
			return err
		}
	}

	return nil
}
