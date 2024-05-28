package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseDuration(durationStr string) (time.Duration, error) {
	var duration time.Duration
	var err error

	if strings.Contains(durationStr, "m") {
		duration, err = time.ParseDuration(durationStr)
	} else if strings.Contains(durationStr, "h") {
		duration, err = time.ParseDuration(durationStr)
	} else if strings.Contains(durationStr, "d") {
		days, err := strconv.Atoi(strings.TrimSuffix(durationStr, "d"))
		if err != nil {
			return 0, fmt.Errorf("invalid days in duration: %s", durationStr)
		}
		duration = time.Duration(days) * 24 * time.Hour
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
