package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseDuration(durationStr string) (time.Duration, error) {
	if durationStr == "Never" {
		return 0, fmt.Errorf("interface never flapped or no data available")
	}

	if strings.Contains(durationStr, "w") || strings.Contains(durationStr, "d") {
		return parseWeeksAndDays(durationStr)
	}

	return time.ParseDuration(durationStr)
}

// parseWeeksAndDays parses a string like "33w4d 12:30", "4d 09:14", or "6w0d 17:30" into a time.Duration.
func parseWeeksAndDays(durationStr string) (time.Duration, error) {
	var totalDuration time.Duration
	// Regular expression to match the format: "33w4d 12:30"
	re := regexp.MustCompile(`(?:(\d+)w)?(?:(\d+)d)?\s?(?:(\d+):(\d+))?`)

	matches := re.FindStringSubmatch(durationStr)
	if matches == nil {
		return 0, fmt.Errorf("invalid duration format: %s", durationStr)
	}

	weeks, _ := strconv.Atoi(matches[1])
	days, _ := strconv.Atoi(matches[2])
	hours, _ := strconv.Atoi(matches[3])
	minutes, _ := strconv.Atoi(matches[4])

	totalDuration = time.Duration(weeks)*7*24*time.Hour +
		time.Duration(days)*24*time.Hour +
		time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute

	return totalDuration, nil
}
