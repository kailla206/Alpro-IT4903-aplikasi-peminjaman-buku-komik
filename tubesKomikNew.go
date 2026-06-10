package main

import "fmt"

type Komik struct {
	ID    string
	Judul string
	Stok  int
}

type Member struct {
	ID   string
	Nama string
}

type Peminjaman struct {
	ID         string
	IDKomik    string
	IDMember   string
	TanggalPin int // hari ke-
	TanggalKem int // hari ke- (0 = belum dikembalikan)
	Denda      float64
}

// ===================== DATA GLOBAL =====================

var daftarKomik []Komik
var daftarMember []Member
var daftarPinjam []Peminjaman
var dendaPerHari float64 = 1000
var biayaDaftar float64 = 5000

// ===================== KOMIK =====================

func tambahKomik(id, judul string, stok int) {
	daftarKomik = append(daftarKomik, Komik{id, judul, stok})
	fmt.Println("Komik berhasil ditambahkan.")
}

func ubahKomik(id, judulBaru string, stokBaru int) {
	for i := 0; i < len(daftarKomik); i++ {
		if daftarKomik[i].ID == id {
			daftarKomik[i].Judul = judulBaru
			daftarKomik[i].Stok = stokBaru
			fmt.Println("Data komik berhasil diubah.")
			return
		}
	}
	fmt.Println("Komik dengan ID tersebut tidak ditemukan.")
}

func hapusKomik(id string) {
	for i := 0; i < len(daftarKomik); i++ {
		if daftarKomik[i].ID == id {
			daftarKomik = append(daftarKomik[:i], daftarKomik[i+1:]...)
			fmt.Println("Komik berhasil dihapus.")
			return
		}
	}
	fmt.Println("Komik dengan ID tersebut tidak ditemukan.")
}

func tampilKomik() {
	fmt.Println("\n===== Daftar Komik =====")
	if len(daftarKomik) == 0 {
		fmt.Println("(Belum ada data komik)")
		return
	}
	for _, k := range daftarKomik {
		fmt.Printf("ID: %s | Judul: %s | Stok: %d\n", k.ID, k.Judul, k.Stok)
	}
}

func cariKomik(id string) {
	for _, k := range daftarKomik {
		if k.ID == id {
			fmt.Printf("Ditemukan → ID: %s | Judul: %s | Stok: %d\n", k.ID, k.Judul, k.Stok)
			return
		}
	}
	fmt.Println("Komik tidak ditemukan.")
}

// ===================== MEMBER =====================

func tambahMember(id, nama string) {
	daftarMember = append(daftarMember, Member{id, nama})
	fmt.Println("Member berhasil ditambahkan.")
}

func ubahMember(id, namaBaru string) {
	for i := 0; i < len(daftarMember); i++ {
		if daftarMember[i].ID == id {
			daftarMember[i].Nama = namaBaru
			fmt.Println("Data member berhasil diubah.")
			return
		}
	}
	fmt.Println("Member dengan ID tersebut tidak ditemukan.")
}

func hapusMember(id string) {
	for i := 0; i < len(daftarMember); i++ {
		if daftarMember[i].ID == id {
			daftarMember = append(daftarMember[:i], daftarMember[i+1:]...)
			fmt.Println("Member berhasil dihapus.")
			return
		}
	}
	fmt.Println("Member dengan ID tersebut tidak ditemukan.")
}

func tampilMember() {
	fmt.Println("\n===== Daftar Member =====")
	if len(daftarMember) == 0 {
		fmt.Println("(Belum ada data member)")
		return
	}
	for _, m := range daftarMember {
		fmt.Printf("ID: %s | Nama: %s\n", m.ID, m.Nama)
	}
}

func cariMember(id string) {
	for _, m := range daftarMember {
		if m.ID == id {
			fmt.Printf("Ditemukan → ID: %s | Nama: %s\n", m.ID, m.Nama)
			return
		}
	}
	fmt.Println("Member tidak ditemukan.")
}

// ===================== PEMINJAMAN =====================

func tambahPeminjaman(idPinjam, idKomik, idMember string, hariPinjam int) {
	daftarPinjam = append(daftarPinjam, Peminjaman{
		ID:         idPinjam,
		IDKomik:    idKomik,
		IDMember:   idMember,
		TanggalPin: hariPinjam,
		TanggalKem: 0,
		Denda:      0,
	})
	fmt.Println("Peminjaman berhasil dicatat.")
}

