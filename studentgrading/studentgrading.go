package main

import (
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

var myHashMap = make(map[string]Student)

func main() {

	students := make([]Student, 0)

	universities := make([]string, 0)
	students, universities = readDataFromCSV("grades.csv")

	getAverageScoreInfo(students)

	overallTopper := getOverallTopper(students)
	fmt.Println("The overall topper is:", overallTopper.FirstName, overallTopper.LastName, "with average score of", overallTopper.AverageScore, "and grade", overallTopper.Grade)

	universityWiseTopper := make([]Student, 0)
	universityWiseTopper = getUniversityWiseTopper(universities, students)
	fmt.Println("University wise toppersss list is: ", universityWiseTopper)

}

func (student Student) getAverage() float64 { //this is method
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

func getAverageScoreInfo(students []Student) {
	for i := 0; i < len(students); i++ {
		var average float64
		average = students[i].getAverage()
		students[i].AverageScore = average
		grade := calculateGrade(average)
		students[i].Grade = grade

		fmt.Println("Average of", students[i].FirstName, students[i].LastName, "from university", students[i].University, "is:", average)
		fmt.Println("The grade is:", grade)
	}
}

func getOverallTopper(students []Student) Student {
	topper := Student{"", "", "", 0.0, 0.0, 0.0, 0.0, 0.0, ""}
	var maxScore float64 = 0.0
	for i := 0; i < len(students); i++ {

		if students[i].AverageScore > maxScore {
			maxScore = students[i].AverageScore
			topper = students[i]
		}
	}
	return topper
}

func getUniversityWiseTopper(universities []string, students []Student) []Student {
	universityWiseTopperSlice := make([]Student, 0)
	for i := 0; i < len(universities); i++ {
		topper := Student{"", "", "", 0.0, 0.0, 0.0, 0.0, 0.0, ""}
		if universities[i] == "RCOEM" {
			topper = getIndividualUniversityTopper(students, universities[i])
			universityWiseTopperSlice = append(universityWiseTopperSlice, topper)
		} else if universities[i] == "Mumbai University" {
			topper = getIndividualUniversityTopper(students, universities[i])
			universityWiseTopperSlice = append(universityWiseTopperSlice, topper)
		} else if universities[i] == "Delhi University" {
			topper = getIndividualUniversityTopper(students, universities[i])
			universityWiseTopperSlice = append(universityWiseTopperSlice, topper)
		} else {
			topper = getIndividualUniversityTopper(students, universities[i])
			universityWiseTopperSlice = append(universityWiseTopperSlice, topper)
		}
	}

	return universityWiseTopperSlice
}

func getIndividualUniversityTopper(students []Student, university string) Student {
	var topper Student
	var maxAverageScore float64

	for j := 0; j < len(students); j++ {
		if university == "RCOEM" {
			if students[j].AverageScore > maxAverageScore && students[j].University == university {
				maxAverageScore = students[j].AverageScore
				topper = students[j]
				myHashMap[students[j].University] = students[j]
			}
		} else if university == "Mumbai University" {
			if students[j].AverageScore > maxAverageScore && students[j].University == university {
				maxAverageScore = students[j].AverageScore
				topper = students[j]
				myHashMap[students[j].University] = students[j]
			}
		} else if university == "Delhi University" {
			if students[j].AverageScore > maxAverageScore && students[j].University == university {
				maxAverageScore = students[j].AverageScore
				topper = students[j]
				myHashMap[students[j].University] = students[j]
			}
		} else {
			if students[j].AverageScore > maxAverageScore && students[j].University == university {
				maxAverageScore = students[j].AverageScore
				topper = students[j]
				myHashMap[students[j].University] = students[j]
			}
		}
	}
	return topper
}

func readDataFromCSV(path string) ([]Student, []string) {

	studentsList := make([]Student, 0)
	universities := make([]string, 0)
	csvfile, err := os.Open("grades.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader((csvfile))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		studentsList = append(studentsList, Student{
			FirstName:    record[0],
			LastName:     record[1],
			University:   record[2],
			Test1:        parseToInt(record[3]),
			Test2:        parseToInt(record[4]),
			Test3:        parseToInt(record[5]),
			Test4:        parseToInt(record[6]),
			AverageScore: 0.0,
			Grade:        "",
		})

		universities = append(universities, record[2])
	}

	fmt.Println(studentsList)
	fmt.Println(universities)

	return studentsList, universities
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
