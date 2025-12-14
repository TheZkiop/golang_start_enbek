package main

/**
 *	E-Mail: thezkiop@proton.me
 **/

import (
	"fmt"
)

func main() {
	// Вызов функции для вывода чисел
	PrintNumbers()
}

// PrintNumbers выводит числа от 1 до 10, каждое на новой строке
func PrintNumbers() {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
}
