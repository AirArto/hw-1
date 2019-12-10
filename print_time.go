package main

import (
	"fmt"

	"github.com/beevik/ntp"
)

func PrintDt() bool {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err == nil {
		fmt.Printf("time now is %s\n", time)
	} else {
		return false
	}
	return true
}
