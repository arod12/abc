package main

import "fmt"

func main() {
    fmt.Printf("Hello, World.\n")
	c := make(chan string)
	c <- "abc"
}

