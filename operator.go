package main

import "fmt"

func main(){
	/*
	operator aritmatika, perbandingan, logika  sama dengan umumnya,
	memiliki makna dan fungsi yang sama
	 */
	var value = ((( 2 + 6) % 3) * 4 - 2 )/ 3
	var isEqual = (value == 2)
	var	left	=	false
	var	right	=	true
	fmt.Print("hasil nilai value ", value, isEqual)

	var	leftAndRight	=	left	&&	right
	fmt.Printf("left	&& right \t(%t) \n", leftAndRight)
	var	leftOrRight	=	left || right
	fmt.Printf("left	|| right \t(%t) \n", leftOrRight)
	var	leftReverse	=	!left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}
