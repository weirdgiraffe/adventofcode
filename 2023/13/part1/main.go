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
func reflectionAt(l []string) int {
	// spew.Dump(l)
	i, j := 0, len(l)-1
	for ; i < j; i, j = i+1, j-1 {
		if l[i] != l[j] {
			return 0
		}
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
	if pos := front(l); pos > 0 {
		return pos * 100
	}
	if pos := back(l); pos > 0 {
		return pos * 100
	}
	return 0
}

func reflections(pattern []string) int {
	if n := vertical(pattern); n > 0 {
		return n
	}
	return horizontal(pattern)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	pattern := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			sum += reflections(pattern)
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
