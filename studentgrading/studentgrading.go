package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Student struct {
	FirstName    string
	LastName     string
	University   string
	Test1        float64
	Test2        float64
	Test3        float64
	Test4        float64
	AverageScore float64
	Grade        string
}

func main() {

	students := make([]Student, 0)
	toppersList := make(map[string]Student)

	//universities := make([]string, 0)
	students = readDataFromCSV("grades.csv")

	averageScoreInfo(students)

	overallTopper := getOverallTopper(students)
	fmt.Println("The overall topper is:", overallTopper.FirstName, overallTopper.LastName, "with average score of", overallTopper.AverageScore, "and grade", overallTopper.Grade)

	toppersList = getUniversityWiseTopper(students)

	for k, v := range toppersList {
		fmt.Println("University is:", k, "topper is:", v)
	}

	// universityWiseTopper := make([]Student, 0)
	// universityWiseTopper = getUniversityWiseTopper(universities, students)
	// fmt.Println("University wise toppersss list is: ", universityWiseTopper)

}

func getAverage(student Student) float64 { //this is method
	average := (student.Test1 + student.Test2 + student.Test3 + student.Test4) / 4
	return float64(average)
}

func calculateGrade(average float64) string { //this is a function

	var grade string

	if average < 35.0 {
		grade = "F"
	} else if average >= 35.0 && average < 50 {
		grade = "C"
	} else if average >= 50.0 && average < 70.0 {
		grade = "B"
	} else {
		grade = "A"
	}

	return grade
}

func averageScoreInfo(students []Student) {
	for i := 0; i < len(students); i++ {
		var average float64
		average = getAverage(students[i])
		students[i].AverageScore = average
		grade := calculateGrade(average)
		students[i].Grade = grade

		fmt.Println("Average of", students[i].FirstName, students[i].LastName, "from university", students[i].University, "is:", average)
		fmt.Println("The grade is:", grade)
	}
}

func getOverallTopper(students []Student) Student {
	//variable and assign max stud[0]
	var topper Student = students[0]

	var maxScore float64 = 0.0
	for i := 0; i < len(students); i++ {

		if students[i].AverageScore > maxScore {
			maxScore = students[i].AverageScore
			topper = students[i]
		}
	}
	return topper
}

func getUniversityWiseTopper(students []Student) map[string]Student {
	universityWiseTopperList := make(map[string]Student)

	for i := 0; i < len(students); i++ {
		if val, ok := universityWiseTopperList[students[i].University]; ok {
			if students[i].AverageScore > universityWiseTopperList[students[i].University].AverageScore {
				universityWiseTopperList[students[i].University] = val
			}
		} else {
			universityWiseTopperList[students[i].University] = students[i]
		}
	}

	return universityWiseTopperList
}

func readDataFromCSV(path string) []Student {

	studentsList := make([]Student, 0)
	csvfile, err := os.Open("grades.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer csvfile.Close()

	reader := csv.NewReader(bufio.NewReader(csvfile))
	reader.Read()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			fmt.Println(err)
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		studentsList = append(studentsList, Student{
			FirstName:    record[0],
			LastName:     record[1],
			University:   record[2],
			Test1:        parseToFloat(record[3]),
			Test2:        parseToFloat(record[4]),
			Test3:        parseToFloat(record[5]),
			Test4:        parseToFloat(record[6]),
			AverageScore: 0.0,
			Grade:        "",
		})
	}

	return studentsList
}

func parseToFloat(input string) float64 {
	number, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return float64(number)
}

//prefer functions over methods
//step down
//not using data structure names inside variable name
//test cases

//create a map and key should be university, and val will be students struct
