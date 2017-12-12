package os

import "time"

type Now struct {}

func (*Now) Time() time.Time {
	return time.Now()
}
