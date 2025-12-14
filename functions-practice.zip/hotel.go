package main

/**
 *	E-Mail: thezkiop@proton.me
 **/

import (
  "fmt"
)

// CalculateRoomPrice вычисляет стоимость проживания
// на основе тарифа за ночь и количества ночей
func CalculateRoomPrice(rate float64, nights uint) float64 {
	return rate * float64(nights)
}

// ApplyDiscount применяет скидку к стоимости
// price - исходная стоимость
// discountPercent - процент скидки (от 0 до 100)
func ApplyDiscount(price float64, discountPercent float64) float64 {
  if discountPercent < 0 && discountPercent > 100 {
	discountPercent = 0
  }

  return price * (1 - discountPercent / 100)
}

// FormatRoomInfo возвращает информацию о номере в виде строки
// Формат: "Номер [номер]: вместимость [количество] человек(а), [количество] ночь/ночей"
// Примечание: Используйте "ночь" для 1 ночи и "ночей" для 0, 2, 3, 4... ночей
// Используйте "человек" для 1, 5, 6... человек и "человека" для 2, 3, 4 человек
func FormatRoomInfo(roomNumber int, capacity int, nights int) string {
	capacityFormat := "человек"
	switch capacity % 10 {
		case 2, 3, 4:
			capacityFormat = "человека"
	}
	
	nightsFormat := "ночей"
	switch nights % 10 {
		case 1:
			nightsFormat = "ночь"
	}

	return fmt.Sprintf("Номер %d: вместимость %d %s, %d %s", roomNumber, capacity, capacityFormat, nights, nightsFormat)
}
