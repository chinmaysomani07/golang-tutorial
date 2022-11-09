package main

import (
	"fmt"
)

type Student struct {
	FirstName  string
	LastName   string
	University string
	Test1      float64
	Test2      float64
	Test3      float64
	Test4      float64
}

func main() {

	students := make([]Student, 0)
	students = append(students, Student{"Chinmay", "Somani", "RCOEM Uni.", 76, 87, 66, 37})
	students = append(students, Student{"Rutvik", "Bhute", "RCOEM Uni.", 76, 72, 66, 36})

	for i := 0; i < len(students); i++ {
		var average float64
		average = students[i].getAverage()

		grade := calculateGrade(average)

		fmt.Println("Average of", students[i].FirstName, students[i].LastName, "from university", students[i].University, "is:", average)
		fmt.Println("The grade is:", grade)
	}

}

func (student Student) getAverage() float64 { //this is function
	average := (student.Test1 + student.Test2 + student.Test3 + student.Test4) / 4
	return float64(average)
}

func calculateGrade(average float64) string { //this is a method

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
