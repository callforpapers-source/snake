package game

import (
	"sync"
	"time"
)

type Keys int

const (
	Up Keys = iota
	Down
	Left
	Right
)

type Shift struct {
	up bool
	down bool
	left  bool
	right bool
}

var (
	Wall = int('#')
	Snake = int('x')
	Tail = int('o')
	Food = int('*')
	Floor = int(' ')
	MaxWidth = 0
	MaxHeight = 0
	CurrentFoodX = -1
	CurrentFoodY = -1
	mutex = &sync.Mutex{}
	Speed = 80
	minSpeed = 160
	maxSpeed = 80
)


type GameProperty struct {
	shift Shift
	snake [][]int
	HeadX int
	HeadY int
	Score int
	Lost  int
}

type Game struct {
	Level int
	Board [][]int
	DataBase *levels
	Property GameProperty
}

/*
	Initialize the game.
*/
func NewGame(level int) *Game {
	game := &Game{
		Level: level,
		DataBase: &levels{},
		Property: GameProperty{},
	}
	game.DataBase.LoadLevels()
	// board is 2D slice that store the current level
	game.Board = game.DataBase.GetLevel(game.Level)
	return game
}

/*
	Going to the next level.
*/
func (game *Game) Next() *Game {
	game.Level++
	if game.Level > game.DataBase.MaxLevel {
		game.Level = 1
	}
	game.Board = game.DataBase.GetLevel(game.Level)
	game.Property = GameProperty{}
	return game
}

/*
	Turn off moves.
*/
func (game *Game) shiftReset() {
	game.Property.shift = Shift{
		up: false,
		down: false,
		left: false,
		right: false,
	}
}

/*
	Manage the snake's speed.
*/
func (game *Game) SetSpeed(i int) {
	if i == -1 {
		if Speed < minSpeed {
			Speed += 40
		}
	} else {
		if Speed > maxSpeed {
			Speed -= 40
		}
	}
}

/*
	Which direction is on.
*/
func (game *Game) rotate(i int, value, set bool) bool {
	switch i {
	case 1:
		if set {
			game.Property.shift.up = value
		}
		return game.Property.shift.up
	case 2:
		if set {
			game.Property.shift.down = value
		}
		return game.Property.shift.down
	case 3:
		if set {
			game.Property.shift.left = value
		}
		return game.Property.shift.left
	case 4:
		if set {
			game.Property.shift.right = value
		}
		return game.Property.shift.right
	}
	return false
}

/*
	Manage the snake position. Eat, Move, Failure.
*/
func (game *Game) checkMove(x, y, shift int) {
	if game.rotate(shift, false, false) {
		return
	}
	game.shiftReset()
	game.rotate(shift, true, true)
	firstLap := 0
	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		for game.rotate(shift, false, false) {
			// Check the falure
			if game.gameOver(x, y) {
				firstLap = 1
				return
			}
			// Check the food and appling changes by eating food
			game.hasEaten()
			// Change the snake position
			game.snakePosition(x, y)
			firstLap = 1
			// Refresh the screen by new changes
			Render(game)
			time.Sleep(time.Duration(Speed) * time.Millisecond)
		}
	}()
	// Wait until the end of first lap; could be done by channels
	for firstLap == 0 {
	}
}

/*
	Manage moves.
*/
func (game *Game) Move(dir Keys) {
	switch dir {
	case Up:
		if !game.Property.shift.down {
			game.checkMove(-1, 0, 1)
		}
	case Down:
		if !game.Property.shift.up {
			game.checkMove(1, 0, 2)
		}
	case Left:
		if !game.Property.shift.right {
			game.checkMove(0, -1, 3)
		}
	case Right:
		if !game.Property.shift.left {
			game.checkMove(0, 1, 4)
		}
	}
}

/*
	Check the position of the head with the food.
*/
func (game *Game) hasEaten() {
	if CurrentFoodX == game.Property.HeadX && CurrentFoodY == game.Property.HeadY {
		game.initSnake()
		game.Board = InitFood(game.Board)
		game.Property.Score += 5
	}
}

/*
	Circuits.
*/
func (game *Game) restruct(x, y int) (int, int) {
	tempX, tempY := game.Property.HeadX, game.Property.HeadY
	if (tempY+y) == MaxWidth {
		tempY = -1
	} else if tempY+y == -1 {
		tempY = MaxWidth
	}

	if (tempX+x) == MaxHeight {
		tempX = -1
	} else if tempX+x == -1 {
		tempX = MaxHeight
	}
	return tempX, tempY
}

/*
	Move the snake and its tail.
*/
func (game *Game) snakePosition(x, y int) {
	tempX, tempY := game.restruct(x, y)
	prevX, prevY := game.Property.HeadX, game.Property.HeadY
	game.Board[game.Property.HeadX][game.Property.HeadY] = Floor
	tempX = tempX + x
	tempY = tempY + y
	game.Board[tempX][tempY] = Snake
	game.Property.HeadX, game.Property.HeadY = tempX, tempY
	for tailItem, cell := range game.Property.snake {
		game.Board[cell[0]][cell[1]] = Floor
		game.Board[prevX][prevY] = Tail
		game.Property.snake[tailItem] = []int{prevX, prevY}
		prevX, prevY = cell[0], cell[1]
	}
}

/*
	If the snake goes through the wall or eats its tail.
*/
func (game *Game) gameOver(x, y int) bool {
	tempX, tempY := game.restruct(x, y)
	head := game.Board[tempX+x][tempY+y]
	if head == Wall || head == Tail {
		// Preparing a new game
		game.Board = game.DataBase.GetLevel(game.Level)
		game.Property.Lost++
		game.Property.Score = 0
		game.Property.snake = [][]int{}
		game.shiftReset()
		Render(game)
		return true
	}
	return false
}

/*
	Add a new tail to the snake.
*/
func (game *Game) initSnake() {
	var prev, nprev []int
	var x, y int
	level := len(game.Property.snake)
	if level == 0 {
		prev = []int{game.Property.HeadX, game.Property.HeadY}
		n1, n2 := game.dimension()
		nprev = []int{game.Property.HeadX + n2, game.Property.HeadY + n1}
	} else {
		if level >= 2 {
			prev = game.Property.snake[level-2]
			nprev = game.Property.snake[level-1]
		} else {
			prev = game.Property.snake[level-1]
			nprev = []int{game.Property.HeadX, game.Property.HeadY}
		}
	}
	if prev[0] == nprev[0] {
		x = prev[0]
		if prev[1] < nprev[1] {
			y = prev[1]
		} else {
			y = nprev[1]
		}
	} else {
		y = prev[1]
		if prev[0] < nprev[0] {
			x = prev[0]
		} else {
			x = nprev[0]
		}
	}
	game.Board[x][y] = Tail
	game.Property.snake = append(game.Property.snake, []int{x, y})
}

/*
	Get the direction of the snake.
*/
func (game *Game) dimension() (x, y int) {
	switch true {
	case game.Property.shift.up:
		return 0, -1
	case game.Property.shift.down:
		return 0, 1
	case game.Property.shift.left:
		return -1, 0
	case game.Property.shift.right:
		return 1, 0
	}
	return 0, 0
}
