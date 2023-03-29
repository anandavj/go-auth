package helper

import (
	"time"
)

func GetCurrentTime() int64 {
	return time.Now().UnixMilli()
}
