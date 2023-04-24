package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

// Each speaker must wait for a go-ahead to guarantee the sequencing.
func main() {

}

/*func boring(c <-chan ) {
	//c <- Message{ fmt.Sprintf("%s: %d", msg, i), waitForIt }

}*/

func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	waitForIt := make(chan bool) // shared between all messages.
	go func() { // We launch the goroutine from inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}
