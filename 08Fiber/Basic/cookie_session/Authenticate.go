package cookie_session

import "time"

type Authenticate struct {
	UserId      string
	Keys        []string
	ExpiredTime time.Time
}
