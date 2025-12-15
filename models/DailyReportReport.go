package models

import (
	"fmt"
	"log"
	"time"

	"tronroll21-dev/yudoksystem/models/helpers"
)

type ReportRecordData struct {
	DailyReportRaw
	CashCountGoukei          int
	CashAmountGoukei         int
	SettleCountGoukei        int
	SettleAmountGoukei       int
	Machine1CashCountGoukei  int
	Machine2CashCountGoukei  int
	Machine3CashCountGoukei  int
	Machine4CashCountGoukei  int
	Machine5CashCountGoukei  int
	Machine1CashAmountGoukei int
	Machine2CashAmountGoukei int
	Machine3CashAmountGoukei int
	Machine4CashAmountGoukei int
	Machine5CashAmountGoukei int
	CashCountGoukeiGoukei    int
	CashAmountGoukeiGoukei   int
	UnsettledCountGoukei     int
	UnsettledAmountGoukei    int
	UriageGoukeiNyuuyoku     int
	UriageGoukeiInshoku      int
	UriageGoukeiSeitai       int
	UriageGoukeiTotal        int
	Machine1ECountGoukei     int
	Machine2ECountGoukei     int
	Machine3ECountGoukei     int
	Machine4ECountGoukei     int
	Machine5ECountGoukei     int
	Machine1EAmountGoukei    int
	Machine2EAmountGoukei    int
	Machine3EAmountGoukei    int
	Machine4EAmountGoukei    int
	Machine5EAmountGoukei    int
	Machine1CCountGoukei     int
	Machine2CCountGoukei     int
	Machine3CCountGoukei     int
	Machine4CCountGoukei     int
	Machine5CCountGoukei     int
	Machine1CAmountGoukei    int
	Machine2CAmountGoukei    int
	Machine3CAmountGoukei    int
	Machine4CAmountGoukei    int
	Machine5CAmountGoukei    int
	ECountGoukei             int
	EAmountGoukei            int
	ESettleCountGoukei       int
	ESettleAmountGoukei      int
	ECountGoukeiGoukei       int
	EAmountGoukeiGoukei      int
	CCountGoukei             int
	CAmountGoukei            int
	CSettleCountGoukei       int
	CSettleAmountGoukei      int
	CCountGoukeiGoukei       int
	CAmountGoukeiGoukei      int
	Machine1QrCountGoukei    int
	Machine2QrCountGoukei    int
	Machine3QrCountGoukei    int
	Machine4QrCountGoukei    int
	Machine5QrCountGoukei    int
	Machine1QrAmountGoukei   int
	Machine2QrAmountGoukei   int
	Machine3QrAmountGoukei   int
	Machine4QrAmountGoukei   int
	Machine5QrAmountGoukei   int
	QrCountGoukei            int
	QrAmountGoukei           int
	QrSettleCountGoukei      int
	QrSettleAmountGoukei     int
	QrCountGoukeiGoukei      int
	QrAmountGoukeiGoukei     int
	Machine1CashlessGoukei   int
	Machine2CashlessGoukei   int
	Machine3CashlessGoukei   int
	Machine4CashlessGoukei   int
	Machine5CashlessGoukei   int
	CashlessGoukei           int
	Machine1TotalGoukei      int
	Machine2TotalGoukei      int
	Machine3TotalGoukei      int
	Machine4TotalGoukei      int
	Machine5TotalGoukei      int
	TotalGoukei              int
	NyuuyokuWariai           float64
	InshokuWariai            float64
	SeitaiWariai             float64
	Miseisan_Tousha          int
	TounyuuGoukei            int
	Nyuukinyoteikingaku      int
	Nyuukinyoteibi           string
}

// Define a struct to pass data to the template
type ReportData struct {
	Date   string
	Record *ReportRecordData
	Found  bool
}

