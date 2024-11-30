package main

import "fmt"

// Toplama işlemi yapan bir fonksiyon
func add(a int, b int) int {
	return a + b
}

// Çoklu değer döndüren bir fonksiyon
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	// add() fonksiyonunu çağırma
	result := add(10, 5)
	fmt.Println("Addition Result:", result)

	// divide() fonksiyonunu çağırma
	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient:", quotient, "Remainder:", remainder)
}
