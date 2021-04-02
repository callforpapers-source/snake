package game

import (
	"bufio"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

const gameDataFile = "snake_levels.txt"

type levels struct {
	// Number of levels
	MaxLevel int
	// Levels
	data [][]string
}

/*
	Initialize a new food randomly.
*/
func InitFood(board [][]int) [][]int {
	rand.Seed(time.Now().UnixNano())
	for {
		randX, randY := rand.Intn(MaxHeight), rand.Intn(MaxWidth)
		if board[randX][randY] == Floor {
			board[randX][randY] = Food
			CurrentFoodX = randX
			CurrentFoodY = randY
			break
		}
	}
	return board
}

/*
	Convert the string to a 2D slice.
*/
func (DataBase *levels) GetLevel(l int) [][]int {
	level := DataBase.data[l-1]
	board := make([][]int, len(level))
	MaxWidth = len(level[0])
	MaxHeight = len(level)
	for i, str := range level {
		board[i] = make([]int, len(str))
		for j := 0; j < len(str); j++ {
			var c int
			switch int(str[j]) {
			case Wall:
				c = Wall
			case Floor:
				c = Floor
			case Snake:
				c = Snake
			}
			board[i][j] = c
		}
	}
	board = InitFood(board)
	return board
}

/*
	Read levels from snake_levels.txt.
*/
func (DataBase *levels) LoadLevels() {
	file, err := os.Open(gameDataFile)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	level := 0
	matrix := make([]string, 0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				DataBase.data = append(DataBase.data, matrix)
				DataBase.MaxLevel = level + 1
				return
			}
			panic(err)
		}
		line = strings.TrimRight(line, "\t\n\f\r")
		if len(line) != 0 {
			matrix = append(matrix, line)
		} else {
			DataBase.data = append(DataBase.data, matrix)
			matrix = make([]string, 0)
			level++
		}
	}
}
