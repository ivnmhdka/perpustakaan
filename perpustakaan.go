package main

import (
	"fmt"
	"os"
	"strconv"
)

const NMAX int = 100
const dendaFix float64 = 5000

type Buku struct {
	idBuku, judul, kategori string
	statusPinjam            bool
	totalPeminjam           int
}

type Pengguna struct {
	nama, username, password string
	denda                    float64
}

type PinjamBuku struct {
	idPeminjaman, idBuku, pengguna string
	tanggalPinjam, tanggalKembali  string
	tarif                          float64
}

type arrBuku [NMAX]Buku
type arrPengguna [NMAX]Pengguna
type arrPeminjaman [NMAX]PinjamBuku

var listBuku arrBuku
var listPengguna arrPengguna
var listPeminjaman arrPeminjaman

var penggunaAktif Pengguna
var jumlahBuku int = 0
var jumlahPengguna int = 0
var jumlahPeminjaman int = 0
var totalTarif float64 = 0
var totalDenda float64 = 0

func main() {
	menu()
}

func menu() {
	fmt.Println("------------------------------------------------")
	fmt.Println("*           Aplikasi Perpustakaan          *")
	fmt.Println("* Anggota 1 *") //Isi dengan nama dan NIM anggota
	fmt.Println("* Anggota 2 *") //Isi dengan nama dan NIM anggota
	fmt.Println("* Anggota 3 *") //Isi dengan nama dan NIM anggota
	fmt.Println("* Menu *")
	fmt.Println("1. Registrasi")
	fmt.Println("2. Login")
	fmt.Println("3. Admin")
	fmt.Println("4. Keluar")

	var username, password string

	var pilihan int
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		registrasi(&listPengguna, &jumlahPengguna)
	case 2:
		fmt.Println("*   Halaman Login   *")
		fmt.Print("Username: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		fmt.Scan(&password)
		login(username, password, listPengguna, jumlahPengguna)
	case 3:
		fmt.Println("")
		menuAdmin()
	case 4:
		fmt.Println("Terima Kasih telah menggunakan aplikasi kami!")
		os.Exit(0)
		return
	default:
		fmt.Println("Pilihan Anda tidak valid!")
		fmt.Println("")
		menu()
	}
}

func registrasi(A *arrPengguna, n *int) {
	if *n >= NMAX {
		fmt.Println("Jumlah Pengguna Penuh")
		fmt.Println("")
		menu()
		return
	}
	fmt.Println("Masukkan Data Diri:")
	fmt.Print("Nama: ")
	fmt.Scan(&A[*n].nama)
	fmt.Print("Username: ")
	fmt.Scan(&A[*n].username)
	fmt.Print("Password: ")
	fmt.Scan(&A[*n].password)
	A[*n].denda = 0.0
	*n++
	fmt.Println("Registrasi Berhasil!")
	fmt.Println("")
	menu()
	return
}

func login(user string, pw string, A arrPengguna, n int) {
	var success bool = false
	var i int = 0
	for i < n && !success {
		if user == A[i].username && pw == A[i].password {
			penggunaAktif = A[i]
			fmt.Println("Login Pengguna Berhasil")
			success = true
			fmt.Println("")
			menuPengguna()
		}
		i++
	}
	if !success {
		fmt.Println("Login Pengguna Gagal")
		fmt.Println("")
		menu()
		return
	}
}

func menuAdmin() {
	fmt.Println("*   Menu Admin   *")
	fmt.Println("1. Cek Buku Terurut Berdasarkan Keyword")
	fmt.Println("2. Tambah Buku")
	fmt.Println("3. Edit Buku")
	fmt.Println("4. Hapus Buku")
	fmt.Println("5. Cek List Peminjaman")
	fmt.Println("6. Total Denda")
	fmt.Println("7. Keluar dari Halaman Admin")

	var pilihan int
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		cekListBuku(&listBuku, jumlahBuku)
		fmt.Println("")
	case 2:
		tambahBuku(&listBuku, &jumlahBuku)
		fmt.Println("")
	case 3:
		editBuku(&listBuku, jumlahBuku)
		fmt.Println("")
	case 4:
		hapusBuku(&listBuku, &jumlahBuku)
		fmt.Println("")
	case 5:
		cekPeminjaman(listPeminjaman, jumlahPeminjaman)
		fmt.Println("")
	case 6:
		cekPenghasilan()
		fmt.Println("")
	case 7:
		fmt.Println("Kembali ke Halaman Utama")
		fmt.Println("")
		menu()
		return
	default:
		fmt.Println("Pilihan tidak valid!")
		fmt.Println("")
		menuAdmin()
	}
}

