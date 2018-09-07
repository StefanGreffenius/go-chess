package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"string_utils"
)

var X_NAMES = [8]string{"A", "B", "C", "D", "E", "F", "G", "H"}
var Y_NAMES = [8]string{"1", "2", "3", "4", "5", "6", "7", "8"}
var chessboard = [][]string{
	[]string{"WT", "WL", "WP" ,"WD" ,"WK" ,"WP" ,"WL", "WT"},
	[]string{"WB", "WB", "WB", "WB", "WB", "WB", "WB", "WB"},
  []string{"##", "##", "##", "##", "##", "##", "##", "##"},
	[]string{"##", "##", "##", "##", "##", "##", "##", "##"},
	[]string{"##", "##", "##", "##", "##", "##", "##", "##"},
	[]string{"##", "##", "##", "##", "##", "##", "##", "##"},
	[]string{"SB", "SB", "SB", "SB", "SB", "SB", "SB", "SB"},
  []string{"ST", "SL", "SP", "SD", "SK", "SP", "SL", "ST"},
}

type Movement struct {
	fromX, fromY int
	toX, toY int
}

func FieldName(x int, y int) string {
	return string(X_NAMES[x] + Y_NAMES[y])
}

func ShowIntro() {
	fmt.Println("Welcome to CHESS GO")
	fmt.Println(string_utils.Reverse("A-T-E-B"))
	fmt.Println()
}

func ShowChessboard() {

	fmt.Println()
	fmt.Println("    ###########################")
	for x := 7; x >= 0; x-= 1 {
		fmt.Printf(X_NAMES[x] + "   | ")
		for y := 0; y <= 7 ; y+= 1 {
      fmt.Printf(chessboard[y][x] + " ")
		}
		fmt.Printf("|")
		fmt.Println()
  }
	fmt.Println("    ###########################")
	fmt.Println("\n      1  2  3  4  5  6  7  8")
	fmt.Println()
}

func GetMovement(turnnumber int) Movement {
	reader := bufio.NewReader(os.Stdin)

  if turnnumber%2 == 0 {
		fmt.Println("BLACK its your turn!")
	} else {
		fmt.Println("WHITE its your turn!")
	}

	fmt.Print("Move from: ")
  from, _ := reader.ReadString('\n')
	from = strings.Replace(from, "\n", "", -1)
	fmt.Print("Move to: ")
  to, _ := reader.ReadString('\n')
	to = strings.Replace(to, "\n", "", -1)

  movement := Movement{}
	for x := 7; x >= 0; x-= 1 {
		for y := 0; y <= 7 ; y+= 1 {
      if FieldName(x, y) == from {
				movement.fromX = y
				movement.fromY = x
			}

			if FieldName(x, y) == to {
				movement.toX = y
				movement.toY = x
			}
		}
	}

  return movement
}

func ExecuteMovement(movement Movement) {
	chessboard[movement.toX][movement.toY] = chessboard[movement.fromX][movement.fromY]
	chessboard[movement.fromX][movement.fromY] = "##"
}

func KingsAreAlive() bool {
	var kings int = 0

	for x := 0; x <= 7 ; x+= 1 {
		for y := 0; y <= 7 ; y+= 1 {
			if chessboard[x][y] == "SK" || chessboard[x][y] == "WK" {
        kings += 1
			}
		}
	}

  return kings == 2
}

func ShowWinner(turnnumber int) {
	fmt.Println()
	for x := 0; x <= 7 ; x+= 1 {
		for y := 0; y <= 7 ; y+= 1 {
			if chessboard[x][y] == "SK" {
        fmt.Print("\nBLACK HAS WON")
			}
			if chessboard[x][y] == "WK" {
        fmt.Print("\nWHITE HAS WON")
			}
		}
	}
	fmt.Print(" after ", turnnumber, " turns\n\n")
}

func main() {
  ShowIntro()
	ShowChessboard()

  turnnumber := 0
	for KingsAreAlive() {
		turnnumber += 1
		movement := GetMovement(turnnumber)
		ExecuteMovement(movement)
	  ShowChessboard()
	}

  ShowWinner(turnnumber)
}
