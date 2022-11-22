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
	pathfileslice := strings.Split(pathfile, "/")
	dirandfiles := pathfile + "\n"

	err := filepath.Walk(pathfile,
		func(path string, info os.FileInfo, err error) error {

			if info.Name() == pathfileslice[len(pathfileslice)-1] {
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
				dirandfiles = ""
				dirandfiles = recGetInJson(pathfile, temp, 0, treestruct, inforoot)
				return nil
			}
			if treestruct.xml {
				dirandfiles = ""
				dirandfiles = recGetInXML(pathfile, temp, 0, treestruct, inforoot)
				return nil
			}

			relPath, err := filepath.Rel(pathfile, path)
			currLevel := len(strings.Split(relPath, "/"))

			if treestruct.level > 0 && currLevel-1 == treestruct.level {
				return filepath.SkipDir
			}

			if err != nil {
				return err
			}

			if treestruct.dironly {
				if !info.IsDir() {
					return nil //to stop the execution there itself if it finds file
				}
			}

			dirandfiles += strings.Repeat(Spaces4, currLevel-1) + BoxVH + " "

			if treestruct.permission {
				dirandfiles += OpenBrkt + info.Mode().String() + CloseBrkt + " "
			}

			if treestruct.relpath {
				dirandfiles += pathfile + "/"
			}
			dirandfiles += info.Name() + "\n"
			return nil
		})

	if err != nil {
		fmt.Println(err)
	}

	filesDir := getFilesAndDirCount(treestruct, noOfDir, noOfFiles)
	dirandfiles += filesDir
	return dirandfiles
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
			str += fmt.Sprintf("%v<directories>%v</directories>%v", strings.Repeat(Space, 4), dir-1, NewLine)
			str += "  </report>"
			str += "\n" + "</tree>"
		} else if treestruct.json {
			str += fmt.Sprintf(",%v%v{\"type\":\"report\",\"directories\":%v}%v]", NewLine, strings.Repeat(Space, 3), dir-1, NewLine)
		} else {
			str = fmt.Sprintf("%v directories ", dir-1)
		}

	} else {
		if treestruct.xml {
			str += "  <report>" + NewLine
			str += fmt.Sprintf("%v<directories>%v</directories>%v", strings.Repeat(Space, 4), dir-1, NewLine)
			str += fmt.Sprintf("%v<files>%v</files>%v", strings.Repeat(Space, 4), files, NewLine)
			str += "  </report>"
			str += "\n" + "</tree>"
		} else if treestruct.json {
			str += fmt.Sprintf(",%v%v{\"type\":\"report\",\"directories\":%v,\"files\":%v}%v]", NewLine, strings.Repeat(Space, 3), dir-1, files, NewLine)
		} else {
			str = fmt.Sprintf("\n%v directories, %v files\n", dir-1, files)
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

// func getJson(directoriesinfo []os.FileInfo, paths []string, treestruct TreeStruct) string {

// 	noofdir := 0
// 	nooffiles := 0
// 	root := os.Args[len(os.Args)-1]
// 	files, err := os.ReadDir(paths[0]) //it is the parent directory
// 	res := "[\n"
// 	res += fmt.Sprintf("%v{\"type\":\"directory\",\"name\":\"%v\",\"contents\":[", " ", os.Args[len(os.Args)-1])

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	res, noofdir, nooffiles = recDir(root, res, files, noofdir, nooffiles)
// 	res += fmt.Sprintf("\n]}\n,\n%v{\"type\":\"report\",\"directories\":%v,\"files\":%v}\n]", " ", noofdir, nooffiles)
// 	return res
// }

// func recDir(root string, res string, files []fs.DirEntry, noofdir int, nooffiles int) (string, int, int) {
// 	lengthroot := len(strings.Split(os.Args[len(os.Args)-1], "/"))
// 	for i := 0; i < len(files); i++ {
// 		lengthdir := len(strings.Split(root, "/"))
// 		if files[i].IsDir() {
// 			noofdir++
// 			res += fmt.Sprintf("\n%v{\"type\":\"directory\",\"name\":\"%v\",\"contents\":[", strings.Repeat(" ", lengthdir-lengthroot+2), files[i].Name())
// 			files2, err := os.ReadDir(root + "/" + files[i].Name())
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			root = root + "/" + files[i].Name()
// 			res, noofdir, nooffiles = recDir(root, res, files2, noofdir, nooffiles)
// 			res += fmt.Sprintf("\n%v]}", strings.Repeat(" ", lengthdir-lengthroot+2))
// 		} else {
// 			nooffiles++
// 			roottemp := root + "/" + files[i].Name()
// 			lengthdir := len(strings.Split(roottemp, "/"))
// 			res += fmt.Sprintf("\n%v{\"type\":\"file\",\"name\":\"%v\"}", strings.Repeat(" ", lengthdir-lengthroot+1), files[i].Name())
// 		}
// 	}
// 	return res, noofdir, nooffiles
// }

// func getDirectoriesAndPaths(file string) ([]os.FileInfo, []string) {
// 	directoriesinfo := make([]os.FileInfo, 0)
// 	paths := make([]string, 0)
// 	err := filepath.Walk(file,
// 		func(path string, info os.FileInfo, err error) error {

// 			if info.IsDir() && info.Name() == ".git" {
// 				return filepath.SkipDir
// 			}

// 			if err != nil {
// 				return err
// 			}
// 			paths = append(paths, path)
// 			directoriesinfo = append(directoriesinfo, info)

// 			return nil
// 		})

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return directoriesinfo, paths //do not return 2 things
// }

// func getLevels(directoriesinfo []os.FileInfo, paths []string, treestruct TreeStruct) string {

// 	res := ""
// 	temp := strings.Split(os.Args[len(os.Args)-1], "/")
// 	templen := len(temp)

// 	noofdir := 0
// 	nooffiles := 0
// 	if treestruct.relpath && treestruct.permission && treestruct.dironly {
// 		for i := 0; i < len(paths); i++ {
// 			l := getLengthOfPath(paths[i])
// 			temp1 := strings.Split(paths[i], "/")
// 			if len(temp1) <= templen+treestruct.level {
// 				if directoriesinfo[i].IsDir() {
// 					res += fmt.Sprintf("%v%v [%v]%v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), paths[i])
// 					noofdir++
// 				}
// 			}
// 		}
// 		res += fmt.Sprintf("\n %v directories", noofdir-1)
// 	} else if treestruct.relpath && treestruct.permission {
// 		for i := 0; i < len(paths); i++ {
// 			l := getLengthOfPath(paths[i])
// 			temp1 := strings.Split(paths[i], "/")
// 			if len(temp1) <= templen+treestruct.level {
// 				if directoriesinfo[i].IsDir() {
// 					res += fmt.Sprintf("%v%v[%v] %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), paths[i])
// 					noofdir++
// 				} else {
// 					res += fmt.Sprintf("%v%v[%v] %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), paths[i])
// 					nooffiles++
// 				}
// 			}
// 		}
// 		res += fmt.Sprintf("\n %v directories, %v files", noofdir-1, nooffiles)
// 	} else if treestruct.relpath && treestruct.dironly {
// 		for i := 0; i < len(paths); i++ {
// 			l := getLengthOfPath(paths[i])
// 			temp1 := strings.Split(paths[i], "/")
// 			if len(temp1) <= templen+treestruct.level {
// 				if directoriesinfo[i].IsDir() {
// 					res += fmt.Sprintf("%v%v %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, paths[i])
// 					noofdir++
// 				}
// 			}
// 		}
// 		res += fmt.Sprintf("\n %v directories", noofdir-1)
// 	} else if treestruct.permission && treestruct.dironly {
// 		for i := 0; i < len(paths); i++ {
// 			l := getLengthOfPath(paths[i])
// 			temp1 := strings.Split(paths[i], "/")
// 			if len(temp1) <= templen+treestruct.level {
// 				min := int64(math.Round(math.Min(float64(len(temp1)), float64(templen+treestruct.level))))
// 				dirname := temp1[:min]
// 				if directoriesinfo[i].IsDir() {
// 					res += fmt.Sprintf("%v%v [%v]%v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), dirname[len(dirname)-1])
// 					noofdir++
// 				}
// 			}
// 		}
// 		res += fmt.Sprintf("\n %v directories", noofdir-1)
// 	} else if treestruct.permission {
// 		for i := 0; i < len(paths); i++ {
// 			l := getLengthOfPath(paths[i])
// 			temp1 := strings.Split(paths[i], "/")
// 			if len(temp1) <= templen+treestruct.level {
// 				min := int64(math.Round(math.Min(float64(len(temp1)), float64(templen+treestruct.level))))
// 				fileordirname := temp1[:min]

// 				if directoriesinfo[i].IsDir() {
// 					res += fmt.Sprintf("%v%v [%v]%v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), fileordirname[len(fileordirname)-1])
// 					noofdir++
// 				} else {
// 					nooffiles++
// 					res += fmt.Sprintf("%v%v [%v]%v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), temp1[len(fileordirname)-1])
// 				}
// 			}
// 		}
// 		res += fmt.Sprintf("\n %v directories, %v files", noofdir-1, nooffiles)
// 	} else if treestruct.relpath {
// 		for i := 0; i < len(paths); i++ {
// 			l := getLengthOfPath(paths[i])
// 			temp1 := strings.Split(paths[i], "/")
// 			if len(temp1) <= templen+treestruct.level {
// 				if directoriesinfo[i].IsDir() {
// 					noofdir++
// 					res += fmt.Sprintf("%v%v%v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, paths[i])
// 				} else {
// 					nooffiles++
// 					res += fmt.Sprintf("%v%v%v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, paths[i])
// 				}
// 			}
// 		}
// 		res += fmt.Sprintf("\n %v directories, %v files", noofdir-1, nooffiles)
// 	} else if treestruct.dironly {
// 		for i := 0; i < len(paths); i++ {
// 			l := getLengthOfPath(paths[i])
// 			temp1 := strings.Split(paths[i], "/")
// 			if len(temp1) <= templen+treestruct.level {
// 				min := int64(math.Round(math.Min(float64(len(temp1)), float64(templen+treestruct.level))))
// 				dirname := temp1[:min]

// 				if directoriesinfo[i].IsDir() {
// 					noofdir++
// 					res += fmt.Sprintf("%v%v %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, dirname[len(dirname)-1])
// 				}
// 			}
// 		}
// 		res += fmt.Sprintf("\n %v directories", noofdir-1)
// 	} else {
// 		for i := 0; i < len(paths); i++ {
// 			l := getLengthOfPath(paths[i])
// 			temp1 := strings.Split(paths[i], "/")
// 			if len(temp1) <= templen+treestruct.level {
// 				min := int64(math.Round(math.Min(float64(len(temp1)), float64(templen+treestruct.level))))
// 				fileordirname := temp1[:min]

// 				if directoriesinfo[i].IsDir() {
// 					noofdir++
// 					res += fmt.Sprintf("%v%v %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, fileordirname[len(fileordirname)-1])
// 				} else {
// 					nooffiles++
// 					res += fmt.Sprintf("%v%v %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, temp1[len(fileordirname)-1])
// 				}
// 			}
// 		}
// 		res += fmt.Sprintf("\n %v directories, %v files", noofdir-1, nooffiles)
// 	}

// 	return res
// }

// func getDirectoryOnly(directoriesinfo []os.FileInfo, paths []string) string {
// 	noofdir := 0
// 	restring := ""
// 	restring += fmt.Sprintf("%v\n", os.Args[len(os.Args)-1])
// 	temp := strings.Split(os.Args[len(os.Args)-1], "/")
// 	templen := len(temp)

// 	for i := 0; i < len(directoriesinfo); i++ {
// 		l := getLengthOfPath(paths[i])
// 		if directoriesinfo[i].IsDir() && l == templen+1 {
// 			restring += fmt.Sprintf("%v%v%v\n", BoxVer, BoxHor, directoriesinfo[i].Name())
// 			noofdir++
// 		} else if directoriesinfo[i].IsDir() && l != templen {
// 			restring += fmt.Sprintf("%v%v%v\n", strings.Repeat(" ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Name())
// 			noofdir++
// 		}
// 	}
// 	restring += fmt.Sprintf("\n%v directiories", noofdir)
// 	return restring
// }

// func getModTime(directoriesinfo []os.FileInfo, paths []string, restring string) string {

// 	var files []fs.FileInfo
// 	var err error

// 	res := ""
// 	res += fmt.Sprintf("%v\n", os.Args[len(os.Args)-1])
// 	for i := 0; i < len(directoriesinfo); i++ {
// 		l := getLengthOfPath(paths[i])
// 		if directoriesinfo[i].IsDir() {
// 			if len(restring) > 0 {
// 				res += fmt.Sprintf("%v%v%v\n", strings.Repeat("  ", l-5), BoxUpAndRig+BoxHor, directoriesinfo[i].Name())
// 				continue
// 			} else {
// 				res += fmt.Sprintf("   %v\n", directoriesinfo[i].Name())
// 			}
// 			p := paths[i]
// 			files, err = ioutil.ReadDir(p)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 		} else {
// 			sort.Slice(files, func(i, j int) bool {
// 				return files[i].ModTime().Before(files[j].ModTime())
// 			})

// 			for i := 0; i < len(files); i++ {
// 				if !files[i].IsDir() {
// 					res += fmt.Sprintf("%v%v%v \n", strings.Repeat("  ", l-4), BoxUpAndRig+BoxHor, files[i].Name())
// 				}
// 			}

// 			files = nil
// 		}
// 	}
// 	return res
// }

// func getWithPermissions(directoriesinfo []os.FileInfo, paths []string, treestruct TreeStruct) string {
// 	res := ""
// 	noofdire := 0
// 	//nooffile := 0
// 	temp := strings.Split(os.Args[len(os.Args)-1], "/")
// 	templen := len(temp)

// 	var files []fs.FileInfo
// 	var err error

// 	if treestruct.modtime && treestruct.dironly {
// 		for i := 0; i < len(directoriesinfo); i++ {
// 			l := getLengthOfPath(paths[i])
// 			//	temp1 := strings.Split(paths[i], "/")
// 			if directoriesinfo[i].IsDir() {
// 				noofdire++
// 				res += fmt.Sprintf("%v%v[%v] %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), directoriesinfo[i].Name())

// 				//res += fmt.Sprintf("[%v]%v\n", directoriesinfo[i].Mode(), directoriesinfo[i].Name())
// 				p := paths[i]
// 				files, err = ioutil.ReadDir(p)
// 				if err != nil {
// 					fmt.Println(err)
// 				}
// 			}
// 		}
// 		res += fmt.Sprintf("\n%v directiories", noofdire-1)
// 	} else if treestruct.dironly {
// 		for i := 0; i < len(directoriesinfo); i++ {
// 			l := getLengthOfPath(paths[i])
// 			if directoriesinfo[i].IsDir() {
// 				res += fmt.Sprintf("%v%v[%v] %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), directoriesinfo[i].Name())
// 			}
// 		}
// 	} else if treestruct.modtime {
// 		res = ""
// 		for i := 0; i < len(directoriesinfo); i++ {
// 			l := getLengthOfPath(paths[i])
// 			if directoriesinfo[i].IsDir() {
// 				res += fmt.Sprintf("   [%v]%v %v\n", directoriesinfo[i].Mode(), directoriesinfo[i].ModTime(), directoriesinfo[i].Name())
// 				p := paths[i]
// 				files, err = ioutil.ReadDir(p)
// 				if err != nil {
// 					fmt.Println(err)
// 				}
// 			} else {
// 				sort.Slice(files, func(i, j int) bool {
// 					return files[i].ModTime().Before(files[j].ModTime())
// 				})

// 				for i := 0; i < len(files); i++ {
// 					res += fmt.Sprintf("%v%v[%v]%v %v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, files[i].Mode(), files[i].ModTime(), files[i].Name())
// 				}
// 				files = nil
// 			}
// 		}
// 	} else {
// 		res = ""
// 		for i := 0; i < len(directoriesinfo); i++ {
// 			l := getLengthOfPath(paths[i])
// 			if directoriesinfo[i].IsDir() {
// 				res += fmt.Sprintf("%v\n", directoriesinfo[i].Name())
// 			} else {
// 				res += fmt.Sprintf("%v%v[%v]%v\n", strings.Repeat("  ", l-templen), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), directoriesinfo[i].Name())
// 			}
// 		}
// 	}

// 	return res
// }

// func getRelativePaths(directoriesinfo []os.FileInfo, paths []string, treestruct TreeStruct) string {
// 	res := ""
// 	var files []fs.FileInfo
// 	var err error

// 	if treestruct.permission && treestruct.modtime && treestruct.dironly {
// 		res = ""
// 		for i := 0; i < len(directoriesinfo); i++ {
// 			l := getLengthOfPath(paths[i])
// 			if directoriesinfo[i].IsDir() {
// 				res += fmt.Sprintf("%v%v[%v] %v %v\n", strings.Repeat("  ", l-5), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), paths[i], directoriesinfo[i].ModTime())
// 			}
// 		}
// 	} else if treestruct.permission && treestruct.dironly {
// 		res = ""
// 		for i := 0; i < len(directoriesinfo); i++ {
// 			l := getLengthOfPath(paths[i])
// 			if directoriesinfo[i].IsDir() {
// 				res += fmt.Sprintf("%v%v[%v] %v\n", strings.Repeat("  ", l-5), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), paths[i])
// 			}
// 		}
// 	} else if treestruct.modtime && treestruct.dironly {
// 		res = ""
// 		for i := 0; i < len(directoriesinfo); i++ {
// 			l := getLengthOfPath(paths[i])
// 			if directoriesinfo[i].IsDir() {
// 				res += fmt.Sprintf("%v%v%v %v\n", strings.Repeat("  ", l-5), BoxUpAndRig+BoxHor, paths[i], directoriesinfo[i].ModTime())
// 			}
// 		}
// 	} else if treestruct.modtime && treestruct.permission {
// 		res = ""
// 		for i := 0; i < len(directoriesinfo); i++ {
// 			l := getLengthOfPath(paths[i])
// 			if directoriesinfo[i].IsDir() {
// 				res += fmt.Sprintf("  [%v] %v %v \n", directoriesinfo[i].Mode(), paths[i], directoriesinfo[i].ModTime())
// 				p := paths[i]
// 				files, err = ioutil.ReadDir(p)
// 				if err != nil {
// 					fmt.Println(err)
// 				}
// 			} else {
// 				sort.Slice(files, func(i, j int) bool {
// 					return files[i].ModTime().Before(files[j].ModTime())
// 				})

// 				for i := 0; i < len(files); i++ {
// 					res += fmt.Sprintf("%v%v [%v]%v %v\n", strings.Repeat("  ", l-4), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), paths[i], files[i].ModTime())
// 				}
// 				files = nil
// 			}
// 		}
// 	} else if treestruct.dironly {
// 		res = ""
// 		for i := 0; i < len(directoriesinfo); i++ {
// 			l := getLengthOfPath(paths[i])
// 			if directoriesinfo[i].IsDir() {
// 				res += fmt.Sprintf("%v%v%v\n", strings.Repeat("  ", l-5), BoxUpAndRig+BoxHor, paths[i])
// 			}
// 		}
// 	} else if treestruct.permission {
// 		res = ""
// 		for i := 0; i < len(directoriesinfo); i++ {
// 			l := getLengthOfPath(paths[i])
// 			if directoriesinfo[i].IsDir() {
// 				res += fmt.Sprintf("   [%v]%v\n", directoriesinfo[i].Mode(), paths[i])
// 			} else {
// 				res += fmt.Sprintf("%v%v[%v]%v\n", strings.Repeat("  ", l-4), BoxUpAndRig+BoxHor, directoriesinfo[i].Mode(), paths[i])
// 			}
// 		}
// 	} else if treestruct.modtime {
// 		res = ""
// 		for i := 0; i < len(directoriesinfo); i++ {
// 			l := getLengthOfPath(paths[i])
// 			if directoriesinfo[i].IsDir() {
// 				res += fmt.Sprintf("   %v %v \n", paths[i], directoriesinfo[i].ModTime())
// 				p := paths[i]
// 				files, err = ioutil.ReadDir(p)
// 				if err != nil {
// 					fmt.Println(err)
// 				}
// 			} else {
// 				sort.Slice(files, func(i, j int) bool {
// 					return files[i].ModTime().Before(files[j].ModTime())
// 				})

// 				for i := 0; i < len(files); i++ {
// 					res += fmt.Sprintf("%v%v %v %v\n", strings.Repeat("  ", l-4), BoxUpAndRig+BoxHor, paths[i], files[i].ModTime())
// 				}
// 				files = nil
// 			}
// 		}
// 	}

// 	return res
// }
