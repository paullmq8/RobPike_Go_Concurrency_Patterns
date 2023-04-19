package main

func main() {
	//var c chan int
	//c = make(chan int)
	// or
	c := make(chan int)
	//c <- 1
	go func() {
		value := <-c
		println(value)
	}()
	c <- 1
	println("End")
}
