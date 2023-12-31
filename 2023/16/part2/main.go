package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

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

func pass(l, o [][]byte, i, j int, direction byte) {
	stack := &Stack{}
	for {
		if (i < 0) || (j < 0) || (i == len(l)) || (j == len(l[0])) {
			item := stack.Pop()
			if item == nil {
				return
			}
			i, j, direction = item.i, item.j, item.direction
		}

		if o[i][j]&byte(direction) != 0 {
			i = -1
			j = -1
			continue
		}

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

func variant(l, o [][]byte, i, j int, direction byte) int {
	pass(l, o, i, j, direction)
	return count(o)
}

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

func solve(lines []string) int {
	l := make([][]byte, len(lines))
	for i, line := range lines {
		l[i] = []byte(line)
	}
	pool := sync.Pool{
		New: func() any {
			o := make([][]byte, len(lines))
			for i, line := range lines {
				o[i] = bytes.Repeat([]byte{0}, len(line))
			}
			return &o
		},
	}

	var wg sync.WaitGroup
	out := make(chan int)
	do := func(i, j int, direction byte) {
		wg.Add(1)
		go func() {
			v := pool.Get()
			o := *v.(*[][]byte)
			out <- variant(l, o, i, j, direction)

			for i := range o {
				for j := range o[i] {
					o[i][j] = 0
				}
			}
			pool.Put(v)
			wg.Done()
		}()
	}

	cols := len(lines[0])
	rows := len(lines)
	for i := 0; i < rows; i++ {
		do(i, 0, R)
		do(i, cols-1, L)
	}
	for i := 0; i < cols; i++ {
		do(0, i, D)
		do(rows-1, i, U)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var result int
	for v := range out {
		if v > result {
			result = v
		}
	}

	return result

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
