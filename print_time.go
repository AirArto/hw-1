package printtime

import (
	"time"

	"github.com/beevik/ntp"
)

// Do for datetime printing
func Do() (time.Time, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	return time, err
}
