package main

import (
	"strings"
	"testing"
)

func TestPrintDirectoryOnly(t *testing.T) {
	treestruct := TreeStruct{true, false, true, false, false, 0}
	dirinfo, paths := getDirectoriesAndPaths("/Users/chinmaysomani/Desktop/gocodes/tree", treestruct)
	want := `/Users/chinmaysomani/Desktop/gocodes/tree
	│──tempfolder
	  └──tempfolder2
	   └──tempfolder3
	
	3 directiories`
	got := getDirectoryOnly(dirinfo, paths)

	bool1 := strings.Contains(want, "tempfolder3")
	bool2 := strings.Contains(got, "tempfolder3")
	if bool1 != bool2 {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestDirOnlyWithPermission(t *testing.T) {
	treestruct := TreeStruct{true, false, true, false, false, 0}
	dirinfo, paths := getDirectoriesAndPaths("/Users/chinmaysomani/Desktop/gocodes/tree", treestruct)

	want := `└──[drwxr-xr-x] tree
	└──[drwxr-xr-x] tempfolder
	  └──[drwxr-xr-x] tempfolder2
		└──[drwxr-xr-x] tempfolder3`
	got := getWithPermissions(dirinfo, paths, treestruct)

	bool1 := strings.Contains(want, "[drwxr-xr-x] tree")
	bool2 := strings.Contains(got, "[drwxr-xr-x] tree")

	if bool1 != bool2 {
		t.Errorf("want %v, got %v", want, got)
	}
}
