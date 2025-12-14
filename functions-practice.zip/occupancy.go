package main

/**
 *	E-Mail: thezkiop@proton.me
 **/

// CalculateOccupancyRate вычисляет процент занятых номеров
// roomsOccupied - количество занятых номеров
// totalRooms - общее количество номеров
// Возвращает процент занятых номеров (от 0 до 100)
func CalculateOccupancyRate(roomsOccupied, totalRooms int) float64 {
  roomsCount := totalRooms + roomsOccupied
  roomsOccupiedPercent := (float64 (roomsCount) / float64 (totalRooms)) - 1 // ????, ну ладно...
  return float64 (roomsOccupiedPercent * 100)
}

// DetermineOccupancyLevel определяет уровень загрузки отеля:
// "Низкий" - если загрузка менее 30%
// "Средний" - если загрузка от 30% до 60% включительно
// "Высокий" - если загрузка более 60%
func DetermineOccupancyLevel(rate float64) string {
  level := ""
  if rate < 30 {
	level = "Низкий"	
  } else if rate >= 30 && rate <= 60 {
	level = "Средний"
  } else if rate > 60 && rate <= 100 {
	level = "Высокий"
  }
  return level
}

// GetRecommendation возвращает рекомендацию, основанную на уровне загрузки:
// "Низкий" → "Запустите программу скидок"
// "Средний" → "Обычный режим работы"
// "Высокий" → "Повысьте цены"
func GetRecommendation(level string) string {
  recomm := ""
  if level == "Низкий" {
  	recomm = "Запустите программу скидок"
  } else if level == "Средний" {
	recomm = "Обычный режим работы"
  } else if level == "Высокий" {
	recomm = "Повысьте цены"
  }
  return recomm
}
