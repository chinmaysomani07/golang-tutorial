package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeStruct(t *testing.T) {
	//assert := assert.New(t)

	want := "/Users/chinmaysomani/Desktop/gocodes/tree\n│── tempfolder\n    │── tempfolder2\n        " +
		"│── tempfolder3\n3 directories"
	treestruct := TreeStruct{true, false, false, false, false, false, 0}
	got := getTree(treestruct, "/Users/chinmaysomani/Desktop/gocodes/tree")
	assert.Equal(t, want, got, "Not equal")

	want = "/Users/chinmaysomani/Desktop/gocodes/tree\n│── tempfolder\n    │── tempfolder2\n2 directories"
	treestruct = TreeStruct{true, false, false, false, false, false, 2}
	got = getTree(treestruct, "/Users/chinmaysomani/Desktop/gocodes/tree")
	assert.Equal(t, want, got, "Not equal")

	want = "/Users/chinmaysomani/Desktop/gocodes/tree\n│── [-rw-r--r--] go.mod\n│── [-rw-r--r--] go.sum\n│── [drwxr-xr-x] tempfolder\n    " +
		"│── [drwxr-xr-x] tempfolder2\n│── [-rw-r--r--] tree.go\n│── [-rw-r--r--] tree_test.go\n\n2 directories, 5 files\n"
	treestruct = TreeStruct{false, false, true, false, false, false, 2}
	got = getTree(treestruct, "/Users/chinmaysomani/Desktop/gocodes/tree")
	assert.Equal(t, want, got, "Not equal")

	want = "/Users/chinmaysomani/Desktop/gocodes/tree\n│── /Users/chinmaysomani/Desktop/gocodes/tree/tempfolder\n    " +
		"│── /Users/chinmaysomani/Desktop/gocodes/tree/tempfolder2\n        " +
		"│── /Users/chinmaysomani/Desktop/gocodes/tree/tempfolder3\n3 directories"
	treestruct = TreeStruct{true, false, false, true, false, false, 0}
	got = getTree(treestruct, "/Users/chinmaysomani/Desktop/gocodes/tree")
	assert.Equal(t, want, got, "Not equal")

	want = "[\n  {\"type\":\"directory\",\"name\":/Users/chinmaysomani/Desktop/gocodes/tree,\"contents:\"[\n     " +
		"{\"type\":\"file\",\"name\":\"go.mod\"}\n     {\"type\":\"file\",\"name\":\"go.sum\"}\n     " +
		"{\"type\":\"directory\",\"name\":\"tempfolder\",\"contents\":[\n      " +
		"{\"type\":\"directory\",\"name\":\"tempfolder2\",\"contents\":[\n       " +
		"{\"type\":\"file\",\"name\":\"tempfile\"}\n       {\"type\":\"directory\",\"name\":\"tempfolder3\",\"contents\":[\n        " +
		"{\"type\":\"file\",\"name\":\"tempfile1\"}\n       }]\n      }]\n     }]\n     {\"type\":\"file\",\"name\":\"tree.go\"}\n     " +
		"{\"type\":\"file\",\"name\":\"tree_test.go\"}\n,\n   {\"type\":\"report\",\"directories\":3,\"files\":6}\n]"
	treestruct = TreeStruct{false, false, false, false, true, false, 0}
	got = getTree(treestruct, "/Users/chinmaysomani/Desktop/gocodes/tree")
	assert.Equal(t, want, got, "Not equal")

	want = "/Users/chinmaysomani/Desktop/gocodes/tree\n│── [drwxr-xr-x] /Users/chinmaysomani/Desktop/gocodes/tree/tempfolder\n2 directories"
	treestruct = TreeStruct{true, false, true, true, false, false, 1}
	got = getTree(treestruct, "/Users/chinmaysomani/Desktop/gocodes/tree")
	assert.Equal(t, want, got, "Not equal")
}
