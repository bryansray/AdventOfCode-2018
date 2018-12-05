package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type sleepLog struct {
	minutes int
	time []int
}

type event struct {
	timestamp time.Time
	message string
}

type sortedEvents []event

func (slice sortedEvents) Less(i, j int) bool 	{ return slice[i].timestamp.Before(slice[j].timestamp) }
func (slice sortedEvents) Swap(i, j int) 		{ slice[i], slice[j] = slice[j], slice[i] }
func (slice sortedEvents) Len() int 			{ return len(slice) }

func getId(message string) (id int) {
	regex := regexp.MustCompile(`\d{1,4}`)

	id, _ = strconv.Atoi(regex.FindString(message))

	return
}

func find(message string, search string) bool {
	return strings.Contains(message, search)
}

func parse(filename string) (results sortedEvents) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	regex := regexp.MustCompile(`\[(\d{4}\-\d{2}\-\d{2}\s\d{2}:\d{2})\]\s(.*)`)


	var events []event
	for scanner.Scan() {
		line := scanner.Text()

		timestamp, _ := time.Parse("2006-01-02 15:04", regex.FindStringSubmatch(line)[1])
		message := regex.FindStringSubmatch(line)[2]

		events = append(events, event{ timestamp: timestamp, message: message })
	}

	sort.Sort(sortedEvents(events))

	results = events

	return
}

func main() {
	events := parse("./day-4/input.txt")

	guards := make(map[int]*sleepLog)

	var id int
	var start time.Time
	var end time.Time

	for _, event := range events {
		if find(event.message, "Guard") {
			id = getId(event.message)
			_, ok := guards[id]
			if ok {
				continue
			} else {
				guards[id] = &sleepLog{
					time: make([]int, 60),
					minutes: 0,
				}
			}
		}

		if find(event.message, "falls") {
			start = event.timestamp
		}

		if find(event.message, "wakes") {
			end = event.timestamp

			e := end.Sub(start)
			s := int(time.Duration(start.Minute()))
			d := int(e / time.Minute)

			for j := s; j < s + d; j++ {
				guards[id].time[j]++
				guards[id].minutes++
			}
		}
	}

	//fmt.Println(len(guards))
	var longest, foundId int

	for key, value := range guards {
		if value.minutes > longest {
			longest = value.minutes
			foundId = key
		}
	}

	var max, index int
	for i, value := range guards[foundId].time {
		if value > max {
			max = value
			index = i
		}
	}

	fmt.Println(foundId, index, foundId* index)

	// Part 2
	longest = 0
	for key, value := range guards {
		for i, j := range value.time {
			if j > longest {
				longest = j
				index = i
				foundId = key
			}
		}
	}

	fmt.Println(foundId, index, foundId* index)
}
