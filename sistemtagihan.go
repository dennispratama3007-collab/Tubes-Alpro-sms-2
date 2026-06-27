package main

import (
	"fmt"
	"strings"
)

const NMAX int = 100

type Tanggal struct {
	Hari  int
	Bulan int
	Tahun int
}

type Tagihan struct {
	ID         int
	Nama       string
	Kategori   string
	Nominal    float64
	Status     string
	JatuhTempo Tanggal
}

type DataTagihan struct {
	data   [NMAX]Tagihan
	jumlah int
}

func lebihAwal(t1, t2 Tanggal) bool {
	if t1.Tahun != t2.Tahun {
		return t1.Tahun < t2.Tahun
	}
	if t1.Bulan != t2.Bulan {
		return t1.Bulan < t2.Bulan
	}
	return t1.Hari < t2.Hari
}

func tambahTagihan(T *DataTagihan, nextID *int) {
	if T.jumlah >= NMAX {
		fmt.Println("Kapasitas penyimpanan tagihan sudah penuh!")
		return
	}

	var tagihan Tagihan
	tagihan.ID = *nextID

	fmt.Println("\n=== Tambah Tagihan Baru ===")
	fmt.Print("Nama Tagihan (Gunakan _ untuk spasi): ")
	fmt.Scanln(&tagihan.Nama)

	fmt.Print("Kategori (Gunakan _ untuk spasi)    : ")
	fmt.Scanln(&tagihan.Kategori)

	fmt.Print("Nominal Biaya                       : ")
	fmt.Scanln(&tagihan.Nominal)

	fmt.Println("Tanggal Jatuh Tempo:")
	fmt.Print("  Hari (DD)    : ")
	fmt.Scanln(&tagihan.JatuhTempo.Hari)
	fmt.Print("  Bulan (MM)   : ")
	fmt.Scanln(&tagihan.JatuhTempo.Bulan)
	fmt.Print("  Tahun (YYYY) : ")
	fmt.Scanln(&tagihan.JatuhTempo.Tahun)

	tagihan.Status = "Belum"

	T.data[T.jumlah] = tagihan
	T.jumlah++
	*nextID++

	fmt.Println("Tagihan berhasil ditambahkan!")
}

func tampilTagihan(T DataTagihan) {
	fmt.Println("\n=== Daftar Tagihan (SIMTAB) ===")
	if T.jumlah == 0 {
		fmt.Println("Belum ada data tagihan tercatat.")
		return
	}

	for i := 0; i < T.jumlah; i++ {
		fmt.Printf("ID         : %d\n", T.data[i].ID)
		fmt.Printf("Nama       : %s\n", T.data[i].Nama)
		fmt.Printf("Kategori   : %s\n", T.data[i].Kategori)
		fmt.Printf("Nominal    : Rp %.0f\n", T.data[i].Nominal)
		fmt.Printf("Status     : %s\n", T.data[i].Status)
		fmt.Printf("Jatuh Tempo: %02d-%02d-%d\n", T.data[i].JatuhTempo.Hari, T.data[i].JatuhTempo.Bulan, T.data[i].JatuhTempo.Tahun)
		fmt.Println("------------------------------------")
	}
}

