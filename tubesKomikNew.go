package main

import "fmt"

const NMAX = 100
const BIAYA_DAFTAR = 10000
const DENDA_PER_HARI = 2000

type Komik struct {
	ID   int
	Nama string
	Stok int
}

type Member struct {
	ID   int
	Nama string
}

type Peminjaman struct {
	IDPinjam  int
	IDMember  int
	IDKomik   int
	LamaHari  int
	Terlambat int
	Denda     int
}

var komik [NMAX]Komik
var member [NMAX]Member
var pinjam [NMAX]Peminjaman

var nKomik, nMember, nPinjam int

// ================= KOMIK =================

func tambahKomik() {
	fmt.Print("ID Komik : ")
	fmt.Scan(&komik[nKomik].ID)

	fmt.Print("Nama Komik : ")
	fmt.Scan(&komik[nKomik].Nama)

	fmt.Print("Stok : ")
	fmt.Scan(&komik[nKomik].Stok)

	nKomik++
	fmt.Println("Data komik berhasil ditambahkan")
}

func tampilKomik() {
	fmt.Println("\n=== DATA KOMIK ===")
	for i := 0; i < nKomik; i++ {
		fmt.Println(
			komik[i].ID,
			komik[i].Nama,
			komik[i].Stok)
	}
}

func cariKomik(id int) int {
	for i := 0; i < nKomik; i++ {
		if komik[i].ID == id {
			return i
		}
	}
	return -1
}

func editKomik() {
	var id int

	fmt.Print("Masukkan ID Komik : ")
	fmt.Scan(&id)

	idx := cariKomik(id)

	if idx != -1 {
		fmt.Print("Nama Baru : ")
		fmt.Scan(&komik[idx].Nama)

		fmt.Print("Stok Baru : ")
		fmt.Scan(&komik[idx].Stok)

		fmt.Println("Data berhasil diubah")
	} else {
		fmt.Println("Komik tidak ditemukan")
	}
}

func hapusKomik() {
	var id int

	fmt.Print("ID Komik : ")
	fmt.Scan(&id)

	idx := cariKomik(id)

	if idx != -1 {
		for i := idx; i < nKomik-1; i++ {
			komik[i] = komik[i+1]
		}
		nKomik--
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Komik tidak ditemukan")
	}
}

// ================= MEMBER =================

func tambahMember() {
	fmt.Print("ID Member : ")
	fmt.Scan(&member[nMember].ID)

	fmt.Print("Nama Member : ")
	fmt.Scan(&member[nMember].Nama)

	nMember++

	fmt.Println("Member berhasil ditambahkan")
}

func tampilMember() {
	fmt.Println("\n=== DATA MEMBER ===")
	for i := 0; i < nMember; i++ {
		fmt.Println(member[i].ID, member[i].Nama)
	}
}

func cariMember(id int) int {
	for i := 0; i < nMember; i++ {
		if member[i].ID == id {
			return i
		}
	}
	return -1
}

func editMember() {
	var id int

	fmt.Print("ID Member : ")
	fmt.Scan(&id)

	idx := cariMember(id)

	if idx != -1 {
		fmt.Print("Nama Baru : ")
		fmt.Scan(&member[idx].Nama)

		fmt.Println("Data berhasil diubah")
	} else {
		fmt.Println("Member tidak ditemukan")
	}
}

func hapusMember() {
	var id int

	fmt.Print("ID Member : ")
	fmt.Scan(&id)

	idx := cariMember(id)

	if idx != -1 {
		for i := idx; i < nMember-1; i++ {
			member[i] = member[i+1]
		}
		nMember--
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Member tidak ditemukan")
	}
}

// ================= PEMINJAMAN =================

func tambahPeminjaman() {
	fmt.Print("ID Peminjaman : ")
	fmt.Scan(&pinjam[nPinjam].IDPinjam)

	fmt.Print("ID Member : ")
	fmt.Scan(&pinjam[nPinjam].IDMember)

	fmt.Print("ID Komik : ")
	fmt.Scan(&pinjam[nPinjam].IDKomik)

	fmt.Print("Lama Pinjam (hari): ")
	fmt.Scan(&pinjam[nPinjam].LamaHari)

	fmt.Print("Terlambat (hari): ")
	fmt.Scan(&pinjam[nPinjam].Terlambat)

	pinjam[nPinjam].Denda =
		pinjam[nPinjam].Terlambat * DENDA_PER_HARI

	nPinjam++

	fmt.Println("Peminjaman berhasil ditambahkan")
}

