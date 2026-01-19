package maps

import (
	"fmt"
	"maps"
	"slices"
)

func GoMaps() {
	m := make(map[string]int)
	m["a1"] = 9
	m["a2"] = 10

	fmt.Println("Map : ", m )
	// mengambil value dari sebuah map menggunakan key
	v1 := m["a1"]
    fmt.Println("v1:", v1)

	// jika mengambil key yang tidak ada. akan otomatis o
	v3 := m["k3"]
    fmt.Println("v3:", v3)
	//manggil panjang map
	fmt.Println("len", len(m))
	// value dari map bisa diolah
	m["a3"] = m["a2"] + m["a1"]
	delete(m,"a2")
	fmt.Println("maps setelah ditambah A3 dan A2 dihapus : ",m)
	// clear untuk menghapus seluruh isi dari map
	clear(m)
	fmt.Println("map setelah clear(m): ",m)
	// mengecek apakah m["a2"] ada 
	_, prs := m["a3"]
    fmt.Println("prs:", prs)
	//inisialisasi map dalam satu baris 

	n:=map[string]int{"a" : 0, "b" :5}
	fmt.Println(n)
	n2 := map[string]int{"a" : 888, "b" :344}
	// func copy yang hanya bisa digunakan apabila nama key nya sama 
	maps.Copy(n2,n)
	fmt.Println(n)
	fmt.Println(n2)

	// menggunakan fungsi equals
	if maps.Equal(n,n2) {
		fmt.Println("n == n2")
	}
	// fungsi maps.keys untul mengambil key nya saja
	keys := slices.Sorted(maps.Keys(n))
	fmt.Println(keys)
	n["c"] = 8
	// fungsi maps.DeleteFunc menghapus isi dari data map dengan kondisi
	maps.DeleteFunc(n, func(k string, v int) bool{
		//menghapus data ganjil
		return v%2 != 0
	})
	fmt.Println(n)
}