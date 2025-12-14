package main

/**
 *	E-Mail: thezkiop@proton.me
 **/

// Sum возвращает сумму двух целых чисел
func Sum(a, b int) int {
    return a + b
}

// Average возвращает среднее арифметическое трех чисел
func Average(a, b, c float64) float64 {
    return (a + b + c) / 3
}

// Greeting возвращает приветствие в формате "Привет, [имя]!"
func Greeting(name string) string {
    return ("Привет, " + name + "!")
}
