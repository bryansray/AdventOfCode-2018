package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readChanges() (changes []int) {
	file, err := os.Open("./day-1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		changes = append(changes, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func sum(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func findFirstDuplicate(changes []int) (result int) {
	seen := make(map[int]int)

	//var result int;
	for found := false; found == false; {
		for _, value := range changes {
			result += value
			seen[result] += 1

			if seen[result] > 1 {
				//fmt.Println("Duplicate: ", result)
				found = true
				break
			}
		}
	}

	return
}

func main() {
	changes := readChanges()
	sum := sum(changes)
	duplicate := findFirstDuplicate(changes)

	fmt.Println("Sum: ", sum, " - len: ", len(changes))
	fmt.Println("Duplicate: ", duplicate)
}
