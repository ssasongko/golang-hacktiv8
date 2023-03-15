package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// mendifinisikan struct
type siswa struct {
	no        int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	// cek apakah ada argumen yang diinputkan
	if len(os.Args) == 2 {
		pengendalianSukses()
		return
	} else if len(os.Args) == 1 {
		pengendalianGagal(0)
		return
	}
	pengendalianGagal(1)
}

func pengendalianGagal(jenis int) {
	if jenis == 0 {
		fmt.Println("Tolong masukan argumen berupa nomor absen ketika menjalankan perintah, contoh :")
		fmt.Println("go run biodata.go 1")
		return
	}

	fmt.Println("Hanya boleh memasukan satu argumen")
}

func pengendalianSukses() {
	// arg sebagi inputan dari user
	var arg = os.Args[1:][0]
	// konversi string ke int
	var intArg, err = strconv.Atoi(arg)
	_ = err

	// panggil fungsi cariSiswa
	cariSiswa(intArg)
}

// fungsi untuk mencari data siswa berdasarakan index
func cariSiswa(noAbsen int) {
	var dataSiswa = []siswa{
		{no: 1, nama: "Nur Sasongko", alamat: "Kota Bandung", pekerjaan: "Mahasiswa", alasan: "Mau belajar"},
		{no: 2, nama: "Nur Kholis Majid", alamat: "Kota Pangandaran", pekerjaan: "Teller", alasan: "Ingin merubah nasib"},
		{no: 3, nama: "Wahyu", alamat: "Kota Pati", pekerjaan: "Pengangguran", alasan: "Memiliki minat terhadap pemrograman"},
		{no: 4, nama: "Sarah Purba", alamat: "Kota Purwakarta", pekerjaan: "Vlogger", alasan: "Ingin membangun channel pemrograman bahasa Go"},
	}

	// cek apakah no absen ada di dalam data siswa
	for _, item := range dataSiswa {
		if item.no == noAbsen {
			pemisah(20)
			fmt.Println("Data siswa ditemukan")
			pemisah(20)
			fmt.Printf("No Absen: %d\n", item.no)
			fmt.Printf("Nama Siswa: %s\n", item.nama)
			fmt.Printf("Alamat: %s\n", item.alamat)
			fmt.Printf("Pekerjaan: %s\n", item.pekerjaan)
			fmt.Printf("Alasan: %s\n", item.alasan)
			pemisah(20)
			return
		}
	}
	fmt.Println("Data siswa tidak ditemukan")
}

// fungsi untuk membuat pemisah
func pemisah(jumlah int) {
	fmt.Println(strings.Repeat("-", 20))
}