func ubahTagihan(T *DataTagihan) {
	var id, idx int
	fmt.Println("\n=== Ubah Data Tagihan ===")
	fmt.Print("Masukkan ID Tagihan yang ingin diubah: ")
	fmt.Scanln(&id)

	idx = -1
	for i := 0; i < T.jumlah; i++ {
		if T.data[i].ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("Tagihan dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Printf("Nama Baru [%s] (Gunakan _ untuk spasi): ", T.data[idx].Nama)
	fmt.Scanln(&T.data[idx].Nama)

	fmt.Printf("Kategori Baru [%s] (Gunakan _ untuk spasi): ", T.data[idx].Kategori)
	fmt.Scanln(&T.data[idx].Kategori)

	fmt.Printf("Nominal Baru [%.0f]: ", T.data[idx].Nominal)
	fmt.Scanln(&T.data[idx].Nominal)

	fmt.Println("Data tagihan berhasil diperbarui.")
}

func hapusTagihan(T *DataTagihan) {
	var id, idx int
	fmt.Println("\n=== Hapus Data Tagihan ===")
	fmt.Print("Masukkan ID Tagihan yang ingin dihapus: ")
	fmt.Scanln(&id)

	idx = -1
	for i := 0; i < T.jumlah; i++ {
		if T.data[i].ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("Tagihan dengan ID tersebut tidak ditemukan.")
		return
	}

	for i := idx; i < T.jumlah-1; i++ {
		T.data[i] = T.data[i+1]
	}
	T.jumlah--

	fmt.Println("Tagihan berhasil dihapus dari sistem.")
}

func bayarTagihan(T *DataTagihan) {
	var id int
	fmt.Println("\n=== Update Status Pembayaran ===")
	fmt.Print("Masukkan ID Tagihan yang dibayar: ")
	fmt.Scanln(&id)

	for i := 0; i < T.jumlah; i++ {
		if T.data[i].ID == id {
			if T.data[i].Status == "Lunas" {
				fmt.Println("Pemberitahuan: Tagihan ini sudah berstatus Lunas.")
			} else {
				T.data[i].Status = "Lunas"
				fmt.Println("Sukses! Status tagihan diubah menjadi Lunas.")
			}
			return
		}
	}
	fmt.Println("ID Tagihan tidak ditemukan.")
}

func cariSequential(T DataTagihan, kataKunci string) {
	var ditemukan bool = false
	fmt.Println("\n--- Hasil Pencarian (Sequential Search) ---")

	for i := 0; i < T.jumlah; i++ {
		if strings.Contains(strings.ToLower(T.data[i].Nama), strings.ToLower(kataKunci)) ||
			strings.Contains(strings.ToLower(T.data[i].Kategori), strings.ToLower(kataKunci)) {
			ditemukan = true
			fmt.Printf("ID: %d | %s [%s] | Rp %.0f | Status: %s\n",
				T.data[i].ID, T.data[i].Nama, T.data[i].Kategori, T.data[i].Nominal, T.data[i].Status)
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada nama tagihan atau kategori yang cocok.")
	}
}

func urutNamaSaja(T *DataTagihan) {
	for i := 1; i < T.jumlah; i++ {
		key := T.data[i]
		j := i - 1
		for j >= 0 && strings.ToLower(T.data[j].Nama) > strings.ToLower(key.Nama) {
			T.data[j+1] = T.data[j]
			j--
		}
		T.data[j+1] = key
	}
}

func cariBinary(T DataTagihan, nama string) {
	urutNamaSaja(&T)

	var low, high, mid, idx int
	low = 0
	high = T.jumlah - 1
	idx = -1

	for low <= high && idx == -1 {
		mid = (low + high) / 2
		if strings.ToLower(T.data[mid].Nama) == strings.ToLower(nama) {
			idx = mid
		} else if strings.ToLower(T.data[mid].Nama) < strings.ToLower(nama) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println("\n--- Hasil Pencarian (Binary Search) ---")
	if idx != -1 {
		fmt.Printf("Ditemukan! ID: %d | Nama: %s | Kategori: %s | Nominal: Rp %.0f | Status: %s\n",
			T.data[idx].ID, T.data[idx].Nama, T.data[idx].Kategori, T.data[idx].Nominal, T.data[idx].Status)
	} else {
		fmt.Println("Data tidak ditemukan. (Catatan: Masukkan nama lengkap sesuai yang terdaftar).")
	}
}

func urutSelection(T *DataTagihan) {
	var minIdx int
	for i := 0; i < T.jumlah-1; i++ {
		minIdx = i
		for j := i + 1; j < T.jumlah; j++ {
			if lebihAwal(T.data[j].JatuhTempo, T.data[minIdx].JatuhTempo) {
				minIdx = j
			}
		}
		T.data[i], T.data[minIdx] = T.data[minIdx], T.data[i]
	}
	fmt.Println("Data berhasil diurutkan berdasarkan tanggal jatuh tempo terdekat (Selection Sort).")
}

func urutInsertion(T *DataTagihan) {
	for i := 1; i < T.jumlah; i++ {
		key := T.data[i]
		j := i - 1
		for j >= 0 && lebihAwal(key.JatuhTempo, T.data[j].JatuhTempo) {
			T.data[j+1] = T.data[j]
			j--
		}
		T.data[j+1] = key
	}
	fmt.Println("Data berhasil diurutkan berdasarkan tanggal jatuh tempo terdekat (Insertion Sort).")
}

func tampilStatistik(T DataTagihan) {
	fmt.Println("\n=== Statistik Keuangan SIMTAB ===")
	if T.jumlah == 0 {
		fmt.Println("Data kosong. Statistik belum dapat dihitung.")
		return
	}

	var totalBiaya float64 = 0
	var jumlahLunas int = 0

	for i := 0; i < T.jumlah; i++ {
		totalBiaya += T.data[i].Nominal
		if T.data[i].Status == "Lunas" {
			jumlahLunas++
		}
	}

	persentase := (float64(jumlahLunas) / float64(T.jumlah)) * 100

	fmt.Printf("Total Biaya Seluruh Tagihan : Rp %.2f\n", totalBiaya)
	fmt.Printf("Persentase Tagihan Lunas    : %.2f%%\n", persentase)
}

func main() {
	var daftar DataTagihan
	var pilihan, subPilihan int
	var nextID int = 1
	var kataKunci string

	for pilihan != 9 {
		fmt.Println("\n=============================================")
		fmt.Println("     APLIKASI MANAJEMEN TAGIHAN BULANAN      ")
		fmt.Println("=============================================")
		fmt.Println("1. Tambah Tagihan")
		fmt.Println("2. Tampilkan Semua Tagihan")
		fmt.Println("3. Ubah Informasi Tagihan")
		fmt.Println("4. Hapus Data Tagihan")
		fmt.Println("5. Bayar Tagihan (Update Status)")
		fmt.Println("6. Cari Data Tagihan")
		fmt.Println("7. Urutkan Tagihan (Jatuh Tempo Terdekat)")
		fmt.Println("8. Tampilkan Statistik Aplikasi")
		fmt.Println("9. Keluar")
		fmt.Println("=============================================")
		fmt.Print("Pilih Menu (1-9): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahTagihan(&daftar, &nextID)
		case 2:
			tampilTagihan(daftar)
		case 3:
			ubahTagihan(&daftar)
		case 4:
			hapusTagihan(&daftar)
		case 5:
			bayarTagihan(&daftar)
		case 6:
			fmt.Println("\nPilihan Metode Pencarian:")
			fmt.Println("1. Sequential Search (Bisa potongan kata)")
			fmt.Println("2. Binary Search (Harus kata lengkap)")
			fmt.Print("Pilih metode (1/2): ")
			fmt.Scanln(&subPilihan)

			fmt.Print("Masukkan kata kunci / nama: ")
			fmt.Scanln(&kataKunci)

			if subPilihan == 1 {
				cariSequential(daftar, kataKunci)
			} else if subPilihan == 2 {
				cariBinary(daftar, kataKunci)
			} else {
				fmt.Println("Pilihan metode pencarian tidak valid.")
			}
		case 7:
			fmt.Println("\nPilihan Algoritma Pengurutan:")
			fmt.Println("1. Selection Sort")
			fmt.Println("2. Insertion Sort")
			fmt.Print("Pilih algoritma (1/2): ")
			fmt.Scanln(&subPilihan)

			if subPilihan == 1 {
				urutSelection(&daftar)
			} else if subPilihan == 2 {
				urutInsertion(&daftar)
			} else {
				fmt.Println("Pilihan algoritma sorting tidak valid.")
			}
		case 8:
			tampilStatistik(daftar)
		case 9:
			fmt.Println("\nTerima kasih! Selesai menggunakan SIMTAB.")
		default:
			fmt.Println("Pilihan salah. Harap masukkan nomor menu antara 1 sampai 9.")
		}
	}
}
