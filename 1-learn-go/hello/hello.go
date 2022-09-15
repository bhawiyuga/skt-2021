package main

import (
	"fmt"
	"time"
)

type Mahasiswa struct {
	nama string
	ipk  float32
	kota string
}

func ubahNilai(a *int) {
	*a = 1000
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	/*
		var a int = 10
		b := a + 100
		c := "Selamat datang"
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)

		d := 100
		ubahNilai(&d)
		fmt.Println(d)

		mhs1 := Mahasiswa{"Andi", 3.8, "Malang"}
		fmt.Println("Nama mahasiswa : ", mhs1.nama)

		myArray := [5]int{1, 2, 3, 4, 5}
		fmt.Println(myArray[0])

		var dynamicArray []Mahasiswa
		dynamicArray = append(dynamicArray, Mahasiswa{"Budi", 3.8, "Surabaya"})
		dynamicArray = append(dynamicArray, Mahasiswa{"Charlie", 3.9, "Jakarta"})
		fmt.Println(dynamicArray)

		fmt.Println("Hello")*/

	go say("world")
	say("hello")
}
