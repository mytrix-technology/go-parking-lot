package parkinglot

import (
	"fmt"
)

var Type_Of_Input int

func Init() {
	fmt.Println("*****************************Selamat Datang di Sistem Parking Lot*****************************")
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("Pilih salah satu type masukkan ? Tulis 1 untuk command shell | 2 untuk file")
	fmt.Scanf("%d", &Type_Of_Input)
}
