package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func UnixToStr(timeUnix int64, layout string) string {
	rand.Seed(time.Now().Unix())
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	return timeStr
}

func GetRandomTime(timeType int) string {
	rand.Seed(time.Now().Unix())
	hour := rand.Intn(12)
	minute := rand.Intn(59)
	second := rand.Intn(59)
	if timeType == 1 {
		return strconv.Itoa(hour) + ":" + strconv.Itoa(minute) + ":" + strconv.Itoa(second)
	} else {
		return strconv.Itoa(hour+12) + ":" + strconv.Itoa(minute) + ":" + strconv.Itoa(second)
	}
}
