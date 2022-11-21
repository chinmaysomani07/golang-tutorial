package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type TreeStruct struct {
	dironly, modtime, permission, relpath bool
	level                                 int
}

const (
	BoxVer      = "│"
	BoxHor      = "──"
	BoxVH       = BoxVer + BoxHor
	BoxUpAndRig = "└"

	OpenBrkt      = "["
	CloseBrkt     = "]"
	Command       = "tree"
	PathSeperator = string(os.PathSeparator)
	NewLine       = "\n"
	Space         = " "
	Spaces3       = "   "
	Spaces4       = "    "
)

func main() {
	fmt.Println(makeTree())
}

func makeTree() string {
	treestruct := parseInput()
	dirinfo, paths := getDirectoriesAndPaths(os.Args[len(os.Args)-1], treestruct)

	restring := ""
	if treestruct.level > 0 {
		restring = getLevels(dirinfo, paths, treestruct)
	} else if treestruct.relpath {
		restring = relativepathssss(dirinfo, paths, treestruct)
	} else if treestruct.permission {
		restring = permissionsss(dirinfo, paths, treestruct)
	} else if treestruct.modtime {
		restring = getModTime(dirinfo, paths, restring)
	} else if treestruct.dironly {
		restring = printDirectoryOnly(dirinfo, paths)
	}

	return restring
}

func parseInput() TreeStruct {

	treestruct := TreeStruct{false, false, false, false, false, 0}
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "-t" {
			treestruct.modtime = true
		} else if os.Args[i] == "-f" {
			treestruct.relpath = true
		} else if os.Args[i] == "-p" {
			treestruct.permission = true
		} else if os.Args[i] == "-d" {
			treestruct.dironly = true
		} else if os.Args[i] == "-L" {
			treestruct.level = parseToInt(os.Args[i+1])
		}
	}
	return treestruct
}

