package main

import (
	"strconv"
	"testing"
)

var testBoards = []Board{
	{
		// Blank
		MarkBlank, MarkBlank, MarkBlank,
		MarkBlank, MarkBlank, MarkBlank,
		MarkBlank, MarkBlank, MarkBlank,
	},
	{
		// In progress
		MarkX, MarkBlank, MarkBlank,
		MarkX, MarkO, MarkBlank,
		MarkBlank, MarkBlank, MarkO,
	},
	{
		// Horizontal
		MarkO, MarkO, MarkO,
		MarkBlank, MarkBlank, MarkBlank,
		MarkBlank, MarkBlank, MarkBlank,
	},
	{
		// Vertical
		MarkX, MarkO, MarkO,
		MarkX, MarkBlank, MarkBlank,
		MarkX, MarkBlank, MarkBlank,
	},
	{
		// Diagonal
		MarkO, MarkX, MarkO,
		MarkBlank, MarkO, MarkBlank,
		MarkBlank, MarkBlank, MarkO,
	},
	{
		// Draw
		MarkX, MarkO, MarkO,
		MarkO, MarkX, MarkX,
		MarkX, MarkX, MarkO,
	},
}

func TestGame_next(t *testing.T) {
	cases := []struct {
		Game
		expected bool
	}{
		{Game{b: testBoards[0]}, true},
		{Game{b: testBoards[1]}, true},
		{Game{b: testBoards[2]}, false},
		{Game{b: testBoards[3]}, false},
		{Game{b: testBoards[4]}, false},
		{Game{b: testBoards[5]}, false},
	}
	for i := range cases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			tc := &cases[i]
			got := tc.Game.next()
			if got != tc.expected {
				t.Errorf("expected %v but got %v", tc.expected, got)
			}
		})
	}
}

func TestBoard_status(t *testing.T) {
	cases := []struct {
		Board
		expected int
	}{
		{testBoards[0], -1},
		{testBoards[1], -1},
		{testBoards[2], 2},
		{testBoards[3], 1},
		{testBoards[4], 2},
		{testBoards[5], 0},
	}
	for i := range cases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			tc := &cases[i]
			got := tc.Board.status()
			if got != tc.expected {
				t.Errorf("expected %q but got %q", tc.expected, got)
			}
		})
	}
}

func TestBoard_Print(t *testing.T) {
	cases := []struct {
		Board
		expected string
	}{
		{
			Board:    testBoards[0],
			expected: "  |   |  \n---------\n  |   |  \n---------\n  |   |  ",
		},
		{
			Board:    testBoards[1],
			expected: "X |   |  \n---------\nX | O |  \n---------\n  |   | O",
		},
		{
			Board:    testBoards[2],
			expected: "O | O | O\n---------\n  |   |  \n---------\n  |   |  ",
		},
	}
	for i := range cases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			tc := &cases[i]
			got := tc.String()
			if got != tc.expected {
				t.Errorf("expected:\n%v\nbut got:\n%v", tc.expected, got)
			}
		})
	}
}

func TestMark_String(t *testing.T) {
	cases := []struct {
		Mark
		expected string
	}{
		{MarkBlank, " "},
		{MarkX, "X"},
		{MarkO, "O"},
		{234, " "},
	}
	for i := range cases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			tc := &cases[i]
			got := tc.Mark.String()
			if got != tc.expected {
				t.Errorf("expected %q but got %q", tc.expected, got)
			}
		})
	}
}
