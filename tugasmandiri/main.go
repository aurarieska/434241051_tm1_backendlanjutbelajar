package main

import "fmt"

func main() {
	println("string dengan tipe data dan value ", akuStringTp)
	println("string tanpa tipe data dengan value ", akuString)
	println("string tanpa tipe data dan value ", stringKosong)
	println("int dengan tipe data dan value ", iniIntTp)
	println("int tanpa tipe data dengan value ", iniInt)
	println("int tanpa tipe data dan value ", intKosyong)
	println("float dengan tipe data dan value ", itufloatTP)
	println("float tanpa tipe data dengan value ", yofloat)
	println("float tanpa tipe data dan value ", floatKosong)
	println("boolean dengan tipe data dan value ", buyiyanTp)
	println("bolean tanpa tipe data dengan value ", buliyan)
	println("boolean tanpa tipe data dan value ", boolKosong)
	fmt.Println("slice dengan tipe data dan value ", slcliteral)
	fmt.Println("slice tanpa tipe data dengan value ", slicemake)
	fmt.Println("slice tanpa tipe data dan value ", slicekosong)

	a := 10
	b := 100

	println("sebelum swapbypointer", a, b)
	swapbypointer(&a,&b)
	println(a, b)

	//buat variabel baru karena yg atas sudah kena overwrite
	c := 10
	d := 100
	println("sebelum swapbyvalue", c, d)
	swapbyvalue(c,d)
	println(c,d)

	e := []string{"aura", "rieska"}
	fmt.Println("sebelum updateslide", e)
	updateslice(&e, "maharani")
	fmt.Println("sesudah updateslide", e)

	println(GetInfo())

	println("masukkan nilai baru")
	var nilai float64
	fmt.Scan(&nilai)
	UpdateGrade(nilai)
	println("status setelah melalui update nilai")
	println(GetInfo())

	println("status setelah melalui activate")
	activate()
	println(GetInfo())

	println("status setelah melalui deactive")
	deactive()
	println(GetInfo())
}