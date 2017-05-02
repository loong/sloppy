package sloppy

import "log"

// Suggest will suggest a path if possible
func Suggest(t *pathTree, path string) (string, bool) {
	parts := splitPath(path)

	if len(t.root.children) == 0 {
		log.Fatal("TODO: No paths added yet")
	}

	suggested := ""

	curr := t.root
	for _, p := range parts {
		nextSuggestion, next := suggestNext(p, curr)

		if next == nil {
			return "", false
		}

		if nextSuggestion != "" {
			suggested += "/" + nextSuggestion
		}
		curr = next
	}

	return suggested, true
}

func suggestNext(part string, n *node) (string, *node) {
	suggestion := part

	// if this part of the path is correct, move on to the next
	next, ok := n.children[part]
	if ok {
		return suggestion, next
	}

	min := 100
	for k, v := range n.children {
		// This is good for the following case:
		//
		// Let there be /likes and /comments. If we request
		// /c, this function will suggest /likes due to the
		// fact that comments edit distance is larger due to
		// the length. Hence we cut the lenght to a similar
		// size first.
		comp := k
		if len(k) > len(part) {
			comp = k[:len(part)]
		}

		dist := levDist(comp, part)
		if dist < min {
			min = dist
			next = v
			suggestion = k
		}
	}

	return suggestion, next
}
