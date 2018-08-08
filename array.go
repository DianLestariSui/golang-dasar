package main

import "fmt"

func main()  {
	/*
	Variabel	 	names		dideklarasikan	sebagai	 	array	string		dengan	alokasi	elemen	 	4		slot.	Cara
	mengisi	slot	elemen	array	bisa	dilihat	di	kode	di	atas,	yaitu	dengan	langsung	mengakses
	elemen	menggunakan	indeks,	lalu	mengisinya.
	 */
	var	names	[4]string
	names[0]	=	"trafalgar"
	names[1]	=	"d"
	names[2]	=	"water"
	names[3]	=	"law"
	fmt.Println(names[0],	names[1],	names[2],	names[3])

}
