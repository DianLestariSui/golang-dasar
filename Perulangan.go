package main

import "fmt"

func main()  {
	for	i	:=	0;	i	<	5;	i++	{
		fmt.Println("Angka",	i)
	}
//Argumen	Hanya	Kondisi
	var	a	=	0
	for	a	<	5	{
		fmt.Println("Angka",	a)
		a++
	}
//	for	Tanpa Argumen
	var	b	=	0
	for	{
		fmt.Println("Angka",	b)
		b++
		if	b	==	5	{
			break
		}
	}
//	break		&	 	continue
	for	i	:=	1;	i	<=	10;	i++	{
		if	i	%	2	==	1	{
			continue
		}
		if	i	>	8	{
			break
		}
		fmt.Println("Angka",	i)
	}
	//Perulangan	Bersarang
	for	i	:=	0;	i	<	5;	i++	{
		for	j	:=	i;	j	<	5;	j++	{
			fmt.Print(j,	"	")
		}
		fmt.Println()
	}
//Pemanfaatan	Label	Dalam	Perulangan
/*
Tepat	sebelum	keyword	 	for		terluar,	terdapat	baris	kode	 	outerLoop:	.	Maksud	dari	kode
tersebut	adalah	disiapkan	sebuah	label	bernama	 	outerLoop		untuk	 	for		dibawahnya.	Nama
label	bisa	diganti	dengan	nama	lain	(dan	harus	diakhiri	dengan	tanda	titik	dua	atau	colon
( 	:	)	).
 */
	outerLoop:
	for	i	:=	0;	i	<	5;	i++	{
		for	j	:=	0;	j	<	5;	j++	{
			if	i	==	3	{
				break	outerLoop
			}
			fmt.Print("matriks	[",	i,	"][",	j,	"]",	"\n")
		}
	}
}