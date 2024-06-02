package main

import (
	"fmt"
	"strings"
)

const NMAX = 1000

type Patient struct {
	namapasien string
	asal       string
	umur       int
}

type Package struct {
	namapaket string
	tanggal   int
}

type rekaphasil struct {
	tanggal    int
	namapaket  string
	namapasien string
}

var pasien [NMAX]Patient
var rekap [NMAX]rekaphasil
var paket [NMAX]Package

var nPasien int

func insert(tambah int) {
	fmt.Println("Nama Pasien, Asal, Umur, Nama Paket, Tahun Periode")
	for i := nPasien; i < (nPasien + tambah); i++ {
		fmt.Scan(&pasien[i].namapasien, &pasien[i].asal, &pasien[i].umur, &paket[i].namapaket, &paket[i].tanggal)
		rekap[i].tanggal = paket[i].tanggal
		rekap[i].namapaket = paket[i].namapaket
		rekap[i].namapasien = pasien[i].namapasien
		rekap[i].tanggal = paket[i].tanggal
	}
	nPasien += tambah
	fmt.Println("Data berhasil dimasukkan")
}

func edit(x string) {
	var pilihan int
	var find bool
	var i int

	for i < nPasien && !find {
		if pasien[i].namapasien == x {
			find = true
		} else {
			i++
		}
	}

	if find {
		header()
		fmt.Println("Ingin mengubah data apa?")
		fmt.Printf("1. Nama Pasien\n2. Asal Pasien\n3. Umur Pasien\n4. Nama Paket\n5. Tahun Periode\n")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			header()
			fmt.Println("Masukkan nama pasien baru: ")
			fmt.Scan(&pasien[i].namapasien)
			rekap[i].namapasien = pasien[i].namapasien
			fmt.Println("Data berhasil diubah")
		case 2:
			header()
			fmt.Println("Masukkan asal pasien baru: ")
			fmt.Scan(&pasien[i].asal)
			fmt.Println("Data berhasil diubah")
		case 3:
			header()
			fmt.Println("Masukkan umur pasien baru: ")
			fmt.Scan(&pasien[i].umur)
			fmt.Println("Data berhasil diubah")
		case 4:
			header()
			fmt.Println("Masukkan nama paket baru: ")
			fmt.Scan(&paket[i].namapaket)
			rekap[i].namapaket = paket[i].namapaket
			fmt.Println("Data berhasil diubah")
		case 5:
			header()
			fmt.Println("Masukkan tahun MCU yang baru: ")
			fmt.Scan(&paket[i].tanggal)
			rekap[i].tanggal = paket[i].tanggal
			fmt.Println("Data berhasil diubah")
		default:
			fmt.Println("Menu tidak tersedia")
		}
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func delete(nama string) {
	var a int
	if BinarySearch(nama, &a) {
		for i := a; i < nPasien-1; i++ {
			pasien[i] = pasien[i+1]
			paket[i] = paket[i+1]
			rekap[i] = rekap[i+1]
		}
		nPasien--
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func laporanPemasukan(periode int) {
	var total int
	var found bool
	fmt.Printf("Data pasien yang melakukan MCU pada periode %d:\n", periode)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("%-20s %-10s %-5s %-15s %-20s\n", "Nama Pasien", "Asal", "Umur", "Nama Paket", "Tahun Periode")
	fmt.Println("----------------------------------------------------------------------")
	for i := 0; i < nPasien; i++ {
		if rekap[i].tanggal == periode {
			fmt.Printf("%-20s %-10s %-5d %-15s %-20d\n", rekap[i].namapasien, pasien[i].asal, pasien[i].umur, rekap[i].namapaket, rekap[i].tanggal)
			found = true
			switch rekap[i].namapaket {
			case "A":
				total += 50000
			case "B":
				total += 75000
			case "C":
				total += 100000
			}
		}
	}
	if !found {
		fmt.Println("Periode tidak ditemukan")
	} else {
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println("Total pemasukan pada tahun", periode, "adalah", total)
	}
}

func cetak() {
	header()
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("%-20s %-10s %-5s %-15s %-20s\n", "Nama Pasien", "Asal", "Umur", "Nama Paket", "Tahun Periode")
	fmt.Println("----------------------------------------------------------------------")
	for i := 0; i < nPasien; i++ {
		fmt.Printf("%-20s %-10s %-5d %-15s %-20d\n", pasien[i].namapasien, pasien[i].asal, pasien[i].umur, paket[i].namapaket, paket[i].tanggal)
	}
}

func sortPaket() {
	var j int
	var tempPaket Package
	var tempPasien Patient
	var tempRekap rekaphasil
	var pass int = 1
	var pilihan int
	fmt.Printf("1.Ascending\n2.Descending\n")
	fmt.Scan(&pilihan)
	for pass <= nPasien-1 {
		for j = 0; j <= nPasien-pass-1; j++ {
			if (pilihan == 1 && strings.Compare(paket[j].namapaket, paket[j+1].namapaket) > 0) || (pilihan == 2 && strings.Compare(paket[j].namapaket, paket[j+1].namapaket) < 0) {
				tempPaket = paket[j]
				paket[j] = paket[j+1]
				paket[j+1] = tempPaket
				tempPasien = pasien[j]
				pasien[j] = pasien[j+1]
				pasien[j+1] = tempPasien
				tempRekap = rekap[j]
				rekap[j] = rekap[j+1]
				rekap[j+1] = tempRekap
			}
		}
		pass++
	}
	cetak()
}

func sortPeriode() {
	var j int
	var tempPaket Package
	var tempPasien Patient
	var tempRekap rekaphasil
	var pass int = 1
	var pilihan int
	fmt.Printf("1.Ascending\n2.Descending\n")
	fmt.Scan(&pilihan)
	for pass <= nPasien-1 {
		for j = 0; j <= nPasien-pass-1; j++ {
			if (pilihan == 1 && paket[j].tanggal > paket[j+1].tanggal) || (pilihan == 2 && paket[j].tanggal < paket[j+1].tanggal) {
				tempPaket = paket[j]
				paket[j] = paket[j+1]
				paket[j+1] = tempPaket
				tempPasien = pasien[j]
				pasien[j] = pasien[j+1]
				pasien[j+1] = tempPasien
				tempRekap = rekap[j]
				rekap[j] = rekap[j+1]
				rekap[j+1] = tempRekap
			}
		}
		pass++
	}
	cetak()
}

func sortNama() {
	var j int
	var tempPaket Package
	var tempPasien Patient
	var tempRekap rekaphasil
	var pass int = 1
	var pilihan int
	fmt.Printf("1.Ascending\n2.Descending\n")
	fmt.Scan(&pilihan)
	for pass <= nPasien-1 {
		for j = 0; j <= nPasien-pass-1; j++ {
			if (pilihan == 1 && strings.Compare(pasien[j].namapasien, pasien[j+1].namapasien) > 0) || (pilihan == 2 && strings.Compare(pasien[j].namapasien, pasien[j+1].namapasien) < 0) {
				tempPaket = paket[j]
				paket[j] = paket[j+1]
				paket[j+1] = tempPaket
				tempPasien = pasien[j]
				pasien[j] = pasien[j+1]
				pasien[j+1] = tempPasien
				tempRekap = rekap[j]
				rekap[j] = rekap[j+1]
				rekap[j+1] = tempRekap
			}
		}
		pass++
	}
	cetak()
}

func searchPaket(nama string) {
	var found bool
	fmt.Printf("Data pasien yang melakukan MCU dengan paket %s:\n", nama)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("%-20s %-10s %-5s %-15s %-20s\n", "Nama Pasien", "Asal", "Umur", "Nama Paket", "Tahun Periode")
	fmt.Println("----------------------------------------------------------------------")
	for i := 0; i < nPasien; i++ {
		if rekap[i].namapaket == nama {
			fmt.Printf("%-20s %-10s %-5d %-15s %-20d\n", rekap[i].namapasien, pasien[i].asal, pasien[i].umur, rekap[i].namapaket, rekap[i].tanggal)
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

func searchPeriode(periode int) {
	var found bool
	fmt.Printf("Data pasien yang melakukan MCU pada periode %d:\n", periode)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("%-20s %-10s %-5s %-15s %-20s\n", "Nama Pasien", "Asal", "Umur", "Nama Paket", "Tahun Periode")
	fmt.Println("----------------------------------------------------------------------")
	for i := 0; i < nPasien; i++ {
		if rekap[i].tanggal == periode {
			fmt.Printf("%-20s %-10s %-5d %-15s %-20d\n", rekap[i].namapasien, pasien[i].asal, pasien[i].umur, rekap[i].namapaket, rekap[i].tanggal)
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

func searchNama(nama string) {
	var found bool
	fmt.Printf("Data pasien dengan nama %s:\n", nama)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("%-20s %-10s %-5s %-15s %-20s\n", "Nama Pasien", "Asal", "Umur", "Nama Paket", "Tahun Periode")
	fmt.Println("----------------------------------------------------------------------")
	for i := 0; i < nPasien; i++ {
		if rekap[i].namapasien == nama {
			fmt.Printf("%-20s %-10s %-5d %-15s %-20d\n", rekap[i].namapasien, pasien[i].asal, pasien[i].umur, rekap[i].namapaket, rekap[i].tanggal)
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

func header() {
	fmt.Printf("==============================\n")
	fmt.Printf("Aplikasi Medical Check Up\n")
	fmt.Printf("==============================\n\n")
}

func BinarySearch(nama string, mid *int) bool {
	//sort nama pasien menggunakan selection sort
	var i, j, minIdx int
	var tempPasien Patient
	var tempPaket Package
	var tempRekap rekaphasil
	for i = 0; i < nPasien-1; i++ {
		minIdx = i
		for j = i + 1; j < nPasien; j++ {
			if strings.ToLower(pasien[j].namapasien) < strings.ToLower(pasien[minIdx].namapasien) {
				minIdx = j
			}
		}
		tempPasien = pasien[minIdx]
		pasien[minIdx] = pasien[i]
		pasien[i] = tempPasien

		tempPaket = paket[minIdx]
		paket[minIdx] = paket[i]
		paket[i] = tempPaket

		tempRekap = rekap[minIdx]
		rekap[minIdx] = rekap[i]
		rekap[i] = tempRekap
	}

	low, high := 0, nPasien-1
	found := false
	for low <= high && !found {
		*mid = (low + high) / 2
		if pasien[*mid].namapasien == nama {
			found = true
		} else if pasien[*mid].namapasien < nama {
			low = *mid + 1
		} else {
			high = *mid - 1
		}
	}

	return found
}

func main() {
	var tambah int
	var pilihan int

	for {
		header()
		fmt.Printf("1. Penambahan Data\n2. Pengubahan Data\n3. Penghapusan Data\n4. Laporan Pemasukan\n5. Pencarian Pasien\n6. Tampilkan Data Pasien Terurut\n7. Keluar\n")
		fmt.Printf("\nMenu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("Berapa pasien yang ingin ditambahkan? ")
			fmt.Scan(&tambah)
			fmt.Printf("\nSilahkan masukkan data pasien sesuai urutan yang diminta\n")
			insert(tambah)
		case 2:
			fmt.Println("Ingin merubah data pasien atas nama siapa? ")
			var nama string
			fmt.Scan(&nama)
			edit(nama)
		case 3:
			fmt.Println("Ingin menghapus data pasien atas nama siapa? ")
			var nama string
			fmt.Scan(&nama)
			delete(nama)
		case 4:
			fmt.Println("Ingin menampilkan laporan pemasukan MCU pada periode apa? ")
			var periode int
			fmt.Scan(&periode)
			laporanPemasukan(periode)
		case 5:
			fmt.Printf("Ingin mencari daftar pasien berdasarkan apa?\n1. Berdasarkan paket MCU\n2. Berdasarkan Periode\n3. Berdasarkan nama pasien\n4. Kembali\n")
			fmt.Scan(&pilihan)
			switch pilihan {
			case 1:
				fmt.Println("Masukkan nama paket")
				var nama string
				fmt.Scan(&nama)
				searchPaket(nama)
			case 2:
				fmt.Println("Masukkan periode")
				var periode int
				fmt.Scan(&periode)
				searchPeriode(periode)
			case 3:
				fmt.Println("Masukkan nama pasien")
				var nama string
				fmt.Scan(&nama)
				searchNama(nama)
			case 4:
			}
		case 6:
			fmt.Printf("Ingin menampilakan data pasien terurut berdasarkan apa?\n1. Periode MCU\n2. Paket MCU\n3. Nama\n4. Kembali\n")
			fmt.Scan(&pilihan)
			switch pilihan {
			case 1:
				header()
				sortPeriode()
			case 2:
				header()
				sortPaket()
			case 3:
				header()
				sortNama()
			case 4:
			default:
				fmt.Printf("Menu tidak tersedia\n")
			}
		case 7:

		default:

			fmt.Printf("Menu tidak tersedia\n")
		}
		if pilihan == 7 {

			fmt.Println("Aplikasi tertutup")
			break
		}
	}
}
