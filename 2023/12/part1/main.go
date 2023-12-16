package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Operational = '.'
	Broken      = '#'
	Unknown     = '?'
)

func record(line string) (conditions string, checksum []int) {
	l := strings.Split(line, " ")
	for _, s := range strings.Split(l[1], ",") {
		n, _ := strconv.Atoi(s)
		checksum = append(checksum, n)
	}
	return l[0], checksum
}

type Seq struct {
	b []byte
}

func (s *Seq) Clear() {
	s.b = s.b[:0]
}

func (s *Seq) Add(c byte) {
	s.b = append(s.b, c)
}

func (s Seq) Empty() bool {
	return len(s.b) == 0
}

func arrangements(line string) int {
	l, cl := record(line)
	return count("", 0, l, cl)
}

const Ident = "|   "

func count(prefix string, broken int, line string, rest []int) int {
	if len(line) == 0 {
		if len(rest) == 1 && rest[0] == broken {
			return 1
		}
		if len(rest) == 0 {
			return 1
		}
		return 0 // if line is empty then there are no arrangements
	}

	fmt.Printf("%sL %s B:%d R:%v", prefix, line, broken, rest)

	n := 0
	switch line[0] {
	case Broken:
		fmt.Printf(" -> #\n")
		n += count(prefix+Ident, broken+1, line[1:], rest)
	case Unknown:
		fmt.Printf(" -> ?[#]\n")
		n += count(prefix+Ident, broken+1, line[1:], rest)

		fmt.Printf("%sL %s B:%d R:%v -> ?[.]", prefix, line, broken, rest)
		if len(rest) > 0 {
			if broken == rest[0] {
				fmt.Printf(" POP: %d", broken)
				broken = 0
				rest = rest[1:]
			}
		}
		fmt.Println()
		if broken == 0 {
			n += count(prefix+Ident, 0, line[1:], rest)
		}
	case Operational:
		if len(rest) > 0 {
			if broken == rest[0] {
				fmt.Printf(" POP: %d", broken)
				broken = 0
				rest = rest[1:]
			}
		}
		fmt.Println()
		if broken == 0 {
			n += count(prefix+Ident, 0, line[1:], rest)
		}
	}
	fmt.Printf("%sEXIT: %d\n", prefix, n)
	return n
}

func part1(data string) int {
	sum := 0
	for _, line := range strings.Split(data, "\n") {
		sum += arrangements(line)
	}
	return 0
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic("failed to read input.txt")
	}
	sum := part1(string(data))
	fmt.Println("sum=", sum)
}
