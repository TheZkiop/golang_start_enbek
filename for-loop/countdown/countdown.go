package main

/**
 *	E-Mail: thezkiop@proton.me
 **/

import (
	"fmt"
	"time"
)

func main() {
	Countdown()
}

func ConvertToHumanTime(time int) string {
	hour := time / 3600
	minutes := (time % 3600) / 60
	seconds := time % 60
	formattedTime := fmt.Sprintf("%d:", hour)
	if minutes < 10 {
		formattedTime += fmt.Sprintf("0%d:", minutes)
	} else {
		formattedTime += fmt.Sprintf("%d:", minutes)
	}
	if seconds < 10 {
		formattedTime += fmt.Sprintf("0%d", seconds)
	} else {
		formattedTime += fmt.Sprintf("%d", seconds)
	}
	return formattedTime
}

// For cool animation uncomment next...
func Countdown() {
	// Set your time here
	timerTime := 10
	
	// Untouchable stuff goes next
	elapsedTime := timerTime


	for {
		// | This...
		// fmt.Println("\033cТаймер на ", ConvertToHumanTime(timerTime), ", Осталось: ", ConvertToHumanTime(elapsedTime))

		fmt.Println(elapsedTime)
		elapsedTime = elapsedTime - 1
		
		if elapsedTime < 1 {
			fmt.Println("Пуск!")
			break
		}

		time.Sleep(time.Second)
	}

	// | And this
	/*
	for {
		fmt.Println("\033c\033[5mТаймер истёк!, CTRL + C для выхода")
		time.Sleep(time.Minute)
	}
	*/
}

