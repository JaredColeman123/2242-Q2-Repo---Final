package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go Doctor1(&wg)
	go Doctor2(&wg)

	wg.Wait()
}

func Doctor1(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Thank you for waiting, the Doctor will see you now.")
}

func Doctor2(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Sorry to keep you waiting, please come in.")
}
