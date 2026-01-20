package collections

import "fmt"

func RunRange() {

	nums := []int{2, 3, 4}
	sum := 0
	// penggunan range pada slice
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum : ", sum)
	// jangan lupa range mengembalikan 2 nilai yaitu index dan value
	// karena itu seringkali penggunaan range memakai "_,a = range nums"
	//_, adalah cara untuk membuang index jika tidak dibutuhkan
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//penggunaan range dalam perulangan data map
	// yang akan menghasil kan key dan value
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key : ", k)
	}
	// range juga bisa dipakai untuk string yang akan mengembalikan data yang berisi
	// index dan ascii character dengan panjang yang sama dengan jumlah char string
	for i, j := range "go" {
		fmt.Println(i, j)
	}

}
