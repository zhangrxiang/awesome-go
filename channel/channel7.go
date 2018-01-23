package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     uint8
	Address Addr
}

type Addr struct {
	city     string
	district string
}

func main() {
	persionChan := make(chan Person, 1)

	p1 := Person{"Harry", 32, Addr{"Shanxi", "Xian"}}
	fmt.Printf("P1 (1): %v\n", p1)

	persionChan <- p1

	p1.Address.district = "shijingshan"
	fmt.Printf("P2 (2): %v\n", p1)

	p1_copy := <-persionChan
	fmt.Printf("p1_copy: %v\n", p1_copy)
}
