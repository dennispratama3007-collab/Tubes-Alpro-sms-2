package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Tagihan struct {
	ID       int
	Nama     string
	Kategori string
	Nominal  float64
	Status   string
}

var daftarTagihan []Tagihan
var nextID int = 1

func tambahTagihan() {
	var nama string
	var kategori string
	var nominal float64
	var status string

	fmt.Print("Nama Tagihan : ")
	fmt.Scanln(&nama)

	fmt.Print("Kategori : ")
	fmt.Scanln(&kategori)

	fmt.Print("Nominal : ")
	fmt.Scanln(&nominal)

	fmt.Print("Status (Lunas/Belum) : ")
	fmt.Scanln(&status)

	tagihan := Tagihan{
		ID:       nextID,
		Nama:     nama,
		Kategori: kategori,
		Nominal:  nominal,
		Status:   status,
	}

	daftarTagihan = append(daftarTagihan, tagihan)
	nextID++

	fmt.Println("Data berhasil ditambahkan")
}

func tampilTagihan() {
	if len(daftarTagihan) == 0 {
		fmt.Println("Belum ada data")
		return
	}

	fmt.Println("\nDAFTAR TAGIHAN")

	for _, t := range daftarTagihan {
		fmt.Printf(
			"ID:%d | %s | %s | %.0f | %s\n",
			t.ID,
			t.Nama,
			t.Kategori,
			t.Nominal,
			t.Status,
		)
	}
}

func ubahStatus() {
	var id int
	var status string

	fmt.Print("Masukkan ID : ")
	fmt.Scanln(&id)

	for i := range daftarTagihan {
		if daftarTagihan[i].ID == id {
			fmt.Print("Status Baru : ")
			fmt.Scanln(&status)

			daftarTagihan[i].Status = status

			fmt.Println("Status berhasil diperbarui")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func hapusTagihan() {
	var id int

	fmt.Print("Masukkan ID : ")
	fmt.Scanln(&id)

	for i := range daftarTagihan {
		if daftarTagihan[i].ID == id {
			daftarTagihan = append(
				daftarTagihan[:i],
				daftarTagihan[i+1:]...,
			)

			fmt.Println("Data berhasil dihapus")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func totalTagihan() {
	var total float64

	for _, t := range daftarTagihan {
		if strings.ToLower(t.Status) != "lunas" {
			total += t.Nominal
		}
	}

	fmt.Printf("Total Tagihan Belum Lunas : %.0f\n", total)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== SISTEM INFORMASI MANAJEMEN TAGIHAN BULANAN ===")
		fmt.Println("1. Tambah Tagihan")
		fmt.Println("2. Lihat Tagihan")
		fmt.Println("3. Ubah Status")
		fmt.Println("4. Hapus Tagihan")
		fmt.Println("5. Total Tagihan")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih Menu : ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahTagihan()
		case 2:
			tampilTagihan()
		case 3:
			ubahStatus()
		case 4:
			hapusTagihan()
		case 5:
			totalTagihan()
		case 6:
			fmt.Println("Program selesai")
			return
		default:
			fmt.Println("Menu tidak tersedia")
		}

		reader.ReadString('\n')
	}
}
