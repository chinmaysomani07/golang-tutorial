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

type Command struct {
	command  string
	filepath string
}

func main() {

	inputCommand := enterCommand()
	//printFilePath(inputCommand.command, inputCommand.filepath)

	// inputCommand = enterCommand()
	// searchDirectoryOnly(inputCommand.command, inputCommand.filepath)

	//printRelativeFilePath(inputCommand.command, inputCommand.filepath)

	//printFilePathWithPermission(inputCommand.command, inputCommand.filepath)

	sortWithModificationTime(inputCommand.command, inputCommand.filepath)
}

func enterCommand() Command {

	fmt.Println("Enter the command:")
	var command string
	var filepath string
	fmt.Scan(&command, &filepath)

	data := Command{command, filepath}
	return data
}

func printFilePath(command string, file string) {

	noOfDirectories := 0
	noOfFiles := 0

	checkValidCommand(command)

	err := filepath.Walk(file,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if info.IsDir() {
				fmt.Println("|   ")
				fmt.Printf("|-- %s\n", info.Name())
				noOfDirectories++

			} else {
				fmt.Println("|   ")
				fmt.Printf("|------- %s\n", info.Name())
				noOfFiles++
			}

			return nil
		})

	fmt.Printf("%v directories, %v files\n", noOfDirectories, noOfFiles)
	if err != nil {
		log.Println(err)
	}
}

func printRelativeFilePath(command string, file string) {
	noOfDirectories := 0
	noOfFiles := 0

	checkValidCommand(command)

	err := filepath.Walk(file,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if info.IsDir() {
				fmt.Println("|   ")
				fmt.Printf("|-- %s\n", path)
				noOfDirectories++

			} else {
				fmt.Println("|   ")
				fmt.Printf("|------- %s\n", path)
				noOfFiles++
			}

			return nil
		})

	fmt.Printf("%v directories, %v files\n", noOfDirectories, noOfFiles)
	if err != nil {
		log.Println(err)
	}
}

func printDirectoryOnly(command string, file string) {
	noOfDirectories := 0

	checkValidCommand(command)

	err := filepath.Walk(file,
		func(path string, info os.FileInfo, err error) error {

			if info.IsDir() {
				fmt.Println("|   ")
				fmt.Printf("|-- %s\n", info.Name())
				noOfDirectories++
			}
			return nil
		})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(noOfDirectories, " directories")
}

func printFilePathWithPermission(command string, file string) {

	noOfDirectories := 0
	noOfFiles := 0

	checkValidCommand(command)

	err := filepath.Walk(file,
		func(path string, info os.FileInfo, err error) error {

			info.ModTime()

			if err != nil {
				return err
			}

			if info.IsDir() {
				fmt.Println("|   ")
				fmt.Printf("|--[%v] %v\n", info.Mode(), info.Name())
				noOfDirectories++

			} else {
				fmt.Println("|   ")
				fmt.Printf("|------- [%v]%v\n", info.Mode(), info.Name())
				noOfFiles++
			}

			return nil
		})

	fmt.Printf("%v directories, %v files\n", noOfDirectories, noOfFiles)
	if err != nil {
		log.Println(err)
	}
}

func sortWithModificationTime(command string, file string) {
	noOfDirectories := 0
	noOfFiles := 0
	files := make([]fs.FileInfo, 0)

	checkValidCommand(command)

	err := filepath.Walk(file,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if info.IsDir() {
				fmt.Println("|   ")
				fmt.Printf("|-- %s\n", info.Name())
				files, err = ioutil.ReadDir(path)
				if err != nil {
					fmt.Println(err.Error())
				}
				noOfDirectories++

			} else {

				sort.Slice(files, func(i, j int) bool {
					return files[i].ModTime().Before(files[j].ModTime())
				})

				for i := 0; i < len(files); i++ {
					fmt.Printf("|------- %s %v\n", files[i].Name(), files[i].ModTime())
				}

				noOfFiles++
				files = nil
			}

			return nil
		})

	fmt.Printf("%v directories, %v files\n", noOfDirectories, noOfFiles)
	if err != nil {
		log.Println(err)
	}
}

func checkValidCommand(command string) bool {

	validCommand := true
	if !strings.EqualFold(command, "TREE") {
		validCommand = false
		log.Fatal("WRONG COMMAND ENTERED. PLEASE ENTER tree <filepath>")
	}
	return validCommand
}
