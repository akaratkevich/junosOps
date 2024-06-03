package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseFlappedStamp(lastFlapped string) (time.Duration, error) {
	if lastFlapped == "Never" {
		return 0, nil
	}

	re := regexp.MustCompile(`\((\d+w)?(\d+d)?\s?(\d+:\d+)?\sago\)`)
	// should match :
	//"(3w ago)"
	//"(2d ago)"
	//"(5w2d ago)"
	//"(3w 12:45 ago)"
	//"(4d 13:20 ago)"
	//"(1w3d 14:05 ago)""

	matches := re.FindStringSubmatch(lastFlapped)
	if matches == nil {
		return 0, fmt.Errorf("unexpected format for duration: %s", lastFlapped)
	}

	var totalSeconds int64
	for i, match := range matches[1:] {
		if match == "" {
			continue
		}
		switch i {
		case 0: // weeks
			weeks, _ := strconv.Atoi(strings.TrimSuffix(match, "w"))
			totalSeconds += int64(weeks * 7 * 24 * 60 * 60)
		case 1: // days
			days, _ := strconv.Atoi(strings.TrimSuffix(match, "d"))
			totalSeconds += int64(days * 24 * 60 * 60)
		case 2: // hh:mm
			parts := strings.Split(match, ":")
			hours, _ := strconv.Atoi(parts[0])
			minutes, _ := strconv.Atoi(parts[1])
			totalSeconds += int64(hours*60*60 + minutes*60)
		}
	}

	return time.Duration(totalSeconds) * time.Second, nil
}
