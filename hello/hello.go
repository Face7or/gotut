package main

import (
	"fmt"

	"github.com/Face7or/gotut/greetings"
)

func main() {
	//fmt.Printf("Hello, GO!")
	message := greetings.Hello("Gladys")
    fmt.Println(message)
}