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
	KaisuukenLoss            int
	SixKaisuukenLoss         int
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
	SCutGoukei               int
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

	Machine1CashCountGoukei = record.Machine1CashCount - record.Machine1SettleCount
	Machine2CashCountGoukei = record.Machine2CashCount - record.Machine2SettleCount
	Machine3CashCountGoukei = record.Machine3CashCount - record.Machine3SettleCount
	Machine4CashCountGoukei = record.Machine4CashCount - record.Machine4SettleCount
	Machine5CashCountGoukei = record.Machine5CashCount - record.Machine5SettleCount
	Machine1CashAmountGoukei = record.Machine1CashAmount - record.Machine1SettleAmount
	Machine2CashAmountGoukei = record.Machine2CashAmount - record.Machine2SettleAmount
	Machine3CashAmountGoukei = record.Machine3CashAmount - record.Machine3SettleAmount
	Machine4CashAmountGoukei = record.Machine4CashAmount - record.Machine4SettleAmount
	Machine5CashAmountGoukei = record.Machine5CashAmount - record.Machine5SettleAmount

	Machine1CashAmountGoukeiWithUnsettled := Machine1CashAmountGoukei - record.Machine1UnsettledAmount
	Machine2CashAmountGoukeiWithUnsettled := Machine2CashAmountGoukei - record.Machine2UnsettledAmount
	Machine3CashAmountGoukeiWithUnsettled := Machine3CashAmountGoukei - record.Machine3UnsettledAmount
	Machine4CashAmountGoukeiWithUnsettled := Machine4CashAmountGoukei - record.Machine4UnsettledAmount
	Machine5CashAmountGoukeiWithUnsettled := Machine5CashAmountGoukei - record.Machine5UnsettledAmount

	Machine1ECountGoukei = record.Machine1ECount - record.Machine1ESettleCount
	Machine2ECountGoukei = record.Machine2ECount - record.Machine2ESettleCount
	Machine3ECountGoukei = record.Machine3ECount - record.Machine3ESettleCount
	Machine4ECountGoukei = record.Machine4ECount - record.Machine4ESettleCount
	Machine5ECountGoukei = record.Machine5ECount - record.Machine5ESettleCount
	Machine1EAmountGoukei = record.Machine1EAmount - record.Machine1ESettleAmount
	Machine2EAmountGoukei = record.Machine2EAmount - record.Machine2ESettleAmount
	Machine3EAmountGoukei = record.Machine3EAmount - record.Machine3ESettleAmount
	Machine4EAmountGoukei = record.Machine4EAmount - record.Machine4ESettleAmount
	Machine5EAmountGoukei = record.Machine5EAmount - record.Machine5ESettleAmount

	Machine1CCountGoukei = record.Machine1CCount - record.Machine1CSettleCount
	Machine2CCountGoukei = record.Machine2CCount - record.Machine2CSettleCount
	Machine3CCountGoukei = record.Machine3CCount - record.Machine3CSettleCount
	Machine4CCountGoukei = record.Machine4CCount - record.Machine4CSettleCount
	Machine5CCountGoukei = record.Machine5CCount - record.Machine5CSettleCount
	Machine1CAmountGoukei = record.Machine1CAmount - record.Machine1CSettleAmount
	Machine2CAmountGoukei = record.Machine2CAmount - record.Machine2CSettleAmount
	Machine3CAmountGoukei = record.Machine3CAmount - record.Machine3CSettleAmount
	Machine4CAmountGoukei = record.Machine4CAmount - record.Machine4CSettleAmount
	Machine5CAmountGoukei = record.Machine5CAmount - record.Machine5CSettleAmount

	Machine1QrCountGoukei = record.Machine1QrCount - record.Machine1QrSettleCount
	Machine2QrCountGoukei = record.Machine2QrCount - record.Machine2QrSettleCount
	Machine3QrCountGoukei = record.Machine3QrCount - record.Machine3QrSettleCount
	Machine4QrCountGoukei = record.Machine4QrCount - record.Machine4QrSettleCount
	Machine5QrCountGoukei = record.Machine5QrCount - record.Machine5QrSettleCount
	Machine1QrAmountGoukei = record.Machine1QrAmount - record.Machine1QrSettleAmount
	Machine2QrAmountGoukei = record.Machine2QrAmount - record.Machine2QrSettleAmount
	Machine3QrAmountGoukei = record.Machine3QrAmount - record.Machine3QrSettleAmount
	Machine4QrAmountGoukei = record.Machine4QrAmount - record.Machine4QrSettleAmount
	Machine5QrAmountGoukei = record.Machine5QrAmount - record.Machine5QrSettleAmount

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
		record.Machine1UnsettledAmount +
			record.Machine2UnsettledAmount +
			record.Machine3UnsettledAmount +
			record.Machine4UnsettledAmount +
			record.Machine5UnsettledAmount

	TounyuuGoukei := TotalGoukei + record.Change - Miseisan_Tousha + record.PhoneFee - record.HonjitsuMitounyuuAmountUncertain - record.HonjitsuMitounyuuAmountCertain + record.Deficiency + record.ZenjitsuMitounyuuAmount

	originalDate, err := time.Parse("2006-01-02", record.DateString)
	if err != nil {
		return ReportData{}, fmt.Errorf("failed to parse date: %w", err)
	}

	Nyuukinyoteibi := getNyuukinyouteibi(originalDate)

	Nyuukinyoteikingaku := TounyuuGoukei - record.Change

	reportRecordData := ReportRecordData{
		DailyReportRaw: *record,
		CashCountGoukei: record.Machine1CashCount +
			record.Machine2CashCount +
			record.Machine3CashCount +
			record.Machine4CashCount +
			record.Machine5CashCount,
		CashAmountGoukei: record.Machine1CashAmount +
			record.Machine2CashAmount +
			record.Machine3CashAmount +
			record.Machine4CashAmount +
			record.Machine5CashAmount,
		SettleCountGoukei: record.Machine1SettleCount +
			record.Machine2SettleCount +
			record.Machine3SettleCount +
			record.Machine4SettleCount +
			record.Machine5SettleCount,
		SettleAmountGoukei: record.Machine1SettleAmount +
			record.Machine2SettleAmount +
			record.Machine3SettleAmount +
			record.Machine4SettleAmount +
			record.Machine5SettleAmount,
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
		UnsettledCountGoukei: record.Machine1UnsettledCount +
			record.Machine2UnsettledCount +
			record.Machine3UnsettledCount +
			record.Machine4UnsettledCount +
			record.Machine5UnsettledCount,
		UnsettledAmountGoukei: record.Machine1UnsettledAmount +
			record.Machine2UnsettledAmount +
			record.Machine3UnsettledAmount +
			record.Machine4UnsettledAmount +
			record.Machine5UnsettledAmount,
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
		ECountGoukei: record.Machine1ECount +
			record.Machine2ECount +
			record.Machine3ECount +
			record.Machine4ECount +
			record.Machine5ECount,
		EAmountGoukei: record.Machine1EAmount +
			record.Machine2EAmount +
			record.Machine3EAmount +
			record.Machine4EAmount +
			record.Machine5EAmount,
		ESettleCountGoukei: record.Machine1ESettleCount +
			record.Machine2ESettleCount +
			record.Machine3ESettleCount +
			record.Machine4ESettleCount +
			record.Machine5ESettleCount,
		ESettleAmountGoukei: record.Machine1ESettleAmount +
			record.Machine2ESettleAmount +
			record.Machine3ESettleAmount +
			record.Machine4ESettleAmount +
			record.Machine5ESettleAmount,
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
		CCountGoukei: record.Machine1CCount +
			record.Machine2CCount +
			record.Machine3CCount +
			record.Machine4CCount +
			record.Machine5CCount,
		CAmountGoukei: record.Machine1CAmount +
			record.Machine2CAmount +
			record.Machine3CAmount +
			record.Machine4CAmount +
			record.Machine5CAmount,
		CSettleCountGoukei: record.Machine1CSettleCount +
			record.Machine2CSettleCount +
			record.Machine3CSettleCount +
			record.Machine4CSettleCount +
			record.Machine5CSettleCount,
		CSettleAmountGoukei: record.Machine1CSettleAmount +
			record.Machine2CSettleAmount +
			record.Machine3CSettleAmount +
			record.Machine4CSettleAmount +
			record.Machine5CSettleAmount,
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
		QrCountGoukei: record.Machine1QrCount +
			record.Machine2QrCount +
			record.Machine3QrCount +
			record.Machine4QrCount +
			record.Machine5QrCount,
		QrAmountGoukei: record.Machine1QrAmount +
			record.Machine2QrAmount +
			record.Machine3QrAmount +
			record.Machine4QrAmount +
			record.Machine5QrAmount,
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
		KaisuukenLoss:       record.TicketSalesCount - (record.SalesNoEnd - record.SalesNoStart + 1),
		SixKaisuukenLoss:    record.SixTicketSalesCount - (record.SixSalesNoEnd - record.SixSalesNoStart + 1),
		SCutGoukei:          record.SCutMale + record.SCutFemale + record.SCutChild,
	}

	reportRecordData.DailyReportRaw.JapaneseWeekday, _ = helpers.GetJapaneseWeekdayKanji(reportRecordData.DailyReportRaw.Date)
	//record.JapaneseWeekday, _ = helpers.GetJapaneseWeekdayKanji(record.Date)

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
