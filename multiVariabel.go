package main

import "fmt"

func main()  {
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	var var1, var2, var3 string = "empat", "lima", "enam"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	//variabel dengan tipe data yang berbeda
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	//contoh variabel underscore
	name , _ :="dian","variabel dalam keranjang sampah"

	//DEKLARASI VARIABEL MENGGUNAKAN KEYWORD NEW
	nilai := new (string)

	//contoh multi variabel ke 1
	fmt.Printf("Multi variabel %s!\n ", first)
	fmt.Printf("Multi variabel %s!\n ", second)
	fmt.Printf("Multi variabel %s!\n ", third)
	//contoh multi variabel ke 2
	fmt.Printf("Multi variabel %s!\n ", var1)
	fmt.Printf("Multi variabel %s!\n ", var2)
	fmt.Printf("Multi variabel %s!\n ", var3)
	//contoh multi variabel ke 3
	fmt.Printf("Multi variabel %s!\n ", seventh)
	fmt.Printf("Multi variabel %s!\n ", eight)
	fmt.Printf("Multi variabel %s!\n ", ninth)
	//contoh multi variabel ke 4 dengan tipe data yang berbeda
	fmt.Printf("Multi variabel tipe data int %s!\n ", one)
	fmt.Printf("Multi variabel tioe data boolean %s!\n ", isFriday)
	fmt.Printf("Multi variabel tipe data doble %s!\n ", twoPointTwo)
	fmt.Printf("Multi variabel tipe data string %s!\n ", say)
	fmt.Printf("variabel underscore %s!\n ", name)
	//deklasari variabel menggunakan keyword new
	fmt.Println(nilai)
}
