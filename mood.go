package main

import (
	"fmt"
	"strings"
	"time"
)

func mood() string {
	today := time.Now()
	year := today.Year()
	month := today.Month().String()
	day := today.Day()
	weekDay := today.Weekday().String()

	prefix := fmt.Sprintf(
		"[%d/%s/%d] %s > ",
		year,
		getAbbreviation(month),
		day,
		getAbbreviation(weekDay),
	)

	switch strings.ToLower(weekDay) {
	case "sunday":
		return prefix + "Last day of freedom!"
	case "monday":
		return prefix + "Why is Monday a thing?"
	case "tuesday":
		return prefix + "At least it’s not Monday."
	case "wednesday":
		return prefix + "Halfway to freedom!"
	case "thursday":
		return prefix + "So close, yet so far."
	case "friday":
		return prefix + "Work today, party tomorrow!"
	case "saturday":
		return prefix + "No alarms, no stress!"
	default:
		return "Chill mode: ON."
	}
}

func getAbbreviation(weekDay string) string {
	return string([]rune(weekDay)[:3])
}
