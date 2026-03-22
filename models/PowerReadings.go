package models

import (
	"database/sql"
	"fmt"
	"log"
)

/*
データベースから取得する元データを格納する構造体
*/
type PowerReading struct {
	ID           int            `db:"id" json:"id"`
	Year         int            `db:"year" json:"year"`
	Month        int            `db:"month" json:"month"`
	Day          int            `db:"day" json:"day"`
	PowerReading float64        `db:"power_reading" json:"power_reading"`
	Author       string         `db:"author" json:"author"`
	Memo         sql.NullString `db:"memo" json:"memo"`
}

// GetPowerReadingByDate queries the mock database for all records on the specified date
func GetPowerReadingsByYearAndMonth(yearStr string, monthStr string) ([]PowerReading, bool, error) {
	var year, month int
	fmt.Sscanf(yearStr, "%d", &year)
	fmt.Sscanf(monthStr, "%d", &month)

	// Calculate previous month and year for the 21st baseline
	prevMonth := month - 1
	prevYear := year
	if prevMonth == 0 {
		prevMonth = 12
		prevYear = year - 1
	}

	query := `
	SELECT
		T1.ID,
		T1.year,
		T1.month,
		T1.day,
		T1.power_reading,
		T1.author,
		T1.memo
	FROM
		daily_power_readings AS T1
	WHERE
		(T1.year = ? AND T1.month = ?)
		OR (T1.year = ? AND T1.month = ? AND T1.day = 21)
	ORDER BY
		T1.year, T1.month, T1.day;
	`
	rows, err := db.Query(query, year, month, prevYear, prevMonth)
	if err != nil {
		log.Fatalf("An SQL error occurred: %v\n", err)
		return nil, false, err
	}
	defer rows.Close()

	var readings []PowerReading
	for rows.Next() {
		var raw PowerReading
		err := rows.Scan(
			&raw.ID,
			&raw.Year,
			&raw.Month,
			&raw.Day,
			&raw.PowerReading,
			&raw.Author,
			&raw.Memo,
		)
		if err != nil {
			log.Fatalf("An error occurred while scanning row: %v\n", err)
			return nil, false, err
		}
		readings = append(readings, raw)
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("An error occurred after scanning rows: %v\n", err)
		return nil, false, err
	}

	if len(readings) == 0 {
		fmt.Printf("No records with date %d-%02d found\n", year, month)
		return readings, false, nil
	}

	fmt.Printf("Found %d records for Date: %d-%02d\n", len(readings), year, month)
	return readings, true, nil
}

// SavePowerReading inserts or updates a power reading record
func SavePowerReading(r *PowerReading) error {
	log.Printf("Saving power reading: %f", r.PowerReading)
	log.Printf("Saving power reading: %s", r.Author)
	log.Printf("Saving power reading: %s", r.Memo.String)
	log.Printf("Saving power reading: %d", r.Day)
	log.Printf("Saving power reading: %d", r.Month)
	log.Printf("Saving power reading: %d", r.Year)

	query := `
	INSERT INTO daily_power_readings (ID, year, month, day, power_reading, author, memo)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
		power_reading = VALUES(power_reading),
		author = VALUES(author),
		memo = VALUES(memo);
	`
	_, err := db.Exec(query, r.ID, r.Year, r.Month, r.Day, r.PowerReading, r.Author, r.Memo.String)
	if err != nil {
		log.Printf("Error saving power reading: %v", err)
		return err
	}
	return nil
}
