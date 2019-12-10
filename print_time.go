package print_time

import (
	"fmt"
	"log"
	"github.com/beevik/ntp"
)

func PrintDt() bool {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err == nil {
		fmt.Printf("Time now is %s\n", time)
	} else {
		log.Println("Time displaying error: %s\n", err)
		return false
	}
	return true
}
