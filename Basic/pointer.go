package main

import "fmt"

func main() {

	var p *int
	i := 42
	p = &i
	fmt.Printf("value of p: %d\n", *p)
	fmt.Printf("address of P:%d\n",&p)
	*p=2
}
