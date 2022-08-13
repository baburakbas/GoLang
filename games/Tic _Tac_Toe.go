//Welcome to TicTacToe

package main

import (
	"errors"
	"fmt"
)

func main() {
	var playWith int

	var player int = 1
	var CPU int = 2
	var mapXox [3][3]string = [3][3]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}} //Our map for playground
	fmt.Print("\n\n")
	fmt.Println("        ******************************")
	fmt.Println("        Welcome to the Tic Tac Toe !! ")

	fmt.Println("        ******************************")
	fmt.Print("\n\n")
	fmt.Println("Do you want play with player or CPU , \n***press 1 for player\n***press 2 for CPU \n\n\n ")

	fmt.Scan(&playWith) //Choose opponent to play

	if playWith == CPU {
		fmt.Println("****this option is still in development. You can try player mode.\n****if you want to be part of the team ,contact baburakbas4646@gmail.com")
	}
	if playWith == player {

		fmt.Println("Press where you want to mark")
		fmt.Println("__1___|___2_____|___3__ ")
		fmt.Println("__4___|___5_____|___6__ ")
		fmt.Printf("__7___|___8_____|___9__ \n\n\n")

		// after that part game algortm is begin
		gameXox(mapXox)

	}
}

func gameXox(mapXox [3][3]string) (massage string, err error) {

	var i, j, k int
	var move string
	var typeOfMove string
	massage = "Game Over , wp!!"
	for k = 0; k < 9; k++ { //The game can be up to 9 rounds
		//If the elements are in the same row or column, the game is over. Can also be diagonal
		//Lines
		if mapXox[0][0] == mapXox[0][1] && mapXox[0][1] == mapXox[0][2] {
			fmt.Print("\n\nGame Over , wp!!")
			break
		}
		if mapXox[1][0] == mapXox[1][1] && mapXox[1][1] == mapXox[1][2] {
			fmt.Print("\n\nGame Over , wp!!")
			break
		}
		if mapXox[2][0] == mapXox[2][1] && mapXox[2][1] == mapXox[2][2] {
			fmt.Print("\n\nGame Over , wp!!")
			break
		}
		//Columns
		if mapXox[0][0] == mapXox[1][0] && mapXox[1][0] == mapXox[2][0] {
			fmt.Print("\n\nGame Over , wp!!")
			break
		}
		if mapXox[0][1] == mapXox[1][1] && mapXox[1][1] == mapXox[2][1] {
			fmt.Print("\n\nGame Over , wp!!")
			break
		}
		if mapXox[0][2] == mapXox[1][2] && mapXox[1][2] == mapXox[2][2] {
			fmt.Print("\n\nGame Over , wp!!")
			break
		}
		//Diogonal
		if mapXox[0][0] == mapXox[1][1] && mapXox[1][1] == mapXox[2][2] {
			fmt.Print("\n\nGame Over , wp!!")
			break
		}
		if mapXox[0][2] == mapXox[1][1] && mapXox[1][1] == mapXox[2][0] {
			fmt.Print("\n\nGame Over , wp!!")
			break
		}
		if k%2 != 1 {
			typeOfMove = "X"
		} else {
			typeOfMove = "O"
		}
		fmt.Print("\n\n")
		fmt.Scan(&move)

		if move != "1" && move != "2" && move != "3" && move != "4" && move != "5" && move != "6" && move != "7" && move != "8" && move != "9" { //Player must choose from map
			err = errors.New("Input must be one of the maps number!!")
			fmt.Print(err)
			break

		}

		if k == 8 { //If there played 8 rounds, that means no winners for this game
			fmt.Print("\n\n")
			fmt.Print("It's a draw wp!")
		}
		fmt.Print("\n")
		for i = 0; i < 3; i++ {

			fmt.Print("\n")
			// drawing map and move
			for j = 0; j < 3; j++ {
				if move != mapXox[i][j] {
					fmt.Print(" | ")
					fmt.Print(mapXox[i][j])
				} else {
					fmt.Printf(" | %s", typeOfMove)
					mapXox[i][j] = typeOfMove

				}

			}

		}

	}
	return massage, err
}
