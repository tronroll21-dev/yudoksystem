package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

/*
データベースから取得する元データを格納する構造体
*/
type DailyReportRaw struct {
	ID                               int     `db:"ID"`
	Date                             []uint8 `db:"日付" json:"-"`
	DateString                       string
	JapaneseWeekday                  string `json:"-"`
	WeatherCode                      int    `db:"天気ｺｰﾄﾞ"`
	WeatherCondition                 string `db:"天気状況" json:"-"`
	WeatherMark                      string `db:"天気ﾏｰｸ" json:"-"`
	StaffCode                        int    `db:"担当ｺｰﾄﾞ"`
	StaffName                        string `db:"担当" json:"-"`
	Machine1CashCount                int    `db:"1号機現金枚数"`
	Machine1CashAmount               int    `db:"1号機現金金額"`
	Machine1SettleCount              int    `db:"1号機精算枚数"`
	Machine1SettleAmount             int    `db:"1号機精算金額"`
	Machine2CashCount                int    `db:"2号機現金枚数"`
	Machine2CashAmount               int    `db:"2号機現金金額"`
	Machine2SettleCount              int    `db:"2号機精算枚数"`
	Machine2SettleAmount             int    `db:"2号機精算金額"`
	Machine3CashCount                int    `db:"3号機現金枚数"`
	Machine3CashAmount               int    `db:"3号機現金金額"`
	Machine3SettleCount              int    `db:"3号機精算枚数"`
	Machine3SettleAmount             int    `db:"3号機精算金額"`
	Machine4CashCount                int    `db:"4号機現金枚数"`
	Machine4CashAmount               int    `db:"4号機現金金額"`
	Machine4SettleCount              int    `db:"4号機精算枚数"`
	Machine4SettleAmount             int    `db:"4号機精算金額"`
	Machine5CashCount                int    `db:"5号機現金枚数"`
	Machine5CashAmount               int    `db:"5号機現金金額"`
	Machine5SettleCount              int    `db:"5号機精算枚数"`
	Machine5SettleAmount             int    `db:"5号機精算金額"`
	Machine1UnsettledCount           int    `db:"1号機未精算枚数"`
	Machine1UnsettledAmount          int    `db:"1号機未精算金額"`
	Machine2UnsettledCount           int    `db:"2号機未精算枚数"`
	Machine2UnsettledAmount          int    `db:"2号機未精算金額"`
	Machine3UnsettledCount           int    `db:"3号機未精算枚数"`
	Machine3UnsettledAmount          int    `db:"3号機未精算金額"`
	Machine4UnsettledCount           int    `db:"4号機未精算枚数"`
	Machine4UnsettledAmount          int    `db:"4号機未精算金額"`
	Machine5UnsettledCount           int    `db:"5号機未精算枚数"`
	Machine5UnsettledAmount          int    `db:"5号機未精算金額"`
	Machine1QrCount                  int    `db:"1号機QR件数"`
	Machine1QrAmount                 int    `db:"1号機QR金額"`
	Machine2QrCount                  int    `db:"2号機QR件数"`
	Machine2QrAmount                 int    `db:"2号機QR金額"`
	Machine3QrCount                  int    `db:"3号機QR件数"`
	Machine3QrAmount                 int    `db:"3号機QR金額"`
	Machine4QrCount                  int    `db:"4号機QR件数"`
	Machine4QrAmount                 int    `db:"4号機QR金額"`
	Machine5QrCount                  int    `db:"5号機QR件数"`
	Machine5QrAmount                 int    `db:"5号機QR金額"`
	Machine1QrSettleCount            int    `db:"1号機QR精算件数"`
	Machine1QrSettleAmount           int    `db:"1号機QR精算金額"`
	Machine2QrSettleCount            int    `db:"2号機QR精算件数"`
	Machine2QrSettleAmount           int    `db:"2号機QR精算金額"`
	Machine3QrSettleCount            int    `db:"3号機QR精算件数"`
	Machine3QrSettleAmount           int    `db:"3号機QR精算金額"`
	Machine4QrSettleCount            int    `db:"4号機QR精算件数"`
	Machine4QrSettleAmount           int    `db:"4号機QR精算金額"`
	Machine5QrSettleCount            int    `db:"5号機QR精算件数"`
	Machine5QrSettleAmount           int    `db:"5号機QR精算金額"`
	Machine1ECount                   int    `db:"1号機電子マネ件数"`
	Machine1EAmount                  int    `db:"1号機電子マネ金額"`
	Machine2ECount                   int    `db:"2号機電子マネ件数"`
	Machine2EAmount                  int    `db:"2号機電子マネ金額"`
	Machine3ECount                   int    `db:"3号機電子マネ件数"`
	Machine3EAmount                  int    `db:"3号機電子マネ金額"`
	Machine4ECount                   int    `db:"4号機電子マネ件数"`
	Machine4EAmount                  int    `db:"4号機電子マネ金額"`
	Machine5ECount                   int    `db:"5号機電子マネ件数"`
	Machine5EAmount                  int    `db:"5号機電子マネ金額"`
	Machine1ESettleCount             int    `db:"1号機電子マネ精算件数"`
	Machine1ESettleAmount            int    `db:"1号機電子マネ精算金額"`
	Machine2ESettleCount             int    `db:"2号機電子マネ精算件数"`
	Machine2ESettleAmount            int    `db:"2号機電子マネ精算金額"`
	Machine3ESettleCount             int    `db:"3号機電子マネ精算件数"`
	Machine3ESettleAmount            int    `db:"3号機電子マネ精算金額"`
	Machine4ESettleCount             int    `db:"4号機電子マネ精算件数"`
	Machine4ESettleAmount            int    `db:"4号機電子マネ精算金額"`
	Machine5ESettleCount             int    `db:"5号機電子マネ精算件数"`
	Machine5ESettleAmount            int    `db:"5号機電子マネ精算金額"`
	Machine1CCount                   int    `db:"1号機クレジット件数"`
	Machine1CAmount                  int    `db:"1号機クレジット金額"`
	Machine2CCount                   int    `db:"2号機クレジット件数"`
	Machine2CAmount                  int    `db:"2号機クレジット金額"`
	Machine3CCount                   int    `db:"3号機クレジット件数"`
	Machine3CAmount                  int    `db:"3号機クレジット金額"`
	Machine4CCount                   int    `db:"4号機クレジット件数"`
	Machine4CAmount                  int    `db:"4号機クレジット金額"`
	Machine5CCount                   int    `db:"5号機クレジット件数"`
	Machine5CAmount                  int    `db:"5号機クレジット金額"`
	Machine1CSettleCount             int    `db:"1号機クレジット精算件数"`
	Machine1CSettleAmount            int    `db:"1号機クレジット精算金額"`
	Machine2CSettleCount             int    `db:"2号機クレジット精算件数"`
	Machine2CSettleAmount            int    `db:"2号機クレジット精算金額"`
	Machine3CSettleCount             int    `db:"3号機クレジット精算件数"`
	Machine3CSettleAmount            int    `db:"3号機クレジット精算金額"`
	Machine4CSettleCount             int    `db:"4号機クレジット精算件数"`
	Machine4CSettleAmount            int    `db:"4号機クレジット精算金額"`
	Machine5CSettleCount             int    `db:"5号機クレジット精算件数"`
	Machine5CSettleAmount            int    `db:"5号機クレジット精算金額"`
	AdultTicketCount                 int    `db:"大人入浴券枚数"`
	AdultSetTicketCount              int    `db:"大人入浴セット券枚数"`
	ChildTicketCount                 int    `db:"小人入浴券枚数"`
	TicketCount                      int    `db:"回数券回収"`
	InvitationTicketCount            int    `db:"招待券回収"`
	YuutaikenTicketCount             int    `db:"優待券回収"`
	InfantTicketCount                int    `db:"感謝祭招待券回収"`
	PointCardAdultCount              int    `db:"ﾎﾟｲﾝﾄｶｰﾄﾞ大人回収"`
	PointCardChildCount              int    `db:"ﾎﾟｲﾝﾄｶｰﾄﾞﾞ小人回収"`
	TicketSalesCount                 int    `db:"回数券売数"`
	OldTicketCount                   int    `db:"過去回数券回収"`
	SalesNoStart                     int    `db:"販売NO（始まり）"`
	SalesNoEnd                       int    `db:"販売NO（終わり）"`
	Change                           int    `db:"釣銭（ﾌﾛﾝﾄ含む）"`
	PhoneFee                         int    `db:"電話料投入"`
	Teuri                            int    `db:"手売（金額）"`
	YuutaikenSalesCount              int    `db:"優待券販売（枚数）"`
	YuutaikenSalesAmount             int    `db:"優待券販売（金額）"`
	HonjitsuMitounyuuAmountUncertain int    `db:"本日未投入金額（未確定）"`
	HonjitsuMitounyuuAmountCertain   int    `db:"本日未投入金額（確定）"`
	Deficiency                       int    `db:"過不足"`
	ZenjitsuMitounyuuAmount          int    `db:"前日未投入金額"`
	ReceiptTotalAmount               int    `db:"投入ﾚｼｰﾄ合計金額"`
	Remarks                          string `db:"備考"`
	SCutMale                         int    `db:"ｴｽｶｯﾄ男性"`
	SCutFemale                       int    `db:"ｴｽｶｯﾄ女性"`
	SCutChild                        int    `db:"ｴｽｶｯﾄ子供"`
	SixTicketCount                   int    `db:"6回数券回収"`
	SixTicketSalesCount              int    `db:"6回数券売数"`
	SixSalesNoStart                  int    `db:"6販売NO（始まり）"`
	SixSalesNoEnd                    int    `db:"6販売NO（終わり）"`
	MaleTicketCount                  int    `db:"男回数券枚数"`
	FemaleTicketCount                int    `db:"女回数券枚数"`
	SixMaleTicketCount               int    `db:"6男回数券枚数"`
	SixFemaleTicketCount             int    `db:"6女回数券枚数"`
	CouponCount                      int    `db:"ｸｰﾎﾟﾝ回収"`
	RearegiAmount                    int    `db:"リアレジ金額"`
	RearegiTicketAmount              int    `db:"リアレジ金額回数券系"`
	RearegiRelaxAmount               int    `db:"リアレジ金額リラク系"`
	ReportSpace                      string `db:"報告ｽﾍﾟｰｽ"`
}

