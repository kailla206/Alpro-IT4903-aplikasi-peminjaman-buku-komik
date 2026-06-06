package main
import "fmt"
	const NMAX int = 100
	type komik struct {
		id, stok int
		nama string
	}
	type member struct {
		id   int
		nama string
	}

	type peminjaman struct {
		idPinjam int
		idKomik  int
		idMember int
		lamaHari int
	}
type tabKomik [NMAX]komik
type tabMember [NMAX]member
type tabPinjam [NMAX]peminjaman

func tambahKomik(A *tabKomik, n *int) {
	fmt.Print("ID Komik : ")
	fmt.Scan(&A[*n].id)

	fmt.Print("Nama Komik : ")
	fmt.Scan(&A[*n].nama)

	fmt.Print("Stok : ")
	fmt.Scan(&A[*n].stok)
	*n = *n + 1
}

func tampilKomik(A tabKomik, n int) {
	var i int

	fmt.Println("\nDATA KOMIK")
	for i = 0; i < n; i++ {
		fmt.Println(A[i].id, A[i].nama, A[i].stok)
	}
}

func cariKomik(A tabKomik, n, id int) int {
	var i int

	for i = 0; i < n; i++ {
		if A[i].id == id {
			return i
		}
	}
	return -1
}

func editKomik(A *tabKomik, n int) {
	var id, idx int

	fmt.Print("ID Komik yang diedit : ")
	fmt.Scan(&id)

	idx = cariKomik(*A, n, id)

	if idx != -1 {
		fmt.Print("Nama Baru : ")
		fmt.Scan(&A[idx].nama)

		fmt.Print("Stok Baru : ")
		fmt.Scan(&A[idx].stok)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusKomik(A *tabKomik, n *int) {
	var id, idx, i int

	fmt.Print("ID Komik yang dihapus : ")
	fmt.Scan(&id)

	idx = cariKomik(*A, *n, id)

	if idx != -1 {
		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n = *n - 1
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func tambahMember(B *tabMember, n *int) {
	fmt.Print("ID Member : ")
	fmt.Scan(&B[*n].id)

	fmt.Print("Nama Member : ")
	fmt.Scan(&B[*n].nama)

	*n = *n + 1
}

func tampilMember(B tabMember, n int) {
	var i int

	fmt.Println("\nDATA MEMBER")
	for i = 0; i < n; i++ {
		fmt.Println(B[i].id, B[i].nama)
	}
}

func cariMember(B tabMember, n, id int) int {
	var i int

	for i = 0; i < n; i++ {
		if B[i].id == id {
			return i
		}
	}
	return -1
}

func editMember(B *tabMember, n int) {
	var id, idx int

	fmt.Print("ID Member yang diedit : ")
	fmt.Scan(&id)

	idx = cariMember(*B, n, id)

	if idx != -1 {
		fmt.Print("Nama Baru : ")
		fmt.Scan(&B[idx].nama)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusMember(B *tabMember, n *int) {
	var id, idx, i int

	fmt.Print("ID Member yang dihapus : ")
	fmt.Scan(&id)

	idx = cariMember(*B, *n, id)

	if idx != -1 {
		for i = idx; i < *n-1; i++ {
			B[i] = B[i+1]
		}
		*n = *n - 1
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func tambahPinjam(C *tabPinjam, n *int) {
	fmt.Print("ID Peminjaman : ")
	fmt.Scan(&C[*n].idPinjam)

	fmt.Print("ID Komik : ")
	fmt.Scan(&C[*n].idKomik)

	fmt.Print("ID Member : ")
	fmt.Scan(&C[*n].idMember)

	fmt.Print("Lama Pinjam (hari) : ")
	fmt.Scan(&C[*n].lamaHari)

	*n = *n + 1
}

func tampilPinjam(C tabPinjam, n int) {
	var i int

	fmt.Println("\nDATA PEMINJAMAN")
	for i = 0; i < n; i++ {
		fmt.Println(C[i].idPinjam, C[i].idKomik, C[i].idMember, C[i].lamaHari)
	}
}

func cariPinjam(C tabPinjam, n, id int) int {
	var i int

	for i = 0; i < n; i++ {
		if C[i].idPinjam == id {
			return i
		}
	}
	return -1
}

func editPinjam(C *tabPinjam, n int) {
	var id, idx int

	fmt.Print("ID Peminjaman yang diedit : ")
	fmt.Scan(&id)

	idx = cariPinjam(*C, n, id)

	if idx != -1 {
		fmt.Print("ID Komik Baru : ")
		fmt.Scan(&C[idx].idKomik)

		fmt.Print("ID Member Baru : ")
		fmt.Scan(&C[idx].idMember)

		fmt.Print("Lama Hari Baru : ")
		fmt.Scan(&C[idx].lamaHari)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusPinjam(C *tabPinjam, n *int) {
	var id, idx, i int

	fmt.Print("ID Peminjaman yang dihapus : ")
	fmt.Scan(&id)

	idx = cariPinjam(*C, *n, id)

	if idx != -1 {
		for i = idx; i < *n-1; i++ {
			C[i] = C[i+1]
		}
		*n = *n - 1
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func hitungDenda() {
	var hari, denda int

	fmt.Print("Jumlah hari terlambat : ")
	fmt.Scan(&hari)

	denda = hari * 2000

	fmt.Println("Denda =", denda)
}

func totalUang() {
	var jumlahDaftar, biayaDaftar int
	var totalHariTerlambat int
	var total int

	fmt.Print("Jumlah Pendaftaran : ")
	fmt.Scan(&jumlahDaftar)

	fmt.Print("Biaya Pendaftaran : ")
	fmt.Scan(&biayaDaftar)

	fmt.Print("Total Hari Terlambat : ")
	fmt.Scan(&totalHariTerlambat)

	total = (jumlahDaftar * biayaDaftar) + (totalHariTerlambat * 2000)

	fmt.Println("Total Uang =", total)
}

func main() {
	var dataKomik tabKomik
	var dataMember tabMember
	var dataPinjam tabPinjam
	var nKomik, nMember, nPinjam, pilih, id, idx int

	for pilih != 0 {
		fmt.Println("\n===== APLIKASI PEMINJAMAN KOMIK =====")
		fmt.Println("1. Tambah Komik")
		fmt.Println("2. Edit Komik")
		fmt.Println("3. Hapus Komik")
		fmt.Println("4. Tampil Komik")
		fmt.Println("5. Tambah Member")
		fmt.Println("6. Edit Member")
		fmt.Println("7. Hapus Member")
		fmt.Println("8. Tampil Member")
		fmt.Println("9. Tambah Peminjaman")
		fmt.Println("10. Edit Peminjaman")
		fmt.Println("11. Hapus Peminjaman")
		fmt.Println("12. Tampil Peminjaman")
		fmt.Println("13. Cari Komik")
		fmt.Println("14. Cari Member")
		fmt.Println("15. Hitung Denda")
		fmt.Println("16. Hitung Total Uang")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahKomik(&dataKomik, &nKomik)
		} else if pilih == 2 {
			editKomik(&dataKomik, nKomik)
		} else if pilih == 3 {
			hapusKomik(&dataKomik, &nKomik)
		} else if pilih == 4 {
			tampilKomik(dataKomik, nKomik)
		} else if pilih == 5 {
			tambahMember(&dataMember, &nMember)
		} else if pilih == 6 {
			editMember(&dataMember, nMember)
		} else if pilih == 7 {
			hapusMember(&dataMember, &nMember)
		} else if pilih == 8 {
			tampilMember(dataMember, nMember)
		} else if pilih == 9 {
			tambahPinjam(&dataPinjam, &nPinjam)
		} else if pilih == 10 {
			editPinjam(&dataPinjam, nPinjam)
		} else if pilih == 11 {
			hapusPinjam(&dataPinjam, &nPinjam)
		} else if pilih == 12 {
			tampilPinjam(dataPinjam, nPinjam)
		} else if pilih == 13 {
			fmt.Print("Masukkan ID Komik : ")
			fmt.Scan(&id)

			idx = cariKomik(dataKomik, nKomik, id)

			if idx != -1 {
				fmt.Println("ID   :", dataKomik[idx].id)
				fmt.Println("Nama :", dataKomik[idx].nama)
				fmt.Println("Stok :", dataKomik[idx].stok)
			} else {
				fmt.Println("Komik tidak ditemukan")
			}

		} else if pilih == 14 {
			fmt.Print("Masukkan ID Member : ")
			fmt.Scan(&id)

			idx = cariMember(dataMember, nMember, id)

			if idx != -1 {
				fmt.Println("ID   :", dataMember[idx].id)
				fmt.Println("Nama :", dataMember[idx].nama)
			} else {
				fmt.Println("Member tidak ditemukan")
			}
		} else if pilih == 15 {
			hitungDenda()
		} else if pilih == 16 {
			totalUang()
		} else if pilih != 0 {
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}
