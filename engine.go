package main

import termbox "github.com/nsf/termbox-go"

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
			query: e.query,
			rows:  e.queryResult,
		}
		err = e.drawer.Draw(dp)
		if err != nil {
			panic(err)
		}

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			// case 0:
			// 	e.addChar(ev.Ch)
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
