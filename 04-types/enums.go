package types

import "fmt"

type ServerState int
// <-- Di Go, tidak ada keyword khusus bernama enum 
// seperti di bahasa Java atau C. Sebaliknya, Enum di Go dibuat dengan 
// cara memanipulasi Tipe Data Kustom (Type) dan Konstanta (Const).-->

// kata kuci iota otomatis generate nilai konstanta berurutan
// mulai dari 0,1,2 dan seterusnya
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}
// Di Go, jika sebuah tipe memiliki method bernama String() 
// yang mengembalikan string, maka fungsi fmt.Println() 
// akan otomatis memanggil method ini.
// Efeknya: Saat Anda melakukan fmt.Println(StateIdle), 
// outputnya bukan angka 0, melainkan tulisan "idle".
// Jika method ini dihapus, outputnya akan kembali menjadi angka.
func (ss ServerState) String() string {
	return stateName[ss]
}
func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:
        return StateIdle
    case StateError:
        return StateError
    default:
		// panic: Ini jaring pengaman. Jika ada status aneh yang masuk 
		// (misal angka 99), 
		// program akan berhenti paksa dan memberi tahu errornya.
        panic(fmt.Errorf("unknown state: %s", s))
    }
}
func GoEnums() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	ns2 := transition(ns)
	fmt.Println(ns2)
}
