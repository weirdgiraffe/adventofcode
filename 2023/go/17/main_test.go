package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
 */

func TestFoo(t *testing.T) {
	input := []string{
		"2413432311323",
		"3215453535623",
		"3255245654254",
		"3446585845452",
		"4546657867536",
		"1438598798454",
		"4457876987766",
		"3637877979653",
		"4654967986887",
		"4564679986453",
		"1224686865563",
		"2546548887735",
		"4322674655533",
	}
	b := NewBoard(input)
	fmt.Println(b)
	fmt.Println()

	sum := b.FindPath(Point{0, 0}, Point{12, 12}, 1, 3)
	require.Equal(t, 102, sum)
}
