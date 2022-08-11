package main

import "fmt"

func main() {
	var i, j, playWith, k int
	var move string
	var player int = 1
	var CPU int = 2
	var typeOfMove string
	var mapXox [3][3]string = [3][3]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	fmt.Println("******************************")
	fmt.Println("Welcome to the Tic Tac Toe !! ")

	fmt.Println("******************************")
	fmt.Println("Do you want play with player or CPU , \n***press 1 for player\n***press 2 for CPU ")

	fmt.Scan(&playWith)

	if playWith == CPU {
		fmt.Println("****this option is still in development. You can try player mode.\n****if you want to be part of the team ,contact baburakbas4646@gmail.com")
	}
	if playWith == player {

		fmt.Println("Press where you want to mark")
		fmt.Println("__1___|___2_____|___3__ ")
		fmt.Println("__4___|___5_____|___6__ ")
		fmt.Printf("__7___|___8_____|___9__ \n\n\n")

		// after that part game algortm is begin
		for k = 0; k < 9; k++ {
			//If the elements are in the same row or column, the game is over. Can also be diagonal.

			//Lines
			if mapXox[0][0] == mapXox[0][1] && mapXox[0][1] == mapXox[0][2] {
				fmt.Print("\n\nGame Over")
				break
			}
			if mapXox[1][0] == mapXox[1][1] && mapXox[1][1] == mapXox[1][2] {
				fmt.Print("\n\nGame Over")
				break
			}
			if mapXox[2][0] == mapXox[2][1] && mapXox[2][1] == mapXox[2][2] {
				fmt.Print("\n\nGame Over")
				break

			}
			//Columns
			if mapXox[0][0] == mapXox[1][0] && mapXox[1][0] == mapXox[2][0] {
				fmt.Print("\n\nGame Over")
				break
			}
			if mapXox[0][1] == mapXox[1][1] && mapXox[1][1] == mapXox[2][1] {
				fmt.Print("\n\nGame Over")
				break
			}
			if mapXox[0][2] == mapXox[1][2] && mapXox[1][2] == mapXox[2][2] {
				fmt.Print("\n\nGame Over")
				break
			}
			//Diogonal
			if mapXox[0][0] == mapXox[1][1] && mapXox[1][1] == mapXox[2][2] {
				fmt.Print("\n\nGame Over")
				break
			}
			if mapXox[0][2] == mapXox[1][1] && mapXox[1][1] == mapXox[2][0] {
				fmt.Print("\n\nGame Over")
				break
			}

			if k%2 != 1 {
				typeOfMove = "X"
			} else {
				typeOfMove = "O"
			}
			fmt.Print("\n\n")
			fmt.Scan(&move)

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
	}
}
