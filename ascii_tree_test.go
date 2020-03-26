package asciitree

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	fmt.Println("test")
	rand.Seed(time.Now().UnixNano())
}

func TestAdd(t *testing.T) {
	root := ASCIITree{}
	child := ASCIITree{}
	root.Add(&child)

	assert.Contains(t, root.children, &child)
	assert.Equal(t, &root, child.Parent)
}

func TestIndex(t *testing.T) {
	root := ASCIITree{}
	N := rand.Intn(10)
	for i := 0; i < N; i++ {
		root.Add(&ASCIITree{})
	}
	child := ASCIITree{}
	root.Add(&child)
	M := rand.Intn(10)
	for i := 0; i < M; i++ {
		root.Add(&ASCIITree{})
	}
	assert.Equal(t, N, child.Index())
}

func TestLast(t *testing.T) {
	root := ASCIITree{}
	N := rand.Intn(10)
	for i := 0; i < N; i++ {
		root.Add(&ASCIITree{})
	}
	middleChild := ASCIITree{}
	root.Add(&middleChild)
	M := rand.Intn(10)
	lastChild := ASCIITree{}
	for i := 0; i < M; i++ {
		root.Add(&ASCIITree{})
	}
	root.Add(&lastChild)
	assert.Equal(t, true, root.Last())
	assert.Equal(t, false, middleChild.Last())
	assert.Equal(t, true, lastChild.Last())
}

func TestLevel(t *testing.T) {
	// root not without text
	{
		level0 := ASCIITree{}
		level1 := ASCIITree{}
		level2 := ASCIITree{}
		level1.Add(&level2)
		level0.Add(&level1)

		assert.Equal(t, -1, level0.Level())
		assert.Equal(t, 0, level1.Level())
		assert.Equal(t, 1, level2.Level())
	}

	// root not with text
	{
		level0 := ASCIITree{Text: "root"}
		level1 := ASCIITree{}
		level2 := ASCIITree{}
		level1.Add(&level2)
		level0.Add(&level1)

		assert.Equal(t, 0, level0.Level())
		assert.Equal(t, 1, level1.Level())
		assert.Equal(t, 2, level2.Level())
	}
}

func TestEmpty(t *testing.T) {
	notEmptyTree := ASCIITree{
		children: []*ASCIITree{&ASCIITree{}},
	}
	emptyTree := ASCIITree{}

	assert.Equal(t, false, notEmptyTree.Empty())
	assert.Equal(t, true, emptyTree.Empty())
}

func TestAncestors(t *testing.T) {
	root := ASCIITree{}
	parent1 := ASCIITree{}
	parent2 := ASCIITree{}
	child := ASCIITree{}
	root.Add(&parent1)
	parent1.Add(&parent2)
	parent2.Add(&child)
	want := []*ASCIITree{
		&root,
		&parent1,
		&parent2,
	}
	assert.Equal(t, want, child.Ancestors())
}

func TestPrintTree(t *testing.T) {
	basicTree := New(
		"root",
		New("child1"),
		New("child2",
			New("grandchild1"),
			New("grandchild2"),
		),
		New("child3"),
	)

	bigTree := New("root",
		New("child1",
			New("child1.1"),
			New("child1.2"),
			New("child1.3"),
			New("child1.4",
				New("child1.4.1"),
				New("child1.4.2"),
				New("child1.4.3"),
			),
			New("child1.6"),
			New("child1.7"),
			New("child1.8"),
		),
		New("child2",
			New("child2.1",
				New("child2.1.1",
					New("child.2.1.1.1"),
					New("child.2.1.1.2"),
					New("child.2.1.1.3",
						New("child.2.1.1.3.1"),
						New("child.2.1.1.3.2"),
					),
					New("child.2.1.1.4"),
					New("child.2.1.1.5"),
					New("child.2.1.1.5"),
				),
				New("child2.1.2"),
			),
		),
		New("child3"),
	)

	diff(t, basicTree, "basic.expected.txt")
	diff(t, bigTree, "big.expected.txt")
}

func diff(t *testing.T, tree *ASCIITree, filename string) {
	var basicTreeGot bytes.Buffer
	tree.PrintTree(&basicTreeGot)
	basicTreeWant, err := ioutil.ReadFile(filepath.Join("testdata", filename))
	if err != nil {
		t.Errorf("Couldn't open '%s', error: %v\n", filename, err)
	}
	assert.Equal(t, basicTreeWant, basicTreeGot.Bytes())
}

func TestString(t *testing.T) {
	root := New("root",
		New("1"),
		New("2",
			New("2.1"),
			New("2.2"),
		),
		New("3"),
	)
	got := root.String()
	want := `root
 ├── 1
 ├─┬ 2
 │ ├── 2.1
 │ └── 2.2
 └── 3
`
	assert.Equal(t, want, got)
}