// GetSalesReportByDate fetches data and generates an HTML report.
func GetSalesReportByDate(targetDate time.Time) (ReportData, error) {
	// Helper function to convert []uint8 to string (YYYY-MM-DD)
	convertDateUint8ToString := func(dateUint8 []uint8) string {
		if len(dateUint8) == 0 {
			return ""
		}
		return string(dateUint8)
	}

	record, found, err := GetSalesRecordByDate(targetDate)
	if err != nil {
		return ReportData{}, fmt.Errorf("failed to get sales record: %w", err)
	}

	var dateStr string
	if record != nil && len(record.Date) > 0 {
		dateStr = convertDateUint8ToString(record.Date)
	} else {
		dateStr = targetDate.Format("2006-01-02")
	}
	log.Printf("Record date: %s", dateStr)

	var Machine1CashCountGoukei, Machine2CashCountGoukei, Machine3CashCountGoukei, Machine4CashCountGoukei, Machine5CashCountGoukei,
		Machine1CashAmountGoukei, Machine2CashAmountGoukei, Machine3CashAmountGoukei, Machine4CashAmountGoukei, Machine5CashAmountGoukei int
	var Machine1ECountGoukei, Machine2ECountGoukei, Machine3ECountGoukei, Machine4ECountGoukei, Machine5ECountGoukei,
		Machine1EAmountGoukei, Machine2EAmountGoukei, Machine3EAmountGoukei, Machine4EAmountGoukei, Machine5EAmountGoukei int
	var Machine1QrCountGoukei, Machine2QrCountGoukei, Machine3QrCountGoukei, Machine4QrCountGoukei, Machine5QrCountGoukei,
		Machine1QrAmountGoukei, Machine2QrAmountGoukei, Machine3QrAmountGoukei, Machine4QrAmountGoukei, Machine5QrAmountGoukei int
	var Machine1CCountGoukei, Machine2CCountGoukei, Machine3CCountGoukei, Machine4CCountGoukei, Machine5CCountGoukei,
		Machine1CAmountGoukei, Machine2CAmountGoukei, Machine3CAmountGoukei, Machine4CAmountGoukei, Machine5CAmountGoukei int

	Machine1CashCountGoukei = record.Machine1CashCountNum - record.Machine1SettleCountNum
	Machine2CashCountGoukei = record.Machine2CashCountNum - record.Machine2SettleCountNum
	Machine3CashCountGoukei = record.Machine3CashCountNum - record.Machine3SettleCountNum
	Machine4CashCountGoukei = record.Machine4CashCountNum - record.Machine4SettleCountNum
	Machine5CashCountGoukei = record.Machine5CashCountNum - record.Machine5SettleCountNum
	Machine1CashAmountGoukei = record.Machine1CashAmountNum - record.Machine1SettleAmountNum
	Machine2CashAmountGoukei = record.Machine2CashAmountNum - record.Machine2SettleAmountNum
	Machine3CashAmountGoukei = record.Machine3CashAmountNum - record.Machine3SettleAmountNum
	Machine4CashAmountGoukei = record.Machine4CashAmountNum - record.Machine4SettleAmountNum
	Machine5CashAmountGoukei = record.Machine5CashAmountNum - record.Machine5SettleAmountNum

	Machine1CashAmountGoukeiWithUnsettled := Machine1CashAmountGoukei - record.Machine1UnsettledAmountNum
	Machine2CashAmountGoukeiWithUnsettled := Machine2CashAmountGoukei - record.Machine2UnsettledAmountNum
	Machine3CashAmountGoukeiWithUnsettled := Machine3CashAmountGoukei - record.Machine3UnsettledAmountNum
	Machine4CashAmountGoukeiWithUnsettled := Machine4CashAmountGoukei - record.Machine4UnsettledAmountNum
	Machine5CashAmountGoukeiWithUnsettled := Machine5CashAmountGoukei - record.Machine5UnsettledAmountNum

	Machine1ECountGoukei = record.Machine1ECountNum - record.Machine1ESettleCountNum
	Machine2ECountGoukei = record.Machine2ECountNum - record.Machine2ESettleCountNum
	Machine3ECountGoukei = record.Machine3ECountNum - record.Machine3ESettleCountNum
	Machine4ECountGoukei = record.Machine4ECountNum - record.Machine4ESettleCountNum
	Machine5ECountGoukei = record.Machine5ECountNum - record.Machine5ESettleCountNum
	Machine1EAmountGoukei = record.Machine1EAmountNum - record.Machine1ESettleAmountNum
	Machine2EAmountGoukei = record.Machine2EAmountNum - record.Machine2ESettleAmountNum
	Machine3EAmountGoukei = record.Machine3EAmountNum - record.Machine3ESettleAmountNum
	Machine4EAmountGoukei = record.Machine4EAmountNum - record.Machine4ESettleAmountNum
	Machine5EAmountGoukei = record.Machine5EAmountNum - record.Machine5ESettleAmountNum

	Machine1CCountGoukei = record.Machine1CCountNum - record.Machine1CSettleCountNum
	Machine2CCountGoukei = record.Machine2CCountNum - record.Machine2CSettleCountNum
	Machine3CCountGoukei = record.Machine3CCountNum - record.Machine3CSettleCountNum
	Machine4CCountGoukei = record.Machine4CCountNum - record.Machine4CSettleCountNum
	Machine5CCountGoukei = record.Machine5CCountNum - record.Machine5CSettleCountNum
	Machine1CAmountGoukei = record.Machine1CAmountNum - record.Machine1CSettleAmountNum
	Machine2CAmountGoukei = record.Machine2CAmountNum - record.Machine2CSettleAmountNum
	Machine3CAmountGoukei = record.Machine3CAmountNum - record.Machine3CSettleAmountNum
	Machine4CAmountGoukei = record.Machine4CAmountNum - record.Machine4CSettleAmountNum
	Machine5CAmountGoukei = record.Machine5CAmountNum - record.Machine5CSettleAmountNum

	Machine1QrCountGoukei = record.Machine1QrCountNum - record.Machine1QrSettleCountNum
	Machine2QrCountGoukei = record.Machine2QrCountNum - record.Machine2QrSettleCountNum
	Machine3QrCountGoukei = record.Machine3QrCountNum - record.Machine3QrSettleCountNum
	Machine4QrCountGoukei = record.Machine4QrCountNum - record.Machine4QrSettleCountNum
	Machine5QrCountGoukei = record.Machine5QrCountNum - record.Machine5QrSettleCountNum
	Machine1QrAmountGoukei = record.Machine1QrAmountNum - record.Machine1QrSettleAmountNum
	Machine2QrAmountGoukei = record.Machine2QrAmountNum - record.Machine2QrSettleAmountNum
	Machine3QrAmountGoukei = record.Machine3QrAmountNum - record.Machine3QrSettleAmountNum
	Machine4QrAmountGoukei = record.Machine4QrAmountNum - record.Machine4QrSettleAmountNum
	Machine5QrAmountGoukei = record.Machine5QrAmountNum - record.Machine5QrSettleAmountNum

	Machine1CashlessGoukei := Machine1QrAmountGoukei + Machine1EAmountGoukei + Machine1CAmountGoukei
	Machine2CashlessGoukei := Machine2QrAmountGoukei + Machine2EAmountGoukei + Machine2CAmountGoukei
	Machine3CashlessGoukei := Machine3QrAmountGoukei + Machine3EAmountGoukei + Machine3CAmountGoukei
	Machine4CashlessGoukei := Machine4QrAmountGoukei + Machine4EAmountGoukei + Machine4CAmountGoukei
	Machine5CashlessGoukei := Machine5QrAmountGoukei + Machine5EAmountGoukei + Machine5CAmountGoukei

	Machine1TotalGoukei := Machine1CashAmountGoukeiWithUnsettled + Machine1EAmountGoukei + Machine1QrAmountGoukei + Machine1CAmountGoukei
	Machine2TotalGoukei := Machine2CashAmountGoukeiWithUnsettled + Machine2EAmountGoukei + Machine2QrAmountGoukei + Machine2CAmountGoukei
	Machine3TotalGoukei := Machine3CashAmountGoukeiWithUnsettled + Machine3EAmountGoukei + Machine3QrAmountGoukei + Machine3CAmountGoukei
	Machine4TotalGoukei := Machine4CashAmountGoukeiWithUnsettled + Machine4EAmountGoukei + Machine4QrAmountGoukei + Machine4CAmountGoukei
	Machine5TotalGoukei := Machine5CashAmountGoukeiWithUnsettled + Machine5EAmountGoukei + Machine5QrAmountGoukei + Machine5CAmountGoukei

	TotalGoukei := Machine1TotalGoukei + Machine2TotalGoukei + Machine3TotalGoukei + Machine4TotalGoukei + Machine5TotalGoukei

	Miseisan_Tousha :=
		record.Machine1UnsettledAmountNum +
			record.Machine2UnsettledAmountNum +
			record.Machine3UnsettledAmountNum +
			record.Machine4UnsettledAmountNum +
			record.Machine5UnsettledAmountNum

	TounyuuGoukei := TotalGoukei + record.Change - Miseisan_Tousha + record.PhoneFee - record.HonjitsuMitounyuuAmountUncertain - record.HonjitsuMitounyuuAmountCertain + record.Deficiency + record.ZenjitsuMitounyuuAmount

	originalDate, err := time.Parse("2006-01-02", record.DateString)
	if err != nil {
		return ReportData{}, fmt.Errorf("failed to parse date: %w", err)
	}

	Nyuukinyoteibi := getNyuukinyouteibi(originalDate)

	Nyuukinyoteikingaku := TounyuuGoukei - record.Change

	reportRecordData := ReportRecordData{
		DailyReportRaw: *record,
		CashCountGoukei: record.Machine1CashCountNum +
			record.Machine2CashCountNum +
			record.Machine3CashCountNum +
			record.Machine4CashCountNum +
			record.Machine5CashCountNum,
		CashAmountGoukei: record.Machine1CashAmountNum +
			record.Machine2CashAmountNum +
			record.Machine3CashAmountNum +
			record.Machine4CashAmountNum +
			record.Machine5CashAmountNum,
		SettleCountGoukei: record.Machine1SettleCountNum +
			record.Machine2SettleCountNum +
			record.Machine3SettleCountNum +
			record.Machine4SettleCountNum +
			record.Machine5SettleCountNum,
		SettleAmountGoukei: record.Machine1SettleAmountNum +
			record.Machine2SettleAmountNum +
			record.Machine3SettleAmountNum +
			record.Machine4SettleAmountNum +
			record.Machine5SettleAmountNum,
		Machine1CashCountGoukei:  Machine1CashCountGoukei,
		Machine2CashCountGoukei:  Machine2CashCountGoukei,
		Machine3CashCountGoukei:  Machine3CashCountGoukei,
		Machine4CashCountGoukei:  Machine4CashCountGoukei,
		Machine5CashCountGoukei:  Machine5CashCountGoukei,
		Machine1CashAmountGoukei: Machine1CashAmountGoukei,
		Machine2CashAmountGoukei: Machine2CashAmountGoukei,
		Machine3CashAmountGoukei: Machine3CashAmountGoukei,
		Machine4CashAmountGoukei: Machine4CashAmountGoukei,
		Machine5CashAmountGoukei: Machine5CashAmountGoukei,
		CashCountGoukeiGoukei: Machine1CashCountGoukei +
			Machine2CashCountGoukei +
			Machine3CashCountGoukei +
			Machine4CashCountGoukei +
			Machine5CashCountGoukei,
		CashAmountGoukeiGoukei: Machine1CashAmountGoukei +
			Machine2CashAmountGoukei +
			Machine3CashAmountGoukei +
			Machine4CashAmountGoukei +
			Machine5CashAmountGoukei,
		UnsettledCountGoukei: record.Machine1UnsettledCountNum +
			record.Machine2UnsettledCountNum +
			record.Machine3UnsettledCountNum +
			record.Machine4UnsettledCountNum +
			record.Machine5UnsettledCountNum,
		UnsettledAmountGoukei: record.Machine1UnsettledAmountNum +
			record.Machine2UnsettledAmountNum +
			record.Machine3UnsettledAmountNum +
			record.Machine4UnsettledAmountNum +
			record.Machine5UnsettledAmountNum,
		UriageGoukeiNyuuyoku:  Machine1CashAmountGoukeiWithUnsettled + Machine2CashAmountGoukeiWithUnsettled,
		UriageGoukeiInshoku:   Machine3CashAmountGoukeiWithUnsettled + Machine4CashAmountGoukeiWithUnsettled,
		UriageGoukeiSeitai:    Machine5CashAmountGoukeiWithUnsettled,
		UriageGoukeiTotal:     Machine1CashAmountGoukei + Machine2CashAmountGoukei + Machine3CashAmountGoukei + Machine4CashAmountGoukei + Machine5CashAmountGoukei,
		Machine1ECountGoukei:  Machine1ECountGoukei,
		Machine2ECountGoukei:  Machine2ECountGoukei,
		Machine3ECountGoukei:  Machine3ECountGoukei,
		Machine4ECountGoukei:  Machine4ECountGoukei,
		Machine5ECountGoukei:  Machine5ECountGoukei,
		Machine1EAmountGoukei: Machine1EAmountGoukei,
		Machine2EAmountGoukei: Machine2EAmountGoukei,
		Machine3EAmountGoukei: Machine3EAmountGoukei,
		Machine4EAmountGoukei: Machine4EAmountGoukei,
		Machine5EAmountGoukei: Machine5EAmountGoukei,
		ECountGoukei: record.Machine1ECountNum +
			record.Machine2ECountNum +
			record.Machine3ECountNum +
			record.Machine4ECountNum +
			record.Machine5ECountNum,
		EAmountGoukei: record.Machine1EAmountNum +
			record.Machine2EAmountNum +
			record.Machine3EAmountNum +
			record.Machine4EAmountNum +
			record.Machine5EAmountNum,
		ESettleCountGoukei: record.Machine1ESettleCountNum +
			record.Machine2ESettleCountNum +
			record.Machine3ESettleCountNum +
			record.Machine4ESettleCountNum +
			record.Machine5ESettleCountNum,
		ESettleAmountGoukei: record.Machine1ESettleAmountNum +
			record.Machine2ESettleAmountNum +
			record.Machine3ESettleAmountNum +
			record.Machine4ESettleAmountNum +
			record.Machine5ESettleAmountNum,
		ECountGoukeiGoukei: Machine1ECountGoukei +
			Machine2ECountGoukei +
			Machine3ECountGoukei +
			Machine4ECountGoukei +
			Machine5ECountGoukei,
		EAmountGoukeiGoukei: Machine1EAmountGoukei +
			Machine2EAmountGoukei +
			Machine3EAmountGoukei +
			Machine4EAmountGoukei +
			Machine5EAmountGoukei,
		Machine1CCountGoukei:  Machine1CCountGoukei,
		Machine2CCountGoukei:  Machine2CCountGoukei,
		Machine3CCountGoukei:  Machine3CCountGoukei,
		Machine4CCountGoukei:  Machine4CCountGoukei,
		Machine5CCountGoukei:  Machine5CCountGoukei,
		Machine1CAmountGoukei: Machine1CAmountGoukei,
		Machine2CAmountGoukei: Machine2CAmountGoukei,
		Machine3CAmountGoukei: Machine3CAmountGoukei,
		Machine4CAmountGoukei: Machine4CAmountGoukei,
		Machine5CAmountGoukei: Machine5CAmountGoukei,
		CCountGoukei: record.Machine1CCountNum +
			record.Machine2CCountNum +
			record.Machine3CCountNum +
			record.Machine4CCountNum +
			record.Machine5CCountNum,
		CAmountGoukei: record.Machine1CAmountNum +
			record.Machine2CAmountNum +
			record.Machine3CAmountNum +
			record.Machine4CAmountNum +
			record.Machine5CAmountNum,
		CSettleCountGoukei: record.Machine1CSettleCountNum +
			record.Machine2CSettleCountNum +
			record.Machine3CSettleCountNum +
			record.Machine4CSettleCountNum +
			record.Machine5CSettleCountNum,
		CSettleAmountGoukei: record.Machine1CSettleAmountNum +
			record.Machine2CSettleAmountNum +
			record.Machine3CSettleAmountNum +
			record.Machine4CSettleAmountNum +
			record.Machine5CSettleAmountNum,
		CCountGoukeiGoukei: Machine1CCountGoukei +
			Machine2CCountGoukei +
			Machine3CCountGoukei +
			Machine4CCountGoukei +
			Machine5CCountGoukei,
		CAmountGoukeiGoukei: Machine1CAmountGoukei +
			Machine2CAmountGoukei +
			Machine3CAmountGoukei +
			Machine4CAmountGoukei +
			Machine5CAmountGoukei,
		Machine1QrCountGoukei:  Machine1QrCountGoukei,
		Machine2QrCountGoukei:  Machine2QrCountGoukei,
		Machine3QrCountGoukei:  Machine3QrCountGoukei,
		Machine4QrCountGoukei:  Machine4QrCountGoukei,
		Machine5QrCountGoukei:  Machine5QrCountGoukei,
		Machine1QrAmountGoukei: Machine1QrAmountGoukei,
		Machine2QrAmountGoukei: Machine2QrAmountGoukei,
		Machine3QrAmountGoukei: Machine3QrAmountGoukei,
		Machine4QrAmountGoukei: Machine4QrAmountGoukei,
		Machine5QrAmountGoukei: Machine5QrAmountGoukei,
		QrCountGoukei: record.Machine1QrCountNum +
			record.Machine2QrCountNum +
			record.Machine3QrCountNum +
			record.Machine4QrCountNum +
			record.Machine5QrCountNum,
		QrAmountGoukei: record.Machine1QrAmountNum +
			record.Machine2QrAmountNum +
			record.Machine3QrAmountNum +
			record.Machine4QrAmountNum +
			record.Machine5QrAmountNum,
		QrCountGoukeiGoukei: Machine1QrCountGoukei +
			Machine2QrCountGoukei +
			Machine3QrCountGoukei +
			Machine4QrCountGoukei +
			Machine5QrCountGoukei,
		QrAmountGoukeiGoukei: Machine1QrAmountGoukei +
			Machine2QrAmountGoukei +
			Machine3QrAmountGoukei +
			Machine4QrAmountGoukei +
			Machine5QrAmountGoukei,
		Machine1CashlessGoukei: Machine1CashlessGoukei,
		Machine2CashlessGoukei: Machine2CashlessGoukei,
		Machine3CashlessGoukei: Machine3CashlessGoukei,
		Machine4CashlessGoukei: Machine4CashlessGoukei,
		Machine5CashlessGoukei: Machine5CashlessGoukei,
		CashlessGoukei: Machine1CashlessGoukei +
			Machine2CashlessGoukei +
			Machine3CashlessGoukei +
			Machine4CashlessGoukei +
			Machine5CashlessGoukei,
		Machine1TotalGoukei: Machine1TotalGoukei,
		Machine2TotalGoukei: Machine2TotalGoukei,
		Machine3TotalGoukei: Machine3TotalGoukei,
		Machine4TotalGoukei: Machine4TotalGoukei,
		Machine5TotalGoukei: Machine5TotalGoukei,
		TotalGoukei:         TotalGoukei,
		NyuuyokuWariai:      helpers.CalculateWariai(float64(Machine1TotalGoukei+Machine2TotalGoukei), float64(TotalGoukei)),
		InshokuWariai:       helpers.CalculateWariai(float64(Machine3TotalGoukei+Machine4TotalGoukei), float64(TotalGoukei)),
		SeitaiWariai:        helpers.CalculateWariai(float64(Machine5TotalGoukei), float64(TotalGoukei)),
		Miseisan_Tousha:     Miseisan_Tousha,
		TounyuuGoukei:       TounyuuGoukei,
		Nyuukinyoteikingaku: Nyuukinyoteikingaku,
		Nyuukinyoteibi:      Nyuukinyoteibi,
	}

	record.JapaneseWeekday, _ = helpers.GetJapaneseWeekdayKanji(record.Date)

	// Prepare the data for the template
	data := ReportData{
		Date:   dateStr,
		Record: &reportRecordData,
		Found:  found,
	}

	return data, nil
}

func getNyuukinyouteibi(originalDate time.Time) string {
	var daysToAdd int

	// The time.Weekday type represents the day of the week.
	// time.Friday is "金" (Kinyoubi - Friday)
	// time.Thursday is "木" (Mokuyoubi - Thursday)
	switch originalDate.Weekday() {
	case time.Friday: // "金" - Friday
		daysToAdd = 3
	case time.Thursday: // "木" - Thursday
		daysToAdd = 4
	default: // All other days (Mon, Tue, Wed, Sat, Sun)
		daysToAdd = 2
	}

	// Add the calculated number of days to the original date.
	// time.Add(duration) is a good way to add days (24 * time.Hour).
	durationToAdd := time.Duration(daysToAdd) * 24 * time.Hour
	newDate := originalDate.Add(durationToAdd)

	dateBytes := []uint8(newDate.Format("2006-01-02"))
	JapaneseWeekdayKanji, err := helpers.GetJapaneseWeekdayKanji(dateBytes)
	if err != nil {
		return ""
	}

	return newDate.Format("2006年01月02日") + "（" + JapaneseWeekdayKanji + "）"

}
