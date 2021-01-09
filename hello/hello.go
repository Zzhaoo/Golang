package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Edison Zhao")
	fmt.Println(message)
}
