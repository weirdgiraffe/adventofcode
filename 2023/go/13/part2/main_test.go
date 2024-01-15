package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVertical(t *testing.T) {
	pattern := []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
	}
	n := reflections(pattern)
	require.Equal(t, 300, n)
}

func TestHorizontal(t *testing.T) {
	pattern := []string{
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}
	n := reflections(pattern)
	require.Equal(t, 100, n)
}
