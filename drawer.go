package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Drawer struct {
	prompt string
}

type DrawerParameter struct {
	query string
	rows  []string
}

func NewDrawer(prompt string) *Drawer {
	d := &Drawer{
		prompt: prompt,
	}
	return d
}

func (d *Drawer) Draw(param *DrawerParameter) error {

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	err := d.drawQueryRow(param.query)
	if err != nil {
		return err
	}

	err = d.drawResultRows(param.rows)
	if err != nil {
		return err
	}

	termbox.Flush()

	return nil
}

func (d *Drawer) drawQueryRow(query string) error {

	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault

	str := d.prompt + query

	var cells []termbox.Cell

	for _, c := range str {
		cells = append(cells, termbox.Cell{
			Ch: c,
			Fg: color,
			Bg: backgroundColor,
		})
	}

	d.drawRow(0, 0, cells)
	return nil
}

func (d *Drawer) drawResultRows(rows []string) error {

	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault

	for i, row := range rows {
		var cells []termbox.Cell

		for _, c := range row {
			cells = append(cells, termbox.Cell{
				Ch: c,
				Fg: color,
				Bg: backgroundColor,
			})
		}

		d.drawRow(0, 1+i, cells)
	}
	return nil
}

func (d *Drawer) drawRow(x int, y int, cells []termbox.Cell) {
	i := 0
	for _, c := range cells {
		termbox.SetCell(x+i, y, c.Ch, c.Fg, c.Bg)

		w := runewidth.RuneWidth(c.Ch)
		if w == 0 || (w == 2 && runewidth.IsAmbiguousWidth(c.Ch)) {
			w = 1
		}

		i += w
	}
}
