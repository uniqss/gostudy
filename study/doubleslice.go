package main

import "fmt"

type Table struct {
	value       [][]interface{}
	rowCount    int
	columnCount int
}

func NewTable(columnCount int) *Table {
	ret := &Table{rowCount: 0, columnCount: columnCount}
	return ret
}

// row column   0 indexed
func (t *Table) SetValue(row int, column int, val interface{}) bool {
	if column >= t.columnCount {
		return false
	} else if row < t.rowCount {
		t.value[row][column] = valn
	} else if row == t.rowCount {
		t.AppendRow()
		t.value[row][column] = val
	} else {
		return false
	}
	return true
}

func (t *Table) AppendRow() {
	t.rowCount++
	var newRow []interface{}

	// append according to row type
	for i := 0; i < t.columnCount; i++ {
		if i%2 == 0 {
			newRow = append(newRow, 0)
		} else {
			newRow = append(newRow, "")
		}
	}

	t.value = append(t.value, newRow)
}

func main() {
	var oneSlice []int
	for i := 0; i < 4; i++ {
		oneSlice = append(oneSlice, i)
	}
	fmt.Println(oneSlice)

	columnCount := 4
	var value [1][4]int
	for i := 0; i < columnCount; i++ {
		value[0][i] = i
	}
	fmt.Println(value)

	t := NewTable(4)
	fmt.Println(t)
	t.SetValue(0, 1, 1234)
	fmt.Println(t)
	t.SetValue(0, 2, "asdf")
	fmt.Println(t)
	t.SetValue(0, 3, 333)
	fmt.Println(t)
	t.SetValue(0, 4, 333)
	fmt.Println(t)
}
