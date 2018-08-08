package main

import "fmt"

func main(){
	const firstName string = "john"
	const lastName  = "wick"

	fmt.Print("halo ", firstName, "!\n")
	fmt.Print("nice to meet you", lastName, "!\n")

	/*
	fungsi fmt.Print() sama dengan fungsi fmt.Println() pembedanya
	fungsi fmt.print() tidak memiliki garis baru dan perlu ditambahkan spasi
	karena fungsi ini tidak menambahkan spasi di sela-sela nilai parameter yang digabungkan
	*/
}
