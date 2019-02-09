package main

import (
	"github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

const coldef = termbox.ColorDefault

func drawBox() {
	termbox.Clear(coldef, coldef)

	// インプット部分の作成
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

	// アウトプット部分の作成
	var outputCells [][]termbox.Cell
	for i, row := range outputRows {
		outputCells = append(outputCells, []termbox.Cell{})
		for _, c := range row {
			outputCells[i] = append(outputCells[i], termbox.Cell{
				Ch: c,
				Fg: attr,
				Bg: termbox.ColorDefault,
			})
		}
	}

	// インプット部分の書き込み
	i := 0
	for _, c := range cells {
		termbox.SetCell(0+i, 0, c.Ch, c.Fg, c.Bg)

		w := runewidth.RuneWidth(c.Ch)
		if w == 0 || w == 2 && runewidth.IsAmbiguousWidth(c.Ch) {
			w = 1
		}
		i += w
	}

	w := 0
	for i, row := range outputCells {
		for _, c := range row {
			termbox.SetCell(0+w, 1+i, c.Ch, c.Fg, c.Bg)
			width := runewidth.RuneWidth(c.Ch)
			if width == 0 || width == 2 && runewidth.IsAmbiguousWidth(c.Ch) {
				width = 1
			}
			w += width
		}
		w = 0
	}

	// カーソルの書き込み
	termbox.SetCursor(RuneLength(inputString), 0)

	termbox.Flush()
}

var inputString []rune
var outputRows []string

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
			case termbox.KeyEnter:
				addRow(inputString)
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

func addRow(runes []rune) {
	outputRows = append(outputRows, string(runes))
}
