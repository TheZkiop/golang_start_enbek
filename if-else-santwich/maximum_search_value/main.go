package main

/**
 *	E-Mail: thezkiop@proton.me
 **/

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	a := rand.Intn(100) + 1
	b := rand.Intn(100) + 1
	c := rand.Intn(100) + 1
	
	fmt.Printf("Числа: %d, %d, %d\n", a, b, c)
	
	max := FindMax(a, b, c)
	
	fmt.Printf("Максимальное значение: %d\n", max)
}

// FindMax возвращает максимальное значение из трёх чисел
func FindMax(a, b, c int) int {
	return max(a, b, c)
}