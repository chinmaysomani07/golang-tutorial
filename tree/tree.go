package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// create struct of option and value and desc

// tree -p -L 3 /Users/chinmaysomani/Desktop/gocodes
// create a slice of this struct

// tree -p -d /Users/chinmaysomani/Desktop/gocodes

type Command struct {
	option1, option2, filepath string
}

const (
	BoxHor      = "──"
	BoxUpAndRig = "└"
)

func main() {

	command := Command{}
	if strings.Contains(os.Args[1], "/") {
		command = Command{"", "", os.Args[1]}
	} else if strings.Contains(os.Args[2], "-") {
		command = Command{os.Args[1], os.Args[2], os.Args[3]}
	} else {
		command = Command{os.Args[1], "", os.Args[2]}
	}

	dirinfo, paths := getDirectoriesAndPaths(command.filepath)
	if command.option1 == "-t" {
		printByModificationTime(dirinfo, paths, command)
	} else if command.option1 == "-f" {
		printRelativePath(dirinfo, paths)
	} else if command.option1 == "-p" {
		printWithPermission(dirinfo, paths)
	} else if command.option1 == "-d" {
		printDirectoryOnly(dirinfo, paths)
	} else if command.option1 == "" {
		printPathOfFiles(dirinfo, paths)
	} else {
		log.Fatal("This command not defined")
	}
}

func getDirectoriesAndPaths(file string) ([]os.FileInfo, []string) {

	directoriesinfo := make([]os.FileInfo, 0)
	paths := make([]string, 0)
	err := filepath.Walk(file,
		func(path string, info os.FileInfo, err error) error {

			paths = append(paths, path)

			if err != nil {
				return err
			}

			directoriesinfo = append(directoriesinfo, info)
			return nil
		})

	if err != nil {
		fmt.Println(err)
	}
	return directoriesinfo, paths
}

func printPathOfFiles(directoriesinfo []os.FileInfo, paths []string) {
	for i := 0; i < len(directoriesinfo); i++ {
		l := getLengthOfPath(paths[i])
		if directoriesinfo[i].IsDir() {
			fmt.Printf("%v%v\n", strings.Repeat("  ", l-4), directoriesinfo[i].Name())
		} else {
			fmt.Printf("%v%v%v\n", strings.Repeat("  ", l-4), BoxUpAndRig+BoxHor, directoriesinfo[i].Name())
		}
	}
}

func printDirectoryOnly(directoriesinfo []os.FileInfo, paths []string) {
	for i := 0; i < len(directoriesinfo); i++ {
		l := getLengthOfPath(paths[i])
		if directoriesinfo[i].IsDir() {
			fmt.Printf("%v%v%v\n", strings.Repeat("  ", l-5), BoxUpAndRig+BoxHor, directoriesinfo[i].Name())
		}
	}
}

func printRelativePath(directoriesinfo []os.FileInfo, paths []string) {

	for i := 0; i < len(directoriesinfo); i++ {
		l := getLengthOfPath(paths[i])
		if directoriesinfo[i].IsDir() {
			fmt.Printf("   %v\n", directoriesinfo[i].Name())
		} else {
			fmt.Printf("%v%v%v\n", strings.Repeat("  ", l-4), BoxUpAndRig+BoxHor, paths[i])
		}
	}
}

func printWithPermission(directoriesinfo []os.FileInfo, paths []string) {

	for i := 0; i < len(directoriesinfo); i++ {
		l := getLengthOfPath(paths[i])
		if directoriesinfo[i].IsDir() {
			fmt.Printf("   %v\n", directoriesinfo[i].Name())
		} else {
			fmt.Printf("%v%v[%v]%v\n", strings.Repeat("  ", l-4), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), directoriesinfo[i].Name())
		}
	}
}

func printByModificationTime(directoriesinfo []os.FileInfo, paths []string, command Command) {

	var files []fs.FileInfo
	var err error

	for i := 0; i < len(directoriesinfo); i++ {
		l := getLengthOfPath(paths[i])
		if directoriesinfo[i].IsDir() {
			if command.option2 == "-p" {
				fmt.Printf("   [%v %v] %v\n", directoriesinfo[i].Mode(), directoriesinfo[i].ModTime(), directoriesinfo[i].Name())
			} else if command.option2 == "-d" {
				fmt.Printf("%v%v[%v] %v\n", strings.Repeat("  ", l-5), BoxUpAndRig+BoxHor, directoriesinfo[i].ModTime(), directoriesinfo[i].Name())
				continue
			} else {
				fmt.Printf("   %v%v\n", directoriesinfo[i].ModTime(), directoriesinfo[i].Name())
			}
			p := paths[i]
			files, err = ioutil.ReadDir(p)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			sort.Slice(files, func(i, j int) bool {
				return files[i].ModTime().Before(files[j].ModTime())
			})

			for i := 0; i < len(files); i++ {
				if command.option2 == "-p" {
					fmt.Printf("%v%v[%v%v] %v\n", strings.Repeat("  ", l-4), BoxUpAndRig+BoxHor, files[i].Mode(), files[i].ModTime(), files[i].Name())
				} else {
					fmt.Printf("%v%v%v %v\n", strings.Repeat("  ", l-4), BoxUpAndRig+BoxHor, files[i].Name(), files[i].ModTime())
				}
			}

			files = nil
		}
	}
}

func getLengthOfPath(path string) int {
	return len(strings.Split(path, "/"))
}
