package main

import (
	"strings"
	"testing"
)

func TestPrintDirectoryOnly(t *testing.T) {
	treestruct := TreeStruct{true, false, true, false, 0}
	dirinfo, paths := getDirectoriesAndPaths("/Users/chinmaysomani/Desktop/gocodes/tree", treestruct)
	want := ``
	got := printDirectoryOnly(dirinfo, paths)

	bool1 := strings.Contains(want, "tree")
	bool2 := strings.Contains(got, "tree")
	if bool1 != bool2 {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestDirOnlyWithPermission(t *testing.T) {
	treestruct := TreeStruct{true, false, true, false, 0}
	dirinfo, paths := getDirectoriesAndPaths("/Users/chinmaysomani/Desktop/gocodes/tree", treestruct)

	want := "└──[drwxr-xr-x] tree"
	got := permissionsss(dirinfo, paths, treestruct)

	bool1 := strings.Contains(want, "[drwxr-xr-x] tree")
	bool2 := strings.Contains(got, "[drwxr-xr-x] tree")

	if bool1 != bool2 {
		t.Errorf("want %v, got %v", want, got)
	}
}
