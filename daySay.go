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
		"\033[38;5;196;1m--\033[0m \033[38;5;45;1m%s \033[38;5;196;1m--\033[0m ",
		getAbbreviation(weekDay),
	)

	switch strings.ToLower(weekDay) {
	case "sunday":
		return getColoredMessage(prefix, "Last day of freedom!")
	case "monday":
		return getColoredMessage(prefix, "Why is Monday a thing?")
	case "tuesday":
		return getColoredMessage(prefix, "At least it’s not Monday.")
	case "wednesday":
		return getColoredMessage(prefix, "Halfway to freedom!")
	case "thursday":
		return getColoredMessage(prefix, "So close, yet so far.")
	case "friday":
		return getColoredMessage(prefix, "Work today, party tomorrow!")
	case "saturday":
		return getColoredMessage(prefix, "No alarms, no stress!")
	default:
		return getColoredMessage(prefix, "Chill mode: ON.")
	}
}

func getAbbreviation(weekDay string) string {
	return string([]rune(weekDay)[:3])
}

func getColoredMessage(prefix, message string) string {
	return prefix + "\033[38;5;213;1m" + message + "\033[0m"
}
