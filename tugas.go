// // Manajemen parkir dalam sebuah gedung perkantoran
package main

import "fmt"

const NMAX = 25

type kendaraan struct {
	status string
	nomorPlat, jenisKendaraan string
	jamMasuk, menitMasuk, jamKeluar, menitKeluar int
}


type tabMotor [NMAX] kendaraan
type tabMobil [NMAX] kendaraan

func main() {
	var dataMotor tabMotor
	var dataMobil tabMobil
	var nMotor int
	var nMobil int
	var selection int
	var nomorPlat string

	// set empty elemen untuk array pada motor & mobil
	loadStatusParkiran(&dataMotor, &dataMobil)

	for {
		// Tampilan Menu
		tampilanMenu()
		// User Activity List
		fmt.Print("Silahkan Pilih Menu: ")
		fmt.Scan(&selection)

		if selection == 1 { // tambah data kendaraan
			addVehicle(&dataMotor, &dataMobil, &nMotor, &nMobil)
		} else if selection == 2 { // menampilkan data parkiran
			tampilkanDataParkiran(dataMotor, dataMobil)
		} else if selection == 3 { // hapus kendaraan
			fmt.Println("Masukkan nomor plat dari kendaraan yang ingin di hapus: ")
			fmt.Scan(&nomorPlat)
			hapusDataKendaraan(&dataMotor, &dataMobil, &nMotor, &nMobil, nomorPlat)
		} else if selection == 4 { // edit kendaraan
			fmt.Println("Masukkan nomor plat dari kendaraan yang ingin di edit: ")
			fmt.Scan(&nomorPlat)
			editKendaraan(&dataMotor, &dataMobil, nomorPlat)
		} else if selection == 5 { // cari kendaraan
			fmt.Println("Masukkan nomor plat dari kendaraan yang ingin dicari: ")
			fmt.Scan(&nomorPlat)
			searchByLicensePlateNumber(dataMotor, dataMobil, nomorPlat)
		} else if selection == 6 { // exit program
			return 
		}
	}
}

// selection 1
func addVehicle(kMotor *tabMotor, kMobil *tabMobil, nMotor *int, nMobil *int) {

	if *nMotor >= NMAX {
		fmt.Println("Data kendaraan penuh")
		return
	}

	if *nMobil >= NMAX {
		fmt.Println("Data kendaraan penuh")
		return
	} 

	var k kendaraan

	// menampilkan ketersediaan slot ketika ingin menambahkan data kendaraan
	cekKetersediaanSlotParkir(*nMotor, *nMobil)

	fmt.Print("Masukkan jenis kendaraan (motor/mobil): ")
	fmt.Scan(&k.jenisKendaraan)

	if k.jenisKendaraan == "motor" {

	// add data motor
	fmt.Print("Masukkan plat nomor kendaraan: ")
	fmt.Scan(&k.nomorPlat)
	fmt.Print("Masukkan jam masuk kendaraan: ")
	fmt.Scan(&k.jamMasuk)
	fmt.Print("Masukkan menit masuk kendaraan: ")
	fmt.Scan(&k.menitMasuk)
	kMotor[*nMotor].nomorPlat = k.nomorPlat
	kMotor[*nMotor].jenisKendaraan = k.jenisKendaraan
	kMotor[*nMotor].jamMasuk = k.jamMasuk
	kMotor[*nMotor].menitMasuk = k.menitMasuk
	kMotor[*nMotor].status = "Terisi"
	*nMotor++

	} else if k.jenisKendaraan == "mobil" {

	// add data mobil
	fmt.Print("Masukkan plat nomor kendaraan: ")
	fmt.Scan(&k.nomorPlat)
	fmt.Print("Masukkan jam masuk kendaraan: ")
	fmt.Scan(&k.jamMasuk)
	fmt.Print("Masukkan menit masuk kendaraan: ")
	fmt.Scan(&k.menitMasuk)
	kMobil[*nMobil].nomorPlat = k.nomorPlat
	kMobil[*nMobil].jenisKendaraan = k.jenisKendaraan
	kMobil[*nMobil].jamMasuk = k.jamMasuk
	kMobil[*nMobil].menitMasuk = k.menitMasuk
	kMobil[*nMobil].status = "Terisi"
	*nMobil++
	} else {
		fmt.Println("Jenis kendaraan tidak valid")
	}

	fmt.Println()
	fmt.Println("Kendaraan berhasil ditambahkan")
}


