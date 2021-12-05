package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

type pair struct {
	x, y int64
}

func getLineEndPoints(line string) [2]pair {
	var x, y int64
	parts := strings.Split(line, "->")
	// fmt.Println(parts[0], parts[1])
	fs := strings.Split(strings.Trim(parts[0], " "), ",")
	x, _ = strconv.ParseInt(fs[0], 10, 64)
	y, _ = strconv.ParseInt(fs[1], 10, 64)
	first := pair{x, y}
	// fmt.Println(first)
	sc := strings.Split(strings.Trim(parts[1], " "), ",")
	x, _ = strconv.ParseInt(sc[0], 10, 64)
	y, _ = strconv.ParseInt(sc[1], 10, 64)
	second := pair{x, y}
	// fmt.Println(second)

	return [...]pair{first, second}
}

func allInBetweenPoints(ends [2]pair) []pair {
	var pts []pair
	if ends[0].x == ends[1].x {
		var minY, maxY int64
		if ends[0].y < ends[1].y {
			minY = ends[0].y
			maxY = ends[1].y
		} else {
			minY = ends[1].y
			maxY = ends[0].y
		}
		for i := minY; i <= maxY; i++ {
			pts = append(pts, pair{ends[0].x, i})
		}
	} else if ends[0].y == ends[1].y {
		var minX, maxX int64
		if ends[0].x < ends[1].x {
			minX = ends[0].x
			maxX = ends[1].x
		} else {
			minX = ends[1].x
			maxX = ends[0].x
		}
		for i := minX; i <= maxX; i++ {
			pts = append(pts, pair{i, ends[0].y})
		}
	} else {
		// Diagonal Case
		var minX, maxX, mul, sY int64
		if ends[0].x < ends[1].x {
			minX = ends[0].x
			maxX = ends[1].x
			sY = ends[0].y
			mul = -1
			if sY < ends[1].y {
				mul = 1
			}
		} else {
			minX = ends[1].x
			maxX = ends[0].x
			sY = ends[1].y
			mul = -1
			if sY < ends[0].y {
				mul = 1
			}
		}
		for i := minX; i <= maxX; i++ {
			pts = append(pts, pair{i, sY + mul*(i-minX)})
		}

	}
	return pts
}

func main() {
	var vents [1000][1000]int

	for scanner.Scan() {
		line := scanner.Text()
		endPoints := getLineEndPoints(line)
		allPoints := allInBetweenPoints(endPoints)

		fmt.Println(allPoints)
		for _, p := range allPoints {
			vents[p.x][p.y] += 1
		}
	}
	total := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if vents[i][j] > 1 {
				total += 1
			}
		}
	}
	fmt.Println(total)
}
