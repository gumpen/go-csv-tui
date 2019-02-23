package main

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

const (
	// Prompt string
	Prompt string = "> "
)

// Engine manage the whole
type Engine struct {
	query       string
	queryResult []string
	drawer      *Drawer
	table       *TableManager
	cursorX     int
}

// EngineParameter is parameter required for NewEngine()
type EngineParameter struct {
	rows [][]string
}

type EngineResult struct {
	err     error
	content string
}

// NewEngine initialize Engine struct
func NewEngine(param *EngineParameter) (*Engine, error) {

	e := &Engine{
		query:       "",
		queryResult: []string{"", ""},
		drawer:      NewDrawer(Prompt),
		table:       NewTableManager(param.rows),
		cursorX:     0,
	}
	return e, nil
}

// Run
func (e *Engine) Run() *EngineResult {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	e.queryResult = e.table.getOriginRows()
	// MAINLOOP:
	for {

		// クエリの文法チェック
		// クエリの実行→queryResult

		dp := &DrawerParameter{
			query:   e.query,
			rows:    e.queryResult,
			cursorX: e.cursorX,
		}
		err = e.drawer.Draw(dp)
		if err != nil {
			panic(err)
		}

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case 0:
				e.addCharToQuery(ev.Ch)
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				e.deleteCharFromQuery()
			case termbox.KeyArrowLeft:
				e.moveCursorLeft()
			case termbox.KeyArrowRight:
				e.moveCursorRight()
			case termbox.KeyEsc:
				return &EngineResult{content: "esc!!!\n"}
			}
		case termbox.EventError:
			panic(err)
		default:
		}
	}
	return &EngineResult{}
}

func (e *Engine) addCharToQuery(ch rune) {
	// 文字の挿入はカーソルのインデックス位置に合わせないといけない
	queryRune := []rune(e.query)
	e.query = string(append(queryRune, ch))

	e.cursorX = e.cursorX + runeWidth(ch)

	// TODO: string, []rune, []termbox.Cellのindexをよしなにmappingする必要がある
	// マルチバイト文字対応のため
	fmt.Println("")
	fmt.Printf("len(e.query):%#v\n", len(e.query))
	fmt.Printf("len([]rune(e.query)):%#v\n", len([]rune(e.query)))
	fmt.Printf("cellLen([]rune(e.query)):%#v\n", cellLen([]rune(e.query)))
	fmt.Printf("e.cursorX:%#v\n", e.cursorX)
}

func (e *Engine) deleteCharFromQuery() {

}

func (e *Engine) moveCursorLeft() {
	if e.cursorX > 0 {
		e.cursorX--
	}
}

func (e *Engine) moveCursorRight() {
	if cellLen([]rune(e.query)) > e.cursorX {
		e.cursorX++
	}
}
