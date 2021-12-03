package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinates struct {
	horizontal int
	depth      int
	aim        int
}

func main() {
	path, _ := os.Getwd()
	file, err := os.Open(path + "/input.txt")

	if err != nil {
		log.Fatalf("failed to open err=%s", err)

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	position := coordinates{horizontal: 0, depth: 0}
	manualPosition := coordinates{horizontal: 0, depth: 0, aim: 0}

	for scanner.Scan() {
		executeCommand(&position, scanner.Text())
		executeManualCommand(&manualPosition, scanner.Text())
	}

	log.Printf("Final position horizontal=%d, depth=%d, total=%d",
		position.horizontal, position.depth, position.horizontal*position.depth)

	log.Printf("Final manual position horizontal=%d, depth=%d, aim=%d, total=%d",
		manualPosition.horizontal, manualPosition.depth, manualPosition.aim, manualPosition.horizontal*manualPosition.depth)
	file.Close()
}

func executeCommand(position *coordinates, command string) {
	move := strings.Split(command, " ")[0]
	quantity, _ := strconv.Atoi(strings.Split(command, " ")[1])

	switch move {
	case "forward":
		position.horizontal += quantity
	case "up":
		position.depth -= quantity
	case "down":
		position.depth += quantity
	}
}

func executeManualCommand(position *coordinates, command string) {
	move := strings.Split(command, " ")[0]
	quantity, _ := strconv.Atoi(strings.Split(command, " ")[1])
	log.Printf("M=executeManualCommand, action=Moving position=%+v, command=%s", *position, command)
	switch move {
	case "forward":
		position.horizontal += quantity
		position.depth += position.aim * quantity
	case "up":
		position.aim -= quantity
	case "down":
		position.aim += quantity
	}
}
