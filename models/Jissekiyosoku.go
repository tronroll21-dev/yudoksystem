package models

import (
	"fmt"
	"os"
	"time"
)

func calculateRecent21st(t time.Time) time.Time {
	// 1. Start by finding the 21st of the current month.
	// We use time.Date to reset the hour, minute, second, and nanosecond fields
	// for clean date comparison.
	recent21st := time.Date(t.Year(), t.Month(), 21, 0, 0, 0, 0, t.Location())

	// 2. Check if the calculated 21st is in the future relative to the current time (t).
	// If the current day is less than the 21st, then the 21st of the current month
	// is still in the future. In this case, we need to look back at the 21st
	// of the previous month.
	if t.Day() < 21 {
		// Subtract one month from the 21st of the current month.
		// AddDate(year_offset, month_offset, day_offset)
		recent21st = recent21st.AddDate(0, -1, 0)
	}

	return recent21st
}

func calculateYesterday(t time.Time) time.Time {
	// 1. Start by finding the 21st of the current month.
	// We use time.Date to reset the hour, minute, second, and nanosecond fields
	// for clean date comparison.
	yesterday := t.AddDate(0, 0, -1)

	return yesterday
}

func GetJissekiyosokuData() ([][]string, error) {
	// Read the SQL file
	queryBytes, err := os.ReadFile("sql/jissekiyosoku.sql")
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL file: %w", err)
	}
	query := string(queryBytes)

	// Get the current date and time
	currentTime := time.Now()

	//CalculateYesterday(currentTime)
	yesterday := calculateYesterday(currentTime)

	// Calculate the most recent 21st
	mostRecent21st := calculateRecent21st(currentTime)

	// Execute the query
	rows, err := db.Query(query, mostRecent21st.Format("2006-01-02"), yesterday.Format("2006-01-02"))
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Get column names for headers
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	// Initialize the result with headers
	var data [][]string
	data = append(data, columns)

	// Prepare slice for scanning values
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range columns {
		valuePtrs[i] = &values[i]
	}

	// Scan all rows
	for rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Convert values to strings
		var row []string
		for _, val := range values {
			if val == nil {
				row = append(row, "")
			} else if b, ok := val.([]byte); ok {
				row = append(row, string(b))
			} else {
				row = append(row, fmt.Sprintf("%v", val))
			}
		}
		data = append(data, row)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return data, nil
}
