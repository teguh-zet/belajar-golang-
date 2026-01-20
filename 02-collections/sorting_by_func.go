package collections // Package sama dengan file lain di folder ini

import (
    "cmp"
    "fmt"
    "slices"
)

// Ganti nama func main menjadi Public Function
func RunSortFuncExample() {
    fmt.Println("--- Sorting by Function ---")

    // CONTOH 1: Mengurutkan String berdasarkan PANJANGNYA (Length)
    // Bukan berdasarkan abjad
    fruits := []string{"peach", "banana", "kiwi"}

    // Fungsi komparator kustom
    // return negatif jika a < b
    // return 0 jika a == b
    // return positif jika a > b
    lenCmp := func(a, b string) int {
        return cmp.Compare(len(a), len(b))
    }

    slices.SortFunc(fruits, lenCmp)
    fmt.Println("By Length:", fruits) 
    // Output: [kiwi peach banana] (4 huruf, 5 huruf, 6 huruf)


    // CONTOH 2: Mengurutkan Struct User berdasarkan UMUR (Age)
    type Person struct {
        name string
        age  int
    }

    people := []Person{
        {name: "Jax", age: 37},
        {name: "TJ", age: 25},
        {name: "Alex", age: 72},
    }
    person := []Person{
        {name: "wadd", age: 100},
        {name: "TJ", age: 25},
        {name: "Alex", age: 72},
    }
	slices.SortFunc(person, func(a,b Person)int{
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(person)
    // Menggunakan Anonymous Function langsung di dalam parameter
    slices.SortFunc(people, func(a, b Person) int {
        // Kita bandingkan field 'age' milik a dan b
        return cmp.Compare(a.age, b.age)
    })

    fmt.Println("By Age:   ", people)
    // Output: TJ (25), Jax (37), Alex (72)
}