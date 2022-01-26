package main

import "fmt"

func main() {
	var takim1, takim2 = "A takımı", "B takımı"
	fmt.Println(takim1, "ile", takim2, "akşam 19'da karşılaşacaklar")
	//diğer yazım şekli
	fmt.Printf("%s ile %s akşam 19'da karşılaşacaklar\n", takim1, takim2)
	var takim1gol, takim2gol = 3, 2

	fmt.Printf("Maç sonu skoru %d-%d", takim1gol, takim2gol)
}
