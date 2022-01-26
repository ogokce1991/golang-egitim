package main

import "fmt"

func main() {
	/*var a int = 10

	if a > 9 {
		fmt.Println("a sayısı 9 dan büyüktür")
	} else if a == 9 {
		fmt.Println("a sayısı 9 a eşittir.")
	} else {
		fmt.Println("a sayısı 9 dan büyük değildir.")
	}*/

	ortalama := 91

	if ortalama > 90 {
		fmt.Println("A")
	} else if ortalama > 80 {
		fmt.Println("B")
	} else if ortalama > 70 {
		fmt.Println("C")
	} else {
		fmt.Println("Kaldı")
	}

}
