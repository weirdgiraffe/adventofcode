package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

type Point struct {
	X, Y int
}

func (p Point) String() string { return fmt.Sprintf("(%d,%d)", p.X, p.Y) }

type Cell struct {
	Point
	Weight int
}

type Board [][]int

func NewBoard(lines []string) Board {
	b := make([][]int, len(lines))
	for i, line := range lines {
		b[i] = make([]int, len(line))
		for j := range line {
			b[i][j] = int(line[j] - '0')
		}
	}
	return b
}

func (b Board) cell(p Point) *Cell {
	if p.X < 0 || p.X >= len(b) {
		return nil
	}
	if p.Y < 0 || p.Y >= len(b[0]) {
		return nil
	}
	return &Cell{
		Point:  p,
		Weight: b[p.X][p.Y],
	}
}

func (b Board) Up(p Point) *Cell {
	p.X--
	return b.cell(p)
}

func (b Board) Down(p Point) *Cell {
	p.X++
	return b.cell(p)
}

func (b Board) Left(p Point) *Cell {
	p.Y--
	return b.cell(p)
}

func (b Board) Right(p Point) *Cell {
	p.Y++
	return b.cell(p)
}

func (b Board) Move(p Point, d Direction) *Cell {
	switch d {
	case U:
		return b.Up(p)
	case D:
		return b.Down(p)
	case L:
		return b.Left(p)
	case R:
		return b.Right(p)
	default:
		panic("invalid direction")
	}
}

func debug(args ...any) {
	// fmt.Print(args...)
}

func debugln(args ...any) {
	// fmt.Println(args...)
}

func debugf(format string, args ...any) {
	// fmt.Printf(format, args...)
}

type Node struct {
	Point
	Direction
	Repeat int
}

func NewNode(p Point, d Direction) Node {
	return Node{p, d, 1}
}

func (n Node) String() string {
	return strings.Repeat(string(n.Direction), n.Repeat) + n.Point.String()
}

func (b Board) FindPath(src, dst Point, moveMin, moveMax int) int {
	type Item struct {
		Node
		Weight *int
	}
	queue := make([]Item, 0, len(b)*len(b[0]))

	sort := func() {
		sort.Slice(queue, func(i, j int) bool {
			return *queue[i].Weight < *queue[j].Weight
		})
	}

	visited := make(map[Node]*int, len(b)*len(b[0]))
	c := b.Right(src)
	node := NewNode(c.Point, R)
	queue = append(queue, Item{node, &c.Weight})
	visited[node] = &c.Weight
	c = b.Down(src)
	node = NewNode(c.Point, D)
	queue = append(queue, Item{node, &c.Weight})
	visited[node] = &c.Weight
	cost := math.MaxInt

	for len(queue) > 0 {
		item := queue[0]
		if len(queue) > 1 {
			copy(queue, queue[1:])
		}
		queue = queue[:len(queue)-1]
		debugf("pop %s weight: %d\n", item.Node, *item.Weight)

		if item.Point == dst {
			if *item.Weight < cost {
				cost = *item.Weight
				debugln("found path:", item.Node, "weight:", cost)
			}
			continue
		}

		needSort := false
		for _, m := range "<>^v" {
			d := Direction(m)
			if d.Oposite(item.Direction) {
				continue
			}

			i := 0
			if d == item.Direction {
				i = item.Repeat
			}

			point := item.Point
			weight := *item.Weight
			for ; i < moveMin-1; i++ {
				if c = b.Move(point, d); c != nil {
					weight += c.Weight
					point = c.Point
				} else {
					break
				}
			}
			if i < moveMin-1 {
				continue
			}

			if (d == item.Direction) && (i == moveMax) {
				continue
			}

			if c = b.Move(point, d); c != nil {
				weight += c.Weight
				point = c.Point
				i++
			} else {
				continue
			}

			n := NewNode(point, d)
			n.Repeat = i
			if w, ok := visited[n]; ok {
				if weight < *w {
					debugf("  updated %s weight %d->%d\n", n, *w, weight)
					*visited[n] = weight
					needSort = true
				}
				continue
			}
			queue = append(queue, Item{n, &weight})
			visited[n] = &weight
			debugf("  push %s weight %d\n", n, weight)
			needSort = true
		}
		if needSort {
			sort()
		}
	}
	return cost
}

type Direction byte

const (
	U Direction = '^'
	D Direction = 'v'
	L Direction = '<'
	R Direction = '>'
)

func (d Direction) Oposite(other Direction) bool {
	switch d {
	case U:
		return other == D
	case D:
		return other == U
	case L:
		return other == R
	case R:
		return other == L
	default:
		panic("invalid direction")
	}
}

func solve(lines []string) int {
	start := Point{0, 0}
	target := Point{len(lines) - 1, len(lines[0]) - 1}
	b := NewBoard(lines)
	return b.FindPath(start, target, 4, 10)
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %v", err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	t0 := time.Now()
	sum := solve(lines)
	fmt.Println("sum=", sum, "elapsed:", time.Since(t0))
}
