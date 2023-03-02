package main

import "fmt"

type building interface {
	doors() int
	windows() int
	roof() string
}

type warehouse struct {
	location string
	color    string
}

func (w warehouse) doors() int {
	return 3
}

func (w warehouse) windows() int {
	return 10
}

func (w warehouse) roof() string {
	return "wood"
}

type house struct {
	houseNum int
	location string
	color    string
}

func (h house) doors() int {
	return 3
}

func (h house) windows() int {
	return 15
}

func (h house) roof() string {
	return "zinc"
}

func (h house) bedroom() int {
	return 4
}

func main() {
	wh := warehouse{
		location: "Belmopan City",
		color:    "Sky Blue",
	}

	hs := house{
		houseNum: 7,
		location: "Belize City, Belize",
		color:    "Red and Blue",
	}
	printInfo(wh)
	printInfo(hs)

}

func printInfo(b building) {
	fmt.Println("This building has", b.doors(), "doors, with", b.windows(), "windows, whilst the roof is made of", b.roof())
}
