package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parse(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	return line
}

func isReactable(character string, otherCharacter string) bool {
	if character[0] - otherCharacter[0] == 32 || otherCharacter[0] - character[0] == 32 {
		return true
	}

	return false
}

func removeUnit(polymers string, polymer string) string {
	return strings.Replace(strings.Replace(polymers, polymer, "", -1), string(polymer[0] - 32), "", -1)
}

func combust(units []string) []string {
	for i := 0; i < len(units) - 1; i++ {
		if len(units) < i {
			break
		}

		if isReactable(units[i], units[i + 1]) {
			units = append(units[:i], units[i + 2:]...)

			i = i - 2

			if i < -1 { i = -1 }
		}
	}

	return units
}

func main() {
	polymer := parse("./day-5/input.txt")
	units := strings.Split(polymer, "")

	result := combust(units)

	fmt.Println("Part 1: ", len(result), result)

	//polymer = parse("./day-5/input.txt")

	combustionChamber := make(map[int32]int)

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	var adjustedPolymer string
	for _, character := range alphabet {
		adjustedPolymer = removeUnit(polymer, string(character))
		units = strings.Split(adjustedPolymer, "")

		result = combust(units)

		combustionChamber[character] = len(result)
	}

	smallest := int(combustionChamber[97])
	for _, value := range combustionChamber {
		if value < smallest {
			smallest = value
		}
	}

	fmt.Println("Part 2: ", smallest)
}
