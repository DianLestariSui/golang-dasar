package main

import "fmt"

func main()  {
	//tipe data non desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644
	//tipe data numerik desimal
	var decimalNumber = 2.62
	//tipe data boolean
	var exist bool =  true
	//variabel string menggunakan grave accent/backticks(`)
	var text = `ini adalah variabel string 
dengan menggunakan grave accent/backticks`

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
	// tamplate %f digunakan untuk mengformat data numerik desimal menjadi string
	//jumlah digit yang muncul bisa dikontrol menggunakan %.nf, tinggal ganti n dengan angka yang di inginkan
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
	//tipe data boolean
	fmt.Printf("exist? %t \n", exist)
	//string grave accent/backticks(`)
	fmt.Println(text)
}