func ubahPeminjaman(idPinjam, idKomikBaru, idMemberBaru string, hariPinjamBaru int) {
	for i := 0; i < len(daftarPinjam); i++ {
		if daftarPinjam[i].ID == idPinjam {
			daftarPinjam[i].IDKomik = idKomikBaru
			daftarPinjam[i].IDMember = idMemberBaru
			daftarPinjam[i].TanggalPin = hariPinjamBaru
			fmt.Println("Data peminjaman berhasil diubah.")
			return
		}
	}
	fmt.Println("Data peminjaman tidak ditemukan.")
}

func hapusPeminjaman(idPinjam string) {
	for i := 0; i < len(daftarPinjam); i++ {
		if daftarPinjam[i].ID == idPinjam {
			daftarPinjam = append(daftarPinjam[:i], daftarPinjam[i+1:]...)
			fmt.Println("Data peminjaman berhasil dihapus.")
			return
		}
	}
	fmt.Println("Data peminjaman tidak ditemukan.")
}

func tampilPeminjaman() {
	fmt.Println("\n===== Daftar Peminjaman =====")
	if len(daftarPinjam) == 0 {
		fmt.Println("(Belum ada data peminjaman)")
		return
	}
	for _, p := range daftarPinjam {
		status := "Belum dikembalikan"
		if p.TanggalKem > 0 {
			status = fmt.Sprintf("Dikembalikan hari ke-%d", p.TanggalKem)
		}
		fmt.Printf("ID: %s | Komik: %s | Member: %s | Pinjam: hari ke-%d | Status: %s | Denda: Rp%.0f\n",
			p.ID, p.IDKomik, p.IDMember, p.TanggalPin, status, p.Denda)
	}
}

// ===================== DENDA & KEUANGAN =====================

// Batas peminjaman = 7 hari, denda Rp1.000/hari keterlambatan
func hitungDenda(idPinjam string, hariKembali int) {
	for i := 0; i < len(daftarPinjam); i++ {
		if daftarPinjam[i].ID == idPinjam {
			batasHari := daftarPinjam[i].TanggalPin + 7
			terlambat := hariKembali - batasHari
			denda := 0.0
			if terlambat > 0 {
				denda = float64(terlambat) * dendaPerHari
			}
			daftarPinjam[i].TanggalKem = hariKembali
			daftarPinjam[i].Denda = denda
			if terlambat > 0 {
				fmt.Printf("⚠ Terlambat %d hari. Denda: Rp%.0f\n", terlambat, denda)
			} else {
				fmt.Println("Dikembalikan tepat waktu. Tidak ada denda.")
			}
			return
		}
	}
	fmt.Println("Data peminjaman tidak ditemukan.")
}

func hitungTotalPendapatan() {
	totalDenda := 0.0
	for _, p := range daftarPinjam {
		totalDenda += p.Denda
	}
	totalDaftar := float64(len(daftarMember)) * biayaDaftar
	total := totalDenda + totalDaftar
	fmt.Printf("\n===== Total Pendapatan =====\n")
	fmt.Printf("Dari pendaftaran (%d member × Rp%.0f): Rp%.0f\n", len(daftarMember), biayaDaftar, totalDaftar)
	fmt.Printf("Dari denda                           : Rp%.0f\n", totalDenda)
	fmt.Printf("Total                                : Rp%.0f\n", total)
}

// ===================== MENU =====================

func tampilMenu() {
	fmt.Println("\n==============================")
	fmt.Println("  APLIKASI PEMINJAMAN KOMIK  ")
	fmt.Println("==============================")
	fmt.Println("--- Data Komik ---")
	fmt.Println(" 1. Tambah Komik")
	fmt.Println(" 2. Ubah Komik")
	fmt.Println(" 3. Hapus Komik")
	fmt.Println(" 4. Tampilkan Komik")
	fmt.Println(" 5. Cari Komik by ID")
	fmt.Println("--- Data Member ---")
	fmt.Println(" 6. Tambah Member")
	fmt.Println(" 7. Ubah Member")
	fmt.Println(" 8. Hapus Member")
	fmt.Println(" 9. Tampilkan Member")
	fmt.Println("10. Cari Member by ID")
	fmt.Println("--- Peminjaman ---")
	fmt.Println("11. Tambah Peminjaman")
	fmt.Println("12. Ubah Peminjaman")
	fmt.Println("13. Hapus Peminjaman")
	fmt.Println("14. Tampilkan Peminjaman")
	fmt.Println("15. Hitung Denda")
	fmt.Println("16. Hitung Total Pendapatan")
	fmt.Println(" 0. Keluar")
	fmt.Print("Pilihan: ")
}

