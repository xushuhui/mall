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

func Int8Bool(i int8) bool {
	return i == 1
}

func TimeBecomeString(t time.Time) (s string) {
	s = t.Format("2006-01-02 15:04:05")
	return
}
