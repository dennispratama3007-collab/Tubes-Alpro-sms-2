package main

import "fmt"

const NMAX int = 100

type Tagihan struct {
	ID       int
	Nama     string
	Kategori string
	Nominal  float64
	Status   string
}

type DataTagihan struct {
	data   [NMAX]Tagihan
	jumlah int
}

func tambahTagihan(T *DataTagihan, nextID *int) {
	var tagihan Tagihan

	fmt.Println("\n=== Tambah Tagihan ===")

	tagihan.ID = *nextID

	fmt.Print("Nama Tagihan     : ")
	fmt.Scanln(&tagihan.Nama)

	fmt.Print("Kategori         : ")
	fmt.Scanln(&tagihan.Kategori)

	fmt.Print("Nominal          : ")
	fmt.Scanln(&tagihan.Nominal)

	tagihan.Status = "Belum"

	T.data[T.jumlah] = tagihan
	T.jumlah++

	*nextID++

	fmt.Println("Tagihan berhasil ditambahkan.")
}

func tampilTagihan(T DataTagihan) {
	var i int

	fmt.Println("\n=== Daftar Tagihan ===")

	if T.jumlah == 0 {
		fmt.Println("Belum ada data.")
	} else {
		for i = 0; i < T.jumlah; i++ {
			fmt.Println("ID        :", T.data[i].ID)
			fmt.Println("Nama      :", T.data[i].Nama)
			fmt.Println("Kategori  :", T.data[i].Kategori)
			fmt.Printf("Nominal   : Rp %.0f\n", T.data[i].Nominal)
			fmt.Println("Status    :", T.data[i].Status)
			fmt.Println("---------------------------")
		}
	}
}

func cariTagihan(T DataTagihan, nama string) {
	var i int
	var ditemukan bool

	ditemukan = false

	fmt.Println("\n=== Hasil Pencarian ===")

	for i = 0; i < T.jumlah; i++ {

		if T.data[i].Nama == nama {

			ditemukan = true

			fmt.Println("ID       :", T.data[i].ID)
			fmt.Println("Nama     :", T.data[i].Nama)
			fmt.Println("Kategori :", T.data[i].Kategori)
			fmt.Printf("Nominal  : Rp %.0f\n", T.data[i].Nominal)
			fmt.Println("Status   :", T.data[i].Status)
			fmt.Println("---------------------------")
		}
	}

	if !ditemukan {
		fmt.Println("Tagihan tidak ditemukan.")
	}
}

func bayarTagihan(T *DataTagihan) {
	var id int
	var i int
	var ditemukan bool

	fmt.Println("\n=== Bayar Tagihan ===")

	fmt.Print("Masukkan ID tagihan: ")
	fmt.Scanln(&id)

	ditemukan = false

	for i = 0; i < T.jumlah; i++ {

		if T.data[i].ID == id {

			ditemukan = true

			if T.data[i].Status == "Lunas" {
				fmt.Println("Tagihan sudah lunas.")
			} else {
				T.data[i].Status = "Lunas"
				fmt.Println("Tagihan berhasil dibayar.")
			}
		}
	}

	if !ditemukan {
		fmt.Println("ID tagihan tidak ditemukan.")
	}
}

func urutNominal(T *DataTagihan) {
	var i, j, min int
	var temp Tagihan

	for i = 0; i < T.jumlah-1; i++ {
		min = i

		for j = i + 1; j < T.jumlah; j++ {
			if T.data[j].Nominal < T.data[min].Nominal {
				min = j
			}
		}

		temp = T.data[i]
		T.data[i] = T.data[min]
		T.data[min] = temp
	}

	fmt.Println("Data berhasil diurutkan.")
}

func totalBelumBayar(T DataTagihan) float64 {
	var i int
	var total float64

	total = 0

	for i = 0; i < T.jumlah; i++ {
		if T.data[i].Status == "Belum" {
			total += T.data[i].Nominal
		}
	}

	return total
}

func menu() int {
	var pilihan int

	fmt.Println("\n===== APLIKASI MANAJEMEN TAGIHAN BULANAN =====")
	fmt.Println("1. Tambah Tagihan")
	fmt.Println("2. Tampilkan Tagihan")
	fmt.Println("3. Cari Tagihan")
	fmt.Println("4. Urutkan Nominal")
	fmt.Println("5. Bayar Tagihan")
	fmt.Println("6. Total Belum Dibayar")
	fmt.Println("7. Keluar")

	fmt.Print("Pilih menu: ")
	fmt.Scanln(&pilihan)

	return pilihan
}

func main() {
	var daftar DataTagihan
	var pilihan int
	var nextID int
	var nama string

	nextID = 1

	for pilihan != 7 {
		pilihan = menu()

		switch pilihan {

		case 1:
			tambahTagihan(&daftar, &nextID)

		case 2:
			tampilTagihan(daftar)

		case 3:
			fmt.Print("Masukkan nama tagihan: ")
			fmt.Scanln(&nama)

			cariTagihan(daftar, nama)

		case 4:
			urutNominal(&daftar)

		case 5:
			bayarTagihan(&daftar)

		case 6:
			fmt.Printf("Total tagihan belum dibayar: Rp %.2f\n",
				totalBelumBayar(daftar))

		case 7:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")

		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}
