package main

import "fmt"

func main() {
	// menampilkan nilai i
	i := 21
	fmt.Printf("Nilai i : %v\n", i)

	// menampilkan tipe data dari variable i
	fmt.Printf("Tipe data i : %T\n", i)

	// menampilkan tanda #
	fmt.Printf("Menampilkan tanda (persen) : %%\n")

	// menampilkan nilai boolean j
	j := true
	fmt.Printf("Menampilkan nilai j : %v\n", j)

	// menampilkan unicode russia Я (ya)
	ya := 'Я'
	fmt.Printf("Menampilkan unicode Я : %v\n", ya)

	// menampilkan nilai base 10
	fmt.Printf("Menampilkan nilai base 10 : %d \n", i)

	// menampilkan nilai base 8
	fmt.Printf("Menampilkan nilai base 8 : %o \n", i)

	// Menampilkan float :
	nilaiFloat := 123.456
	fmt.Printf("Menampilkan float : %f \n", nilaiFloat)

	// Menampilkan float scientific :
	fmt.Printf("Menampilkan float : %e \n", nilaiFloat)
}
