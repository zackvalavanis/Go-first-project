package main

import (
	"fmt"

	"goFirstFile/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
