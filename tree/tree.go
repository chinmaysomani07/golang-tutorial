package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type TreeStruct struct {
	dironly, modtime, permission, relpath, json, xml bool
	level                                            int
}

const (
	BoxVer        = "│"
	BoxHor        = "──"
	BoxVH         = BoxVer + BoxHor
	BoxUpAndRig   = "└"
	OpenBrkt      = "["
	CloseBrkt     = "]"
	Command       = "tree"
	PathSeperator = string(os.PathSeparator)
	NewLine       = "\n"
	Space         = " "
	Spaces3       = "   "
	Spaces4       = "    "
	OpenTag       = "<"
	Slash         = "/"
	CloseTag      = ">"
)

func main() {
	fmt.Println(makeTree())
}

func makeTree() string {
	treestruct := parseInput()
	restring := getTree(treestruct, os.Args[len(os.Args)-1])
	return restring
}

func parseInput() TreeStruct {
	treestruct := TreeStruct{false, false, false, false, false, false, 0}
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
		} else if os.Args[i] == "-J" {
			treestruct.json = true
		} else if os.Args[i] == "-X" {
			treestruct.xml = true
		}
	}
	return treestruct
}

func getTree(treestruct TreeStruct, pathfile string) string {
	noOfDir := 0
	noOfFiles := 0
	var temp = ""
	var inforoot os.FileInfo
	root := strings.Split(pathfile, "/")
	result := pathfile + "\n"

	err := filepath.Walk(pathfile,
		func(path string, info os.FileInfo, err error) error {

			if info.Name() == root[len(root)-1] {
				inforoot = info
				return nil
			}
			if (pathfile != info.Name()) && info.Name()[0] == '.' && info.IsDir() {
				return filepath.SkipDir
			}

			if info.IsDir() {
				noOfDir++
			} else {
				noOfFiles++
			}

			if treestruct.json {
				result = ""
				result = recGetInJson(pathfile, temp, 0, treestruct, inforoot)
				return nil
			}
			if treestruct.xml {
				result = ""
				result = recGetInXML(pathfile, temp, 0, treestruct, inforoot)
				return nil
			}

			relPath, err := filepath.Rel(pathfile, path)
			if err != nil {
				return err
			}
			currLevel := len(strings.Split(relPath, "/"))

			if treestruct.level > 0 && currLevel-1 == treestruct.level { //make sure we don't go beyond a specified level
				return filepath.SkipDir
			}

			if treestruct.dironly {
				if !info.IsDir() {
					return nil //to stop the execution there itself if it finds file
				}
			}

			result += strings.Repeat(Spaces4, currLevel-1) + BoxVH + " "
			if treestruct.permission {
				result += OpenBrkt + info.Mode().String() + CloseBrkt + " "
			}
			if treestruct.relpath {
				result += pathfile + "/"
			}
			result += info.Name() + "\n"
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}

	filesDir := getFilesAndDirCount(treestruct, noOfFiles, noOfDir)
	result += filesDir
	return result
}

func recGetInJson(root string, res string, n int, treestruct TreeStruct, args os.FileInfo) string {
	files, err := os.ReadDir(root)
	if err != nil {
		fmt.Println(err)
	}

	if treestruct.dironly {
		temp := make([]fs.DirEntry, 0)
		for _, f := range files {
			if f.IsDir() {
				temp = append(temp, f)
			}
		}
		files = temp
	}

	if n == 0 {
		res += OpenBrkt + NewLine
		res += strings.Repeat(Space, n+2) + "{\"type\":\"directory\",\"name\":" + root + "" + getPermissions(treestruct, args) + ",\"contents:\"" + OpenBrkt + NewLine
	}

	if n > 0 && n == treestruct.level {
		return res + strings.Repeat(Space, n+2) + "}]," + "\n"
	}

	for _, f := range files {
		fileinfo, err := f.Info()
		if err != nil {
			fmt.Println(err)
		}
		if !f.IsDir() {
			res += strings.Repeat(Space, n+5) + "{\"type\":\"file\",\"name\":" + "\"" + f.Name() + "\"" + getPermissions(treestruct, fileinfo) + "}" + NewLine
			continue
		}
		res += strings.Repeat(Space, n+5) + "{\"type\":\"directory\",\"name\":" + "\"" + f.Name() + "\"" + getPermissions(treestruct, fileinfo) + ",\"contents\":[" + NewLine
		fileName := root + "/" + f.Name()
		res = recGetInJson(fileName, res, n+1, treestruct, args)
	}

	if n > 0 {
		return res + strings.Repeat(Space, n+4) + "}]" + "\n"
	}
	return res
}

func recGetInXML(root string, line string, n int, treestruct TreeStruct, args os.FileInfo) string {
	files, err := os.ReadDir(root)
	if err != nil {
		fmt.Println(err)
	}

	if treestruct.dironly {
		temp := make([]fs.DirEntry, 0)
		for _, f := range files {
			if f.IsDir() {
				temp = append(temp, f)
			}
		}
		files = temp
	}

	if n == 0 {
		line += "<tree>" + NewLine
		line += strings.Repeat(Space, n+2) + "<directory<directory name= " + root + getPermissions(treestruct, args) + CloseTag + NewLine
	}

	closeDirTag := OpenTag + Slash + "directory" + CloseTag + NewLine
	if n > 0 && n == treestruct.level {
		return line + strings.Repeat(Space, n+3) + closeDirTag
	}

	for _, f := range files {
		if !f.IsDir() {
			line += strings.Repeat(Space, n+4) + OpenTag + "file<file name=" + "\"" + f.Name() + "\"" + getPermissions(treestruct, args) + CloseTag + OpenTag + Slash + "file" + CloseTag + NewLine
			continue
		}
		line += strings.Repeat(Space, n+4) + OpenTag + "directory<directory name=" + "\"" + f.Name() + "\"" + getPermissions(treestruct, args) + CloseTag + NewLine
		line = recGetInXML(root+PathSeperator+f.Name(), line, n+1, treestruct, args)
	}

	if n > 0 {
		return line + strings.Repeat(Space, n+2) + closeDirTag
	}

	return line + strings.Repeat(Space, n+2) + closeDirTag
}

func getFilesAndDirCount(treestruct TreeStruct, files, dir int) string {
	var str string
	if treestruct.dironly {
		if treestruct.xml {
			str += "  <report>" + NewLine
			str += fmt.Sprintf("%v<directories>%v</directories>%v", strings.Repeat(Space, 4), dir, NewLine)
			str += "  </report>"
			str += "\n" + "</tree>"
		} else if treestruct.json {
			str += fmt.Sprintf(",%v%v{\"type\":\"report\",\"directories\":%v}%v]", NewLine, strings.Repeat(Space, 3), dir, NewLine)
		} else {
			str = fmt.Sprintf("%v directories", dir)
		}
	} else {
		if treestruct.xml {
			str += "  <report>" + NewLine
			str += fmt.Sprintf("%v<directories>%v</directories>%v", strings.Repeat(Space, 4), dir, NewLine)
			str += fmt.Sprintf("%v<files>%v</files>%v", strings.Repeat(Space, 4), files, NewLine)
			str += "  </report>"
			str += "\n" + "</tree>"
		} else if treestruct.json {
			str += fmt.Sprintf(",%v%v{\"type\":\"report\",\"directories\":%v,\"files\":%v}%v]", NewLine, strings.Repeat(Space, 3), dir, files, NewLine)
		} else {
			str = fmt.Sprintf("\n%v directories, %v files\n", dir, files)
		}
	}
	return str
}

func getPermissions(treestruct TreeStruct, fileinfo os.FileInfo) string {
	perms := ""
	if treestruct.permission && treestruct.json {
		modes := fileinfo.Mode().String()
		octal := fmt.Sprintf("%#o", fileinfo.Mode().Perm())
		perms = "," + "\"mode\":" + "\"" + octal + "\"" + "," + "\"prot\":" + "\"" + modes + "\""
	} else if treestruct.permission && treestruct.xml {
		modes := fileinfo.Mode().String()
		octal := fmt.Sprintf("%#o", fileinfo.Mode().Perm())
		perms = " mode:" + "\"" + octal + "\"" + " prot:" + "\"" + modes + "\""
	}
	return perms
}

func parseToInt(input string) int {
	number, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return int(number)
}
