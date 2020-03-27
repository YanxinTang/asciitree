# asciitree

[![GoDoc](https://godoc.org/github.com/YanxinTang/asciitree?status.svg)](https://godoc.org/github.com/YanxinTang/asciitree)
[![Go Report Card](https://goreportcard.com/badge/github.com/YanxinTang/asciitree)](https://goreportcard.com/report/github.com/YanxinTang/asciitree)
[![codecov](https://codecov.io/gh/YanxinTang/asciitree/branch/master/graph/badge.svg)](https://codecov.io/gh/YanxinTang/asciitree)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/yanxintang/asciitree/Test%20and%20Build)

Package asciitree render Ascii trees from a tree struct

It is a Golang implement of Javascript package which is [https://www.npmjs.com/package/oo-ascii-tree/v/1.1.0](https://www.npmjs.com/package/oo-ascii-tree/v/1.1.0)

It implements the `oo-ascii-tree` interface as much as possible

## Usage

### Basic Example

```go
package main

import (
	"os"

	"github.com/YanxinTang/asciitree"
)

func main() {
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
}
```

And you will get :
```text
root
 ├── child1       
 ├─┬ child2       
 │ ├── grandchild1
 │ └── grandchild2
 └── child3 
```

### Multiline Support

```go
package main

import (
	"os"

	"github.com/YanxinTang/asciitree"
)

func main() {
	New := asciitree.New
	multiLineTree := New(
		"root",
		New("child1-1\nchild1-2\nchild-3"),
		New("child2",
			New("grandchild1"),
			New("grandchild2-1\ngrandchild2-2"),
		),
		New("child3"),
	)
	multiLineTree.PrintTree(os.Stdout)
}
```
Output:

```text
root
 ├── child1-1
 │   child1-2
 │   child-3
 ├─┬ child2
 │ ├── grandchild1
 │ └── grandchild2-1
 │     grandchild2-2
 └── child3
```

### Advanced Usage

```go
package main

import (
	"os"

	"github.com/YanxinTang/asciitree"
)

func main() {
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
}
```

Output:

```text
root
 ├── CHILD1
 │   ======
 ├─┬ CHILD2
 │ │ ======
 │ ├── grandchild1
 │ └── grandchild2
 └── child3
```

### API

[https://godoc.org/github.com/YanxinTang/asciitree](https://godoc.org/github.com/YanxinTang/asciitree)