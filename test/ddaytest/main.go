package main

import (
	"fmt"
	"time"
)

func main() {

	days := []string{"2018-09-05", "2018-10-05", "2018-11-05", "2018-12-05", "2019-01-05"}
	for i := 0; i < len(days); i++ {
		GetDDay(days[i])
	}

}

func GetDDay(day string) {
	t := time.Now()
	dayTime, _ := time.Parse("2006-01-02", day)
	days := dayTime.Sub(t)

	fmt.Println(int(days.Hours() / 24))
}
