package main

import (
	"log"
	"time"
)

func parseDateTime(dateTimeStr string) time.Time {
	layout := "2006-01-02T15:04"
	t, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		log.Fatal(err)
	}
	return t
}