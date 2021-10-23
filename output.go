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
	return outputTree(w, node, 0, false)
}

func outputTree(w io.Writer, node *Node, margin int, first bool) error {
	var text string
	var indent string
	if margin > 0 {
		if first {
			text += " <-+- "
		} else {
			indent = strings.Repeat(" ", margin)
			text += "   +- "
		}
	}
	text += node.Name

	if node.Looped {
		text += " (looped)"
	}

	if _, err := w.Write([]byte(indent + text)); err != nil {
		return goerr.Wrap(err)
	}

	if len(node.DependedBy) == 0 {
		if _, err := w.Write([]byte("\n")); err != nil {
			return goerr.Wrap(err)
		}
	} else {
		for _, edge := range node.DependedBy {
			if err := outputTree(w, edge, margin+len(text), edge == node.DependedBy[0]); err != nil {
				return err
			}
		}
	}

	return nil
}
