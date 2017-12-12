package today

import "time"

type Now interface {
	Time() time.Time
}
