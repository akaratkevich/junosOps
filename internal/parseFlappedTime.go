package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseFlappedTime(flappedTime string) (time.Duration, error) {
	if strings.Contains(flappedTime, "Never") || strings.TrimSpace(flappedTime) == "" {
		return 0, fmt.Errorf("interface never flapped or no data available")
	}

	// Extract duration part
	parts := strings.Split(flappedTime, "(")
	if len(parts) < 2 {
		return 0, fmt.Errorf("unexpected format for flapped time: %s", flappedTime)
	}
	durationPart := strings.TrimSpace(strings.TrimSuffix(parts[1], " ago)"))

	// Parse duration in HH:MM:SS format
	timeParts := strings.Split(durationPart, ":")
	if len(timeParts) != 3 {
		return 0, fmt.Errorf("unexpected format for duration: %s", durationPart)
	}

	hours, err := strconv.Atoi(timeParts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid hours in duration: %s", durationPart)
	}

	minutes, err := strconv.Atoi(timeParts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid minutes in duration: %s", durationPart)
	}

	seconds, err := strconv.Atoi(timeParts[2])
	if err != nil {
		return 0, fmt.Errorf("invalid seconds in duration: %s", durationPart)
	}

	return time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second, nil
}
