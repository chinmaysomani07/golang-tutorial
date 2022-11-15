package main

import (
	"reflect"
	"testing"
)

// use got and want in golang instead of expected and actual which are more of a java

func TestGetAverage(t *testing.T) {

	student := Student{"Chinmay", "Somani", "RCOEM", "", 99, 99, 99, 99, 0.0}
	want := 99.0
	got := getAverage(student)

	if got != want {
		t.Errorf("Expected output: %v, but found %v", want, got)
	}
}

func TestGetOverallTopper(t *testing.T) {

	students := make([]Student, 0)
	students = append(students, Student{"Chinmay", "Somani", "RCOEM", "A", 99, 99, 99, 99, 99.0})
	students = append(students, Student{"Rishika", "Chhabrani", "Delhi University", "A", 98, 98, 98, 98, 98.0})
	students = append(students, Student{"Anmol", "Gupta", "Pune University", "A", 93, 93, 93, 93, 93.0})
	students = append(students, Student{"Gautami", "Thakare", "No University", "A", 90, 90, 90, 90, 90.0})

	got := getOverallTopper(students)
	want := Student{"Chinmay", "Somani", "RCOEM", "A", 99, 99, 99, 99, 99.0}

	if got != want {
		t.Errorf("The expected topper is: %v, but got %v", want, got)
	}
}

func TestCalculateGrade(t *testing.T) {

	avgs := 93.0
	want := "A"
	got := calculateGrade(avgs)

	if want != got {
		t.Errorf("Expected grade was %v, but got %v", want, got)
	}

	avgs = 69.0
	want = "B"
	got = calculateGrade(avgs)

	if want != got {
		t.Errorf("Expected grade was %v, but got %v", want, got)
	}
}

func TestUniversityWiseTopper(t *testing.T) {

	students := make([]Student, 0)
	students = append(students, Student{"Chinmay", "Somani", "RCOEM", "A", 99, 99, 99, 99, 99.0})
	students = append(students, Student{"Vedant", "Desai", "Amravati University", "A", 89, 54, 76, 98, 79.25})

	want := make(map[string]Student)
	want["Amravati University"] = Student{"Vedant", "Desai", "Amravati University", "A", 89, 54, 76, 98, 79.25}
	want["RCOEM"] = Student{"Chinmay", "Somani", "RCOEM", "A", 99, 99, 99, 99, 99.0}

	got := getUniversityWiseTopper(students)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected value %v, but got %v", want, got)
	}
}

//Table driven tests