func menuPengguna() {
	fmt.Println("* Menu Pengguna *")
	fmt.Println("Selamat Datang,", penggunaAktif.nama)
	fmt.Println("1. Cek Buku Berdasarkan Keyword")
	fmt.Println("2. Pinjam Buku")
	fmt.Println("3. Kembalikan Buku")
	fmt.Println("4. Total Denda Pengguna")
	fmt.Println("5. Buku Terfavorit")
	fmt.Println("6. Logout")

	var pilihan int
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		cekListBuku(&listBuku, jumlahBuku)
	case 2:
		pinjamBuku(&listBuku, jumlahBuku, &listPeminjaman, &jumlahPeminjaman)
	case 3:
		kembaliBuku(&listBuku, jumlahBuku, &listPeminjaman, &jumlahPeminjaman)
	case 4:
		totalDendaUser(listPengguna, jumlahPengguna)
	case 5:
		bukuFavorit(listBuku, jumlahBuku)
	case 6:
		fmt.Println("Anda Telah Logout. Kembali ke Halaman Utama")
		fmt.Println("")
		menu()
		return
	default:
		fmt.Println("Pilihan tidak valid!")
		fmt.Println("")
		menuPengguna()
	}
}

func cekListBuku(A *arrBuku, n int) {
	var pilihan, p1 string
	fmt.Println("Urut Berdasarkan Judul / Kategori? (Input J/K)")
	fmt.Scan(&pilihan)
	if pilihan == "J" || pilihan == "j" {
		sortJudul(A, n)
		p1 = "Judul"
	} else if pilihan == "K" || pilihan == "k" {
		sortKategori(A, n)
		p1 = "Kategori"
	}
	for i := 0; i < n; i++ {
		fmt.Println("List Buku Berdasarkan ", p1, " :")
		fmt.Println("Judul : ", A[i].judul)
		fmt.Println("Kategori : ", A[i].kategori)
		fmt.Println("Status Peminjaman: ", A[i].statusPinjam)
		fmt.Println("Total Peminjam: ", A[i].totalPeminjam)
		fmt.Println("")
	}
	fmt.Println("")
	menuAdmin()
}

