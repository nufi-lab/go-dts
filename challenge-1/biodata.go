package main

import (
	"fmt"
	"os"
)

type Person struct {
	Name, Address, Job, ReasonStudyGo string
}

// membuat data absen dengan map
var data = map[int]Person{
	1:  {"Aisyah Azura", "Jl. Raya Besar No. 5", "Developer", "Suka dengan performa tinggi Golang."},
	2:  {"Annisa Rahma", "Jl. Indah Sekali No. 1", "Designer", "Ingin belajar bahasa pemrograman baru."},
	3:  {"Budi Santoso", "Jl. Harmoni No. 10", "System Analyst", "Pengalaman dalam mengelola sistem besar."},
	4:  {"Cantika Rahayu", "Jl. Sejahtera Semuanya 7", "Data Scientist", "Golang cocok untuk pengolahan data."},
	5:  {"Citra Dewi", "Jl. Nusantara Sejahtera 12", "Frontend Developer", "Minat pada pengembangan antarmuka pengguna yang responsif."},
	6:  {"Dika Pratama", "Jl. Cemerlang No. 8", "Database Administrator", "Keahlian dalam mengoptimalkan performa database."},
	7:  {"Eka Wijaya", "Jl. Perkasa No. 15", "Mobile App Developer", "Pengalaman dalam pengembangan aplikasi seluler."},
	8:  {"Fithri Zulfa", "Jl. Damai Tenteram No. 99", "Software Engineer", "Mendengar banyak tentang efisiensi Golang."},
	9:  {"Fitria Amalia", "Jl. Damai Indah No. 22", "UI/UX Designer", "Senang menciptakan desain antarmuka yang menarik."},
	10: {"Gigi Hadid", "Jl. Mawar No. 3", "Guru", "Ingin switch career."},
	// tambahkan data teman lain jika diperlukan
}

func showBiodata(absen int) {
	friend, found := data[absen]

	if !found {
		fmt.Println("Biodata dengan absen ", absen, "tidak ditemukan!")
		fmt.Println("Absen maksimal adalah nomor absen", len(data))
		return
	}

	fmt.Println("Biodata - Absen ke-", absen)
	fmt.Println("Nama: ", friend.Name)
	fmt.Println("Alamat: ", friend.Address)
	fmt.Println("Pekerjaan: ", friend.Job)
	fmt.Println("Alasan belajar golang: ", friend.ReasonStudyGo)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Cara penggunaan: go run biodata.go <nomor_absen>")
		os.Exit(1)
	}

	absen := os.Args[1]
	// mengubah string menjadi integer
	var absenInt int
	fmt.Sscanf(absen, "%d", &absenInt)

	// menampilkan biodata teman berdasarkan nomor absen
	showBiodata(absenInt)
}