// ===================== MAIN =====================

func main() {
	pilihan := -1
	for pilihan != 0 {
		tampilMenu()
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			var id, judul string
			var stok int
			fmt.Print("ID Komik  : ")
			fmt.Scan(&id)
			fmt.Print("Judul     : ")
			fmt.Scan(&judul)
			fmt.Print("Stok      : ")
			fmt.Scan(&stok)
			tambahKomik(id, judul, stok)

		} else if pilihan == 2 {
			var id, judul string
			var stok int
			fmt.Print("ID Komik  : ")
			fmt.Scan(&id)
			fmt.Print("Judul Baru: ")
			fmt.Scan(&judul)
			fmt.Print("Stok Baru : ")
			fmt.Scan(&stok)
			ubahKomik(id, judul, stok)

		} else if pilihan == 3 {
			var id string
			fmt.Print("ID Komik: ")
			fmt.Scan(&id)
			hapusKomik(id)

		} else if pilihan == 4 {
			tampilKomik()

		} else if pilihan == 5 {
			var id string
			fmt.Print("ID Komik: ")
			fmt.Scan(&id)
			cariKomik(id)

		} else if pilihan == 6 {
			var id, nama string
			fmt.Print("ID Member: ")
			fmt.Scan(&id)
			fmt.Print("Nama     : ")
			fmt.Scan(&nama)
			tambahMember(id, nama)

		} else if pilihan == 7 {
			var id, nama string
			fmt.Print("ID Member : ")
			fmt.Scan(&id)
			fmt.Print("Nama Baru : ")
			fmt.Scan(&nama)
			ubahMember(id, nama)

		} else if pilihan == 8 {
			var id string
			fmt.Print("ID Member: ")
			fmt.Scan(&id)
			hapusMember(id)

		} else if pilihan == 9 {
			tampilMember()

		} else if pilihan == 10 {
			var id string
			fmt.Print("ID Member: ")
			fmt.Scan(&id)
			cariMember(id)

		} else if pilihan == 11 {
			var idPinjam, idKomik, idMember string
			var hariPinjam int
			fmt.Print("ID Peminjaman: ")
			fmt.Scan(&idPinjam)
			fmt.Print("ID Komik     : ")
			fmt.Scan(&idKomik)
			fmt.Print("ID Member    : ")
			fmt.Scan(&idMember)
			fmt.Print("Hari Pinjam  : ")
			fmt.Scan(&hariPinjam)
			tambahPeminjaman(idPinjam, idKomik, idMember, hariPinjam)

		} else if pilihan == 12 {
			var idPinjam, idKomik, idMember string
			var hariPinjam int
			fmt.Print("ID Peminjaman  : ")
			fmt.Scan(&idPinjam)
			fmt.Print("ID Komik Baru  : ")
			fmt.Scan(&idKomik)
			fmt.Print("ID Member Baru : ")
			fmt.Scan(&idMember)
			fmt.Print("Hari Pinjam    : ")
			fmt.Scan(&hariPinjam)
			ubahPeminjaman(idPinjam, idKomik, idMember, hariPinjam)

		} else if pilihan == 13 {
			var id string
			fmt.Print("ID Peminjaman: ")
			fmt.Scan(&id)
			hapusPeminjaman(id)

		} else if pilihan == 14 {
			tampilPeminjaman()

		} else if pilihan == 15 {
			var id string
			var hariKembali int
			fmt.Print("ID Peminjaman  : ")
			fmt.Scan(&id)
			fmt.Print("Hari Kembali   : ")
			fmt.Scan(&hariKembali)
			hitungDenda(id, hariKembali)

		} else if pilihan == 16 {
			hitungTotalPendapatan()

		} else if pilihan == 0 {
			fmt.Println("Terima kasih! Sampai jumpa.")

		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
