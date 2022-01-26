package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	min, max := 1, 1000
	rand.Seed(time.Now().UnixNano()) // bu satırı yazmazsak her zaman aynı sayıyı üretiyor.
	sayi := rand.Intn(max - min)

	fmt.Println("Üretilen sayı: ", sayi)

}
