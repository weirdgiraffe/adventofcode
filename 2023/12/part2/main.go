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

// ???.### 1,1,3
// ???.###????.###????.###????.###????.### 1,1,3,1,1,3,1,1,3,1,1,3,1,1,3
func part2record(line string) (conditions string, checksum []int) {
	l := strings.Split(line, " ")
	damaged := ""
	for i := 0; i < 5; i++ {
		if i != 0 {
			conditions += "?"
			damaged += ","
		}
		conditions += l[0]
		damaged += l[1]
	}
	for _, s := range strings.Split(damaged, ",") {
		n, _ := strconv.Atoi(s)
		checksum = append(checksum, n)
	}
	return conditions, checksum
}

func arrangements(conditions string, checksum []int) int {
	return countArrangements(conditions, 0, checksum)
}

type Cache struct {
	m map[string]int
}

func (c *Cache) Set(line string, broken int, rest []int, v int) {
	if len(line) < 3 {
		return
	}
	key := fmt.Sprintf("%s|%d|%v", line, broken, rest)
	c.m[key] = v
}

func (c *Cache) Get(line string, broken int, rest []int) int {
	key := fmt.Sprintf("%s|%d|%v", line, broken, rest)
	if v, ok := c.m[key]; ok {
		return v
	}
	return -1
}

var cache = &Cache{m: make(map[string]int)}

func countArrangements(line string, broken int, rest []int) int {
	if v := cache.Get(line, broken, rest); v != -1 {
		return v
	}
	if len(line) == 0 {
		if len(rest) == 1 && rest[0] == broken {
			return 1
		}
		if len(rest) == 0 {
			return 1
		}
		return 0
	}
	c := line[0]
	if c == '#' {
		if len(rest) == 0 {
			cache.Set(line, broken, rest, 0)
			return 0
		}

		n := countArrangements(line[1:], broken+1, rest)
		cache.Set(line, broken, rest, n)
		return n
	}
	if c == '?' {
		n := countArrangements("#"+line[1:], broken, rest) +
			countArrangements("."+line[1:], broken, rest)
		cache.Set(line, broken, rest, n)
		return n
	}

	if len(rest) > 0 {
		if broken == rest[0] {
			broken = 0
			rest = rest[1:]
		}
		if broken != 0 {
			return 0
		}
	}
	n := countArrangements(line[1:], 0, rest)
	cache.Set(line, broken, rest, n)
	return n
}

func part1(data string) int {
	sum := 0
	for _, line := range strings.Split(data, "\n") {
		if line != "" {
			sum += arrangements(record(line))
		}
	}
	return sum
}

func part2(data string) int {
	sum := 0
	for _, line := range strings.Split(data, "\n") {
		if line != "" {
			sum += arrangements(part2record(line))
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
	sum := part2(string(data))
	fmt.Println("sum=", sum, "elapsed:", time.Since(t0))
}
