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

	// Split into date and time parts
	parts := strings.Fields(durationStr)

	// Initialize weeks, days, hours, and minutes
	weeks, days, hours, minutes := 0, 0, 0, 0

	for _, part := range parts {
		if strings.HasSuffix(part, "w") {
			weekPart := strings.TrimSuffix(part, "w")
			weeks, _ = strconv.Atoi(weekPart)
		} else if strings.HasSuffix(part, "d") {
			dayPart := strings.TrimSuffix(part, "d")
			days, _ = strconv.Atoi(dayPart)
		} else if strings.Contains(part, ":") {
			timeParts := strings.Split(part, ":")
			if len(timeParts) == 2 {
				hours, _ = strconv.Atoi(timeParts[0])
				minutes, _ = strconv.Atoi(timeParts[1])
			}
		}
	}

	totalDuration = time.Duration(weeks)*7*24*time.Hour +
		time.Duration(days)*24*time.Hour +
		time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute

	return totalDuration, nil
}
