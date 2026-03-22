package models

import (
	"database/sql"
	"fmt"
	"log"
)

/*
データベースから取得する元データを格納する構造体
*/
type GasReading struct {
	ID             int            `db:"id" json:"id"`
	Year           int            `db:"year" json:"year"`
	Month          int            `db:"month" json:"month"`
	Day            int            `db:"day" json:"day"`
	Boiler1Reading int            `db:"boiler1_reading" json:"boiler1_reading"`
	Boiler2Reading int            `db:"boiler2_reading" json:"boiler2_reading"`
	Boiler3Reading int            `db:"boiler3_reading" json:"boiler3_reading"`
	Boiler4Reading int            `db:"boiler4_reading" json:"boiler4_reading"`
	TansanReading  int            `db:"tansan_reading" json:"tansan_reading"`
	Author         string         `db:"author" json:"author"`
	Memo           sql.NullString `db:"memo" json:"memo"`
}

// GetGasReadingsByYearAndMonth queries the mock database for all records on the specified date
func GetGasReadingsByYearAndMonth(yearStr string, monthStr string) ([]GasReading, bool, error) {
	var year, month int
	fmt.Sscanf(yearStr, "%d", &year)
	fmt.Sscanf(monthStr, "%d", &month)

	// Calculate previous month and year for the 21st baseline
	prevMonth := month - 1
	prevYear := year
	if prevMonth == 0 {
		prevMonth = 12
	}
	if prevMonth == 9 {
		prevYear = year - 1
	}

	query := `
	SELECT
		T1.ID,
		T1.year,
		T1.month,
		T1.day,
		T1.boiler1_reading,
		T1.boiler2_reading,
		T1.boiler3_reading,
		T1.boiler4_reading,
		T1.tansan_reading,
		T1.author,
		T1.memo
	FROM
		daily_gas_readings AS T1
	WHERE
		(T1.year = ? AND T1.month = ?)
		OR (T1.year = ? AND T1.month = ? AND T1.day = 21)
	ORDER BY
		T1.year, T1.month, T1.day;
	`

	log.Printf("The query parameters are: %d, %d, %d, %d\n", year, month, prevYear, prevMonth)
	rows, err := db.Query(query, year, month, prevYear, prevMonth)
	if err != nil {
		log.Fatalf("An SQL error occurred: %v\n", err)
		return nil, false, err
	}
	defer rows.Close()

	var readings []GasReading
	for rows.Next() {
		var raw GasReading
		err := rows.Scan(
			&raw.ID,
			&raw.Year,
			&raw.Month,
			&raw.Day,
			&raw.Boiler1Reading,
			&raw.Boiler2Reading,
			&raw.Boiler3Reading,
			&raw.Boiler4Reading,
			&raw.TansanReading,
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

// SaveGasReading inserts or updates a power reading record
func SaveGasReading(r *GasReading) error {
	log.Printf("Saving gas reading: %f", r.TansanReading)
	log.Printf("Saving gas reading: %s", r.Author)
	log.Printf("Saving gas reading: %s", r.Memo.String)
	log.Printf("Saving gas reading: %d", r.Day)
	log.Printf("Saving power reading: %d", r.Month)
	log.Printf("Saving power reading: %d", r.Year)

	query := `
	INSERT INTO daily_gas_readings (ID, year, month, day, boiler1_reading, boiler2_reading, boiler3_reading, boiler4_reading, tansan_reading, author, memo)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
		boiler1_reading = VALUES(boiler1_reading),
		boiler2_reading = VALUES(boiler2_reading),
		boiler3_reading = VALUES(boiler3_reading),
		boiler4_reading = VALUES(boiler4_reading),
		tansan_reading = VALUES(tansan_reading),
		author = VALUES(author),
		memo = VALUES(memo);
	`
	_, err := db.Exec(query, r.ID, r.Year, r.Month, r.Day, r.Boiler1Reading, r.Boiler2Reading, r.Boiler3Reading, r.Boiler4Reading, r.TansanReading, r.Author, r.Memo.String)
	if err != nil {
		log.Printf("Error saving gas reading: %v", err)
		return err
	}
	return nil
}
