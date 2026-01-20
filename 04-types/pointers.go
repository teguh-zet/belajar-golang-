package types
import "fmt"

func GoPointer(){
	i := 1
    fmt.Println("initial:", i)
	//zero tidak mengubah nilai i karena pertukaran valuenya hanya berlaku 
	// di func local 
    zeroval(i)
    fmt.Println("zeroval:", i)
	// jika func yang dipanggil memiliki argumen pointer
	// maka argument yang di kirim harus menggunakan alamat dari variable cth &i
	zeroptr(&i)
    fmt.Println("zeroptr:", i)
    

    fmt.Println("pointer:", &i)

}

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}
