package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	data_student := []student{
		{},
		{nama: "Mohammad Khairil Amin", alamat: "Jalan Gunung Anyar 1", pekerjaan: "Mahasiswa", alasan: "Ingin Mempelajari golang"},
		{nama: "Portia Perkins", alamat: "510-6237 Blandit Road", pekerjaan: "Chef Manager", alasan: "malesuada augue ut lacus. Nulla tincidunt, neque vitae semper egestas,"},
		{nama: "Quon Kane", alamat: "5756 Dui. Ave", pekerjaan: "Global Logistics Supervisor", alasan: "fermentum convallis ligula. Donec luctus aliquet odio. Etiam ligula tortor,"},
		{nama: "Brian Glenn", alamat: "Ap #894-9737 Sed, Street", pekerjaan: "Accountant", alasan: "nisi nibh lacinia orci, consectetuer euismod est arcu ac orci."},
		{nama: "Zephr Cantu", alamat: "963-526 Tellus St.", pekerjaan: "Web Developer", alasan: "libero. Proin mi. Aliquam gravida mauris ut mi. Duis risus"},
		{nama: "Mara Randall", alamat: "8434 Sed St.", pekerjaan: "Healthcare Specialist", alasan: "Suspendisse sagittis. Nullam vitae diam. Proin dolor. Nulla semper tellus"},
		{nama: "Bevis Dejesus", alamat: "5168 Purus, Rd.", pekerjaan: "Investment  Advisor", alasan: "Duis a mi fringilla mi lacinia mattis. Integer eu lacus."},
		{nama: "Dawn Rivers", alamat: "Ap #259-3029 Sit Rd.", pekerjaan: "Business Broker", alasan: "Maecenas ornare egestas ligula. Nullam feugiat placerat velit. Quisque varius."},
		{nama: "Joan Holden", alamat: "Ap #681-9206 Vitae Avenue", pekerjaan: "Doctor", alasan: "Maecenas libero est, congue a, aliquet vel, vulputate eu, odio."},
		{nama: "Arden Wallace", alamat: "Ap #448-6428 Lorem Street", pekerjaan: "Pharmacist", alasan: "Etiam imperdiet dictum magna. Ut tincidunt orci quis lectus. Nullam"},
	}
	if len(os.Args) != 1 {
		for i := 1; i < len(os.Args); i++ {
			absen, _ := strconv.Atoi(os.Args[i])

			if absen <= 0 {
				fmt.Printf("\nMasukan Urutan %v Tidak Valid!\n", absen)
				continue
			}

			if absen >= len(data_student) {
				fmt.Printf("\nBiodata Urutan %v Tidak Ditemukan!\n", absen)
				continue
			}

			fmt.Println("Nama\t\t\t\t: ", data_student[absen].nama)
			fmt.Println("Alamat\t\t\t\t: ", data_student[absen].alamat)
			fmt.Println("Pekerjaan\t\t\t: ", data_student[absen].pekerjaan)
			fmt.Println("Alasan memilih kelas Golang\t: ", data_student[absen].alasan)
		}
	} else {
		fmt.Println("\nMasukan Urutan Biodata Setelah biodata.go!")
	}
}
