package main

import "fmt"

func main() {
	var depan, belakang, kelas, nim string

	fmt.Print("Masukkan Nama  : ")
	fmt.Scanln(&depan, &belakang) // Membaca dua kata untuk nama

	fmt.Print("Masukkan Kelas : ")
	fmt.Scanln(&kelas)

	fmt.Print("Masukkan NIM   : ")
	fmt.Scanln(&nim)

	nama := depan + " " + belakang

	// Menampilkan hasil perkenalan
	fmt.Println("\n===================================")
	fmt.Println("Halo teman-teman! Saya", nama)
	fmt.Println("NIM saya", nim, "dan saya dari kelas", kelas)
	fmt.Println("===================================")
}