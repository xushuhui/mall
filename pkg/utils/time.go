package utils

import "time"

func IsInTime(t time.Time, start time.Time, end time.Time) (in bool) {
	if t.After(start) && t.Before(end) {
		in = true
	}
	return
}

func Int2Bool(i int) bool {
	return i == 1
}