func sortJudul(A *arrBuku, n int) {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if A[i].judul < A[j].judul {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
}

func sortKategori(A *arrBuku, n int) {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (A[i].kategori > A[j].kategori) || (A[i].kategori == A[j].kategori && A[i].judul > A[j].judul) {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
}

func tambahBuku(A *arrBuku, n *int) {
	fmt.Print("Masukkan Judul Buku: ")
	fmt.Scan(&A[*n].judul)
	fmt.Print("Masukkan Kategori Buku: ")
	fmt.Scan(&A[*n].kategori)
	A[*n].idBuku = "Book" + strconv.Itoa(*n+1)
	A[*n].statusPinjam = false
	A[*n].totalPeminjam = 0
	*n++
}

func editBuku(A *arrBuku, n int) {
	var ID string
	fmt.Print("Cari data buku yang ingin diganti: (Input ID Buku) ")
	fmt.Scan(&ID)
	var i int = 0
	var ketemu bool = false
	for i < n && !ketemu {
		if A[i].idBuku == ID && !A[i].statusPinjam {
			ketemu = true
			fmt.Print("Masukkan Judul Buku: ")
			fmt.Scan(&A[i].judul)
			fmt.Print("Masukkan Kategori Buku: ")
			fmt.Scan(&A[i].kategori)
			fmt.Println("")
			menuAdmin()
			return
		}
		if A[i].statusPinjam {
			ketemu = true
			fmt.Println("Buku sedang dipinjam, data tidak bisa diganti!")
			fmt.Println("")
			menuAdmin()
			return
		}
		i++
	}
	if !ketemu {
		fmt.Println("Buku yang dicari tidak ditemukan")
		fmt.Println("")
		menuAdmin()
		return
	}
}

func hapusBuku(A *arrBuku, n *int) {
	var ID string
	fmt.Print("Cari buku yang ingin Anda hapus: (Input ID Buku) ")
	fmt.Scan(&ID)
	var i int = 0
	var ketemu bool = false
	for i < *n && !ketemu {
		if A[i].idBuku == ID && !A[i].statusPinjam {
			ketemu = true
			fmt.Println("Buku telah dihapus!")
			for j := i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			*n--
		}
		if A[i].statusPinjam {
			ketemu = true
			fmt.Println("Buku sedang dipinjam, data tidak bisa dihapus!")
		}
		i++
	}
	if !ketemu {
		fmt.Println("Buku yang dicari tidak ditemukan")
	}
	fmt.Println("")
	menuAdmin()
	return
}

func cekPeminjaman(A arrPeminjaman, n int) {
	fmt.Println("List Peminjaman Buku")
	for i := 0; i < n; i++ {
		fmt.Println("ID Peminjaman: ", A[i].idPeminjaman)
		fmt.Println("ID Buku: ", A[i].idBuku)
		fmt.Println("Nama Peminjam: ", A[i].pengguna)
		fmt.Println("Tanggal Pinjam: ", A[i].tanggalPinjam)
		fmt.Println("Tanggal Kembali: ", A[i].tanggalKembali)
		fmt.Println("Tarif Peminjaman: ", A[i].tarif)
		fmt.Println("")
	}
	fmt.Println("")
	menuAdmin()
	return
}

func cekPenghasilan() {
	fmt.Println("Penghasilan dari Peminjaman Buku adalah Rp. ", totalTarif)
	fmt.Println("Penghasilan dari Denda adalah Rp. ", totalDenda)
	fmt.Println("Total Penghasilan adalah Rp. ", totalTarif+totalDenda)
	fmt.Println("")
	menuAdmin()
	return
}

func pinjamBuku(A *arrBuku, n int, B *arrPeminjaman, o *int) {
	var ID string
	fmt.Println("Masukkan ID Buku yang ingin dipinjam:")
	fmt.Scan(&ID)
	var index int = searchBuku(*A, n, ID)
	if index == -1 {
		fmt.Println("ID Buku tidak ditemukan!")
		fmt.Println("")
		menuPengguna()
		return
	} else {
		B[*o].idBuku = A[index].idBuku
		B[*o].idPeminjaman = "Pinjam" + strconv.Itoa(*o+1)
		B[*o].pengguna = penggunaAktif.nama
		fmt.Print("Masukkan Tanggal Peminjaman (Format DD-MM-YYYY): ")
		fmt.Scan(&B[*o].tanggalPinjam)
		fmt.Print("Masukkan Tanggal Kembali (Format DD-MM-YYYY): ")
		fmt.Scan(&B[*o].tanggalKembali)
		var jumlahHari int = countHari(B[*o].tanggalPinjam, B[*o].tanggalKembali)
		B[*o].tarif = float64(jumlahHari) * 5000
		A[index].statusPinjam = true
		A[index].totalPeminjam++
		*o++
		fmt.Println("")
		menuPengguna()
		return
	}
}

func kembaliBuku(A *arrBuku, n int, B *arrPeminjaman, o *int) {
	var ID, tanggalPengembalian string
	var dendaPeminjaman float64
	fmt.Println("Masukkan ID Peminjaman: ")
	fmt.Scan(&ID)
	var index int = searchPeminjamanBuku(*B, *o, ID)
	if index == -1 {
		fmt.Println("ID Peminjaman tidak ditemukan!")
		fmt.Println("")
		menuPengguna()
		return
	} else {
		fmt.Print("Masukkan Tanggal Pengembalian: (Format DD-MM-YYYY)")
		fmt.Scan(&tanggalPengembalian)
		var jumlahHariPeminjaman int = countHari(B[index].tanggalPinjam, tanggalPengembalian)
		var jumlahSeharusnya int = countHari(B[index].tanggalPinjam, B[index].tanggalKembali)
		var selisih int = jumlahHariPeminjaman - jumlahSeharusnya
		if selisih > 0 {
			dendaPeminjaman = float64(selisih) * dendaFix
			totalDenda += dendaPeminjaman
			fmt.Println("Anda telat mengembalikan selama ", selisih, " hari dan memiliki denda sebesar Rp. ", dendaPeminjaman)
			fmt.Println("")
		}
		var idxBuku int = searchBuku(*A, n, B[index].idBuku)
		A[idxBuku].statusPinjam = false
		for i := index; i < *o; i++ {
			B[i] = B[i+1]
		}
		menuPengguna()
		return
	}
}

func totalDendaUser(A arrPengguna, n int) {
	for i := 0; i < n; i++ {
		totalDenda += A[i].denda
	}
}

func bukuFavorit(A arrBuku, n int) {

}

func searchBuku(A arrBuku, n int, x string) int {
	var idx int = -1
	var i int = 0
	for i < n && idx == -1 {
		if A[i].idBuku == x {
			idx = i
		}
	}
	return idx
}

func searchPeminjamanBuku(A arrPeminjaman, n int, x string) int {
	var idx int = -1
	var i int = 0
	for i < n && idx == -1 {
		if A[i].idBuku == x {
			idx = i
		}
	}
	return idx
}

func countHari(x, y string) int {
	var tP, bP, yP, tK, bK, yK string
	tP = x[0:2]
	bP = x[3:5]
	yP = x[6:]
	tK = y[0:2]
	bK = y[3:5]
	yK = y[6:]
	var tPint, bPint, yPint, tKint, bKint, yBint int
	tPint, _ = strconv.Atoi(tP)
	bPint, _ = strconv.Atoi(bP)
	yPint, _ = strconv.Atoi(yP)
	tKint, _ = strconv.Atoi(tK)
	bKint, _ = strconv.Atoi(bK)
	yBint, _ = strconv.Atoi(yK)
	var hari int
	hari = (yBint-yPint)*360 + (bKint-bPint)*30 + (tKint - tPint)

	return hari
}
