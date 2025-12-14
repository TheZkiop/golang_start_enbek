package main

/**
 *	E-Mail: thezkiop@proton.me
 **/

// FindMax находит максимальное число из переданных аргументов
// Если аргументы не переданы, возвращает 0
func FindMax(numbers ...int) int {
    	if len(numbers) == 0 {
		return 0
	}
	maxVal := numbers[0]
	
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > maxVal {
			maxVal = numbers[i]
		}
	}
	return maxVal
}

// Concatenate соединяет строки через указанный разделитель
// separator - разделитель
// parts - строки для соединения
// Пример: Concatenate(", ", "один", "два", "три") -> "один, два, три"
func Concatenate(separator string, parts ...string) string {
	str := ""
	for i := 0; i < len(parts); i++ {
		separatorOrNot := ""
		if i < len(parts) - 1 {
			separatorOrNot = separator
		}
		str += parts[i] + separatorOrNot
	}
	return str
}
