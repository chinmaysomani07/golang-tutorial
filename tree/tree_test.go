package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeStruct(t *testing.T) {
	//assert := assert.New(t)

	noofdir := 0
	nooffiles := 0
	// want := "/Users/chinmaysomani/Desktop/gocodes/tree\n│── tempfolder\n    │── tempfolder2\n        │── tempfolder3\n"
	// treestruct := TreeStruct{true, false, false, false, false, 0}
	// got := getTree(treestruct, "/Users/chinmaysomani/Desktop/gocodes/tree", &noofdir, &nooffiles)

	// assert.Equal(t, want, got, "Not equal")

	treestruct := TreeStruct{false, false, false, false, false, 3}

	want := `/Users/chinmaysomani/Desktop/gocodes/tree
	│── go.mod
	│── go.sum
	│── tempfolder
		│── tempfolder2
			│── tempfile
			│── tempfolder3
	│── tree.go
	│── tree_test.go
	
	3 directiries, 6 files`

	got := getTree(treestruct, "/Users/chinmaysomani/Desktop/gocodes/tree", &noofdir, &nooffiles)
	assert.Equal(t, want, got, "Not Equal")
}
