package main

import (
	"fmt"
	"strings"
	"time"
)

func daySay() string {
	today := time.Now()
	weekDay := today.Weekday().String()

	prefix := fmt.Sprintf(
		"\033[38;5;45;1m%s \033[38;5;196;1m>\033[0m ",
		getAbbreviation(weekDay),
	)

	switch strings.ToLower(weekDay) {
	case "sunday":
		return prefix + "\033[38;5;213;1mLast day of freedom!\033[0m"
	case "monday":
		return prefix + "\033[38;5;213;1mWhy is Monday a thing?\033[0m"
	case "tuesday":
		return prefix + "\033[38;5;213;1mAt least it’s not Monday.\033[0m"
	case "wednesday":
		return prefix + "\033[38;5;213;1mHalfway to freedom!\033[0m"
	case "thursday":
		return prefix + "\033[38;5;213;1mSo close, yet so far.\033[0m"
	case "friday":
		return prefix + "\033[38;5;213;1mWork today, party tomorrow!\033[0m"
	case "saturday":
		return prefix + "\033[38;5;213;1mNo alarms, no stress!\033[0m"
	default:
		return "\033[38;5;213;1mChill mode: ON.\033[0m"
	}
}

func getAbbreviation(weekDay string) string {
	return string([]rune(weekDay)[:3])
}