func tampilPeminjaman() {
	fmt.Println("\n=== DATA PEMINJAMAN ===")
	for i := 0; i < nPinjam; i++ {
		fmt.Println(
			pinjam[i].IDPinjam,
			pinjam[i].IDMember,
			pinjam[i].IDKomik,
			pinjam[i].Denda)
	}
}

func editPeminjaman() {
	var id int

	fmt.Print("ID Peminjaman : ")
	fmt.Scan(&id)

	for i := 0; i < nPinjam; i++ {
		if pinjam[i].IDPinjam == id {

			fmt.Print("Terlambat Baru : ")
			fmt.Scan(&pinjam[i].Terlambat)

			pinjam[i].Denda =
				pinjam[i].Terlambat * DENDA_PER_HARI

			fmt.Println("Data berhasil diubah")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func hapusPeminjaman() {
	var id int

	fmt.Print("ID Peminjaman : ")
	fmt.Scan(&id)

	for i := 0; i < nPinjam; i++ {
		if pinjam[i].IDPinjam == id {

			for j := i; j < nPinjam-1; j++ {
				pinjam[j] = pinjam[j+1]
			}

			nPinjam--
			fmt.Println("Data berhasil dihapus")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

// ================= SEARCH =================

func searchKomik() {
	var id int

	fmt.Print("ID Komik : ")
	fmt.Scan(&id)

	idx := cariKomik(id)

	if idx != -1 {
		fmt.Println(
			komik[idx].ID,
			komik[idx].Nama,
			komik[idx].Stok)
	} else {
		fmt.Println("Komik tidak ditemukan")
	}
}

func searchMember() {
	var id int

	fmt.Print("ID Member : ")
	fmt.Scan(&id)

	idx := cariMember(id)

	if idx != -1 {
		fmt.Println(
			member[idx].ID,
			member[idx].Nama)
	} else {
		fmt.Println("Member tidak ditemukan")
	}
}

// ================= LAPORAN =================

func totalDenda() {
	total := 0

	for i := 0; i < nPinjam; i++ {
		total += pinjam[i].Denda
	}

	fmt.Println("Total Denda :", total)
}

func totalPemasukan() {
	totalDaftar := nMember * BIAYA_DAFTAR

	totalDenda := 0

	for i := 0; i < nPinjam; i++ {
		totalDenda += pinjam[i].Denda
	}

	fmt.Println("Total Pemasukan =",
		totalDaftar+totalDenda)
}

// ================= MAIN =================

func main() {

	var menu int

	for menu != 15 {

		fmt.Println("\n===== APLIKASI PEMINJAMAN KOMIK =====")
		fmt.Println("1. Tambah Komik")
		fmt.Println("2. Edit Komik")
		fmt.Println("3. Hapus Komik")
		fmt.Println("4. Tambah Member")
		fmt.Println("5. Edit Member")
		fmt.Println("6. Hapus Member")
		fmt.Println("7. Tambah Peminjaman")
		fmt.Println("8. Edit Peminjaman")
		fmt.Println("9. Hapus Peminjaman")
		fmt.Println("10. Tampil Komik")
		fmt.Println("11. Tampil Member")
		fmt.Println("12. Cari Komik")
		fmt.Println("13. Cari Member")
		fmt.Println("14. Total Pemasukan")
		fmt.Println("15. Keluar")
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&menu)

		if menu == 1 {
			tambahKomik()
		} else if menu == 2 {
			editKomik()
		} else if menu == 3 {
			hapusKomik()
		} else if menu == 4 {
			tambahMember()
		} else if menu == 5 {
			editMember()
		} else if menu == 6 {
			hapusMember()
		} else if menu == 7 {
			tambahPeminjaman()
		} else if menu == 8 {
			editPeminjaman()
		} else if menu == 9 {
			hapusPeminjaman()
		} else if menu == 10 {
			tampilKomik()
		} else if menu == 11 {
			tampilMember()
		} else if menu == 12 {
			searchKomik()
		} else if menu == 13 {
			searchMember()
		} else if menu == 14 {
			totalPemasukan()
			totalDenda()
		}
	}
}
