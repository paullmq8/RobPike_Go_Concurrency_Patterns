package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("%s\n", "hello")    // prints hello
	fmt.Printf("%q\n", "hello")    // prints "hello"
	fmt.Printf("%s\n", "hello\n;") // prints hello   ; and \n is not escaped
	fmt.Printf("%q\n", "hello\n;") // prints "hello\n;" \n is escaped here
	fmt.Println("======================")
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		// Receive expression is just a value
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("Your're boring; I'm leaving.")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		// Expression to be sent can be any suitable value.
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
