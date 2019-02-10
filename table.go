package main

import "strings"

type TableManager struct {
	current [][]string
	origin  [][]string
}

func NewTableManager(rows [][]string) *TableManager {
	tm := TableManager{current: rows, origin: rows}
	return &tm
}

func (tm *TableManager) getOriginRows() []string {
	var result []string
	for _, row := range tm.origin {
		result = append(result, strings.Join(row, ","))
	}
	return result
}
