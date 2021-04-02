package game

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

// Colors
const (
	textColor = termbox.ColorWhite
	snakeColor = termbox.ColorGreen
	foodColor = termbox.ColorBlue
	backgroundColor = termbox.ColorBlack
	wallColor = termbox.ColorRed
)

// Help menu
var (
	menu = []string{
		"", "", "", "",
		"Esc: to exit",
		"Tab: to change the level",
		"> <: to speed up and speed up",
		"← ↑ → ↓: Move snake",
		"a w d s: Move snake",
	}
)

/*
	Refresh the screen.
*/
func Render(game *Game) {
	termbox.Clear(textColor, backgroundColor)
	menu[0] = fmt.Sprintf("Level  = %d/%d", game.Level, game.DataBase.MaxLevel)
	menu[1] = fmt.Sprintf("Score  = %d", game.Property.Score)
	menu[2] = fmt.Sprintf("Lost   = %d", game.Property.Lost)
	menu[3] = fmt.Sprintf("Speed  = %d", Speed)
	showMenu()
	for x, cells := range game.Board {
		for y, cell := range cells {
			if cell == Wall {
				termbox.SetCell(y, x, rune(Wall), wallColor, wallColor)
			} else if cell == Snake {
				termbox.SetCell(y, x, rune(Snake), snakeColor, backgroundColor)
				game.Property.HeadX, game.Property.HeadY = x, y
			} else if cell == Tail {
				termbox.SetCell(y, x, rune(Tail), snakeColor, backgroundColor)
			} else if cell == Food {
				termbox.SetCell(y, x, rune(Food), foodColor, backgroundColor)
				CurrentFoodX, CurrentFoodY = x, y
			}
		}
	}
	termbox.Flush()
}

/*
	Show the menu.
*/
func showMenu() {
	x := (MaxHeight / 2) - (len(menu) / 2)
	y := MaxWidth + 2
	for _, expr := range menu {
		expr = fmt.Sprintf(" %s\n", expr)
		for _, char := range expr {
			termbox.SetCell(y, x, char, textColor, backgroundColor)
			y++
		}
		x++
		y = MaxWidth + 2
	}
}
