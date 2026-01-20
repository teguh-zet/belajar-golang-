package errorhandling // Samakan package dengan file lain di folder ini

import (
    "fmt"
    "os"
)

// Ubah nama func main menjadi Public Function
func RunPanicExample() {
    fmt.Println("--- Memulai Contoh Panic ---")

    // CONTOH 1: Panic Manual
    // Kita memaksa program berhenti di sini.
    // Kode di bawahnya TIDAK AKAN PERNAH DIJALANKAN.
    // (Untuk melihat efek kode di bawah, baris ini harus dikomentari dulu)
    // panic("a problem") 

    // --- KODE DI BAWAH INI UNREACHABLE (Tidak Terjangkau) ---
    // Kecuali panic di atas dihapus/dikomentari.

    _, err := os.Create("/tmp/file")
    if err != nil {
        // CONTOH 2: Panic karena Error
        // Ini cara cepat (fail fast) menangani error jika kita malas handle errornya
        // atau jika errornya sangat fatal sehingga program tidak layak lanjut.
		recover()
		panic(err)
    }
    
    fmt.Println("Kode ini tidak akan tercetak karena sudah panic duluan.")
}