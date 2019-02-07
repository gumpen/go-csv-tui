package main

import (
	"github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

const coldef = termbox.ColorDefault

func drawBox() {
	termbox.Clear(coldef, coldef)

	var attr termbox.Attribute
	var cells []termbox.Cell
	for _, c := range inputString {
		attr = termbox.ColorDefault
		cells = append(cells, termbox.Cell{
			Ch: c,
			Fg: attr,
			Bg: termbox.ColorDefault,
		})
	}

	i := 0
	for _, c := range cells {
		termbox.SetCell(0+i, 0, c.Ch, c.Fg, c.Bg)

		w := runewidth.RuneWidth(c.Ch)
		if w == 0 || w == 2 && runewidth.IsAmbiguousWidth(c.Ch) {
			w = 1
		}
		i += w
	}

	termbox.SetCursor(RuneLength(inputString), 0)

	termbox.Flush()
}

var inputString []rune

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	drawBox()
MAINLOOP:
	for {
		drawBox()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case 0:
				inputChar(ev.Ch)
			case termbox.KeyEsc:
				break MAINLOOP
			}
		default:
			drawBox()
		}
	}
}

func inputChar(ch rune) {
	inputString = append(inputString, ch)
}

func RuneLength(str []rune) int {
	var l int
	for _, c := range str {
		w := runewidth.RuneWidth(c)
		if w == 0 || w == 2 && runewidth.IsAmbiguousWidth(c) {
			w = 1
		}
		l += w
	}
	return l
}
