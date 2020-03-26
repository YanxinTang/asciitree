package asciitree_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/YanxinTang/asciitree"
)

func Example() {
	New := asciitree.New
	basicTree := New(
		"root",
		New("child1"),
		New("child2",
			New("grandchild1"),
			New("grandchild2"),
		),
		New("child3"),
	)
	basicTree.PrintTree(os.Stdout)
	// Output:
	// root
	//  ├── child1
	//  ├─┬ child2
	//  │ ├── grandchild1
	//  │ └── grandchild2
	//  └── child3
}

func Example_advancedUsage() {
	TitleNode := func(text string, children ...*asciitree.ASCIITree) *asciitree.ASCIITree {
		title := fmt.Sprintf("%s\n%s", strings.ToUpper(text), strings.Repeat("=", len(text)))
		return asciitree.New(title, children...)
	}

	New := asciitree.New
	titleTree := New(
		"root",
		TitleNode("child1"),
		TitleNode("child2",
			asciitree.New("grandchild1"),
			asciitree.New("grandchild2"),
		),
		asciitree.New("child3"),
	)
	titleTree.PrintTree(os.Stdout)
	// Output:
	// root
	//  ├── CHILD1
	//  │   ======
	//  ├─┬ CHILD2
	//  │ │ ======
	//  │ ├── grandchild1
	//  │ └── grandchild2
	//  └── child3
}

func ExampleAdd() {
	root := asciitree.New("root")
	child1 := asciitree.New("child")
	child2 := asciitree.New("child2")
	root.Add(child1, child2)
	root.PrintTree(os.Stdout)
	// Output:
	// root
	//  ├── child
	//  └── child2
}
