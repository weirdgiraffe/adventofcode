package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func count(lines [][]byte) int {
	n := 0
	for _, line := range lines {
		for _, c := range line {
			if c != 0 {
				n++
			}
		}
	}
	return n
}

const (
	U = 1
	D = 2
	L = 4
	R = 8

	UR   = U + R         // 9
	UL   = U + L         // 5
	DR   = D + R         // 10
	DL   = D + L         // 6
	URD  = U + R + D     // 11
	ULD  = U + L + D     // 7
	URLD = U + R + L + D // 15
)

type Item struct {
	i, j      int
	direction byte
}

type Stack struct {
	l []*Item
}

func (s *Stack) Push(i, j int, d byte) {
	item := &Item{i, j, d}
	s.l = append(s.l, item)
}
func (s *Stack) Pop() *Item {
	if len(s.l) == 0 {
		return nil
	}
	item := s.l[len(s.l)-1]
	s.l = s.l[:len(s.l)-1]
	return item
}

var stack = &Stack{}
var m = make(map[Item]struct{})
var n = 10000000

func pass(l, o [][]byte, i, j int, direction byte) {
	for {
		if n == 0 {
			return
		}
		n--

		if (i < 0) || (j < 0) || (i == len(l)) || (j == len(l[0])) {
			item := stack.Pop()
			if item == nil {
				return
			}
			fmt.Println("n=", n)
			debug(o)
			i, j, direction = item.i, item.j, item.direction
			fmt.Println("POP: i=", i, "j=", j, "d=", direction)
		}

		fmt.Println("GO: i=", i, "j=", j, "d=", direction)
		if o[i][j]&byte(direction) != 0 {
			i = -1
			j = -1
			continue
		}
		fmt.Printf("[%d,%d]: %c %c\n", i, j, direction, l[i][j])

		o[i][j] |= byte(direction)

		switch l[i][j] {
		case '.':
			switch direction {
			case R:
				j++
			case L:
				j--
			case D:
				i++
			case U:
				i--
			}
		case '|':
			switch direction {
			case U:
				i--
			case D:
				i++
			case R, L:
				if i > 0 {
					stack.Push(i-1, j, U)
				}
				i++
				direction = D
			}
		case '-':
			switch direction {
			case U, D:
				if j > 0 {
					stack.Push(i, j-1, L)
				}
				j++
				direction = R
			case R:
				j++
			case L:
				j--
			}
		case '/':
			switch direction {
			case U:
				j++
				direction = R
			case D:
				j--
				direction = L
			case R:
				i--
				direction = U
			case L:
				i++
				direction = D
			}
		case '\\':
			switch direction {
			case U:
				j--
				direction = L
			case D:
				j++
				direction = R
			case R:
				i++
				direction = D
			case L:
				i--
				direction = U
			}
		}
	}
}

func debug(o [][]byte) {
	for i := range o {
		for _, c := range o[i] {
			s := ""
			switch c {
			case U:
				s = "U"
			case R:
				s = "R"
			case L:
				s = "L"
			case D:
				s = "L"
			case UR:
				s = "UR"
			case UL:
				s = "UL"
			case DR:
				s = "DR"
			case DL:
				s = "DL"
			case URD:
				s = "URD"
			case ULD:
				s = "ULD"
			case URLD:
				s = "URLD"
			}
			fmt.Printf("[%4s]", s)
		}
		fmt.Println()
	}
}

func final(o [][]byte) {
	for i := range o {
		for _, c := range o[i] {
			if c != 0 {
				c = '#'
			} else {
				c = '.'
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func solve(lines []string) int {
	l := make([][]byte, len(lines))
	o := make([][]byte, len(lines))
	for i, line := range lines {
		l[i] = []byte(line)
		o[i] = bytes.Repeat([]byte{0}, len(line))
	}
	pass(l, o, 0, 0, R)
	final(o)
	return count(o)
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %v", err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	sum := solve(lines)
	fmt.Println("sum=", sum)
}
