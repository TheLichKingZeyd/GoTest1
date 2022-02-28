package main

import (
	"fmt"

	"rsc.io/quote"
)

func Listquotes() {

	fmt.Println("Hello()\a", quote.Hello())

	fmt.Println("Go()\a", quote.Go())

	fmt.Println("Glass()\a", quote.Glass())

	fmt.Println("Opt()\a", quote.Opt())

}
