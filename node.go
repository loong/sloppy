package sloppy

import "strings"

// Node is a data structure to contain values in the TreePath
type node struct {
	value    string
	children map[string]*node
}

// String recursively converts nodes and its decendants into a string
// representation
func (n *node) string() string {
	return n.stringTraverse("")
}

func (n *node) stringTraverse(prefix string) string {
	output := prefix + "└───" + strings.Replace(n.value, "{}", "{ wildcard }", -1)

	for _, v := range n.children {
		output += "\n" + v.stringTraverse(prefix+"    ")
	}

	return output
}
