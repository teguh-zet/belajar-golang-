package collections // Sesuaikan nama package dengan file lain di folder 02

import (
	"fmt"
	"slices" // Package baru di Run 1.21+
)

// Ganti nama func main() menjadi nama yang deskriptif & Public
func RunSortingExample() {
	fmt.Println("--- Basic Sorting ---")

	// 1. Sorting String
	strs := []string{"z", "a", "b"}

	// slices.Sort mengubah urutan elemen di dalam slice asli (Mutate)
	slices.Sort(strs)
	fmt.Println("Strings:", strs) // Output: [a b c]

	// 2. Sorting Integer
	ints := []int{7, 12, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints) // Output: [2 4 7]

	// 3. Cek apakah sudah terurut
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s) // Output: true
}