// selection 2
func tampilkanDataParkiran(kMotor tabMotor, kMobil tabMobil) {
	// Cetak header tabel

	var jenis string
	fmt.Println("Masukkan jenis kendaraan yang ingin ditampilkan (motor/mobil): ")
	fmt.Scan(&jenis)

	if jenis == "motor" {
	fmt.Println("=====================================================")
	fmt.Println("*		Data Slot Parkiran Motor	  *")
	fmt.Println("=====================================================\n")
	fmt.Printf("%-10s %-10s %-15s %-15s %-10s %-10s\n", "Nomor", "Status", "Plat", "Jenis", "Jam Masuk", "Menit Masuk")
	fmt.Println()

	// Cetak data kendaraan
	for i := 0; i < NMAX; i++ {
		fmt.Printf("%-10d %-10s %-15s %-15s %-10d %-10d\n", i, kMotor[i].status, kMotor[i].nomorPlat, kMotor[i].jenisKendaraan, kMotor[i].jamMasuk, kMotor[i].menitMasuk)
	}
	fmt.Println("===================================================================")
	} else if jenis == "mobil" {
	fmt.Println("=====================================================")
	fmt.Println("*		Data Slot Parkiran Mobil	  *")
	fmt.Println("=====================================================")
	fmt.Printf("%-10s %-10s %-15s %-15s %-10s %-10s\n", "Nomor", "Status", "Plat", "Jenis", "Jam Masuk", "Menit Masuk")
	fmt.Println()

	// Cetak data kendaraan
	for i := 0; i < NMAX; i++ {
		fmt.Printf("%-10d %-10s %-15s %-15s %-10d %-10d\n", i, kMobil[i].status, kMobil[i].nomorPlat, kMobil[i].jenisKendaraan, kMobil[i].jamMasuk, kMobil[i].menitMasuk)
	}
	fmt.Println("===================================================================")
	} else {
		fmt.Println("Jenis yang dimasukkan tidak valid!")
	}
}

// selection 3
func hapusDataKendaraan(kMotor *tabMotor, kMobil *tabMobil, nMotor *int, nMobil *int, nomorPlat string) {

    jenis, index := searchByLicensePlateNumber(*kMotor, *kMobil, nomorPlat)
    var tarifParkir int

    if index == -1 {
        fmt.Println("Data kendaraan tidak ditemukan!")
        return
    }

    if jenis == "motor" {
        // Hapus data
        for i := index; i < *nMotor-1; i++ {
            kMotor[i] = kMotor[i+1]
        }

        tarifParkir = hitungTarif(jenis, index, *kMotor, *kMobil)
        kMotor[*nMotor-1] = kendaraan{}
        *nMotor--
        fmt.Printf("Tarif parkir: Rp %d\n", tarifParkir)
        fmt.Println("Data kendaraan motor berhasil dihapus.")
    } else if jenis == "mobil" {
        // Hapus data
        for i := index; i < *nMobil-1; i++ {
            kMobil[i] = kMobil[i+1]
        }

        tarifParkir = hitungTarif(jenis, index, *kMotor, *kMobil)
        kMobil[*nMobil-1] = kendaraan{}
        *nMobil--
        fmt.Printf("Tarif parkir: Rp %d\n", tarifParkir)
        fmt.Println("Data kendaraan mobil berhasil dihapus.")
    }
}

func hitungTarif(jenis string, indexKendaraan int, kMotor tabMotor, kMobil tabMobil) int {
	var tarifPerJam int
	var totalWaktu, totalJam int

	// Tentukan tarif per jam berdasarkan jenis kendaraan
	if jenis == "motor" {
		tarifPerJam = 2000
	} else if jenis == "mobil" {
		tarifPerJam = 5000
	}

	// Input jam keluar kendaraan
	var jamKeluar, menitKeluar int
	fmt.Println("Masukkan jam keluar kendaraan: ")
	fmt.Scan(&jamKeluar)
	fmt.Println("Masukkan menit keluar kendaraan: ")
	fmt.Scan(&menitKeluar)

	// Hitung total waktu parkir dalam menit
	if jenis == "motor" {
		totalMenit := (jamKeluar*60 + menitKeluar) - (kMotor[indexKendaraan].jamMasuk*60 + kMotor[indexKendaraan].menitMasuk)
		totalJam = totalMenit / 60
		if totalMenit%60 > 0 {
			totalJam++
		}
	} else if jenis == "mobil" {
		totalMenit := (jamKeluar*60 + menitKeluar) - (kMobil[indexKendaraan].jamMasuk*60 + kMobil[indexKendaraan].menitMasuk)
		totalJam = totalMenit / 60
		if totalMenit%60 > 0 {
			totalJam++
		}
	}

	// Hitung total tarif
	totalWaktu = totalJam * tarifPerJam

	return totalWaktu
}


