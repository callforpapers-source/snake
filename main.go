package main

import (
	"github.com/callforpapers-source/snake/game"
	"github.com/nsf/termbox-go"
	"time"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	queue := make(chan termbox.Event)
	go func() {
		for {
			queue <- termbox.PollEvent()
		}
	}()

	control := game.NewGame(1)
	game.Render(control)
	for {
		ev := <-queue
		if ev.Type == termbox.EventKey {
			switch {
			case ev.Key == termbox.KeyArrowUp || ev.Ch == 'w':
				control.Move(game.Up)
			case ev.Key == termbox.KeyArrowDown || ev.Ch == 's':
				control.Move(game.Down)
			case ev.Key == termbox.KeyArrowLeft || ev.Ch == 'a':
				control.Move(game.Left)
			case ev.Key == termbox.KeyArrowRight || ev.Ch == 'd':
				control.Move(game.Right)
			case ev.Ch == ',':
				// Speed down
				control.SetSpeed(-1)
			case ev.Ch == '.':
				// Speed up
				control.SetSpeed(1)
			case ev.Key == termbox.KeyTab:
				// Next level
				control = control.Next()
			case ev.Key == termbox.KeyEsc:
				return
			}
		}
		time.Sleep(time.Millisecond * 25)
		game.Render(control)
	}
}
