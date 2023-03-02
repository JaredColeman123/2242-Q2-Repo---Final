package main

import (
	"fmt"
	"time"
)

func main() {
	go goingToBelizeCity()
	go goingToSanPedro()

	time.Sleep(time.Second * 6)
}

func goingToBelizeCity() {
	fmt.Println("On the way to Belize City")
	time.Sleep(time.Second * 3)
	fmt.Println("we have arrived in Belize City safely!")
}

func goingToSanPedro() {
	fmt.Println("On the way to San Pedro, La Isla Bonita")
	time.Sleep(time.Second * 5)
	fmt.Println("We arrived in San Pedro Safely, Enjoy!")
}
