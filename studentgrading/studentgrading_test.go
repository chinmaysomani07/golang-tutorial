package main

import (
	"reflect"
	"testing"
)

func TestGetAverage(t *testing.T) {

	student := Student{"Chinmay", "Somani", "RCOEM", "", 99, 99, 99, 99, 0.0}
	expectedavgm := 99.0
	avgm := getAverage(student)

	if avgm != expectedavgm {
		t.Errorf("Expected output: %v, but found %v", expectedavgm, avgm)
	}
}

func TestGetOverallTopper(t *testing.T) {

	students := make([]Student, 0)
	students = append(students, Student{"Chinmay", "Somani", "RCOEM", "A", 99, 99, 99, 99, 99.0})
	students = append(students, Student{"Rishika", "Chhabrani", "Delhi University", "A", 98, 98, 98, 98, 98.0})
	students = append(students, Student{"Anmol", "Gupta", "Pune University", "A", 93, 93, 93, 93, 93.0})
	students = append(students, Student{"Gautami", "Thakare", "No University", "A", 90, 90, 90, 90, 90.0})

	topper := getOverallTopper(students)
	exp_topper := Student{"Chinmay", "Somani", "RCOEM", "A", 99, 99, 99, 99, 99.0}

	if topper != exp_topper {
		t.Errorf("The expected topper is: %v, but got %v", exp_topper, topper)
	}
}

func TestCalculateGrade(t *testing.T) {

	avgs := 93.0
	expGrade := "A"
	resgrade := calculateGrade(avgs)

	if expGrade != resgrade {
		t.Errorf("Expected grade was %v, but got %v", expGrade, resgrade)
	}

	avgs = 69.0
	expGrade = "B"
	resgrade = calculateGrade(avgs)

	if expGrade != resgrade {
		t.Errorf("Expected grade was %v, but got %v", expGrade, resgrade)
	}
}

func TestUniversityWiseTopper(t *testing.T) {

	students := make([]Student, 0)
	students = append(students, Student{"Chinmay", "Somani", "RCOEM", "A", 99, 99, 99, 99, 99.0})
	students = append(students, Student{"Vedant", "Desai", "Amravati University", "A", 89, 54, 76, 98, 79.25})

	mtl := make(map[string]Student)
	mtl["Amravati University"] = Student{"Vedant", "Desai", "Amravati University", "A", 89, 54, 76, 98, 79.25}
	mtl["RCOEM"] = Student{"Chinmay", "Somani", "RCOEM", "A", 99, 99, 99, 99, 99.0}

	toppersList := getUniversityWiseTopper(students)

	if !reflect.DeepEqual(mtl, toppersList) {
		t.Errorf("Expected value %v, but got %v", mtl, toppersList)
	}
}
