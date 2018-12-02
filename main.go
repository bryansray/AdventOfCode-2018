package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(array *[]int) int {
	sum := 0
	for _, value := range *array {
		sum += value
	}
	return sum
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var result int;
	var values []int
	seen := make(map[int]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		values = append(values, value)
	}

	fmt.Println("Sum: ", sum(&values), " - len: ", len(values))

	//var found bool

	for found := false; found == false; {
		for _, value := range values {
			result += value
			seen[result] += 1

			if seen[result] > 1 {
				fmt.Println("Duplicate: ", result)
				found = true
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
