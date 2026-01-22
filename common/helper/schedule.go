package helper

import (
	"strings"
	"time"
)

func SplitSchedule(schedule string) []string {
	if schedule == "" {
		return nil
	}

	now := time.Now()
	layout := "15:04:05"

	var result []string
	parts := strings.Split(schedule, ",")

	for _, part := range parts {
		t := strings.TrimSpace(part)

		parsedTime, err := time.Parse(layout, t)
		if err != nil {
			continue
		}

		scheduleTime := time.Date(
			now.Year(), now.Month(), now.Day(),
			parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(),
			0, now.Location(),
		)

		if scheduleTime.After(now) {
			result = append(result, t)
		}
	}

	return result
}
