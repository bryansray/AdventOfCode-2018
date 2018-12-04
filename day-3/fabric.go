package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	X int
	Y int
}

type rectangle struct {
	Width  int
	Height int
}

type fabric struct {
	id        int
	point     point
	rectangle rectangle
}

func (fabric fabric) Right() int {
	return fabric.point.X + fabric.rectangle.Width
}

func (fabric fabric) Bottom() int {
	return fabric.point.Y + fabric.rectangle.Height
}

func min(x, y int)	int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func (fabric fabric) Overlap(other fabric) []point {
	left := max(fabric.point.X, other.point.X)
	top := max(fabric.point.Y, other.point.Y)

	right := min(fabric.Right(), other.Right())
	bottom := min(fabric.Bottom(), other.Bottom())

	var points []point

	for x := left; x < right; x++ {
		for y := top; y < bottom; y++ {
			points = append(points, point{x, y})
		}
	}

	return points
}

func parse(filename string) []fabric {
	var fabrics []fabric

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	regex := regexp.MustCompile(`(?P<id>\d{1,4})\s@\s(?P<x>\d+),(?P<y>\d+):\s(?P<length>\d+)x(?P<width>\d+)`)

	for scanner.Scan() {
		line := scanner.Text()

		claimId, _ := strconv.Atoi( regex.FindStringSubmatch(line)[1] )
		x, _ := strconv.Atoi( regex.FindStringSubmatch(line)[2] )
		y, _ := strconv.Atoi( regex.FindStringSubmatch(line)[3] )

		width, _ := strconv.Atoi( regex.FindStringSubmatch(line)[4] )
		height, _ := strconv.Atoi( regex.FindStringSubmatch(line)[5] )

		fabric := fabric{
			id: claimId,
			point: point{
				X: x,
				Y: y,
			},
			rectangle: rectangle{
				Height: height,
				Width: width,
			},
		}

		fabrics = append(fabrics, fabric)
	}

	return fabrics
}

// 159115 -- too high
// 157210 -- too high
// 116489 -- right :)

func main() {
	fabrics := parse("./day-3/input.txt")

	clothes := make(map[string]int)

	for _, fabric := range fabrics {
		for i := fabric.point.X; i < fabric.point.X + fabric.rectangle.Width; i++ {
			for j := fabric.point.Y; j < fabric.point.Y + fabric.rectangle.Height; j++ {
				clothes[fmt.Sprintf("%d,%d", i, j)]++
			}
		}
	}

	var overlapping int
	for _, j := range clothes {
		if j > 1 {
			overlapping++
		}
	}

	fmt.Println("Part 1: ", overlapping)

	hasOverlaps := make([]bool, len(fabrics))
	for x := 0; x < len(fabrics); x++ {
		for y := x + 1; y < len(fabrics); y++ {
			firstFabric := fabrics[x]
			secondFabric := fabrics[y]
			overlap := firstFabric.Overlap(secondFabric)
			hasOverlap := len(overlap) > 0
			hasOverlaps[x] = hasOverlaps[x] || hasOverlap
			hasOverlaps[y] = hasOverlaps[y] || hasOverlap
		}
	}

	for x := 0; x < len(hasOverlaps); x++ {
		if !hasOverlaps[x] {
			fmt.Println("Part 2: ", fabrics[x].id)
		}
	}
}
