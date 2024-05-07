package utils

import "time"

func Date(day, month, year int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func ToDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}

func DaysBetween(from, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}

	days := make([]time.Time, 0)
	for d := ToDay(from); !d.After(ToDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func IsDateBetween(current, from, to time.Time, isEqual bool) bool {
	if isEqual {
		return (current.After(from) || current.Equal(from)) && (current.Before(to) || current.Equal(to))
	}

	return current.After(from) && current.Before(to)
}
