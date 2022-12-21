package utils

import "time"

type TimeServiceImpl struct {
}

func (*TimeServiceImpl) Now() time.Time {
	return time.Now()
}
