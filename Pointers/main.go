package main

import ("fmt")

func main() {

	a := 42.2
	b := 44
	fmt.Printf("%f -> %x\n", a, &a )
	fmt.Printf("%d -> %x\n", b, &b )
}
