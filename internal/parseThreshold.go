package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Convert the threshold string into seconds.
func ParseThreshold(durationStr string) (time.Duration, error) {
	var totalSeconds int64

	parts := strings.FieldsFunc(durationStr, func(r rune) bool {
		return r == 'd' || r == 'h' || r == 'm' || r == 's' || r == 'w' || r == 'M'
	})
	units := strings.FieldsFunc(durationStr, func(r rune) bool {
		return r >= '0' && r <= '9'
	})

	for i, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			return 0, err
		}
		switch units[i] {
		case "M":
			totalSeconds += int64(value * 30 * 24 * 60 * 60)
		case "w":
			totalSeconds += int64(value * 7 * 24 * 60 * 60)
		case "d":
			totalSeconds += int64(value * 24 * 60 * 60)
		case "h":
			totalSeconds += int64(value * 60 * 60)
		case "m":
			totalSeconds += int64(value * 60)
		case "s":
			totalSeconds += int64(value)
		default:
			return 0, fmt.Errorf("invalid duration format")
		}
	}

	return time.Duration(totalSeconds) * time.Second, nil
}
