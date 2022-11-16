package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

// create struct of option and value and desc

// tree -p -L 3 /Users/chinmaysomani/Desktop/gocodes
// create a slice of this struct

// tree -p -d /Users/chinmaysomani/Desktop/gocodes

func main() {

	info, paths := recursive(os.Args[2])
	if os.Args[1] == "-t" {
		acctomodtime(info, paths)
	} else if os.Args[1] == "-f" {
		printRelativePath(info, paths)
	} else if os.Args[1] == "-p" {
		printWithPermission(info)
	} else if os.Args[1] == "-d" {
		printDirectoryOnly(info)
	} else {
		printPathOfFiles(info)
	}
}

func recursive(file string) ([]os.FileInfo, []string) {

	mydirectoryslice := make([]os.FileInfo, 0)
	pathSlice := make([]string, 0)
	err := filepath.Walk(file,
		func(path string, info os.FileInfo, err error) error {

			pathSlice = append(pathSlice, path)

			if err != nil {
				return err
			}

			mydirectoryslice = append(mydirectoryslice, info)
			return nil
		})

	if err != nil {
		fmt.Println(err)
	}
	return mydirectoryslice, pathSlice
}

func printPathOfFiles(structure []os.FileInfo) {
	for i := 0; i < len(structure); i++ {
		if structure[i].IsDir() {
			fmt.Println("|   ")
			fmt.Printf("|-- %s\n", structure[i].Name())
		} else {
			fmt.Println("|   ")
			fmt.Printf("|------- %s\n", structure[i].Name())
		}
	}
}

func printDirectoryOnly(structure []os.FileInfo) {
	for i := 0; i < len(structure); i++ {
		if structure[i].IsDir() {
			fmt.Println("|   ")
			fmt.Printf("|-- %s\n", structure[i].Name())
		}
	}
}

func printRelativePath(myslice []os.FileInfo, pathSlice []string) {

	for i := 0; i < len(myslice); i++ {
		if myslice[i].IsDir() {
			fmt.Println("|   ")
			fmt.Printf("|-- %s\n", pathSlice[i])
		} else {
			fmt.Println("|   ")
			fmt.Printf("|------- %s\n", pathSlice[i])
		}
	}
}

func printWithPermission(infoslice []os.FileInfo) {
	for i := 0; i < len(infoslice); i++ {
		if infoslice[i].IsDir() {
			fmt.Println("|   ")
			fmt.Printf("|--[%v] %v\n", infoslice[i].Mode(), infoslice[i].Name())
		} else {
			fmt.Println("|   ")
			fmt.Printf("|------- [%v]%v\n", infoslice[i].Mode(), infoslice[i].Name())
		}
	}
}

func acctomodtime(infoslice []os.FileInfo, paths []string) {

	var files []fs.FileInfo
	var err error

	for i := 0; i < len(infoslice); i++ {
		if infoslice[i].IsDir() {
			fmt.Println("|   ")
			fmt.Printf("|-- %s\n", infoslice[i].Name())
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
				fmt.Printf("|------- %s %v\n", files[i].Name(), files[i].ModTime())
			}

			files = nil
		}
	}
}
