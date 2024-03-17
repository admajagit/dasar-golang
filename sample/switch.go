package main

import "fmt"


// pointer 
// passing by value menduplikasi nilai yang mana tidak mempengaruhi nilai aslinya

// pointer (passing by reference)  menduplikasi yang mana mempengaruhi nilai aslinya
// unutk menggunaka pointer kita bisa menggunakan opertaor &+diikiti nama variable
// untuk memodifikasi dari nilai pointer kita pakai *
// jangan lupa pakai *variable untuk mencetak hasil


func main() {
	var pertama = "bandung"
	var kedua = pertama  //poimter


	kedua = "jakarta"

	fmt.Println(pertama)
	fmt.Println(kedua)

	// pointer reference

	siji := "pertama"

	// Mengambil data dengan cara pointer menggunakan operator '&'
	loro:= &siji

	*loro = "begitu ya" // Memodifikasi nilai menggunakan operator '*' untuk mengubah nilai pointer

	
	telu := &loro
	*telu = "alok"

	fmt.Println(*telu)
	fmt.Println(*loro)
	fmt.Println(siji)

}

// operator new 

