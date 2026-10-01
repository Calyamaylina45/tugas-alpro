package main

import (
    "fmt"
    "math"
)

// Fungsi terpisah untuk menghitung luas
func hitungLuas(r float64) float64 {
    return math.Pi * math.Pow(r, 2)
}

// Fungsi terpisah untuk menghitung keliling
func hitungKeliling(r float64) float64 {
    return 2 * math.Pi * r
}

func main() {
    var jariJari float64

    fmt.Print("Input jari-jari lingkaran: ")
    fmt.Scan(&jariJari)

    luas := hitungLuas(jariJari)
    keliling := hitungKeliling(jariJari)

    fmt.Println("--- Hasil Perhitungan ---")
    fmt.Printf("Luas Lingkaran     : %.2f\n", luas)
    fmt.Printf("Keliling Lingkaran : %.2f\n", keliling)
}