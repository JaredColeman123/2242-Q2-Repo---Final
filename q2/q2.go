package main

import "fmt"

func main() {
	message := make(chan string)

	go messenger(message)

	news := <-message
	fmt.Println(news)
}

func messenger(message chan string) {
	message <- "You have an unread message!"

	close(message)
}
