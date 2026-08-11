package habit

import (
	"testing"
	"time"
)

func TestStartOfLocalDayUsesTimestampLocation(t *testing.T) {
	location := time.FixedZone("UTC+08", 8*60*60)
	input := time.Date(2026, time.August, 11, 7, 30, 0, 0, location)

	got := startOfLocalDay(input)
	want := time.Date(2026, time.August, 11, 0, 0, 0, 0, location)
	if !got.Equal(want) {
		t.Fatalf("startOfLocalDay() = %v, want %v", got, want)
	}
	if got.Location() != location {
		t.Fatalf("startOfLocalDay() location = %v, want %v", got.Location(), location)
	}
}
