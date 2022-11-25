package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	treestruct TreeStruct
	path       string
	want       string
}

func TestTree(t *testing.T) {
	tests := []test{
		{treestruct: TreeStruct{true, false, false, false, false, false, 0}, path: "/Users/chinmaysomani/Desktop/gocodes/tree",
			want: "/Users/chinmaysomani/Desktop/gocodes/tree\n│── tempfolder\n    │── tempfolder2\n        " +
				"│── tempfolder3\n3 directories"},
		{treestruct: TreeStruct{true, false, false, false, false, false, 2}, path: "/Users/chinmaysomani/Desktop/gocodes/tree",
			want: "/Users/chinmaysomani/Desktop/gocodes/tree\n│── tempfolder\n    │── tempfolder2\n2 directories"},

		{treestruct: TreeStruct{false, false, true, false, false, false, 2}, path: "/Users/chinmaysomani/Desktop/gocodes/tree",
			want: "/Users/chinmaysomani/Desktop/gocodes/tree\n│── [-rw-r--r--] go.mod\n│── [-rw-r--r--] go.sum\n│── [drwxr-xr-x] tempfolder\n    " + "│── [drwxr-xr-x] tempfolder2\n│── [-rw-r--r--] tree.go\n│── [-rw-r--r--] tree_test.go\n\n2 directories, 5 files\n"},

		{treestruct: TreeStruct{true, false, false, true, false, false, 0}, path: "/Users/chinmaysomani/Desktop/gocodes/tree",
			want: "/Users/chinmaysomani/Desktop/gocodes/tree\n│── /Users/chinmaysomani/Desktop/gocodes/tree/tempfolder\n    " + "│── /Users/chinmaysomani/Desktop/gocodes/tree/tempfolder2\n        " + "│── /Users/chinmaysomani/Desktop/gocodes/tree/tempfolder3\n3 directories"},

		{treestruct: TreeStruct{true, false, true, true, false, false, 1}, path: "/Users/chinmaysomani/Desktop/gocodes/tree",
			want: "/Users/chinmaysomani/Desktop/gocodes/tree\n│── [drwxr-xr-x] /Users/chinmaysomani/Desktop/gocodes/tree/tempfolder\n2 directories"},
	}

	assert := assert.New(t)
	for _, test := range tests {
		got := getTree(test.treestruct, test.path)
		assert.Equal(test.want, got)
	}
}
