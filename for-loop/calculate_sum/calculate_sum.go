package main

/**
 *	E-Mail: thezkiop@proton.me
 **/

import (
	"fmt"
)

func main() {
	// Вызов функции для вычисления суммы
	sum := CalculateSum()
	fmt.Println("Сумма:", sum)
}

// CalculateSum вычисляет сумму чисел от 1 до 100
func CalculateSum() int {
	iter := 1
	val := 1
	stop := 100
	for iter < stop {
		iter = iter + 1
		val = val + iter
	}
	return val
}
