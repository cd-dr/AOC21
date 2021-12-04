package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func isBoardFull(board *[5][5]bool) bool {
	// Check fully marked rows
	for j := 0; j < 5; j++ {
		marked := true
		for k := 0; k < 5; k++ {
			if board[j][k] != true {
				marked = false
				break
			}
		}
		if marked == true {
			return true
		}
	}
	// Check fully marked columns
	for j := 0; j < 5; j++ {
		marked := true
		for k := 0; k < 5; k++ {
			if board[k][j] != true {
				marked = false
				break
			}
		}
		if marked == true {
			return true
		}
	}
	return false
}

func main() {
	var inpSeq []string
	var boards [][5][]string
	nb, i, nblanks := 0, 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			inpSeq = strings.Split(line, ",")
		} else {
			if len(line) == 0 {
				nblanks += 1
			} else {
				bid := (i - 1 - nblanks) / 5
				nr := i - 1 - nblanks - bid*5
				if nb <= bid {
					var tmp [5][]string
					boards = append(boards, tmp)
					nb += 1
				}
				boards[bid][nr] = strings.Fields(line)
			}
		}
		i += 1
	}

	var marked [][5][5]bool
	for i := 0; i < nb; i++ {
		var tmp [5][5]bool
		marked = append(marked, tmp)
	}
	markedBoard, markedInp := -1, int64(-1)
	for _, x := range inpSeq {
		// fmt.Println(x)
		for i := 0; i < nb; i++ {
			for j := 0; j < 5; j++ {
				for k := 0; k < 5; k++ {
					if boards[i][j][k] == x {
						marked[i][j][k] = true
					}
				}
			}
			if isBoardFull(&marked[i]) == true {
				markedBoard = i
				tmp, _ := strconv.ParseInt(x, 10, 64)
				markedInp = tmp
				break
			}
		}
		if markedBoard != -1 {
			break
		}
	}

	unmarkedSum := int64(0)
	for j := 0; j < 5; j++ {
		for k := 0; k < 5; k++ {
			if marked[markedBoard][j][k] == false {
				tmp, _ := strconv.ParseInt(boards[markedBoard][j][k], 10, 64)
				unmarkedSum += tmp
			}
		}
	}

	fmt.Println(unmarkedSum * markedInp)
}
