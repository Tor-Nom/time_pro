package main

import (
	"fmt"
	"github.com/Tor-Nom/time_pro/service"
)

func main() {

	time := "2019-10-11 00:00:00"

	beginTime, endTime := service.GetTime(time)
	fmt.Println("beginTime => ", beginTime)
	fmt.Println("endTime => ", endTime)
}
