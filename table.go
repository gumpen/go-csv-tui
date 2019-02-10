package main

type TableManager struct {
	current [][]string
	origin  [][]string
}

func NewTableManager(rows [][]string) *TableManager {
	tm := TableManager{current: rows, origin: rows}
	return &tm
}
