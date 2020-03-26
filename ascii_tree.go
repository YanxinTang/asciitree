package asciitree

import (
	"fmt"
	"io"
	"strings"
)

// ASCIITree is a tree node
type ASCIITree struct {
	Text     string       // Node text
	Parent   *ASCIITree   // Parent pointer. root node will get nil
	children []*ASCIITree // Pointer set of hildern. root node will get empty slice
}

// New returns a pointer to an ASCIITree struct
// @param text string node text
func New(text string, children ...*ASCIITree) *ASCIITree {
	tree := &ASCIITree{Text: text}
	tree.Add(children...)
	return tree
}

/*
PrintTree print the tree to an output stream
*/
func (t *ASCIITree) PrintTree(w io.Writer) {
	ancestorPrefix := ""
	for _, parent := range t.Ancestors() {
		if parent.Level() <= 0 {
			continue
		}
		if parent.Last() {
			ancestorPrefix += "  "
		} else {
			ancestorPrefix += " │"
		}
	}

	myPrefix := ""
	multilinePrefix := ""
	if t.Level() > 0 {
		if t.Last() {
			if t.Empty() {
				myPrefix += " └── "
				multilinePrefix += "     "
			} else {
				myPrefix += " └─┬ "
				multilinePrefix += " └─┬ "
			}
		} else {
			if t.Empty() {
				myPrefix += " ├── "
				multilinePrefix += " │   "
			} else {
				myPrefix += " ├─┬ "
				multilinePrefix += " │ │ "
			}
		}
	}

	if t.Text != "" {
		lines := strings.Split(t.Text, "\n")
		fmt.Fprintf(w, "%s%s%s\n", ancestorPrefix, myPrefix, lines[0])
		for _, line := range lines[1:] {
			fmt.Fprintf(w, "%s%s%s\n", ancestorPrefix, multilinePrefix, line)
		}
	}

	for _, child := range t.children {
		child.PrintTree(w)
	}
}

// Add appends an ASCIITree node to children property
func (t *ASCIITree) Add(children ...*ASCIITree) {
	for _, child := range children {
		child.Parent = t
	}
	t.children = append(t.children, children...)
}

// Index returns the position of this among its brothers
func (t *ASCIITree) Index() int {
	for i, child := range t.Parent.children {
		if child == t {
			return i
		}
	}
	return -1
}

// Last returns true if this is the last child
func (t *ASCIITree) Last() bool {
	if t.Parent == nil {
		return true
	}
	return t.Index() == len(t.Parent.children)-1
}

// Level returns the node leven
func (t *ASCIITree) Level() int {
	if t.Parent == nil {
		// if the root node does not have text, it will be considered level -1
		// so that all it's children will be roots.
		if t.Text == "" {
			return -1
		}
		return 0
	}
	return t.Parent.Level() + 1
}

// Empty returns
func (t *ASCIITree) Empty() bool {
	return len(t.children) == 0
}

// Ancestors returns parents from root to this node
func (t *ASCIITree) Ancestors() []*ASCIITree {
	if t.Parent == nil {
		return []*ASCIITree{}
	}
	return append(t.Parent.Ancestors(), t.Parent)
}
