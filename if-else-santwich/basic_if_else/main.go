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
	temperature := rand.Intn(41)
	fmt.Println("Температура:", temperature, "°C -", GetTemperatureMessage(temperature))
}

func GetTemperatureMessage(temp int) string {
	msg := "Очень холодно"
	if temp > 30 {
		msg = "Жарко"
	} else if temp >= 16 {
		msg = "Тепло"
	} else if temp >= 0 {
		msg = "Прохладно"
	}
	return msg
}