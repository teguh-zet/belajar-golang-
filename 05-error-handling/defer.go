package errorhandling

import (
    "fmt"
    "os"
)

// Ganti main() jadi Function Public
func RunDeferExample() {
    
    // Kita membuat file temp
    f := createFile("defer.txt")
    
    // --- INI POIN PENTINGNYA ---
    // Kita menjadwalkan penutupan file SEKARANG,
    // tapi eksekusinya nanti SAAT FUNGSI SELESAI.
    defer closeFile(f)

    // Menulis data
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()

    // Hapus file temp agar tidak mengotori komputer (Opsional)
    os.Remove(f.Name()) 

    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}