package helpers

import (
	"fmt"
	"math"
	"time"
)

func GetJapaneseWeekdayKanji(dateBytes []uint8) (string, error) {
	// 1. Convert []uint8 slice to a string
	dateString := string(dateBytes)

	const mysqlDateFormat = "2006-01-02"
	// 2. Parse the string into a time.Time object
	t, err := time.Parse(mysqlDateFormat, dateString)
	if err != nil {
		return "ERR", fmt.Errorf("failed to parse date string '%s': %w", dateString, err)
	}

	// 3. Extract the weekday
	weekday := t.Weekday()

	switch weekday {
	case time.Sunday:
		return "日", nil
	case time.Monday:
		return "月", nil
	case time.Tuesday:
		return "火", nil
	case time.Wednesday:
		return "水", nil
	case time.Thursday:
		return "木", nil
	case time.Friday:
		return "金", nil
	case time.Saturday:
		return "土", nil
	default:
		return "ERR", fmt.Errorf("invalid weekday value: %v", weekday)
	}
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// CalculateRitsu calculates the percentage ratio (as int) of numerator to denominator.
// Returns 0 if denominator is 0.
func CalculateRitsu(numerator, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return int(float64(numerator) * 100 / float64(denominator))
}

// CalculateWariai calculates the percentage ratio (as float64) of numerator to denominator.
func CalculateWariai(numerator, denominator float64) float64 {
	if denominator == 0 {
		return 0.0
	}
	return RoundFloat((numerator*100)/denominator, 1)
}
