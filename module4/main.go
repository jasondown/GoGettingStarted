package main

import (
	"fmt"
)

var (
	plantCapacities = []float64{30, 30, 30, 60, 60, 100}
	activePlants    = []int{0, 1}
	gridLoad        = 75.
)

func main() {
	displayReportMenu()
	var option = getReportMenuOption()
	handleReportMenuOption(option)
}

func displayReportMenu() {
	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Println("Please choose and option: ")
}

func getReportMenuOption() string {
	var option string
	fmt.Scanln(&option)
	return option
}

func handleReportMenuOption(option string) {

	switch option {
	case "1":
		generatePlantCapacityReport(plantCapacities...)
	case "2":
		generatePowerGridReport(activePlants, plantCapacities, gridLoad)
	default:
		fmt.Println("Unknown option, no action taken")
	}
}

func generatePlantCapacityReport(plantCapacities ...float64) {
	for idx, cap := range plantCapacities {
		fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
	}
}

func generatePowerGridReport(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantID := range activePlants {
		capacity += plantCapacities[plantID]
	}

	fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.0f\n", "Load:", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization:", gridLoad/capacity*100)
}
