package main

import "fmt"

func main() {
	/*fmt.Println("İsim giriniz: ")
	var isim string
	fmt.Scanln(&isim)

	fmt.Println("Girdiğiniz isim: ", isim)*/
	fmt.Println("Sayı giriniz: ")
	var sayi int
	fmt.Scanln(&sayi)

	fmt.Println(sayi * sayi)

}
