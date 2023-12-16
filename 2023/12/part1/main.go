package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func record(line string) (conditions string, checksum []int) {
	l := strings.Split(line, " ")
	for _, s := range strings.Split(l[1], ",") {
		n, _ := strconv.Atoi(s)
		checksum = append(checksum, n)
	}
	return l[0], checksum
}

func arrangements(line string) int {
	l, cl := record(line)
	n := countArrangements(l, 0, cl, "")
	return n
}

func countArrangements(line string, broken int, rest []int, v string) int {
	if len(line) == 0 {
		if len(rest) == 1 && rest[0] == broken {
			// fmt.Println("OK:", v)
			return 1
		}
		if len(rest) == 0 {
			// fmt.Println("OK:", v)
			return 1
		}
		// fmt.Println("FAILED:", v)
		return 0
	}
	c := line[0]
	if c == '#' {
		if len(rest) == 0 {
			// fmt.Println("FAILED:", v)
			return 0
		}
		return countArrangements(line[1:], broken+1, rest, v+"#")
	}
	if c == '?' {
		return countArrangements("#"+line[1:], broken, rest, v) +
			countArrangements("."+line[1:], broken, rest, v)
	}

	if len(rest) > 0 {
		if broken == rest[0] {
			broken = 0
			rest = rest[1:]
		}
		if broken != 0 {
			// fmt.Println("FAILED:", v)
			return 0
		}
	}
	return countArrangements(line[1:], 0, rest, v+".")
}

func part1(data string) int {
	sum := 0
	for _, line := range strings.Split(data, "\n") {
		if line != "" {
			sum += arrangements(line)
		}
	}
	return sum
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic("failed to read input.txt")
	}
	t0 := time.Now()
	sum := part1(string(data))
	fmt.Println("sum=", sum, "elapsed:", time.Since(t0))
}
