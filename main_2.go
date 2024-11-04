package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fungsi untuk membalik urutan karakter dalam sebuah kata
func reverseWord(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan kalimat (minimal 3 kata): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Memisahkan kalimat menjadi kata-kata
	words := strings.Fields(input)

	// Mengecek apakah jumlah kata kurang dari 3
	if len(words) < 3 {
		fmt.Println("Katanya kurang ngab, minimal 3")
		return
	}

	// Membalik setiap kata secara individual
	for i, word := range words {
		words[i] = reverseWord(word)
	}

	// Menggabungkan kembali kata-kata yang telah dibalik
	result := strings.Join(words, " ")
	fmt.Println("Hasil:", result)
}

