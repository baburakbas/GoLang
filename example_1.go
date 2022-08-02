//This program calculates your electricity bill.
package main

import "fmt"

func main() {

	var a, b, total float32
	b = dishwashingPrice(5, 31, 7)
	a = loundaryPrice(3, 19, 5, false)

	total = a + b
	fmt.Println("harcanan elektrik = ", total)
	fmt.Println(b, a)

}

func dishwashingPrice(duration float32, watt float32, time int) float32 {
	totalPrice := watt * duration
	if time <= 6 && time >= 0 {                                            
		totalPrice = totalPrice - (totalPrice*10)/100
		return totalPrice
	} else {
		return totalPrice
	}
}

func loundaryPrice(duration float32, watt float32, time float32, supermode bool) float32 {
	totalPrice := watt * duration

	if supermode {
		if time <= 6 && time >= 0 {
			totalPrice = totalPrice - (totalPrice*10)/100 + 50

			return totalPrice
		} else {
			totalPrice += 50
			return totalPrice
		}
	} else {
		if time <= 6 && time >= 0 {
			totalPrice = totalPrice - (totalPrice*10)/100
			return totalPrice

		} else {

			return totalPrice
		}
	}

}
