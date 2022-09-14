package main

import "fmt"

func main() {

	fibonacchi := [2]int{0, 1}

	for i := 2; i < 15; i++ {

		toplam := fibonacchi[0] + fibonacchi[1]
		fmt.Println(toplam)
		fibonacchi[0] = fibonacchi[1]
		fibonacchi[1] = toplam
		fmt.Println("Bunlar hep yalan dolan")
	}

}
