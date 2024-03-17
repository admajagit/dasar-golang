package main

// import package error
import (
	"fmt"
)

// membuat custom error
// error menmapung return string

// jika kit amnegakses sebuah struct otomatis langsung dilempar ke implementasi
type Errorku struct {
	Pesan string
}



func (Jo Errorku) Error() string {
	return Jo.Pesan
}



func Pembagi(pertama int , kedua int) (int , error) {
	if kedua == 0 {

		// 
		return 0,Errorku{
			Pesan: "tidak bisa coy",

		}
	}

	return pertama/kedua,nil
}

func main()  {

	alok,err := Pembagi(100,0)

	fmt.Println(alok)
	fmt.Println(err)
}