func getLevels(directoriesinfo []os.FileInfo, paths []string, treestruct TreeStruct) string {

	res := ""
	temp := strings.Split(os.Args[len(os.Args)-1], "/")
	templen := len(temp)

	noofdir := 0
	nooffiles := 0
	if treestruct.relpath && treestruct.permission && treestruct.dironly {
		for i := 0; i < len(paths); i++ {
			l := getLengthOfPath(paths[i])
			temp1 := strings.Split(paths[i], "/")
			if len(temp1) <= templen+treestruct.level {
				if directoriesinfo[i].IsDir() {
					res += fmt.Sprintf("%v%v [%v]%v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), paths[i])
					noofdir++
				}
			}
		}
		res += fmt.Sprintf("\n %v directories", noofdir-1)
	} else if treestruct.relpath && treestruct.permission {
		for i := 0; i < len(paths); i++ {
			l := getLengthOfPath(paths[i])
			temp1 := strings.Split(paths[i], "/")
			if len(temp1) <= templen+treestruct.level {
				if directoriesinfo[i].IsDir() {
					res += fmt.Sprintf("%v%v[%v] %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), paths[i])
					noofdir++
				} else {
					res += fmt.Sprintf("%v%v[%v] %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), paths[i])
					nooffiles++
				}
			}
		}
		res += fmt.Sprintf("\n %v directories, %v files", noofdir-1, nooffiles)
	} else if treestruct.relpath && treestruct.dironly {
		for i := 0; i < len(paths); i++ {
			l := getLengthOfPath(paths[i])
			temp1 := strings.Split(paths[i], "/")
			if len(temp1) <= templen+treestruct.level {
				if directoriesinfo[i].IsDir() {
					res += fmt.Sprintf("%v%v %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, paths[i])
					noofdir++
				}
			}
		}
		res += fmt.Sprintf("\n %v directories", noofdir-1)
	} else if treestruct.permission && treestruct.dironly {
		for i := 0; i < len(paths); i++ {
			l := getLengthOfPath(paths[i])
			temp1 := strings.Split(paths[i], "/")
			if len(temp1) <= templen+treestruct.level {
				min := int64(math.Round(math.Min(float64(len(temp1)), float64(templen+treestruct.level))))
				tempslice := temp1[:min]
				if directoriesinfo[i].IsDir() {
					res += fmt.Sprintf("%v%v [%v]%v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), tempslice[len(tempslice)-1])
					noofdir++
				}
			}
		}
		res += fmt.Sprintf("\n %v directories", noofdir-1)
	} else if treestruct.permission {
		for i := 0; i < len(paths); i++ {
			l := getLengthOfPath(paths[i])
			temp1 := strings.Split(paths[i], "/")
			if len(temp1) <= templen+treestruct.level {
				min := int64(math.Round(math.Min(float64(len(temp1)), float64(templen+treestruct.level))))
				tempslice := temp1[:min]

				if directoriesinfo[i].IsDir() {
					res += fmt.Sprintf("%v%v [%v]%v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), tempslice[len(tempslice)-1])
					noofdir++
				} else {
					nooffiles++
					res += fmt.Sprintf("%v%v [%v]%v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), temp1[len(tempslice)-1])
				}
			}
		}
		res += fmt.Sprintf("\n %v directories, %v files", noofdir-1, nooffiles)
	} else if treestruct.relpath {
		for i := 0; i < len(paths); i++ {
			l := getLengthOfPath(paths[i])
			temp1 := strings.Split(paths[i], "/")
			if len(temp1) <= templen+treestruct.level {
				if directoriesinfo[i].IsDir() {
					noofdir++
					res += fmt.Sprintf("%v%v%v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, paths[i])
				} else {
					nooffiles++
					res += fmt.Sprintf("%v%v%v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, paths[i])
				}
			}
		}
		res += fmt.Sprintf("\n %v directories, %v files", noofdir-1, nooffiles)
	} else if treestruct.dironly {
		for i := 0; i < len(paths); i++ {
			l := getLengthOfPath(paths[i])
			temp1 := strings.Split(paths[i], "/")
			if len(temp1) <= templen+treestruct.level {
				min := int64(math.Round(math.Min(float64(len(temp1)), float64(templen+treestruct.level))))
				tempslice := temp1[:min]

				if directoriesinfo[i].IsDir() {
					noofdir++
					res += fmt.Sprintf("%v%v %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, tempslice[len(tempslice)-1])
				}
			}
		}
		res += fmt.Sprintf("\n %v directories", noofdir-1)
	} else {
		for i := 0; i < len(paths); i++ {
			l := getLengthOfPath(paths[i])
			temp1 := strings.Split(paths[i], "/")
			if len(temp1) <= templen+treestruct.level {
				min := int64(math.Round(math.Min(float64(len(temp1)), float64(templen+treestruct.level))))
				tempslice := temp1[:min]

				if directoriesinfo[i].IsDir() {
					noofdir++
					res += fmt.Sprintf("%v%v %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, tempslice[len(tempslice)-1])
				} else {
					nooffiles++
					res += fmt.Sprintf("%v%v %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, temp1[len(tempslice)-1])
				}
			}
		}
		res += fmt.Sprintf("\n %v directories, %v files", noofdir-1, nooffiles)
	}

	return res
}

func getDirectoriesAndPaths(file string, treestruct TreeStruct) ([]os.FileInfo, []string) {
	directoriesinfo := make([]os.FileInfo, 0)
	paths := make([]string, 0)
	err := filepath.Walk(file,
		func(path string, info os.FileInfo, err error) error {

			if info.IsDir() && info.Name() == ".git" {
				return filepath.SkipDir
			}

			if err != nil {
				return err
			}
			paths = append(paths, path)
			directoriesinfo = append(directoriesinfo, info)

			return nil
		})

	if err != nil {
		fmt.Println(err)
	}
	return directoriesinfo, paths
}

func printDirectoryOnly(directoriesinfo []os.FileInfo, paths []string) string {

	noofdir := 0
	restring := ""
	restring += fmt.Sprintf("%v\n", os.Args[len(os.Args)-1])
	temp := strings.Split(os.Args[len(os.Args)-1], "/")
	templen := len(temp)

	for i := 0; i < len(directoriesinfo); i++ {
		l := getLengthOfPath(paths[i])
		if directoriesinfo[i].IsDir() && l == templen+1 {
			restring += fmt.Sprintf("%v%v%v\n", BoxVer, BoxHor, directoriesinfo[i].Name())
			noofdir++
		} else if directoriesinfo[i].IsDir() && l != templen {
			restring += fmt.Sprintf("%v%v%v\n", strings.Repeat(" ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Name())
			noofdir++
		}
	}
	restring += fmt.Sprintf("\n%v directiories", noofdir)
	return restring
}

func getModTime(directoriesinfo []os.FileInfo, paths []string, restring string) string {

	var files []fs.FileInfo
	var err error

	res := ""
	res += fmt.Sprintf("%v\n", os.Args[len(os.Args)-1])
	for i := 0; i < len(directoriesinfo); i++ {
		l := getLengthOfPath(paths[i])
		if directoriesinfo[i].IsDir() {
			if len(restring) > 0 {
				res += fmt.Sprintf("%v%v%v\n", strings.Repeat("  ", l-5), BoxUpAndRig+BoxHor, directoriesinfo[i].Name())
				continue
			} else {
				res += fmt.Sprintf("   %v\n", directoriesinfo[i].Name())
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
				if !files[i].IsDir() {
					res += fmt.Sprintf("%v%v%v \n", strings.Repeat("  ", l-4), BoxUpAndRig+BoxHor, files[i].Name())
				}
			}

			files = nil
		}
	}
	return res
}

func permissionsss(directoriesinfo []os.FileInfo, paths []string, treestruct TreeStruct) string {
	res := ""
	noofdire := 0
	//nooffile := 0
	temp := strings.Split(os.Args[len(os.Args)-1], "/")
	templen := len(temp)

	var files []fs.FileInfo
	var err error

	if treestruct.modtime && treestruct.dironly {
		for i := 0; i < len(directoriesinfo); i++ {
			l := getLengthOfPath(paths[i])
			//	temp1 := strings.Split(paths[i], "/")
			if directoriesinfo[i].IsDir() {
				noofdire++
				res += fmt.Sprintf("%v%v[%v] %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), directoriesinfo[i].Name())

				//res += fmt.Sprintf("[%v]%v\n", directoriesinfo[i].Mode(), directoriesinfo[i].Name())
				p := paths[i]
				files, err = ioutil.ReadDir(p)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		res += fmt.Sprintf("\n%v directiories", noofdire-1)
	} else if treestruct.dironly {
		for i := 0; i < len(directoriesinfo); i++ {
			l := getLengthOfPath(paths[i])
			if directoriesinfo[i].IsDir() {
				res += fmt.Sprintf("%v%v[%v] %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), directoriesinfo[i].Name())
			}
		}
	} else if treestruct.modtime {
		res = ""
		for i := 0; i < len(directoriesinfo); i++ {
			l := getLengthOfPath(paths[i])
			if directoriesinfo[i].IsDir() {
				res += fmt.Sprintf("   [%v]%v %v\n", directoriesinfo[i].Mode(), directoriesinfo[i].ModTime(), directoriesinfo[i].Name())
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
					res += fmt.Sprintf("%v%v[%v]%v %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, files[i].Mode(), files[i].ModTime(), files[i].Name())
				}
				files = nil
			}
		}
	} else {
		res = ""
		for i := 0; i < len(directoriesinfo); i++ {
			l := getLengthOfPath(paths[i])
			if directoriesinfo[i].IsDir() {
				res += fmt.Sprintf("%v\n", directoriesinfo[i].Name())
			} else {
				res += fmt.Sprintf("%v%v[%v]%v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), directoriesinfo[i].Name())
			}
		}
	}

	return res
}

func relativepathssss(directoriesinfo []os.FileInfo, paths []string, treestruct TreeStruct) string {
	res := ""
	var files []fs.FileInfo
	var err error

	if treestruct.permission && treestruct.modtime && treestruct.dironly {
		res = ""
		for i := 0; i < len(directoriesinfo); i++ {
			l := getLengthOfPath(paths[i])
			if directoriesinfo[i].IsDir() {
				res += fmt.Sprintf("%v%v[%v] %v %v\n", strings.Repeat("  ", l-5), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), paths[i], directoriesinfo[i].ModTime())
			}
		}
	} else if treestruct.permission && treestruct.dironly {
		res = ""
		for i := 0; i < len(directoriesinfo); i++ {
			l := getLengthOfPath(paths[i])
			if directoriesinfo[i].IsDir() {
				res += fmt.Sprintf("%v%v[%v] %v\n", strings.Repeat("  ", l-5), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), paths[i])
			}
		}
	} else if treestruct.modtime && treestruct.dironly {
		res = ""
		for i := 0; i < len(directoriesinfo); i++ {
			l := getLengthOfPath(paths[i])
			if directoriesinfo[i].IsDir() {
				res += fmt.Sprintf("%v%v%v %v\n", strings.Repeat("  ", l-5), BoxUpAndRig+BoxHor, paths[i], directoriesinfo[i].ModTime())
			}
		}
	} else if treestruct.modtime && treestruct.permission {
		res = ""
		for i := 0; i < len(directoriesinfo); i++ {
			l := getLengthOfPath(paths[i])
			if directoriesinfo[i].IsDir() {
				res += fmt.Sprintf("  [%v] %v %v \n", directoriesinfo[i].Mode(), paths[i], directoriesinfo[i].ModTime())
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
					res += fmt.Sprintf("%v%v [%v]%v %v\n", strings.Repeat("  ", l-4), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), paths[i], files[i].ModTime())
				}
				files = nil
			}
		}
	} else if treestruct.dironly {
		res = ""
		for i := 0; i < len(directoriesinfo); i++ {
			l := getLengthOfPath(paths[i])
			if directoriesinfo[i].IsDir() {
				res += fmt.Sprintf("%v%v%v\n", strings.Repeat("  ", l-5), BoxUpAndRig+BoxHor, paths[i])
			}
		}
	} else if treestruct.permission {
		res = ""
		for i := 0; i < len(directoriesinfo); i++ {
			l := getLengthOfPath(paths[i])
			if directoriesinfo[i].IsDir() {
				res += fmt.Sprintf("   [%v]%v\n", directoriesinfo[i].Mode(), paths[i])
			} else {
				res += fmt.Sprintf("%v%v[%v]%v\n", strings.Repeat("  ", l-4), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), paths[i])
			}
		}
	} else if treestruct.modtime {
		res = ""
		for i := 0; i < len(directoriesinfo); i++ {
			l := getLengthOfPath(paths[i])
			if directoriesinfo[i].IsDir() {
				res += fmt.Sprintf("   %v %v \n", paths[i], directoriesinfo[i].ModTime())
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
					res += fmt.Sprintf("%v%v %v %v\n", strings.Repeat("  ", l-4), BoxUpAndRig+BoxHor, paths[i], files[i].ModTime())
				}
				files = nil
			}
		}
	}

	return res
}

func getLengthOfPath(path string) int {
	return len(strings.Split(path, "/"))
}

func parseToInt(input string) int {
	number, err := strconv.ParseInt(input, 10, 32)

	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return int(number)
}
