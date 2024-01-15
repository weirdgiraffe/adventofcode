package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Item struct {
	Direction byte
	Count     int
	Color     string
}

func (i Item) String() string {
	return fmt.Sprintf("%c %d %s", i.Direction, i.Count, i.Color)
}

func part1(lines []string) []Item {
	item := make([]Item, len(lines))
	for i, line := range lines {
		l := strings.Split(line, " ")
		count, err := strconv.ParseInt(l[1], 10, 64)
		if err != nil {
			panic("wrong number")
		}
		item[i] = Item{
			Direction: l[0][0],
			Count:     int(count),
			Color:     l[2],
		}
	}
	return item
}

func part2(lines []string) []Item {
	item := make([]Item, len(lines))
	for i, line := range lines {
		l := strings.Split(line, " ")
		switch l[2][7] {
		case '0':
			item[i].Direction = 'R'
		case '1':
			item[i].Direction = 'D'
		case '2':
			item[i].Direction = 'L'
		case '3':
			item[i].Direction = 'U'
		}
		count, err := strconv.ParseInt(l[2][2:7], 16, 64)
		if err != nil {
			panic("wrong number")
		}
		item[i].Count = int(count)
	}
	return item
}

type Point struct {
	X, Y int
}

func (p Point) Up() Point {
	p.X--
	return p
}

func (p Point) Down() Point {
	p.X++
	return p
}

func (p Point) Left() Point {
	p.Y--
	return p
}

func (p Point) Right() Point {
	p.Y++
	return p
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func solve(items []Item) int {
	points := make([]Point, 0)

	x, y := 0, 0
	perimeter := 0
	for _, item := range items {
		switch item.Direction {
		case 'R':
			y += item.Count
		case 'D':
			x += item.Count
		case 'L':
			y -= item.Count
		case 'U':
			x -= item.Count
		}
		point := Point{x, y}
		points = append(points, point)
		perimeter += item.Count
	}

	// calculate the area of the polygon
	// https://en.m.wikipedia.org/wiki/Shoelace_formula
	area := 0
	for i := range points {
		j := (i + 1) % len(points)
		p0 := points[i]
		p1 := points[j]
		area += p0.X * p1.Y
		area -= p0.Y * p1.X
	}
	area = int(math.Abs(float64(area / 2)))
	fmt.Println("perimeter=", perimeter)
	fmt.Println("     area=", area)

	// formula to calculat the area of poligon in terms of points
	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	//
	// in our case area = inner + perimiter/2 - 1
	inner := area - perimeter/2 + 1
	return perimeter + inner
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %v", err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	t0 := time.Now()
	sum := solve(part1(lines))
	fmt.Println("part1=", sum, "elapsed:", time.Since(t0))

	t0 = time.Now()
	sum = solve(part2(lines))
	fmt.Println("part2=", sum, "elapsed:", time.Since(t0))
}
