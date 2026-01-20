package types

import "fmt"
// Bagian ini menunjukkan cara membuat fungsi yang bisa menerima 
// slice dengan tipe data apa saja (string, int, float, dll).
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
    for i := range s {
        if v == s[i] {
            return i
        }
    }
    return -1
}
// Bagian ini menunjukkan cara membuat Struktur 
// Data (Linked List) yang fleksibel.
type List[T any] struct { // <-- Generic Types (Struct)
// [T any]:
// T adalah tipe data placeholder.
// any adalah alias untuk interface{}. Artinya T bisa berupa apa saja (int, string, struct user, pointer, dll).
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func GoGenerics() {
    var s = []string{"foo", "bar", "zoo"}

    fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))
	//  _ = SlicesIndex[[]string, string](s, "zoo") <-- contoh asli, tetapi ada error
	// biru yang menunjukkan unnecessary argument
    _ = SlicesIndex(s, "zoo")
	

    // lst := List[int]{}
    // lst.Push(10)
    // lst.Push(13)
    // lst.Push(23)
	// lst.Push(6)
    lst := List[string]{}
    lst.Push("teguh")
    lst.Push("oioi")
    lst.Push("makan")
	lst.Push("nasi")
    fmt.Println("list:", lst.AllElements())
}