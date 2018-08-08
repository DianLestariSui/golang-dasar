package main

import (
	"fmt"
)

func main()  {
	var firstName string = "john"
	var lastName string
	lastName="wick"
	tanpaTipeData :="deklarasi variable tanpa tipe data"

	fmt.Printf("halo	%s	%s!\n",	firstName,	lastName)
	fmt.Printf("halo	john	wick!\n")
	fmt.Printf("halo	%s	%s!\n",	firstName,	lastName)
	fmt.Printf("%s!\n",tanpaTipeData)
	fmt.Println("halo",	firstName,	lastName	+	"!")
}