package main

import "fmt"

func main() {
	fmt.Println("Welcome")
	var name string
	fmt.Printf("What is your name")
	fmt.Scan(&name)
	fmt.Printf("Hello everyone iam  %s,this is my first go programm ", name)

}
