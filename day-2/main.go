package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readBoxIds(input string) (boxIds []string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value := scanner.Text()
		boxIds = append(boxIds, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func group(value string) map[string]int {
	slice := strings.Split(value, "")
	characters := make(map[string]int)

	for _, character := range slice {
		characters[character] += 1
	}

	return characters
}

func analyzeIds(characters map[string]int) (twos int, threes int) {
	for _, value := range characters {
		if value == 2 {
			twos += 1
		}
		if value == 3 {
			threes += 1
		}
	}

	return
}

func main() {
	boxIds := readBoxIds("./day-2/input.txt")

	//i := 0

	var totalTwos, totalThrees int
	for _, boxId := range boxIds {
		grouping := group(boxId)

		twos, threes := analyzeIds(grouping)

		if twos > 0 { totalTwos += 1 }
		if threes > 0 { totalThrees += 1 }
	}

	checksum := totalTwos * totalThrees

	fmt.Println("len: ", len(boxIds))
	fmt.Println("checksum: ", checksum)

	found := false

	for index1, boxId := range boxIds  {
		characters := strings.Split(boxId, "")

		for index, matchingBoxId := range boxIds {
			if boxId == matchingBoxId {
				continue
			}

			matchCharacters := strings.Split(matchingBoxId, "")
			differences := 0

			for index, character := range characters {
				if matchCharacters[index] != character {
					differences += 1
				}
			}

			if differences == 1 {
				fmt.Println(boxId, index1)
				fmt.Println(matchingBoxId, index)
				found = true
			}
		}

		if found == true {
			break
		}
	}
}
