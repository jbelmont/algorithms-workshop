package table

import "fmt"

type Table struct {
	Rows   []Row
	Name   string
	CNames []string
}

type Row struct {
	Columns []Column
	ID      int
}

type Column struct {
	ID    int
	Value string
}

func (t *Table) Print() {
	rows := t.Rows
	fmt.Println(t.Name)

	for _, row := range rows {
		columns := row.Columns
		for i, column := range columns {
			fmt.Println(t.CNames[i], column.ID, column.Value)
		}
	}
}
