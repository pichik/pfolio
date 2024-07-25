package misc

import (
	"fmt"
	"time"
)

func GetTimeStamp(timeStr string) int64 {
	// Define the layout for the input string
	layout := "02.01.2006 15:04:05"

	// Parse the input string into a time.Time object
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return 0
	}

	// Convert the time.Time object to a Unix timestamp (seconds since January 1, 1970 UTC)
	timestamp := t.Unix()

	return timestamp
}
