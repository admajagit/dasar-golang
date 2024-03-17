package helper

import "fmt"

// acces modifier
//  jika nama diawali huruf besar maka artinya bisa diakses dari package lain,jika dimulai dengan huruf kecil ,artinya tidak bisa diakses dari package lain
// biasanya di variable , struct , function

type Person struct{
	Nama string
	Umur int
}

var Alamatmu = "makan"


func Alok(name string){
	fmt.Println(name)
}