// selection 4
func editKendaraan(kMotor *tabMotor, kMobil *tabMobil, nomorPlat string) {
	var userInputString string
	var userInputInteger int

	jenis, index := searchByLicensePlateNumber(*kMotor, *kMobil, nomorPlat)
	if index != -1 {
		fmt.Println("Masukkan nomor plat, masukkan (-) jika tidak ingin merubah data")
		fmt.Scan(&userInputString)
		if userInputString != "-" {
			if jenis == "motor" {
				kMotor[index].nomorPlat = userInputString
			} else if jenis == "mobil" {
				kMobil[index].nomorPlat = userInputString
			}
		}
		fmt.Println("Masukkan jam masuk kendaraan, masukkan '0' jika tidak ingin merubah data")
		fmt.Scan(&userInputInteger)
		if userInputInteger != 0 {
			if jenis == "motor" {
				kMotor[index].jamMasuk = userInputInteger
			} else if jenis == "mobil" {
				kMobil[index].jamMasuk = userInputInteger
			}
		}
		fmt.Println("Masukkan menit masuk kendaraan, masukkan '0' jika tidak ingin merubah data")
		fmt.Scan(&userInputInteger)
		if userInputInteger != 0 {
			if jenis == "motor" {
				kMotor[index].menitMasuk = userInputInteger
			} else if jenis == "mobil" {
				kMobil[index].menitMasuk = userInputInteger
			}
		}
	}
}

func searchByLicensePlateNumber(kMotor tabMotor, kMobil tabMobil, nomorPlat string) (string, int) {
	for i := 0; i < NMAX; i++ {
		if kMotor[i].nomorPlat == nomorPlat {
			// me return jenis kendaraan dan index ketika ditemukan
			return "motor", i
		}
		if kMobil[i].nomorPlat == nomorPlat {
			return "mobil", i
		}
	}
	return "", -1
}

func loadStatusParkiran(kMotor *tabMotor, kMobil *tabMobil) {
	// set default di awal untuk semua data kendaraan pada slot parkiran dengan status dan data kosong
	for i := 0; i < NMAX; i++ {
		//motor
		kMotor[i].status = "Kosong"
		kMotor[i].nomorPlat = "-"
		kMotor[i].jenisKendaraan = "-"
		kMotor[i].jamMasuk = 0
		kMotor[i].menitMasuk = 0
		//mobil
		kMobil[i].status = "Kosong"
		kMobil[i].nomorPlat = "-"
		kMobil[i].jenisKendaraan = "-"
		kMobil[i].jamMasuk = 0
		kMobil[i].menitMasuk = 0
	}
}

func cekKetersediaanSlotParkir(nMotor int, nMobil int) {
	jumlahKetersediaanSlotMotor := NMAX - nMotor
	jumlahKetersediaanSlotMobil := NMAX - nMobil
	fmt.Println("=================================")
	fmt.Println("Jumlah ketersediaan slot motor: ", jumlahKetersediaanSlotMotor)
	fmt.Println("Jumlah ketersediaan slot mobil: ", jumlahKetersediaanSlotMobil)
	fmt.Println("=================================")
}



func tampilanMenu() {
		fmt.Println()
		fmt.Println("==========================")
		fmt.Println("       PILIH MENU         ")
		fmt.Println("==========================")
		fmt.Println("1. Tambah Kendaraan")
		fmt.Println("2. Tampilkan Data Kendaraan")
		fmt.Println("3. Hapus Kendaraan")
		fmt.Println("4. Edit Kendaraan")
		fmt.Println("5. Cari Data Kendaraan")
		fmt.Println("6. Exit")
		fmt.Println("==========================")
		fmt.Println()
}