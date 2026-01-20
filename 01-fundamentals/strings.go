package fundamentals

import (
	"fmt"
	"unicode/utf8"
)

func GoString() {
	// rune yaitu bilangan bulat yang mewakili titik kode Unicode
	// berbeda dengan ascii, unicode tidak hanya terbatas pada char dan
	// simbol yang digunakan
	//unice lebih luas dengan char dan simbol global
	const s = "tสวัสดี"
	fmt.Println("len s : ",len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// digunakan untuk menghitung panjang dari rune yang dimiliki oleh satu string
	fmt.Println("rune count : ", utf8.RuneCountInString(s))
	// dekode setiap string rune beserta offset nya dalam string
	for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

	fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width

        examineRune(runeValue)
    }
}
func examineRune(r rune){
	if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }else if r== 'ว'{
		fmt.Println("found  ว")
	}
}