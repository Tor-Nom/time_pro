package main

import (
	"fmt"
	"github.com/Tor-Nom/time_pro"
)

func main() {

	time := "2019-10-11 00:00:00"

	beginTime, endTime := time_pro.GetTime(time)
	fmt.Println("beginTime => ", beginTime)
	fmt.Println("endTime => ", endTime)
}
