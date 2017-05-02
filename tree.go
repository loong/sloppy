package sloppy

import (
	"strings"
)

// pathTree is a tree data structure for efficient storing and
// retrieving of available endpoint paths
type pathTree struct {
	root *node
}

// newPathTree creates a new and empty pathTree
func newPathTree() *pathTree {
	return &pathTree{
		root: &node{
			children: make(map[string]*node),
		},
	}
}

// addPath parses a path and add it to the pathTree data structure
func (t *pathTree) addPath(path string) {
	curr := t.root
	for _, p := range splitPath(path) {
		next, ok := curr.children[p]
		if !ok {
			next = &node{p, make(map[string]*node)}
			curr.children[p] = next
		}

		curr = next
	}
}

// string converts the pathTree into a string representation and can
// be used to inspect the tree's structure
func (t *pathTree) string() string {
	return t.root.string()
}

func splitPath(path string) []string {
	parts := []string{}
	for _, p := range strings.Split(path, "/") {
		if p != "" {
			parts = append(parts, p)
		}
	}

	return parts
}
