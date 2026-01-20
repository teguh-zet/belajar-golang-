package system // Nama package mengikuti nama folder

import (
	"encoding/json"
	"fmt"
)

// 1. Definisikan Struct dengan "Struct Tags"
// `json:"..."` memberi tahu Go cara menamai field saat diubah jadi JSON.
type Response1 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

// Struct tanpa tag (akan menggunakan nama field asli)
type Response2 struct {
	Page   int
	Fruits []string
}

func RunJSONExample() {
	fmt.Println("--- JSON Encoding/Decoding ---")

	// --- ENCODING (Marshal) : Go -> JSON ---

	// Contoh 1: Tipe data primitif
	bolB, _ := json.Marshal(true)
	fmt.Println("Boolean:", string(bolB))

	intB, _ := json.Marshal(14)
	fmt.Println("Integer:", string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println("Float:", string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println("String:", string(strB))

	// Contoh 2: Slice dan Map
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println("Slice:", string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("Map:", string(mapB))

	// Contoh 3: Struct Custom
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println("Struct dengan Tag:", string(res1B))

	// --- DECODING (Unmarshal) : JSON -> Go ---

	// Anggap ini data dari API (berbentuk bytes/string)
	byt := []byte(`{"page": 1, "fruits": ["apple", "peach"]}`)

	// Kita butuh variabel penampung
	var dat map[string]interface{}

	// Cek error saat decoding
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("Decoded Map:", dat)

	// Mengakses data map (harus di-casting karena interface{})
	num := dat["page"].(float64)
	fmt.Println("Page number:", num)

	// Decode langsung ke Struct (Lebih Aman & Sering Dipakai)
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response1{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println("Decoded Struct:", res)
	fmt.Println("Fruit pertama:", res.Fruits[0])
}