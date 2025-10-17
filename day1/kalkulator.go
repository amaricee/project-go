package main

import "fmt"

func tambah(a, b float64) float64 {
    return a + b
}

func kurang(a, b float64) float64 {
    return a - b
}

func kali(a, b float64) float64 {
    return a * b
}

func bagi(a, b float64) (float64, string) {
    if b == 0 {
        return 0, "Error: Tidak bisa membagi dengan nol"
    }
    return a / b, "Berhasil"
}

func main() {
    var pilihan int
    var x, y float64

    fmt.Println("=== Mini Kalkulator Go ===")
    fmt.Println("1. Tambah")
    fmt.Println("2. Kurang")
    fmt.Println("3. Kali")
    fmt.Println("4. Bagi")
    fmt.Print("Pilih operasi (1-4): ")
    fmt.Scan(&pilihan)

    fmt.Print("Masukkan angka pertama: ")
    fmt.Scan(&x)
    fmt.Print("Masukkan angka kedua: ")
    fmt.Scan(&y)

    switch pilihan {
    case 1:
        fmt.Println("Hasil:", tambah(x, y))
    case 2:
        fmt.Println("Hasil:", kurang(x, y))
    case 3:
        fmt.Println("Hasil:", kali(x, y))
    case 4:
        hasil, status := bagi(x, y)
        fmt.Println(status, "- Hasil:", hasil)
    default:
        fmt.Println("Pilihan tidak valid")
    }
}