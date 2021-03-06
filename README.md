# Snake Game

This is a simple terminal snake game written in Golang.

## How it works?
It used termbox for Event Management and plotting 2D slices.
First of all, it reads levels.txt and imports the levels. Levels separated with a newline.'#' means wall, 'x' means the position of the snake, and ' ' means floor.
It then converts the selected level into a 2D slice. Next, it fills each cell of the terminal screen with one of the array cells. 
The page refreshes with each event.

## Timeline
It took four days
* <b>_First day:_</b> A trip into the termbox and other frameworks and examples and implement the main.go page. Read the Idiomatic Go guide. (2h:45m)
* <b>_Second day:_</b> Implement a basic version without any other options. (3h)
* <b>_Third day:_</b> Add other options and finish the project. (1h:30m)
* <b>_Forth day:_</b> Packaging, minor changes, fix bugs, test the exceptions, remove spaces, comments, writing the README file, and upload. (1h)

## Install

### From GitHub
```
git clone https://github.com/callforpapers-source/snake
cd snake
make run
```

## Commands & Usage

Keybinding                              | Description
----------------------------------------|---------------------------------------
<kbd>Tab</kbd>                          | Next level
<kbd>\<</kbd>       					| Speed down
<kbd>\></kbd>                           | Speed up
<kbd>← ↑ → ↓ OR a w e s</kbd>           | Move
<kbd>Esc</kbd>                          | Exit
