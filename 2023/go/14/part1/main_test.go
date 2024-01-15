package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVertical(t *testing.T) {
	pattern := []string{
		"O....#....",
		"O.OO#....#",
		".....##...",
		"OO.#O....O",
		".O.....O#.",
		"O.#..O.#.#",
		"..O..#O..O",
		".......O..",
		"#....###..",
		"#OO..#....",
	}
	n := solve(pattern)
	require.Equal(t, 136, n)
}
