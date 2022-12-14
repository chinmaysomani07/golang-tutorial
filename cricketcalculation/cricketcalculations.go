package main

import (
	"fmt"
)

func main() {

	var name string
	var surname string
	var numberOfInnings int
	var numberOfNotOut int

	fmt.Print("Enter the player name:")
	fmt.Scan(&name, &surname)

	fmt.Println("Enter the number of innings played by", name, surname, ":")
	fmt.Scan(&numberOfInnings)

	fmt.Println("Enter the nunber of not out innings by", name, surname, ":")
	fmt.Scan(&numberOfNotOut)

	if numberOfNotOut > numberOfInnings {
		fmt.Println("Number of NOT OUT innings cannot be greater than TOTAL NUMBER OF INNINGS played")
	} else {

		runsScoredByPlayer := make([]int, numberOfInnings)
		ballsFacedByPlayer := make([]int, numberOfInnings)

		for i := 0; i < numberOfInnings; i++ {
			fmt.Println("Enter the runs scored by", name, surname, "in innings", (i + 1))
			fmt.Scan(&runsScoredByPlayer[i])
			fmt.Println("Enter the number of balls faced by", name, surname, "in innings", (i + 1))
			fmt.Scan(&ballsFacedByPlayer[i])
		}

		totalRunsScored := sumOfArray(runsScoredByPlayer)
		totalBallsFaced := sumOfArray(ballsFacedByPlayer)

		fmt.Println("Total runs scored by", name, surname, "in", numberOfInnings, "innings are:", totalRunsScored)
		fmt.Println("Total balls faced by", name, surname, "in", numberOfInnings, "innings are:", totalBallsFaced)

		strikeRate := getStrikeRate(float64(totalRunsScored), float64(totalBallsFaced))
		fmt.Println("Strike Rate of", name, surname, "in", numberOfInnings, "innings is:", float32(strikeRate))

		average := getAverage(float64(totalRunsScored), float64(numberOfInnings), float64(numberOfNotOut))
		fmt.Println("Average of", name, surname, "in", numberOfInnings, "innings is:", average)

	}
}

func sumOfArray(arr []int) int {
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}

	return sum
}

func getStrikeRate(runsScored float64, ballsFaced float64) float64 { //this is a method

	strikeRate := (runsScored * 100) / ballsFaced
	return (strikeRate)
}

func getAverage(totalRunsScored float64, numberOfInnings float64, numberOfNotOut float64) float64 {

	average := totalRunsScored / (numberOfInnings - numberOfNotOut)
	return float64(average)
}
