package main

import "fmt"

const (
	// Prompt string
	Prompt string = "> "
)

// Engine manage the whole
type Engine struct {
	query  string
	result []string
	drawer *Drawer
	table  *TableManager
}

// EngineParameter is parameter required for NewEngine()
type EngineParameter struct {
	rows [][]string
}

type EngineResult struct {
}

// NewEngine initialize Engine struct
func NewEngine(param *EngineParameter) (*Engine, error) {

	e := &Engine{
		query:  "",
		result: []string{"", ""},
		drawer: NewDrawer(Prompt),
		table:  NewTableManager(param.rows),
	}
	return e, nil
}

// Run
func (e *Engine) Run() int {
	fmt.Println("hoge")
	return 0
}