// GetSalesRecordByDate queries the mock database for a record on the specified date
func GetSalesRecordByDate(targetDate time.Time) (*DailyReportRaw, bool, error) {

	formattedDate := targetDate.Format("2006-01-02")

	query := `
	SELECT
		T1.ID,
	T1.日付,
	T1.天気コード,
	T1.担当コード,
	T1.1号機現金枚数,
	T1.1号機現金金額,
	T1.1号機精算枚数,
	T1.1号機精算金額,
	T1.2号機現金枚数,
	T1.2号機現金金額,
	T1.2号機精算枚数,
	T1.2号機精算金額,
	T1.3号機現金枚数,
	T1.3号機現金金額,
	T1.3号機精算枚数,
	T1.3号機精算金額,
	T1.4号機現金枚数,
	T1.4号機現金金額,
	T1.4号機精算枚数,
	T1.4号機精算金額,
	T1.5号機現金枚数,
	T1.5号機現金金額,
	T1.5号機精算枚数,
	T1.5号機精算金額,
	T1.1号機未精算枚数,
	T1.1号機未精算金額,
	T1.2号機未精算枚数,
	T1.2号機未精算金額,
	T1.3号機未精算枚数,
	T1.3号機未精算金額,
	T1.4号機未精算枚数,
	T1.4号機未精算金額,
	T1.5号機未精算枚数,
	T1.5号機未精算金額,
	T1.1号機QR件数,
	T1.1号機QR金額,
	T1.2号機QR件数,
	T1.2号機QR金額,
	T1.3号機QR件数,
	T1.3号機QR金額,
	T1.4号機QR件数,
	T1.4号機QR金額,
	T1.5号機QR件数,
	T1.5号機QR金額,
	T1.1号機QR精算件数,
	T1.1号機QR精算金額,
	T1.2号機QR精算件数,
	T1.2号機QR精算金額,
	T1.3号機QR精算件数,
	T1.3号機QR精算金額,
	T1.4号機QR精算件数,
	T1.4号機QR精算金額,
	T1.5号機QR精算件数,
	T1.5号機QR精算金額,
	T1.1号機電子マネ件数,
	T1.1号機電子マネ金額,
	T1.2号機電子マネ件数,
	T1.2号機電子マネ金額,
	T1.3号機電子マネ件数,
	T1.3号機電子マネ金額,
	T1.4号機電子マネ件数,
	T1.4号機電子マネ金額,
	T1.5号機電子マネ件数,
	T1.5号機電子マネ金額,
	T1.1号機電子マネ精算件数,
	T1.1号機電子マネ精算金額,
	T1.2号機電子マネ精算件数,
	T1.2号機電子マネ精算金額,
	T1.3号機電子マネ精算件数,
	T1.3号機電子マネ精算金額,
	T1.4号機電子マネ精算件数,
	T1.4号機電子マネ精算金額,
	T1.5号機電子マネ精算件数,
	T1.5号機電子マネ精算金額,
	T1.1号機クレジット件数,
	T1.1号機クレジット金額,
	T1.2号機クレジット件数,
	T1.2号機クレジット金額,
	T1.3号機クレジット件数,
	T1.3号機クレジット金額,
	T1.4号機クレジット件数,
	T1.4号機クレジット金額,
	T1.5号機クレジット件数,
	T1.5号機クレジット金額,
	T1.1号機クレジット精算件数,
	T1.1号機クレジット精算金額,
	T1.2号機クレジット精算件数,
	T1.2号機クレジット精算金額,
	T1.3号機クレジット精算件数,
	T1.3号機クレジット精算金額,
	T1.4号機クレジット精算件数,
	T1.4号機クレジット精算金額,
	T1.5号機クレジット精算件数,
	T1.5号機クレジット精算金額,
	T1.大人入浴券枚数,
	T1.大人入浴セット券枚数,
	T1.小人入浴券枚数,
	T1.回数券回収,
	T1.招待券回収,
	T1.優待券回収,
	T1.感謝祭招待券回収,
	T1.ﾎﾟｲﾝﾄｶｰﾄﾞ大人回収,
	T1.ﾎﾟｲﾝﾄｶｰﾄﾞﾞ小人回収,
	T1.回数券売数,
	T1.過去回数券回収,
	T1.販売NO（始まり）,
	T1.販売NO（終わり）,
	T1.釣銭（ﾌﾛﾝﾄ含む）,
	T1.電話料投入,
	T1.手売（金額）,
	T1.優待券販売（枚数）,
	T1.優待券販売（金額）,
	T1.本日未投入金額（未確定）,
	T1.本日未投入金額（確定）,
	T1.過不足,
	T1.前日未投入金額,
	T1.投入ﾚｼｰﾄ合計金額,
	T1.備考,
	T1.ｴｽｶｯﾄ男性,
	T1.ｴｽｶｯﾄ女性,
	T1.ｴｽｶｯﾄ子供,
	T1.6回数券回収,
	T1.6回数券売数,
	T1.6販売NO（始まり）,
	T1.6販売NO（終わり）,
	T1.男回数券枚数,
	T1.女回数券枚数,
	T1.6男回数券枚数,
	T1.6女回数券枚数,
	T1.ｸｰﾎﾟﾝ回収,
	T1.リアレジ金額,
	T1.リアレジ金額回数券系,
	T1.リアレジ金額リラク系,
	T1.報告ｽﾍﾟｰｽ
	FROM
		日次報告ﾃｰﾌﾞﾙ AS T1
        INNER JOIN 天気ﾏｽﾀ tm ON T1.天気コード = tm.天気コード
	WHERE
		T1.日付 = ?
	ORDER BY
		T1.ID;
	`
	var raw DailyReportRaw
	var dateString string
	err := db.QueryRow(query, formattedDate).Scan(
		&raw.ID,
		&raw.Date,
		&raw.WeatherCode,
		&raw.StaffCode,
		&raw.Machine1CashCount,
		&raw.Machine1CashAmount,
		&raw.Machine1SettleCount,
		&raw.Machine1SettleAmount,
		&raw.Machine2CashCount,
		&raw.Machine2CashAmount,
		&raw.Machine2SettleCount,
		&raw.Machine2SettleAmount,
		&raw.Machine3CashCount,
		&raw.Machine3CashAmount,
		&raw.Machine3SettleCount,
		&raw.Machine3SettleAmount,
		&raw.Machine4CashCount,
		&raw.Machine4CashAmount,
		&raw.Machine4SettleCount,
		&raw.Machine4SettleAmount,
		&raw.Machine5CashCount,
		&raw.Machine5CashAmount,
		&raw.Machine5SettleCount,
		&raw.Machine5SettleAmount,
		&raw.Machine1UnsettledCount,
		&raw.Machine1UnsettledAmount,
		&raw.Machine2UnsettledCount,
		&raw.Machine2UnsettledAmount,
		&raw.Machine3UnsettledCount,
		&raw.Machine3UnsettledAmount,
		&raw.Machine4UnsettledCount,
		&raw.Machine4UnsettledAmount,
		&raw.Machine5UnsettledCount,
		&raw.Machine5UnsettledAmount,
		&raw.Machine1QrCount,
		&raw.Machine1QrAmount,
		&raw.Machine2QrCount,
		&raw.Machine2QrAmount,
		&raw.Machine3QrCount,
		&raw.Machine3QrAmount,
		&raw.Machine4QrCount,
		&raw.Machine4QrAmount,
		&raw.Machine5QrCount,
		&raw.Machine5QrAmount,
		&raw.Machine1QrSettleCount,
		&raw.Machine1QrSettleAmount,
		&raw.Machine2QrSettleCount,
		&raw.Machine2QrSettleAmount,
		&raw.Machine3QrSettleCount,
		&raw.Machine3QrSettleAmount,
		&raw.Machine4QrSettleCount,
		&raw.Machine4QrSettleAmount,
		&raw.Machine5QrSettleCount,
		&raw.Machine5QrSettleAmount,
		&raw.Machine1ECount,
		&raw.Machine1EAmount,
		&raw.Machine2ECount,
		&raw.Machine2EAmount,
		&raw.Machine3ECount,
		&raw.Machine3EAmount,
		&raw.Machine4ECount,
		&raw.Machine4EAmount,
		&raw.Machine5ECount,
		&raw.Machine5EAmount,
		&raw.Machine1ESettleCount,
		&raw.Machine1ESettleAmount,
		&raw.Machine2ESettleCount,
		&raw.Machine2ESettleAmount,
		&raw.Machine3ESettleCount,
		&raw.Machine3ESettleAmount,
		&raw.Machine4ESettleCount,
		&raw.Machine4ESettleAmount,
		&raw.Machine5ESettleCount,
		&raw.Machine5ESettleAmount,
		&raw.Machine1CCount,
		&raw.Machine1CAmount,
		&raw.Machine2CCount,
		&raw.Machine2CAmount,
		&raw.Machine3CCount,
		&raw.Machine3CAmount,
		&raw.Machine4CCount,
		&raw.Machine4CAmount,
		&raw.Machine5CCount,
		&raw.Machine5CAmount,
		&raw.Machine1CSettleCount,
		&raw.Machine1CSettleAmount,
		&raw.Machine2CSettleCount,
		&raw.Machine2CSettleAmount,
		&raw.Machine3CSettleCount,
		&raw.Machine3CSettleAmount,
		&raw.Machine4CSettleCount,
		&raw.Machine4CSettleAmount,
		&raw.Machine5CSettleCount,
		&raw.Machine5CSettleAmount,
		&raw.AdultTicketCount,
		&raw.AdultSetTicketCount,
		&raw.ChildTicketCount,
		&raw.TicketCount,
		&raw.InvitationTicketCount,
		&raw.YuutaikenTicketCount,
		&raw.InfantTicketCount,
		&raw.PointCardAdultCount,
		&raw.PointCardChildCount,
		&raw.TicketSalesCount,
		&raw.OldTicketCount,
		&raw.SalesNoStart,
		&raw.SalesNoEnd,
		&raw.Change,
		&raw.PhoneFee,
		&raw.Teuri,
		&raw.YuutaikenSalesCount,
		&raw.YuutaikenSalesAmount,
		&raw.HonjitsuMitounyuuAmountUncertain,
		&raw.HonjitsuMitounyuuAmountCertain,
		&raw.Deficiency,
		&raw.ZenjitsuMitounyuuAmount,
		&raw.ReceiptTotalAmount,
		&raw.Remarks,
		&raw.SCutMale,
		&raw.SCutFemale,
		&raw.SCutChild,
		&raw.SixTicketCount,
		&raw.SixTicketSalesCount,
		&raw.SixSalesNoStart,
		&raw.SixSalesNoEnd,
		&raw.MaleTicketCount,
		&raw.FemaleTicketCount,
		&raw.SixMaleTicketCount,
		&raw.SixFemaleTicketCount,
		&raw.CouponCount,
		&raw.RearegiAmount,
		&raw.RearegiTicketAmount,
		&raw.RearegiRelaxAmount,
		&raw.ReportSpace,
	)

	dateString = string(raw.Date)
	raw.DateString = dateString

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("No user found for the date %s\n", formattedDate)
		return &DailyReportRaw{}, false, nil
	case err != nil:
		log.Fatalf("An SQL error occurred: %v\n", err)

		return nil, false, err
	default:
		fmt.Printf("User found: ID: %d, Staff Code: %s, Date: %s\n", raw.ID, raw.StaffCode, raw.Date)

		return &raw, true, nil

	}

	// if err != nil {
	// 	log.Fatal(err)
	// 	return nil, false, nil
	// }
	// return &raw, true, nil
}
