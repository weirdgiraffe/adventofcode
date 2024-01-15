package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hash(s string) int {
	current := 0
	for _, c := range s {
		current += int(c)
		current *= 17
		current = current % 256
	}
	return current
}

func part1(seq string) int {
	sum := 0
	i, j := 0, 0
	for i = range seq {
		if seq[i] == ',' {
			sum += hash(seq[j:i])
			j = i + 1
		}
	}
	return sum + hash(seq[j:])
}

const (
	OpRemove = 1
	OpAdd    = 2
)

func labelByNumber(n int) string {
	return ""
}

func split(s string) (name string, op int, num int) {
	for i, c := range s {
		if c == '-' {
			return s[:i], OpRemove, 0
		}
		if c == '=' {
			n, err := strconv.Atoi(s[i+1:])
			if err != nil {
				panic("not a number")
			}
			return s[:i], OpAdd, n
		}
	}
	panic("incorrect input")
}

type Item struct {
	Label string
	Num   int
}

type Items struct {
	l []Item
}

func (c *Items) Add(label string, num int) {
	for i, item := range c.l {
		if item.Label == label {
			c.l[i].Num = num
			return
		}
	}
	c.l = append(c.l, Item{Label: label, Num: num})
}

func (c *Items) Del(label string) {
	for i, item := range c.l {
		if item.Label == label {
			c.l = append(c.l[:i], c.l[i+1:]...)
			return
		}
	}
}

type Boxes struct {
	m map[int]*Items
}

func newBoxes() *Boxes {
	m := make(map[int]*Items)
	for i := 0; i < 256; i++ {
		m[i] = &Items{}
	}
	return &Boxes{m: m}
}

func (b *Boxes) Get(hash int) *Items {
	return b.m[hash]
}

func (b *Boxes) Sum() int {
	sum := 0
	for i, v := range b.m {
		for j, item := range v.l {
			sum += (i + 1) * (j + 1) * item.Num
		}
	}
	return sum
}

func (b *Boxes) Print() {
	for k, v := range b.m {
		s := fmt.Sprintf("Box %d: ", k)
		for i, item := range v.l {
			s += fmt.Sprintf("[%s %d]", item.Label, item.Num)
			if i < len(v.l)-1 {
				s += " "
			}
		}
		if len(v.l) > 0 {
			fmt.Println(s)
		}
	}
}

func part2(seq string) int {
	m := newBoxes()
	do := func(s string) {
		label, op, num := split(s)
		boxNum := hash(label)
		switch op {
		case OpRemove:
			m.Get(boxNum).Del(label)
		case OpAdd:
			m.Get(boxNum).Add(label, num)
		}
		fmt.Printf("After %q\n", s)
		m.Print()
		fmt.Println()
	}
	i, j := 0, 0
	for i = range seq {
		if seq[i] == ',' {
			do(seq[j:i])
			j = i + 1
		}
	}
	do(seq[j:])
	m.Print()
	return m.Sum()
}

func main() {
	seq, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %v", err)
		os.Exit(1)
	}
	sum := part2(strings.TrimSpace(string(seq)))
	fmt.Println("sum=", sum)
}
