package main

import(
	"fmt"
)

func main() {
	var alas, tutup, tinggi int

	fmt.Scan(&alas)
	fmt.Scan(&tutup)
	fmt.Scan(&tinggi)

	luas := (alas+tutup) * tinggi / 2

	fmt.Printf("luasnya trapesium adalah..")
	fmt.Println(luas)
};