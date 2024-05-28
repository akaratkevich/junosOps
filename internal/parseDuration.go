package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Converts a string representation of a duration into a time.Duration.
func ParseDuration(durationStr string) (time.Duration, error) {
	var duration time.Duration
	var err error

	if strings.Contains(durationStr, "w") || strings.Contains(durationStr, "d") {
		duration, err = parseWeeksAndDays(durationStr)
	} else if strings.Contains(durationStr, "m") || strings.Contains(durationStr, "h") {
		duration, err = time.ParseDuration(durationStr)
	} else if strings.Contains(durationStr, "M") {
		months, err := strconv.Atoi(strings.TrimSuffix(durationStr, "M"))
		if err != nil {
			return 0, fmt.Errorf("invalid months in duration: %s", durationStr)
		}
		duration = time.Duration(months) * 30 * 24 * time.Hour
	} else {
		err = fmt.Errorf("unsupported duration format: %s", durationStr)
	}
	return duration, err
}

// Parses a string like "33w4d" into a time.Duration.
func parseWeeksAndDays(durationStr string) (time.Duration, error) {
	var totalDuration time.Duration

	weekDayParts := strings.Split(durationStr, "w")
	if len(weekDayParts) != 2 {
		return 0, fmt.Errorf("invalid format for weeks and days: %s", durationStr)
	}

	weeks, err := strconv.Atoi(weekDayParts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid weeks in duration: %s", durationStr)
	}

	daysPart := weekDayParts[1]
	days, err := strconv.Atoi(strings.TrimSuffix(daysPart, "d"))
	if err != nil {
		return 0, fmt.Errorf("invalid days in duration: %s", durationStr)
	}

	totalDuration = time.Duration(weeks)*7*24*time.Hour + time.Duration(days)*24*time.Hour
	return totalDuration, nil
}
