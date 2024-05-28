package internal

import (
	"fmt"
	"strings"
	"time"
)

func ParseFlappedTime(flappedTime string) (time.Duration, error) {
	// Split the string to get the duration part
	parts := strings.Split(flappedTime, "(")
	if len(parts) < 2 {
		return 0, fmt.Errorf("unexpected format for flapped time: %s", flappedTime)
	}
	durationPart := strings.TrimSuffix(parts[1], " ago)")

	// Parse the duration
	return time.ParseDuration(durationPart)
}
