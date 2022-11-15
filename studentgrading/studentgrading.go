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
	FirstName, LastName, University, Grade   string
	Test1, Test2, Test3, Test4, AverageScore float64
}

func main() {

	students := make([]Student, 0)
	toppersList := make(map[string]Student)

	students = readDataFromCSV("grades.csv")

	averageScoreInfo(students)

	overallTopper := getOverallTopper(students)
	fmt.Println("The overall topper is:", overallTopper.FirstName, overallTopper.LastName, "with average score of", overallTopper.AverageScore, "and grade", overallTopper.Grade)

	toppersList = getUniversityWiseTopper(students)

	displayUniversityWiseTopper(toppersList)
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

func getAverage(student Student) float64 {
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
	universityWiseTopperList := make(map[string]Student) //change naming

	for i := 0; i < len(students); i++ {
		if _, ok := universityWiseTopperList[students[i].University]; ok {
			if students[i].AverageScore > universityWiseTopperList[students[i].University].AverageScore {
				universityWiseTopperList[students[i].University] = students[i]
			}
		} else {
			universityWiseTopperList[students[i].University] = students[i]
		}
	}

	return universityWiseTopperList
}

func displayUniversityWiseTopper(toppersList map[string]Student) { //change name
	for k, v := range toppersList {
		fmt.Println("University is:", k, "and the topper is:", v)
	}
}

func (student Student) String() string {
	return fmt.Sprintf("%v %v with an average score %v and grade: %v", student.FirstName, student.LastName, student.AverageScore, student.Grade)
}
