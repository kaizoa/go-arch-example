package middleware

import "time"

type Now interface {
	Time() time.Time
}
