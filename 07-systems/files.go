package system // Package harus SAMA

import (
	"bufio"
	"fmt"
	"os"
)

// Helper function untuk cek error biar kode lebih bersih
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RunFileExample() {
	fmt.Println("\n--- File I/O ---")

	filename := "dataku.txt"

	// --- 1. MENULIS FILE (Writing) ---

	// Cara Cepat: Menulis semua data sekaligus (otomatis create/overwrite)
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile(filename, d1, 0644) // 0644 adalah permission file
	check(err)

	// Cara Manual (Lebih Kontrol): Membuka file untuk ditulis
	f, err := os.Create("data2.txt")
	check(err)

	// Penting: Jadwalkan penutupan file
	defer f.Close()

	// Tulis string
	d2 := []byte{115, 111, 109, 101, 10} // bytes untuk "some\n"
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("Menulis %d bytes\n", n2)

	// Tulis string langsung (lebih mudah)
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("Menulis %d bytes\n", n3)

	// Pastikan data benar-benar tertulis dari memori ke harddisk
	f.Sync()

	// --- 2. MEMBACA FILE (Reading) ---

	// Cara Cepat: Baca seluruh isi file ke memori
	dat, err := os.ReadFile(filename)
	check(err)
	fmt.Print("Isi file (semua): ", string(dat))

	// Cara Manual: Membuka file
	f2, err := os.Open(filename)
	check(err)
	defer f2.Close()

	// Membaca sebagian (misal 5 byte pertama)
	b1 := make([]byte, 5)
	n1, err := f2.Read(b1)
	check(err)
	fmt.Printf("Baca %d bytes: %s\n", n1, string(b1[:n1]))

	// Cara Paling Umum: Baca Baris per Baris (Scanner)
	// Ini hemat memori untuk file besar
	fmt.Println("--- Scan per baris ---")
	
	// Reset pointer file ke awal (karena tadi sudah dibaca sebagian)
	f2.Seek(0, 0) 

	scanner := bufio.NewScanner(f2)
	for scanner.Scan() {
		// scanner.Text() mengambil baris saat ini
		fmt.Println("Baris:", scanner.Text())
	}
}