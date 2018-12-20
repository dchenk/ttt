package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to the tic-tac-toe game!")

	// Start off by playing a first game.
	playing := true

	for playing {
		g := setUpGame()

		fmt.Printf("%s vs %s -- Let's go!\n", g.Name1, g.Name2)

		g.Play()

		for {
			// Keep prompting the user until they give a valid response.
			input := strings.ToLower(getString("Would you like to play another game? Enter y or n"))
			if input == "n" {
				playing = false
				break
			}
			if input == "y" {
				playing = true
				break
			}
		}
	}

}

// setUpGame prompts the user to enter the names of the people playing and returns a newly set up Game instance.
func setUpGame() *Game {
	var g Game

	// Get the player names.
	g.Name1 = getString("The name of the first person playing")
	g.Name2 = getString("The name of the second person playing")

	return &g
}

// getString prompts the user with the given prompt, which should no have a colon at the end,
// making sure the user entered a non-empty string.
func getString(prompt string) string {
	var str string
	for str == "" {
		fmt.Printf("%s: ", prompt)
		_, err := fmt.Scan(&str)
		if err != nil {
			errOnInput(err)
			continue
		}
		str = strings.TrimSpace(str)
	}
	return str
}

// errOnInput prints out an error message telling them they've input something invalid.
func errOnInput(err error) {
	fmt.Printf("An error occurred reading your input (%v).\n", err)
}

type Game struct {
	// Name1 and Name2 are the names of the players.
	// The first player uses Xs and the second uses Os.
	Name1, Name2 string
	b            Board
}

func (g *Game) Play() {
	g.PrintBoard()
	turn := uint8(1)
	for g.next() {
		g.getPlayerMove(turn)
		if turn == 1 {
			turn = 2
		} else {
			turn = 1
		}
		g.PrintBoard()
	}
	g.Results()
}

// next indicates whether the next player should be prompted to make their move.
func (g *Game) next() bool {
	return g.b.status() == -1
}

// getPlayerMove gets the player's number, which is either 1 or 2, and prompts them to make a move.
func (g *Game) getPlayerMove(player uint8) {
	playerName := g.Name1
	if player == 2 {
		playerName = g.Name2
	}
	fmt.Printf("It's %s's turn (%s).\n", playerName, Mark(player))

	// row and col start counting from 1.
	var row, col int
	var err error
	var indx int // the index in the underlying array
	for {
		fmt.Println("Please choose a blank cell to mark.")
		for {
			input := getString("Choose which row to place a mark in (between 1 and 3)")
			row, err = strconv.Atoi(input)
			if err == nil && 1 <= row && row <= 3 {
				break
			}
			fmt.Println("There's a problem with your input.")
		}
		for {
			input := getString("Choose which column to place a mark in (between 1 and 3)")
			col, err = strconv.Atoi(input)
			if err == nil && 1 <= col && col <= 3 {
				break
			}
			fmt.Println("There's a problem with your input.")
		}
		indx = (row-1)*3 + col - 1
		if g.b[indx] == MarkBlank {
			break
		}
	}
	g.b[indx] = Mark(player)
}

func (g *Game) PrintBoard() {
	fmt.Println(g.b.String())
}

// Results prints the results of the game, including the winner if there is one.
func (g *Game) Results() {
	const format = "///////\n%s won!\n///////\n"
	switch g.b.status() {
	case 0:
		fmt.Println("It's a draw!")
	case 1:
		fmt.Printf(format, g.Name1)
	case 2:
		fmt.Printf(format, g.Name2)
	}
}

// A Board contains three rows and three columns of cells.
type Board [9]Mark

// Print prints out the board with all marks.
func (b *Board) String() string {
	return fmt.Sprintf(
		"%s | %s | %s\n---------\n%s | %s | %s\n---------\n%s | %s | %s",
		b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8],
	)
}

// status returns:
//   -1 if the game is still going on
//   0 if the game ended in a draw
//   1 if player 1 won
//   2 if player 2 won
func (b *Board) status() int {
	// If there are three in a row anywhere, we're done.
	// Check for completion for each player.
	for player := 1; player <= 2; player++ {
		for cell := 0; cell <= 6; cell++ {
			// If the cell is blank or is marked by the other player, go on to the next cell.
			playerMark := Mark(player)
			if b[cell] == MarkBlank || b[cell] != playerMark {
				continue
			}

			// Check across horizontally.
			switch cell {
			case 0, 3, 6:
				if b[cell+1] == playerMark && b[cell+2] == playerMark {
					return player
				}
			}

			// Check down vertically.
			switch cell {
			case 0, 1, 2:
				if b[cell+3] == playerMark && b[cell+6] == playerMark {
					return player
				}
			}

			// Check diagonally.
			switch cell {
			case 0:
				if b[4] == playerMark && b[8] == playerMark {
					return player
				}
			case 2:
				if b[4] == playerMark && b[6] == playerMark {
					return player
				}
			}
		}
	}

	// If any cells are blank, the game is still going on.
	for _, v := range b {
		if v == MarkBlank {
			return -1
		}
	}

	// All the cells are filled, it's a draw.
	return 0
}

// A Mark indicates the status of a cell.
type Mark uint8

const (
	MarkBlank Mark = 0
	MarkX     Mark = 1
	MarkO     Mark = 2
)

// String implements fmt.Stringer for Mark.
func (m Mark) String() string {
	switch m {
	case MarkX:
		return "X"
	case MarkO:
		return "O"
	default:
		return " "
	}
}
