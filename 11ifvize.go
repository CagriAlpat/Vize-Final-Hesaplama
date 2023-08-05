package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("vize notunuzu giriniz =  ")
	scanner.Scan()
	vize, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("final notunuzu giriniz =  ")
	scanner.Scan()
	final, _ := strconv.ParseFloat(scanner.Text(), 64)

	ortalama := (vize * 0.4) + (final * 0.6)
	fmt.Print("ortalama = \t", ortalama)

	if ortalama >= 85 && ortalama <= 100 {
		fmt.Println("AA")
	} else if ortalama >= 70 && ortalama < 85 {
		fmt.Println("BB")
	} else if ortalama >= 60 && ortalama < 70 {
		fmt.Println("CB")
	} else if ortalama >= 50 && ortalama < 60 {
		fmt.Println("CC")
	} else if ortalama < 50 && ortalama >= 0 {
		fmt.Println("FF")
	} else {
		fmt.Println("0-100 arasında bir değer olmalıdır.")
	}

}
