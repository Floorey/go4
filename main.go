package main

import (
	"fmt"
	"math/rand"
	"time"
)

func turnOnWell() {
	fmt.Println("Brunnen eingeschaltet!")
}
func turnOffWell() {
	fmt.Println("Brunnen ausgeschaltet!")
}
func controlWell(waterLevel float32) {
	if waterLevel > 20.0 {
		turnOnWell()
	} else {
		turnOffWell()
	}
}
func simulateWellControl() {
	waterLevelWell := readWaterLevelSensor("Brunnenbecken")
	controlWell(waterLevelWell)
}
func readWaterLevelSensor(sensorID string) float32 {
	waterLevel := rand.Float32() * 100.0
	fmt.Printf("Wasserstandsessor %s gelesen: %.2f%%\n", sensorID, waterLevel)
	return waterLevel
}
func controlIrrigation(waterLevel float32) {
	if waterLevel < 20.0 {
		fmt.Println("Wasserstand zu niedrig. Bewässerung gestoppt!")
	} else {
		fmt.Println("Wasserstand ausreichend. Bewässerung läuft!")
	}
}
func checkSelection(selection int, waterLevelCistern, waterLevelWell float32) bool {
	switch selection {
	case 1:
		return waterLevelWell > 20.0
	case 2:
		return waterLevelCistern > 20.0
	case 3:
		return true
	default:
		return false
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		waterLevelCistern := readWaterLevelSensor("Zisterne")
		waterLevelWell := readWaterLevelSensor("Brunnenbecken")

		fmt.Println("Please insert your choice:")
		fmt.Println("1. Brunnen")
		fmt.Println("2. Bewässerung")
		fmt.Println("3. Beenden")

		var selection int
		fmt.Scanln(&selection)

		if checkSelection(selection, waterLevelCistern, waterLevelWell) {
			switch selection {
			case 1:
				controlWell(waterLevelWell)
				simulateWellControl()
			case 2:
				controlIrrigation(waterLevelCistern)
			case 3:
				fmt.Println("Pogramm beendet")
				return
			}

		} else {
			fmt.Println("Invalid Choice.")
		}

		time.Sleep(3 * time.Second)
	}
}
