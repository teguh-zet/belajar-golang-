package types

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	//dibawah adalah contoh embedding struct
	base
	str string
}

func RunEmbedding() {
	//contoh deklarasinya
	// co := container{
	// 			base : base{
	// 				num: 19
	// 			}
	// 		, str: "teguh"}  //<-- atau bisa dalam bentu seperti ini
	co := container{base: base{num: 19}, str: "teguh"}
	// pemanggilan langsung dari container
	fmt.Printf("container= {num : %v, str : %v}\n", co.num, co.str)
	// melalui fied base
	fmt.Printf("melalui field base : %v \n", co.base.num)
	//fun describe
	fmt.Println(co.base.describe())
	fmt.Println(co.describe())

	// impelentasi embedding struct dengan interface
	type describer interface {
		describe() string
	}

	//menyalin isi dari container
	var d describer = co
	fmt.Println("describer : ", d.describe())
}
