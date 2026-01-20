package types

import (
    "fmt"
    "iter"
    "slices"
)

type List2[T any] struct {
    head, tail *element2[T]
}

type element2[T any] struct {
    next *element2[T]
    val  T
}

func (lst *List2[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element2[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element2[T]{val: v}
        lst.tail = lst.tail.next
    }
}
// Method ini mengubah Linked List menjadi 
// sesuatu yang bisa di-iterasi (dilakukan perulangan).
func (lst *List2[T]) All() iter.Seq[T] {
	// Mengembalikan iter.Seq[T], ini adalah tipe standar untuk iterator
	// return anynomous func
    return func(yield func(T) bool) {
		//loop untuk membaca linked list
        for e := lst.head; e != nil; e = e.next {
			// yield(e.val) bertugas "menyerahkan" data ke loop 'for range' di main.
            // yield akan mengembalikan false jika loop di main di-break (berhenti).
            if !yield(e.val) {
                return // berhenti ngrim data
            }
        }
    }
}

func genFib() iter.Seq[int] {
    return func(yield func(int) bool) {
        a, b := 1, 1

        for {
            if !yield(a) {
                return
            }
            a, b = b, a+b
        }
    }
}

func GoRangeIterators() {
    lst := List2[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)

    for e := range lst.All() {
        fmt.Println(e)
    }

    all := slices.Collect(lst.All())
    fmt.Println("all:", all)

    for n := range genFib() {
        if n >= 10 {
            break
        }
        fmt.Println(n)
    }

}