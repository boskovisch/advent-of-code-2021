package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

type Window struct {
	index  int
	values [3]int
}

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/input.txt")

	if err != nil {
		log.Fatalf("failed to open err=%s", err)

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	numLargerMeasurements := 0
	lastValue := -1

	numLargerWindowsMeasurements := 0
	lastWindowValue := -1
	window := Window{index: 0, values: [3]int{-1, -1, -1}}

	for scanner.Scan() {
		measurement, _ := strconv.Atoi(scanner.Text())
		compareWithPreviousMeasurement(&numLargerMeasurements, &lastValue, measurement)
		compareWithPreviousMeasurementWindow(&numLargerWindowsMeasurements, &window, &lastWindowValue, measurement)
	}

	log.Printf("Total increased measurements is %d\n", numLargerMeasurements)
	log.Printf("Total increased windows measurements is %d\n", numLargerWindowsMeasurements)
	file.Close()
}

func compareWithPreviousMeasurement(numLargerMeasurements *int, lastValue *int, measurement int) {
	if *lastValue != -1 && measurement > *lastValue {
		//log.Printf("M=main, Action=Value is increased, lastValue=%d, measurement=%d\n", lastValue, measurement)
		*numLargerMeasurements++
	}
	*lastValue = measurement
}

func compareWithPreviousMeasurementWindow(numLargerWindowsMeasurements *int, window *Window, lastWindowValue *int, measurement int) {
	index := window.index
	window.values[index] = measurement
	window.index = nextIndex(index)

	currentWindowValue, err := calculateCurrentWindow(*window, measurement)

	if err == nil {
		if *lastWindowValue != -1 && currentWindowValue > *lastWindowValue {
			log.Printf("M=main, Action=Value is increased, lastWindowValue=%d, currentWindowValue=%d\n", lastWindowValue, currentWindowValue)
			*numLargerWindowsMeasurements++
		}
		*lastWindowValue = currentWindowValue
	}
}

func nextIndex(index int) int {
	if index < 2 {
		return index + 1
	}
	return 0

}

func calculateCurrentWindow(window Window, measurement int) (int, error) {
	currentWindowValue := 0

	for i := range [3]int{} {
		if window.values[i] == -1 {
			return -1, errors.New("not a window yet")
		}
		currentWindowValue += window.values[i]
	}
	return currentWindowValue, nil
}
