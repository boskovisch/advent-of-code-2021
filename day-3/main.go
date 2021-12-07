package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type diagnosticReport struct {
	zeroes            []int
	ones              []int
	data              []string
	mostCommonBinary  string
	leastCommonBinary string
}

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/input.txt")

	if err != nil {
		log.Fatalf("failed to open err=%s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	report := diagnosticReport{ones: make([]int, 12), zeroes: make([]int, 12)}
	for scanner.Scan() {
		extractFromReport(&report, scanner.Text())
		report.data = append(report.data, scanner.Text())
	}
	findFrequencyFromReport(&report)
	gammaRate, _ := strconv.ParseInt(report.mostCommonBinary, 2, 64)
	epsilonRate, _ := strconv.ParseInt(report.leastCommonBinary, 2, 64)
	powerConsumption := gammaRate * epsilonRate

	log.Printf("Finding oxygenGeneratorRating from mostCommonBinary=%s", report.mostCommonBinary)
	oxygenGeneratorRating := findRating(report.data, 0, true)
	oxygenGeneratorRatingDecimal, _ := strconv.ParseInt(oxygenGeneratorRating, 2, 64)
	log.Printf("Finding co2GeneratorRating from leastCommonBinary=%s", report.leastCommonBinary)
	co2GeneratorRating := findRating(report.data, 0, false)
	co2GeneratorRatingDecimal, _ := strconv.ParseInt(co2GeneratorRating, 2, 64)
	lifeSupportRating := oxygenGeneratorRatingDecimal * co2GeneratorRatingDecimal
	log.Printf("mostCommonBinary=%s, leastCommonBinary=%s, gammaRate=%d, epsilonRate=%d, powerConsumption=%d",
		report.mostCommonBinary, report.leastCommonBinary, gammaRate, epsilonRate, powerConsumption)

	log.Printf("oxygenGeneratorRating={bin=%s, dec=%d}, co2GeneratorRating={bin=%s, dec=%d}, lifeSupport=%d",
		oxygenGeneratorRating, oxygenGeneratorRatingDecimal, co2GeneratorRating, co2GeneratorRatingDecimal, lifeSupportRating)
}

func extractFromReport(report *diagnosticReport, reportRow string) {
	for i := 0; i < len(reportRow); i++ {
		if string(reportRow[i]) == "1" {
			report.ones[i]++
			continue
		}
		report.zeroes[i]++
	}
}

func findFrequencyFromReport(report *diagnosticReport) {
	for i := 0; i < len(report.ones); i++ {
		if report.ones[i] > report.zeroes[i] {
			report.mostCommonBinary += "1"
			report.leastCommonBinary += "0"
			continue
		}
		report.mostCommonBinary += "0"
		report.leastCommonBinary += "1"
	}
}

func findRating(report []string, position int, mostCommon bool) string {
	log.Printf("reportSize=%d, postion=%d", len(report), position)
	if len(report) == 1 {
		return report[0]
	}

	filteredSlice := filterSlice(report, position, mostCommon)

	return findRating(filteredSlice, position+1, mostCommon)
}

func filterSlice(slice []string, position int, mostCommon bool) []string {
	var zeroesSlice []string
	var onesSlice []string

	for _, row := range slice {
		if string(row[position]) == "0" {
			zeroesSlice = append(zeroesSlice, row)
			continue
		}
		onesSlice = append(onesSlice, row)
	}
	if mostCommon {
		if len(onesSlice) >= len(zeroesSlice) {
			return onesSlice
		}
		return zeroesSlice
	}

	if len(zeroesSlice) <= len(onesSlice) {
		return zeroesSlice
	}
	return onesSlice
}
