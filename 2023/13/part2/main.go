package main

import (
	"bufio"
	"fmt"
	"os"
)

func transpose(lines []string) []string {
	if len(lines) == 0 {
		return []string{}
	}
	out := make([]string, len(lines[0]))
	b := make([]byte, len(lines))
	for i := 0; i < len(out); i++ {
		for j, line := range lines {
			b[j] = line[i]
		}
		out[i] = string(b)
	}
	return out

}

func smudged(s1, s2 string) bool {
	n := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			if n > 0 {
				return false
			}
			n = i + 1
		}
	}
	return n > 0
}

func reflectionAt(l []string) int {
	// spew.Dump(l)
	i, j := 0, len(l)-1
	ok := false
	for ; i < j; i, j = i+1, j-1 {
		if l[i] != l[j] {
			if !smudged(l[i], l[j]) || ok {
				return 0
			}
			ok = true
		}
	}

	if !ok {
		return 0
	}
	return i // as indexes are 1 based
}

func front(l []string) int {
	d := len(l) % 2
	n := len(l) - d
	for i := n; i >= 0; i -= 2 {
		pos := reflectionAt(l[:i])
		if pos > 0 {
			return pos
		}
	}
	return 0
}

func back(l []string) int {
	d := (len(l) % 2)
	n := len(l)
	for i := 0 + d; i < n; i += 2 {
		pos := reflectionAt(l[i:n])
		if pos > 0 {
			return i + pos
		}
	}
	return 0
}

func vertical(pattern []string) int {
	l := transpose(pattern)
	if len(l) < 2 {
		// need at least two lines for a reflection
		return 0
	}
	fmt.Println("vertical")

	if pos := front(l); pos > 0 {
		return pos
	}

	if pos := back(l); pos > 0 {
		return pos
	}

	return 0
}

func horizontal(pattern []string) int {
	l := pattern
	if len(l) < 2 {
		// need at least two lines for a reflection
		return 0
	}
	fmt.Println("horizontal")

	if pos := front(l); pos > 0 {
		return pos * 100
	}
	if pos := back(l); pos > 0 {
		return pos * 100
	}
	return 0
}

func reflections(pattern []string) int {
	if n := horizontal(pattern); n > 0 {
		return n
	}
	return vertical(pattern)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	pattern := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			n := reflections(pattern)
			if n == 0 {
				fmt.Println("no reflection")
			}
			sum += n
			pattern = pattern[:0]
		} else {
			pattern = append(pattern, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read line: %v", err)
		os.Exit(1)
	}
	if len(pattern) > 0 {
		sum += reflections(pattern)
	}
	fmt.Println("sum=", sum)
}
