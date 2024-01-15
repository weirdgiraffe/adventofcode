package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	input := []string{
		`.|...\....`,
		`|.-.\.....`,
		`.....|-...`,
		`........|.`,
		`..........`,
		`.........\`,
		`..../.\\..`,
		`.-.-/..|..`,
		`.|....-|.\`,
		`..//.|....`,
	}
	n := solve(input)
	require.Equal(t, 46, n)
}
