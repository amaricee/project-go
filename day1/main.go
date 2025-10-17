package main

import "fmt"

func main() {
	umur := 18

	if umur >= 17 {
		fmt.Println("Kamu sudah dewasa.")
	} else if umur >= 13 {
		fmt.Println("Kamu remaja.")
	} else {
		fmt.Println("Kamu masih anak-anak.")
	}
}
