package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	seq := "HASH"
	n := part1(seq)
	require.Equal(t, 52, n)
}

func TestPart1(t *testing.T) {
	seq := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
	n := part1(seq)
	require.Equal(t, 1320, n)
}

func TestPart2(t *testing.T) {
	seq := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
	n := part2(seq)
	require.Equal(t, 145, n)
}
