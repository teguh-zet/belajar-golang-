package functions

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func GoClosureFunctions() {

    nextInt := intSeq()
	// nextInt disini sudah memegang nilai 1 yang diambil dari return i setiap kali dipanggil
    fmt.Println(nextInt())
	// nextInt disini sudah memegang nilai 2
    fmt.Println(nextInt())
	// nextInt disini sudah memegang nilai 3 
    fmt.Println(nextInt())
	// karena di inisialisasi ulang nilai nya berubah menjadi 1
    newInts := intSeq()
    fmt.Println(newInts())
}