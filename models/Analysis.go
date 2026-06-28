package models

import (
	"time"
)

type DailyAnalysisData struct {
	Date              string `json:"date"`
	HotbathRevenue    int    `json:"hotbath_revenue"`
	RestaurantRevenue int    `json:"restaurant_revenue"`
	RelaxationRevenue int    `json:"relaxation_revenue"`
	PowerExpense      int    `json:"power_expense"`
	GasExpense        int    `json:"gas_expense"`
	VisitorCount      int    `json:"visitor_count"`
}

func GetAnalysisDataRange(start, end time.Time) ([]DailyAnalysisData, error) {
	// We need data from start-1 to calculate the first day's diff for utilities
	baselineDate := start.AddDate(0, 0, -1)

	// Fetch all daily reports in range
	reports, err := getDailyReportsInRange(baselineDate, end)
	if err != nil {
		return nil, err
	}

	// Fetch all power readings in range
	powerReadings, err := getPowerReadingsInRange(baselineDate, end)
	if err != nil {
		return nil, err
	}

	// Fetch all gas readings in range
	gasReadings, err := getGasReadingsInRange(baselineDate, end)
	if err != nil {
		return nil, err
	}

	var results []DailyAnalysisData

	// Map data by date string for easy lookup
	reportMap := make(map[string]*DailyReportRaw)
	for _, r := range reports {
		reportMap[string(r.Date)] = r
	}

	powerMap := make(map[string]float64)
	for _, p := range powerReadings {
		dateStr := time.Date(p.Year, time.Month(p.Month), p.Day, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
		powerMap[dateStr] = p.PowerReading
	}

	gasMap := make(map[string]struct {
		b1, b2, b3, b4 int
	})
	for _, g := range gasReadings {
		dateStr := time.Date(g.Year, time.Month(g.Month), g.Day, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
		gasMap[dateStr] = struct{ b1, b2, b3, b4 int }{g.Boiler1Reading, g.Boiler2Reading, g.Boiler3Reading, g.Boiler4Reading}
	}

	// Iterate from start to end
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dateStr := d.Format("2006-01-02")
		prevDateStr := d.AddDate(0, 0, -1).Format("2006-01-02")

		analysis := DailyAnalysisData{
			Date: dateStr,
		}

		// Revenue and Visitors from DailyReportRaw
		if r, ok := reportMap[dateStr]; ok {
			analysis.HotbathRevenue = calculateMachineRevenue(r, 1) + calculateMachineRevenue(r, 2)
			analysis.RestaurantRevenue = calculateMachineRevenue(r, 3) + calculateMachineRevenue(r, 4)
			analysis.RelaxationRevenue = calculateMachineRevenue(r, 5)
			analysis.VisitorCount = r.AdultTicketCount + r.AdultSetTicketCount + r.ChildTicketCount + r.TicketCount + r.InvitationTicketCount + r.YuutaikenTicketCount + r.SixTicketCount + r.InfantTicketCount + r.PointCardAdultCount + r.PointCardChildCount + r.OldTicketCount
		}

		// Power Expense
		if currP, ok := powerMap[dateStr]; ok {
			if prevP, ok := powerMap[prevDateStr]; ok {
				diff := currP - prevP
				// Logic from assets/power-readings/index.js
				// For 2026/04/21 to 2026/05/20, it uses the Math.floor(diff / 10) / 100 logic
				usageDiff := float64(int(diff/10)) / 100.0
				value := usageDiff * 600
				analysis.PowerExpense = int(value * 25.77)
			}
		}

		// Gas Expense
		if currG, ok := gasMap[dateStr]; ok {
			if prevG, ok := gasMap[prevDateStr]; ok {
				diff := (currG.b1 - prevG.b1) + (currG.b2 - prevG.b2) + (currG.b3 - prevG.b3) + (currG.b4 - prevG.b4)
				// Using placeholder price 300 per unit
				analysis.GasExpense = diff * 300
			}
		}

		results = append(results, analysis)
	}

	return results, nil
}

func calculateMachineRevenue(r *DailyReportRaw, machineNum int) int {
	switch machineNum {
	case 1:
		return (r.Machine1CashAmount - r.Machine1CashSettleAmount - r.Machine1UnsettledAmount) +
			(r.Machine1QrAmount - r.Machine1QrSettleAmount) +
			(r.Machine1EAmount - r.Machine1ESettleAmount) +
			(r.Machine1CAmount - r.Machine1CSettleAmount)
	case 2:
		return (r.Machine2CashAmount - r.Machine2CashSettleAmount - r.Machine2UnsettledAmount) +
			(r.Machine2QrAmount - r.Machine2QrSettleAmount) +
			(r.Machine2EAmount - r.Machine2ESettleAmount) +
			(r.Machine2CAmount - r.Machine2CSettleAmount)
	case 3:
		return (r.Machine3CashAmount - r.Machine3CashSettleAmount - r.Machine3UnsettledAmount) +
			(r.Machine3QrAmount - r.Machine3QrSettleAmount) +
			(r.Machine3EAmount - r.Machine3ESettleAmount) +
			(r.Machine3CAmount - r.Machine3CSettleAmount)
	case 4:
		return (r.Machine4CashAmount - r.Machine4CashSettleAmount - r.Machine4UnsettledAmount) +
			(r.Machine4QrAmount - r.Machine4QrSettleAmount) +
			(r.Machine4EAmount - r.Machine4ESettleAmount) +
			(r.Machine4CAmount - r.Machine4CSettleAmount)
	case 5:
		return (r.Machine5CashAmount - r.Machine5CashSettleAmount - r.Machine5UnsettledAmount) +
			(r.Machine5QrAmount - r.Machine5QrSettleAmount) +
			(r.Machine5EAmount - r.Machine5ESettleAmount) +
			(r.Machine5CAmount - r.Machine5CSettleAmount)
	}
	return 0
}

func getDailyReportsInRange(start, end time.Time) ([]*DailyReportRaw, error) {
	query := `
	SELECT
		T1.ID, T1.日付, T1.1号機現金金額, T1.1号機精算金額, T1.1号機未精算金額,
		T1.1号機QR金額, T1.1号機QR精算金額, T1.1号機電子マネ金額, T1.1号機電子マネ精算金額,
		T1.1号機クレジット金額, T1.1号機クレジット精算金額,
		T1.2号機現金金額, T1.2号機精算金額, T1.2号機未精算金額,
		T1.2号機QR金額, T1.2号機QR精算金額, T1.2号機電子マネ金額, T1.2号機電子マネ精算金額,
		T1.2号機クレジット金額, T1.2号機クレジット精算金額,
		T1.3号機現金金額, T1.3号機精算金額, T1.3号機未精算金額,
		T1.3号機QR金額, T1.3号機QR精算金額, T1.3号機電子マネ金額, T1.3号機電子マネ精算金額,
		T1.3号機クレジット金額, T1.3号機クレジット精算金額,
		T1.4号機現金金額, T1.4号機精算金額, T1.4号機未精算金額,
		T1.4号機QR金額, T1.4号機QR精算金額, T1.4号機電子マネ金額, T1.4号機電子マネ精算金額,
		T1.4号機クレジット金額, T1.4号機クレジット精算金額,
		T1.5号機現金金額, T1.5号機精算金額, T1.5号機未精算金額,
		T1.5号機QR金額, T1.5号機QR精算金額, T1.5号機電子マネ金額, T1.5号機電子マネ精算金額,
		T1.5号機クレジット金額, T1.5号機クレジット精算金額,
		T1.大人入浴券枚数, T1.大人入浴セット券枚数, T1.小人入浴券枚数, T1.回数券回収,
		T1.招待券回収, T1.優待券回収, T1.感謝祭招待券回収, T1.ﾎﾟｲﾝﾄｶｰﾄﾞ大人回収,
		T1.ﾎﾟｲﾝﾄｶｰﾄﾞﾞ小人回収, T1.過去回数券回収, T1.6回数券回収
	FROM
		日次報告ﾃｰﾌﾞﾙ AS T1
	WHERE
		T1.日付 BETWEEN ? AND ?
	`
	rows, err := db.Query(query, start.Format("2006-01-02"), end.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []*DailyReportRaw
	for rows.Next() {
		var r DailyReportRaw
		err := rows.Scan(
			&r.ID, &r.Date, &r.Machine1CashAmount, &r.Machine1CashSettleAmount, &r.Machine1UnsettledAmount,
			&r.Machine1QrAmount, &r.Machine1QrSettleAmount, &r.Machine1EAmount, &r.Machine1ESettleAmount,
			&r.Machine1CAmount, &r.Machine1CSettleAmount,
			&r.Machine2CashAmount, &r.Machine2CashSettleAmount, &r.Machine2UnsettledAmount,
			&r.Machine2QrAmount, &r.Machine2QrSettleAmount, &r.Machine2EAmount, &r.Machine2ESettleAmount,
			&r.Machine2CAmount, &r.Machine2CSettleAmount,
			&r.Machine3CashAmount, &r.Machine3CashSettleAmount, &r.Machine3UnsettledAmount,
			&r.Machine3QrAmount, &r.Machine3QrSettleAmount, &r.Machine3EAmount, &r.Machine3ESettleAmount,
			&r.Machine3CAmount, &r.Machine3CSettleAmount,
			&r.Machine4CashAmount, &r.Machine4CashSettleAmount, &r.Machine4UnsettledAmount,
			&r.Machine4QrAmount, &r.Machine4QrSettleAmount, &r.Machine4EAmount, &r.Machine4ESettleAmount,
			&r.Machine4CAmount, &r.Machine4CSettleAmount,
			&r.Machine5CashAmount, &r.Machine5CashSettleAmount, &r.Machine5UnsettledAmount,
			&r.Machine5QrAmount, &r.Machine5QrSettleAmount, &r.Machine5EAmount, &r.Machine5ESettleAmount,
			&r.Machine5CAmount, &r.Machine5CSettleAmount,
			&r.AdultTicketCount, &r.AdultSetTicketCount, &r.ChildTicketCount, &r.TicketCount,
			&r.InvitationTicketCount, &r.YuutaikenTicketCount, &r.InfantTicketCount, &r.PointCardAdultCount,
			&r.PointCardChildCount, &r.OldTicketCount, &r.SixTicketCount,
		)
		if err != nil {
			return nil, err
		}
		reports = append(reports, &r)
	}
	return reports, nil
}

func getPowerReadingsInRange(start, end time.Time) ([]PowerReading, error) {
	query := `
	SELECT year, month, day, power_reading
	FROM daily_power_readings
	WHERE (year * 10000 + month * 100 + day) BETWEEN ? AND ?
	`
	startInt := start.Year()*10000 + int(start.Month())*100 + start.Day()
	endInt := end.Year()*10000 + int(end.Month())*100 + end.Day()

	rows, err := db.Query(query, startInt, endInt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var readings []PowerReading
	for rows.Next() {
		var p PowerReading
		if err := rows.Scan(&p.Year, &p.Month, &p.Day, &p.PowerReading); err != nil {
			return nil, err
		}
		readings = append(readings, p)
	}
	return readings, nil
}

func getGasReadingsInRange(start, end time.Time) ([]GasReading, error) {
	query := `
	SELECT year, month, day, boiler1_reading, boiler2_reading, boiler3_reading, boiler4_reading
	FROM daily_gas_readings
	WHERE (year * 10000 + month * 100 + day) BETWEEN ? AND ?
	`
	startInt := start.Year()*10000 + int(start.Month())*100 + start.Day()
	endInt := end.Year()*10000 + int(end.Month())*100 + end.Day()

	rows, err := db.Query(query, startInt, endInt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var readings []GasReading
	for rows.Next() {
		var g GasReading
		if err := rows.Scan(&g.Year, &g.Month, &g.Day, &g.Boiler1Reading, &g.Boiler2Reading, &g.Boiler3Reading, &g.Boiler4Reading); err != nil {
			return nil, err
		}
		readings = append(readings, g)
	}
	return readings, nil
}
