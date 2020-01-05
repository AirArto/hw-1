package printtime

import (
	"time"

	"github.com/beevik/ntp"
)

// PrintDt for datetime printing
func PrintDt() (time.Time, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	return time, err
}
