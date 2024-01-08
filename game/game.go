package game

import (
	"fmt"
)

type TickTackToe struct {
	Arr [3][3]string
}

func (t *TickTackToe) PrintGame() {
	fmt.Println("\n\n------ GAME -------")
	fmt.Println()
	colLabels := "abc"

	fmt.Print("  ")
	for _, col := range colLabels {
		fmt.Printf(" %c ", col)
	}

	fmt.Println()

	for i, row := range t.Arr {
		fmt.Printf("%d ", i+1)
		for _, val := range row {
			fmt.Printf(" %s ", func() string {
				if val == "" {
					return "-"
				} else {
					return val
				}
			}())
		}
		fmt.Println()
	}
	fmt.Println()
}

func (t *TickTackToe) AddPlay(char string, xPos int8, yPos string) error {
	yNum := func() int8 {
		switch yPos {
		case "a":
			return 0
		case "b":
			return 1
		case "c":
			return 2
		default:
			return -1
		}
	}()
	if t.Arr[xPos][yNum] != "" {
		return fmt.Errorf("[PositionInUse] the selected position (%s, %d) is already filled with %s", yPos, xPos, t.Arr[xPos][yNum])
	}
	t.Arr[xPos][yNum] = char
	return nil
}

func (t *TickTackToe) CheckGame() (bool, string, bool) {
	diagWon, diagChar := checkDiagonals(t)

	rowsWon, rowChar := checkRows(t)

	colsWon, colChar := checkCols(t)

	if rowsWon {
		return true, rowChar, false
	} else if diagWon {
		return true, diagChar, false
	} else if colsWon {
		return true, colChar, false
	} else {
		all := checkAll(t)
		if all {
			return false, "", true
		}
	}

	return false, "", false
}

func checkAll(t *TickTackToe) bool {
	all := true
	for i := 0; i < len(t.Arr); i++ {
		for j := 0; j < len(t.Arr); j++ {
			if t.Arr[i][j] != "x" && t.Arr[i][j] != "o" {
				all = false
			}
		}
	}

	return all
}

func checkDiagonals(t *TickTackToe) (bool, string) {
	valDiag := t.Arr[0][0]
	valAnti := t.Arr[0][len(t.Arr)-1]
	if (valDiag != "x" && valDiag != "o") || (valAnti != "x" && valAnti != "o") {
		return false, ""
	}

	var diag bool = true
	var anti bool = true
	for i := 1; i < len(t.Arr); i++ {
		if t.Arr[i][i] != valDiag {
			diag = false
		}
		if t.Arr[i][len(t.Arr)-1-i] != valAnti {
			anti = false
		}
	}

	if diag {
		return true, valDiag
	} else if anti {
		return true, valAnti
	}

	return false, ""
}

func checkRows(t *TickTackToe) (bool, string) {
	for _, row := range t.Arr {
		val := row[0]

		if val != "x" && val != "o" {
			return false, ""
		}

		var won bool = true
		for _, elem := range row[1:] {
			if elem != val {
				won = false
			}
		}
		if won {
			return true, val
		}
	}

	return false, ""
}

func checkCols(t *TickTackToe) (bool, string) {
	for i := 0; i < len(t.Arr); i++ {
		val := t.Arr[0][i]
		var won bool = true
		if val != "x" && val != "o" {
			return false, ""
		}

		for j := 0; j < len(t.Arr); j++ {
			elem := t.Arr[j][i]
			if elem != val {
				won = false
			}
		}

		if won {
			return true, val
		}
	}

	return false, ""
}
