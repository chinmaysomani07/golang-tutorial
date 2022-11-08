package main

import "fmt"

func main() {

	var name string
	var surname string
	fmt.Print("Enter the player name:")
	fmt.Scan(&name, &surname)
	runsScoredByKohli := [5]int{73, 86, 1, 72, 65}
	ballsFacedByKohli := [5]int{45, 63, 3, 87, 73}
	totalRunsScored := sumOfArray(runsScoredByKohli)
	totalBallsFaced := sumOfArray(ballsFacedByKohli)

	fmt.Println("Total runs scored by", name, surname, "in last 5 innings are:", totalRunsScored)
	fmt.Println("Total balls faced by", name, surname, "in last 5 innings are:", totalBallsFaced)

	strikeRate := getStrikeRate(float64(totalRunsScored), float64(totalBallsFaced))
	fmt.Println("Strike Rate of", name, surname, "in last 5 innings is:", float32(strikeRate))

}

func sumOfArray(arr [5]int) int {
	sum := 0

	for i := 0; i < 5; i++ {
		sum = sum + arr[i]
	}

	return sum
}

func getStrikeRate(runsScored float64, ballsFaced float64) float64 {

	strikeRate := (runsScored * 100) / ballsFaced
	return (strikeRate)
}
