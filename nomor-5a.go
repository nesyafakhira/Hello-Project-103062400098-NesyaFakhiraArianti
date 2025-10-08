package main

import (
	"fmt" // Mengimpor package fmt untuk input-output
)

// Fungsi untuk menghitung volume tabung berdasarkan jari-jari dan tinggi
func hitungVolume(r, t float64) float64 {
	return 3.14 * r * r * t // Rumus volume tabung: π × r² × t
}

// Fungsi untuk menghitung massa zat cair berdasarkan volume dan massa jenis
func hitungMassa(volume, massaJenis float64) float64 {
	return volume * massaJenis // Rumus massa: massa jenis × volume
}

// Prosedur untuk mengecek keseimbangan antara massa kiri dan kanan
func cekBalance(massaKiri, massaKanan float64) {
	if massaKiri == massaKanan { // Jika massa kiri sama dengan massa kanan
		fmt.Println("BALANCE") // Cetak "BALANCE"
	} else { // Jika tidak seimbang
		if massaKiri > massaKanan { // Jika massa kiri lebih besar
			fmt.Printf("%.3f\n", massaKiri-massaKanan) // Cetak selisih massa kiri - kanan
		} else { // Jika massa kanan lebih besar
			fmt.Printf("%.3f\n", massaKanan-massaKiri) // Cetak selisih massa kanan - kiri
		}
	}
}

func main() {
	var r float64                // Variabel untuk menyimpan jari-jari alas tabung
	var tKiri, rhoKiri float64   // Variabel untuk tinggi dan massa jenis zat cair di tabung kiri
	var tKanan, rhoKanan float64 // Variabel untuk tinggi dan massa jenis zat cair di tabung kanan

	fmt.Scan(&r) // Membaca input jari-jari alas tabung
	fmt.Scan(&tKiri, &rhoKiri) // Membaca input tinggi dan massa jenis zat cair di tabung kiri
	fmt.Scan(&tKanan, &rhoKanan) // Membaca input tinggi dan massa jenis zat cair di tabung kanan

	volumeKiri := hitungVolume(r, tKiri) // Menghitung volume zat cair di tabung kiri
	volumeKanan := hitungVolume(r, tKanan) // Menghitung volume zat cair di tabung kanan

	massaKiri := hitungMassa(volumeKiri, rhoKiri) // Menghitung massa zat cair di tabung kiri
	massaKanan := hitungMassa(volumeKanan, rhoKanan) // Menghitung massa zat cair di tabung kanan

	cekBalance(massaKiri, massaKanan) // Memeriksa keseimbangan dan mencetak hasil
}


//////////////////////************nomor 5B****************///////////////////////////////////////

// program Timbangan
// kamus
//     const pi : real = 3.14  // Konstanta nilai π untuk perhitungan volume tabung

// function volume(r, t) -> real
// {mengembalikan volume tabung yang memiliki jari-jari lingkaran r dan tinggi t,
// yang mana volume adalah luas alas x tinggi tabung}
//     return pi * r * r * t  // Menghitung volume tabung dengan rumus π × r² × t

// function massa(r, t, p) -> real
// {mengembalikan massa tabung yang memiliki jari-jari lingkaran r dan tinggi t,
// serta massa jenis p, yang mana massa = volume x massa jenis}
//     return volume(r, t) * p  // Menghitung massa zat cair di dalam tabung

// procedure display(m1, m2)
// {I.S. terdefinisi massa zat cair kiri m1 dan massa zat cair kanan m2 pada timbangan
// F.S. menampilkan "BALANCE" apabila m1 dan m2 adalah sama, atau selisih positif apabila m1 dan m2 berbeda}
//     if m1 = m2 then  // Mengecek apakah massa kiri dan kanan sama
//         output "BALANCE"  // Jika sama, cetak "BALANCE"
//     else
//         output abs(m1 - m2)  // Jika berbeda, cetak selisih massa dalam nilai absolut

// algoritma
//     input r  // Membaca jari-jari alas tabung
//     input t1, p1  // Membaca tinggi dan massa jenis zat cair di tabung kiri
//     input t2, p2  // Membaca tinggi dan massa jenis zat cair di tabung kanan

//     m1 ← massa(r, t1, p1)  // Menghitung massa zat cair di tabung kiri
//     m2 ← massa(r, t2, p2)  // Menghitung massa zat cair di tabung kanan

//     display(m1, m2)  // Menampilkan hasil perbandingan massa

// endprogram