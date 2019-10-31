package main

import (
	"fmt"
	"github.com/Tor-Nom/time_pro/service"
)

func main() {

	endTime := "2019-10-11 00:00:00"
	fmt.Println(service.GetTime(endTime))
}
