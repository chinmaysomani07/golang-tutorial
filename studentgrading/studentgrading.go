package main

import (
	"fmt"
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
	students = addStudents()

	for i := 0; i < len(students); i++ {
		var average float64
		average = students[i].getAverage()
		students[i].AverageScore = average
		grade := calculateGrade(average)
		students[i].Grade = grade

		fmt.Println("Average of", students[i].FirstName, students[i].LastName, "from university", students[i].University, "is:", average)
		fmt.Println("The grade is:", grade)
	}

	overallTopper := getOverallTopper(students)
	fmt.Println("The overall topper is:", overallTopper.FirstName, overallTopper.LastName, "with average score of", overallTopper.AverageScore, "and grade", overallTopper.Grade)

	universityWiseTopper := make([]Student, 0)
	universities := make([]string, 0)
	universities = append(universities, "RCOEM", "Mumbai University", "Delhi University", "Pune University")

	for i := 0; i < len(universities); i++ {
		topper := Student{"", "", "", 0.0, 0.0, 0.0, 0.0, 0.0, ""}
		if universities[i] == "RCOEM" {
			topper = getUniversityTopper(students, universities[i])
			universityWiseTopper = append(universityWiseTopper, topper)
		} else if universities[i] == "Mumbai University" {
			topper = getUniversityTopper(students, universities[i])
			universityWiseTopper = append(universityWiseTopper, topper)
		} else if universities[i] == "Delhi University" {
			topper = getUniversityTopper(students, universities[i])
			universityWiseTopper = append(universityWiseTopper, topper)
		} else {
			topper = getUniversityTopper(students, universities[i])
			universityWiseTopper = append(universityWiseTopper, topper)
		}
	}

	fmt.Println("University wise toppers list is: ", universityWiseTopper)

}

func addStudents() []Student {

	students := make([]Student, 0)
	students = append(students, Student{"Chinmay", "Somani", "RCOEM", 76, 87, 66, 37, 0.0, ""})
	students = append(students, Student{"Rutvik", "Bhute", "RCOEM", 76, 72, 66, 36, 0.0, ""})
	students = append(students, Student{"Rishika", "Chhabrani", "Delhi University", 71, 87, 92, 72, 0.0, ""})
	students = append(students, Student{"Shreya", "Lahoti", "Mumbai University", 87, 83, 92, 54, 0.0, ""})
	students = append(students, Student{"Anmol", "Gupta", "Pune University", 82, 73, 64, 77, 0.0, ""})

	return students
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

func getUniversityTopper(students []Student, university string) Student {
	var topper Student
	var maxAverageScore float64

	for j := 0; j < len(students); j++ {
		if university == "RCOEM" {
			if students[j].AverageScore > maxAverageScore && students[j].University == university {
				maxAverageScore = students[j].AverageScore
				topper = students[j]
				myHashMap[students[j].University] = students[j]
			}
		}
		if university == "Mumbai University" {
			if students[j].AverageScore > maxAverageScore && students[j].University == university {
				maxAverageScore = students[j].AverageScore
				topper = students[j]
				myHashMap[students[j].University] = students[j]
			}
		}
		if university == "Delhi University" {
			if students[j].AverageScore > maxAverageScore && students[j].University == university {
				maxAverageScore = students[j].AverageScore
				topper = students[j]
				myHashMap[students[j].University] = students[j]
			}
		}
		if university == "Pune University" {
			if students[j].AverageScore > maxAverageScore && students[j].University == university {
				maxAverageScore = students[j].AverageScore
				topper = students[j]
				myHashMap[students[j].University] = students[j]
			}
		}
	}
	return topper
}

//prefer functions over methods
