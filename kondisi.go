package main

import "fmt"

func main(){
	/*
	dalam masalah kondisi masih sama pada umumnya
	dengan operator yang sebelumnya dibahas.
	if , else if , & else
	 */
	var	point	=	10
	if	point	==	10	{
		fmt.Println("lulus	dengan	nilai	sempurna")
	}	else	if	point	>	5	{
		fmt.Println("lulus")
	}	else	if	point	==	4	{
		fmt.Println("hampir lulus")
	}	else	{
		fmt.Printf("tidak lulus.	nilai anda %d\n",	point)
	}

	/*
	variabel temporary pada if-else
	 */
	var	pointvar	=	8840.0
	if	percent	:=	pointvar	/	100;	percent	>=	100	{
		fmt.Printf("%.1f%s	perfect!\n",	percent,	"%")
	}	else	if	percent	>=	70	{
		fmt.Printf("%.1f%s	good\n",	percent,	"%")
	}	else	{
		fmt.Printf("%.1f%s	not	bad\n",	percent,	"%")
	}
	//	switch - Case
	var	pointsc	=	6
	switch	pointsc	{
	case	8:
		fmt.Println("perfect")
	case	7:
		fmt.Println("awesome")
	default:
		fmt.Println("not	bad")
	}
	/*
	Pemanfaatan case Untuk Banyak Kondisi
	Sebuah	 	case		dapat	menampung	banyak	kondisi.	Cara	penerapannya	yaitu	dengan
	menuliskan	nilai	pembanding-pembanding	variabel	yang	di-switch	setelah	keyword	 	case
	dipisah	tanda	koma	( 	,	).
	 */
	var	pointcase	=	6
	switch	pointcase	{
	case	8:
		fmt.Println("perfect")
	case	7,	6,	5,	4:
		fmt.Println("awesome")
	default:
		fmt.Println("not	bad")
	}
	/*
	Kurung	Kurawal	Pada	Keyword	 	case		&
	default
	Tanda	kurung	kurawal	( 	{	}	)	bisa	diterapkan	pada	keyword	 	case		dan	 	default	.	Tanda	ini
	opsional,	boleh	dipakai	boleh	tidak.	Bagus	jika	dipakai	pada	blok	kondisi	yang	didalamnya
	ada	banyak	statement,	kode	akan	terlihat	lebih	rapi	dan	mudah	di-maintain.
	 */
	var	pointdefault	=	6
	switch	pointdefault	{
	case	8:
		fmt.Println("perfect")
	case	7,	6,	5,	4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not	bad")
			fmt.Println("you	can	be	better!")
		}
	}

	/*
	Switch	Dengan	Gaya	 	if		-	 	else
	Uniknya	di	Golang,	switch	bisa	digunakan	dengan	gaya	ala	if-else.	Nilai	yang	akan
	dibandingkan	tidak	dituliskan	setelah	keyword	 	switch	,	melainkan	akan	ditulis	langsung
	dalam	bentuk	perbandingan	dalam	keyword	 	case	.
	 */
	var	pointswirch	=	6
	switch	{
	case	pointswirch	==	8:
		fmt.Println("perfect")
	case	(point	<	8)	&&	(point	>	3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not	bad")
			fmt.Println("you	need	to	learn	more")
		}
	}
	/*
	Penggunaan	Keyword	 	fallthrough	Dalam switch
	Seperti	yang	sudah	dijelaskan	sebelumnya,	bahwa	switch	pada	Golang	memiliki	perbedaan
	dengan	bahasa	lain.	Ketika	sebuah	 	case		terpenuhi,	pengecekkan	kondisi	tidak	akan
	diteruskan	ke	case-case	setelahnya.
	Keyword	 	fallthrough		digunakan	untuk	memaksa	proses	pengecekkan	diteruskan	ke	 	case
	selanjutnya.
	 */
	var	pointfallthrough	=	6
	switch	{
	case	pointfallthrough	==	8:
		fmt.Println("perfect")
	case	(pointfallthrough	<	8)	&&	(pointfallthrough	>	3):
		fmt.Println("awesome")
		fallthrough
	case	pointfallthrough	<	5:
		fmt.Println("you	need	to	learn	more")
	default:
		{
			fmt.Println("not	bad")
			fmt.Println("you	need	to	learn	more")
		}
	}
	/*
	Seleksi	Kondisi	Bersarang
	Seleksi	kondisi	bersarang	adalah	seleksi	kondisi,	yang	berada	dalam	seleksi	kondisi,	yang
	mungkin	juga	berada	dalam	seleksi	kondisi,	dan	seterusnya.	Seleksi	kondisi	bersarang	bisa
	dilakukan	pada	 	if		-	 	else	,	 	switch	,	ataupun	kombinasi	keduanya.
	 */
	var	pointbersarang	=	10
	if	pointbersarang	>	7	{
		switch	pointbersarang	{
		case	10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	}	else	{
		if	pointbersarang	==	5	{
			fmt.Println("not	bad")
		}	else	if	pointbersarang	==	3	{
			fmt.Println("keep	trying")
		}	else	{
			fmt.Println("you	can	do	it")
			if	pointbersarang	==	0	{
				fmt.Println("try	harder!")
			}
		}
	}
}
