package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

/*
データベースから取得する元データを格納する構造体
*/
type DailyReportRaw struct {
	ID                               int     `db:"ID"`
	Date                             []uint8 `db:"日付" json:"-"`
	DateString                       string
	JapaneseWeekday                  string
	WeatherCode                      int    `db:"天気ｺｰﾄﾞ"`
	WeatherCondition                 string `db:"天気状況"`
	WeatherMark                      string `db:"天気ﾏｰｸ"`
	StaffCode                        int    `db:"担当ｺｰﾄﾞ"`
	StaffName                        string `db:"担当"`
	Machine1CashCountNum             int    `db:"1号機現金枚数" json:"-"`
	Machine1CashAmountNum            int    `db:"1号機現金金額" json:"-"`
	Machine1SettleCountNum           int    `db:"1号機精算枚数" json:"-"`
	Machine1SettleAmountNum          int    `db:"1号機精算金額" json:"-"`
	Machine2CashCountNum             int    `db:"2号機現金枚数" json:"-"`
	Machine2CashAmountNum            int    `db:"2号機現金金額" json:"-"`
	Machine2SettleCountNum           int    `db:"2号機精算枚数" json:"-"`
	Machine2SettleAmountNum          int    `db:"2号機精算金額" json:"-"`
	Machine3CashCountNum             int    `db:"3号機現金枚数" json:"-"`
	Machine3CashAmountNum            int    `db:"3号機現金金額" json:"-"`
	Machine3SettleCountNum           int    `db:"3号機精算枚数" json:"-"`
	Machine3SettleAmountNum          int    `db:"3号機精算金額" json:"-"`
	Machine4CashCountNum             int    `db:"4号機現金枚数" json:"-"`
	Machine4CashAmountNum            int    `db:"4号機現金金額" json:"-"`
	Machine4SettleCountNum           int    `db:"4号機精算枚数" json:"-"`
	Machine4SettleAmountNum          int    `db:"4号機精算金額" json:"-"`
	Machine5CashCountNum             int    `db:"5号機現金枚数" json:"-"`
	Machine5CashAmountNum            int    `db:"5号機現金金額" json:"-"`
	Machine5SettleCountNum           int    `db:"5号機精算枚数" json:"-"`
	Machine5SettleAmountNum          int    `db:"5号機精算金額" json:"-"`
	Machine1UnsettledCountNum        int    `db:"1号機未精算枚数" json:"-"`
	Machine1UnsettledAmountNum       int    `db:"1号機未精算金額" json:"-"`
	Machine2UnsettledCountNum        int    `db:"2号機未精算枚数" json:"-"`
	Machine2UnsettledAmountNum       int    `db:"2号機未精算金額" json:"-"`
	Machine3UnsettledCountNum        int    `db:"3号機未精算枚数" json:"-"`
	Machine3UnsettledAmountNum       int    `db:"3号機未精算金額" json:"-"`
	Machine4UnsettledCountNum        int    `db:"4号機未精算枚数" json:"-"`
	Machine4UnsettledAmountNum       int    `db:"4号機未精算金額" json:"-"`
	Machine5UnsettledCountNum        int    `db:"5号機未精算枚数" json:"-"`
	Machine5UnsettledAmountNum       int    `db:"5号機未精算金額" json:"-"`
	Machine1QrCountNum               int    `db:"1号機QR件数" json:"-"`
	Machine1QrAmountNum              int    `db:"1号機QR金額" json:"-"`
	Machine2QrCountNum               int    `db:"2号機QR件数" json:"-"`
	Machine2QrAmountNum              int    `db:"2号機QR金額" json:"-"`
	Machine3QrCountNum               int    `db:"3号機QR件数" json:"-"`
	Machine3QrAmountNum              int    `db:"3号機QR金額" json:"-"`
	Machine4QrCountNum               int    `db:"4号機QR件数" json:"-"`
	Machine4QrAmountNum              int    `db:"4号機QR金額" json:"-"`
	Machine5QrCountNum               int    `db:"5号機QR件数" json:"-"`
	Machine5QrAmountNum              int    `db:"5号機QR金額" json:"-"`
	Machine1QrSettleCountNum         int    `db:"1号機QR精算件数" json:"-"`
	Machine1QrSettleAmountNum        int    `db:"1号機QR精算金額" json:"-"`
	Machine2QrSettleCountNum         int    `db:"2号機QR精算件数" json:"-"`
	Machine2QrSettleAmountNum        int    `db:"2号機QR精算金額" json:"-"`
	Machine3QrSettleCountNum         int    `db:"3号機QR精算件数" json:"-"`
	Machine3QrSettleAmountNum        int    `db:"3号機QR精算金額" json:"-"`
	Machine4QrSettleCountNum         int    `db:"4号機QR精算件数" json:"-"`
	Machine4QrSettleAmountNum        int    `db:"4号機QR精算金額" json:"-"`
	Machine5QrSettleCountNum         int    `db:"5号機QR精算件数" json:"-"`
	Machine5QrSettleAmountNum        int    `db:"5号機QR精算金額" json:"-"`
	Machine1ECountNum                int    `db:"1号機電子マネ件数" json:"-"`
	Machine1EAmountNum               int    `db:"1号機電子マネ金額" json:"-"`
	Machine2ECountNum                int    `db:"2号機電子マネ件数" json:"-"`
	Machine2EAmountNum               int    `db:"2号機電子マネ金額" json:"-"`
	Machine3ECountNum                int    `db:"3号機電子マネ件数" json:"-"`
	Machine3EAmountNum               int    `db:"3号機電子マネ金額" json:"-"`
	Machine4ECountNum                int    `db:"4号機電子マネ件数" json:"-"`
	Machine4EAmountNum               int    `db:"4号機電子マネ金額" json:"-"`
	Machine5ECountNum                int    `db:"5号機電子マネ件数" json:"-"`
	Machine5EAmountNum               int    `db:"5号機電子マネ金額" json:"-"`
	Machine1ESettleCountNum          int    `db:"1号機電子マネ精算件数" json:"-"`
	Machine1ESettleAmountNum         int    `db:"1号機電子マネ精算金額" json:"-"`
	Machine2ESettleCountNum          int    `db:"2号機電子マネ精算件数" json:"-"`
	Machine2ESettleAmountNum         int    `db:"2号機電子マネ精算金額" json:"-"`
	Machine3ESettleCountNum          int    `db:"3号機電子マネ精算件数" json:"-"`
	Machine3ESettleAmountNum         int    `db:"3号機電子マネ精算金額" json:"-"`
	Machine4ESettleCountNum          int    `db:"4号機電子マネ精算件数" json:"-"`
	Machine4ESettleAmountNum         int    `db:"4号機電子マネ精算金額" json:"-"`
	Machine5ESettleCountNum          int    `db:"5号機電子マネ精算件数" json:"-"`
	Machine5ESettleAmountNum         int    `db:"5号機電子マネ精算金額" json:"-"`
	Machine1CCountNum                int    `db:"1号機クレジット件数" json:"-"`
	Machine1CAmountNum               int    `db:"1号機クレジット金額" json:"-"`
	Machine2CCountNum                int    `db:"2号機クレジット件数" json:"-"`
	Machine2CAmountNum               int    `db:"2号機クレジット金額" json:"-"`
	Machine3CCountNum                int    `db:"3号機クレジット件数" json:"-"`
	Machine3CAmountNum               int    `db:"3号機クレジット金額" json:"-"`
	Machine4CCountNum                int    `db:"4号機クレジット件数" json:"-"`
	Machine4CAmountNum               int    `db:"4号機クレジット金額" json:"-"`
	Machine5CCountNum                int    `db:"5号機クレジット件数" json:"-"`
	Machine5CAmountNum               int    `db:"5号機クレジット金額" json:"-"`
	Machine1CSettleCountNum          int    `db:"1号機クレジット精算件数" json:"-"`
	Machine1CSettleAmountNum         int    `db:"1号機クレジット精算金額" json:"-"`
	Machine2CSettleCountNum          int    `db:"2号機クレジット精算件数" json:"-"`
	Machine2CSettleAmountNum         int    `db:"2号機クレジット精算金額" json:"-"`
	Machine3CSettleCountNum          int    `db:"3号機クレジット精算件数" json:"-"`
	Machine3CSettleAmountNum         int    `db:"3号機クレジット精算金額" json:"-"`
	Machine4CSettleCountNum          int    `db:"4号機クレジット精算件数" json:"-"`
	Machine4CSettleAmountNum         int    `db:"4号機クレジット精算金額" json:"-"`
	Machine5CSettleCountNum          int    `db:"5号機クレジット精算件数" json:"-"`
	Machine5CSettleAmountNum         int    `db:"5号機クレジット精算金額" json:"-"`
	Machine1CashCount                string `db:"1号機現金枚数"`
	Machine1CashAmount               string `db:"1号機現金金額"`
	Machine1SettleCount              string `db:"1号機精算枚数"`
	Machine1SettleAmount             string `db:"1号機精算金額"`
	Machine2CashCount                string `db:"2号機現金枚数"`
	Machine2CashAmount               string `db:"2号機現金金額"`
	Machine2SettleCount              string `db:"2号機精算枚数"`
	Machine2SettleAmount             string `db:"2号機精算金額"`
	Machine3CashCount                string `db:"3号機現金枚数"`
	Machine3CashAmount               string `db:"3号機現金金額"`
	Machine3SettleCount              string `db:"3号機精算枚数"`
	Machine3SettleAmount             string `db:"3号機精算金額"`
	Machine4CashCount                string `db:"4号機現金枚数"`
	Machine4CashAmount               string `db:"4号機現金金額"`
	Machine4SettleCount              string `db:"4号機精算枚数"`
	Machine4SettleAmount             string `db:"4号機精算金額"`
	Machine5CashCount                string `db:"5号機現金枚数"`
	Machine5CashAmount               string `db:"5号機現金金額"`
	Machine5SettleCount              string `db:"5号機精算枚数"`
	Machine5SettleAmount             string `db:"5号機精算金額"`
	Machine1UnsettledCount           string `db:"1号機未精算枚数"`
	Machine1UnsettledAmount          string `db:"1号機未精算金額"`
	Machine2UnsettledCount           string `db:"2号機未精算枚数"`
	Machine2UnsettledAmount          string `db:"2号機未精算金額"`
	Machine3UnsettledCount           string `db:"3号機未精算枚数"`
	Machine3UnsettledAmount          string `db:"3号機未精算金額"`
	Machine4UnsettledCount           string `db:"4号機未精算枚数"`
	Machine4UnsettledAmount          string `db:"4号機未精算金額"`
	Machine5UnsettledCount           string `db:"5号機未精算枚数"`
	Machine5UnsettledAmount          string `db:"5号機未精算金額"`
	Machine1QrCount                  string `db:"1号機QR件数"`
	Machine1QrAmount                 string `db:"1号機QR金額"`
	Machine2QrCount                  string `db:"2号機QR件数"`
	Machine2QrAmount                 string `db:"2号機QR金額"`
	Machine3QrCount                  string `db:"3号機QR件数"`
	Machine3QrAmount                 string `db:"3号機QR金額"`
	Machine4QrCount                  string `db:"4号機QR件数"`
	Machine4QrAmount                 string `db:"4号機QR金額"`
	Machine5QrCount                  string `db:"5号機QR件数"`
	Machine5QrAmount                 string `db:"5号機QR金額"`
	Machine1QrSettleCount            string `db:"1号機QR精算件数"`
	Machine1QrSettleAmount           string `db:"1号機QR精算金額"`
	Machine2QrSettleCount            string `db:"2号機QR精算件数"`
	Machine2QrSettleAmount           string `db:"2号機QR精算金額"`
	Machine3QrSettleCount            string `db:"3号機QR精算件数"`
	Machine3QrSettleAmount           string `db:"3号機QR精算金額"`
	Machine4QrSettleCount            string `db:"4号機QR精算件数"`
	Machine4QrSettleAmount           string `db:"4号機QR精算金額"`
	Machine5QrSettleCount            string `db:"5号機QR精算件数"`
	Machine5QrSettleAmount           string `db:"5号機QR精算金額"`
	Machine1ECount                   string `db:"1号機電子マネ件数"`
	Machine1EAmount                  string `db:"1号機電子マネ金額"`
	Machine2ECount                   string `db:"2号機電子マネ件数"`
	Machine2EAmount                  string `db:"2号機電子マネ金額"`
	Machine3ECount                   string `db:"3号機電子マネ件数"`
	Machine3EAmount                  string `db:"3号機電子マネ金額"`
	Machine4ECount                   string `db:"4号機電子マネ件数"`
	Machine4EAmount                  string `db:"4号機電子マネ金額"`
	Machine5ECount                   string `db:"5号機電子マネ件数"`
	Machine5EAmount                  string `db:"5号機電子マネ金額"`
	Machine1ESettleCount             string `db:"1号機電子マネ精算件数"`
	Machine1ESettleAmount            string `db:"1号機電子マネ精算金額"`
	Machine2ESettleCount             string `db:"2号機電子マネ精算件数"`
	Machine2ESettleAmount            string `db:"2号機電子マネ精算金額"`
	Machine3ESettleCount             string `db:"3号機電子マネ精算件数"`
	Machine3ESettleAmount            string `db:"3号機電子マネ精算金額"`
	Machine4ESettleCount             string `db:"4号機電子マネ精算件数"`
	Machine4ESettleAmount            string `db:"4号機電子マネ精算金額"`
	Machine5ESettleCount             string `db:"5号機電子マネ精算件数"`
	Machine5ESettleAmount            string `db:"5号機電子マネ精算金額"`
	Machine1CCount                   string `db:"1号機クレジット件数"`
	Machine1CAmount                  string `db:"1号機クレジット金額"`
	Machine2CCount                   string `db:"2号機クレジット件数"`
	Machine2CAmount                  string `db:"2号機クレジット金額"`
	Machine3CCount                   string `db:"3号機クレジット件数"`
	Machine3CAmount                  string `db:"3号機クレジット金額"`
	Machine4CCount                   string `db:"4号機クレジット件数"`
	Machine4CAmount                  string `db:"4号機クレジット金額"`
	Machine5CCount                   string `db:"5号機クレジット件数"`
	Machine5CAmount                  string `db:"5号機クレジット金額"`
	Machine1CSettleCount             string `db:"1号機クレジット精算件数"`
	Machine1CSettleAmount            string `db:"1号機クレジット精算金額"`
	Machine2CSettleCount             string `db:"2号機クレジット精算件数"`
	Machine2CSettleAmount            string `db:"2号機クレジット精算金額"`
	Machine3CSettleCount             string `db:"3号機クレジット精算件数"`
	Machine3CSettleAmount            string `db:"3号機クレジット精算金額"`
	Machine4CSettleCount             string `db:"4号機クレジット精算件数"`
	Machine4CSettleAmount            string `db:"4号機クレジット精算金額"`
	Machine5CSettleCount             string `db:"5号機クレジット精算件数"`
	Machine5CSettleAmount            string `db:"5号機クレジット精算金額"`
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
		T1.ID, T1.日付, T1.天気コード, tm.天気状況, T1.担当コード,
		T1.1号機現金枚数, T1.1号機現金金額, T1.1号機精算枚数, T1.1号機精算金額,
		T1.2号機現金枚数, T1.2号機現金金額, T1.2号機精算枚数, T1.2号機精算金額,
		T1.3号機現金枚数, T1.3号機現金金額, T1.3号機精算枚数, T1.3号機精算金額,
		T1.4号機現金枚数, T1.4号機現金金額, T1.4号機精算枚数, T1.4号機精算金額,
		T1.5号機現金枚数, T1.5号機現金金額, T1.5号機精算枚数, T1.5号機精算金額,
		T1.1号機未精算枚数, T1.1号機未精算金額, T1.2号機未精算枚数, T1.2号機未精算金額,
		T1.3号機未精算枚数, T1.3号機未精算金額, T1.4号機未精算枚数, T1.4号機未精算金額,
		T1.5号機未精算枚数, T1.5号機未精算金額,
		T1.1号機QR件数, T1.1号機QR金額, T1.2号機QR件数, T1.2号機QR金額,
		T1.3号機QR件数, T1.3号機QR金額, T1.4号機QR件数, T1.4号機QR金額,
		T1.5号機QR件数, T1.5号機QR金額, T1.1号機QR精算件数, T1.1号機QR精算金額,
		T1.2号機QR精算件数, T1.2号機QR精算金額, T1.3号機QR精算件数, T1.3号機QR精算金額,
		T1.4号機QR精算件数, T1.4号機QR精算金額, T1.5号機QR精算件数, T1.5号機QR精算金額,
		T1.1号機電子マネ件数, T1.1号機電子マネ金額, T1.2号機電子マネ件数, T1.2号機電子マネ金額,
		T1.3号機電子マネ件数, T1.3号機電子マネ金額, T1.4号機電子マネ件数, T1.4号機電子マネ金額,
		T1.5号機電子マネ件数, T1.5号機電子マネ金額, T1.1号機電子マネ精算件数, T1.1号機電子マネ精算金額,
		T1.2号機電子マネ精算件数, T1.2号機電子マネ精算金額, T1.3号機電子マネ精算件数, T1.3号機電子マネ精算金額,
		T1.4号機電子マネ精算件数, T1.4号機電子マネ精算金額, T1.5号機電子マネ精算件数, T1.5号機電子マネ精算金額,
		T1.1号機クレジット件数, T1.1号機クレジット金額, T1.2号機クレジット件数, T1.2号機クレジット金額,
		T1.3号機クレジット件数, T1.3号機クレジット金額, T1.4号機クレジット件数, T1.4号機クレジット金額,
		T1.5号機クレジット件数, T1.5号機クレジット金額, T1.1号機クレジット精算件数, T1.1号機クレジット精算金額,
		T1.2号機クレジット精算件数, T1.2号機クレジット精算金額, T1.3号機クレジット精算件数, T1.3号機クレジット精算金額,
		T1.4号機クレジット精算件数, T1.4号機クレジット精算金額, T1.5号機クレジット精算件数, T1.5号機クレジット精算金額,
		T1.大人入浴券枚数, T1.大人入浴セット券枚数, T1.小人入浴券枚数, T1.回数券回収,
		T1.招待券回収, T1.優待券回収, T1.感謝祭招待券回収, T1.ﾎﾟｲﾝﾄｶｰﾄﾞ大人回収,
		T1.ﾎﾟｲﾝﾄｶｰﾄﾞﾞ小人回収, T1.回数券売数, T1.過去回数券回収, T1.販売NO（始まり）,
		T1.販売NO（終わり）, T1.釣銭（ﾌﾛﾝﾄ含む）, T1.電話料投入, T1.優待券販売（枚数）,
		T1.優待券販売（金額）, T1.本日未投入金額（未確定）, T1.本日未投入金額（確定）,
		T1.過不足, T1.前日未投入金額, T1.投入ﾚｼｰﾄ合計金額, T1.備考, T1.ｴｽｶｯﾄ男性,
		T1.ｴｽｶｯﾄ女性, T1.ｴｽｶｯﾄ子供, T1.6回数券回収, T1.6回数券売数, T1.6販売NO（始まり）,
		T1.6販売NO（終わり）, T1.男回数券枚数, T1.女回数券枚数, T1.6男回数券枚数, T1.6女回数券枚数, T1.ｸｰﾎﾟﾝ回収,
		T1.リアレジ金額, T1.リアレジ金額回数券系, T1.リアレジ金額リラク系, T1.報告ｽﾍﾟｰｽ
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
		&raw.WeatherCondition,
		&raw.StaffCode,
		&raw.Machine1CashCountNum,
		&raw.Machine1CashAmountNum,
		&raw.Machine1SettleCountNum,
		&raw.Machine1SettleAmountNum,
		&raw.Machine2CashCountNum,
		&raw.Machine2CashAmountNum,
		&raw.Machine2SettleCountNum,
		&raw.Machine2SettleAmountNum,
		&raw.Machine3CashCountNum,
		&raw.Machine3CashAmountNum,
		&raw.Machine3SettleCountNum,
		&raw.Machine3SettleAmountNum,
		&raw.Machine4CashCountNum,
		&raw.Machine4CashAmountNum,
		&raw.Machine4SettleCountNum,
		&raw.Machine4SettleAmountNum,
		&raw.Machine5CashCountNum,
		&raw.Machine5CashAmountNum,
		&raw.Machine5SettleCountNum,
		&raw.Machine5SettleAmountNum,
		&raw.Machine1UnsettledCountNum,
		&raw.Machine1UnsettledAmountNum,
		&raw.Machine2UnsettledCountNum,
		&raw.Machine2UnsettledAmountNum,
		&raw.Machine3UnsettledCountNum,
		&raw.Machine3UnsettledAmountNum,
		&raw.Machine4UnsettledCountNum,
		&raw.Machine4UnsettledAmountNum,
		&raw.Machine5UnsettledCountNum,
		&raw.Machine5UnsettledAmountNum,
		&raw.Machine1QrCountNum,
		&raw.Machine1QrAmountNum,
		&raw.Machine2QrCountNum,
		&raw.Machine2QrAmountNum,
		&raw.Machine3QrCountNum,
		&raw.Machine3QrAmountNum,
		&raw.Machine4QrCountNum,
		&raw.Machine4QrAmountNum,
		&raw.Machine5QrCountNum,
		&raw.Machine5QrAmountNum,
		&raw.Machine1QrSettleCountNum,
		&raw.Machine1QrSettleAmountNum,
		&raw.Machine2QrSettleCountNum,
		&raw.Machine2QrSettleAmountNum,
		&raw.Machine3QrSettleCountNum,
		&raw.Machine3QrSettleAmountNum,
		&raw.Machine4QrSettleCountNum,
		&raw.Machine4QrSettleAmountNum,
		&raw.Machine5QrSettleCountNum,
		&raw.Machine5QrSettleAmountNum,
		&raw.Machine1ECountNum,
		&raw.Machine1EAmountNum,
		&raw.Machine2ECountNum,
		&raw.Machine2EAmountNum,
		&raw.Machine3ECountNum,
		&raw.Machine3EAmountNum,
		&raw.Machine4ECountNum,
		&raw.Machine4EAmountNum,
		&raw.Machine5ECountNum,
		&raw.Machine5EAmountNum,
		&raw.Machine1ESettleCountNum,
		&raw.Machine1ESettleAmountNum,
		&raw.Machine2ESettleCountNum,
		&raw.Machine2ESettleAmountNum,
		&raw.Machine3ESettleCountNum,
		&raw.Machine3ESettleAmountNum,
		&raw.Machine4ESettleCountNum,
		&raw.Machine4ESettleAmountNum,
		&raw.Machine5ESettleCountNum,
		&raw.Machine5ESettleAmountNum,
		&raw.Machine1CCountNum,
		&raw.Machine1CAmountNum,
		&raw.Machine2CCountNum,
		&raw.Machine2CAmountNum,
		&raw.Machine3CCountNum,
		&raw.Machine3CAmountNum,
		&raw.Machine4CCountNum,
		&raw.Machine4CAmountNum,
		&raw.Machine5CCountNum,
		&raw.Machine5CAmountNum,
		&raw.Machine1CSettleCountNum,
		&raw.Machine1CSettleAmountNum,
		&raw.Machine2CSettleCountNum,
		&raw.Machine2CSettleAmountNum,
		&raw.Machine3CSettleCountNum,
		&raw.Machine3CSettleAmountNum,
		&raw.Machine4CSettleCountNum,
		&raw.Machine4CSettleAmountNum,
		&raw.Machine5CSettleCountNum,
		&raw.Machine5CSettleAmountNum,
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

	//数値を書式付き文字列に変換
	pJapanese := message.NewPrinter(language.Japanese)

	raw.Machine1CashAmount = pJapanese.Sprintf("%d", raw.Machine1CashAmountNum)
	raw.Machine1CashCount = pJapanese.Sprintf("%d", raw.Machine1CashCountNum)
	raw.Machine1SettleCount = pJapanese.Sprintf("%d", raw.Machine1SettleCountNum)
	raw.Machine1SettleAmount = pJapanese.Sprintf("%d", raw.Machine1SettleAmountNum)
	raw.Machine2CashCount = pJapanese.Sprintf("%d", raw.Machine2CashCountNum)
	raw.Machine2CashAmount = pJapanese.Sprintf("%d", raw.Machine2CashAmountNum)
	raw.Machine2SettleCount = pJapanese.Sprintf("%d", raw.Machine2SettleCountNum)
	raw.Machine2SettleAmount = pJapanese.Sprintf("%d", raw.Machine2SettleAmountNum)
	raw.Machine3CashCount = pJapanese.Sprintf("%d", raw.Machine3CashCountNum)
	raw.Machine3CashAmount = pJapanese.Sprintf("%d", raw.Machine3CashAmountNum)
	raw.Machine3SettleCount = pJapanese.Sprintf("%d", raw.Machine3SettleCountNum)
	raw.Machine3SettleAmount = pJapanese.Sprintf("%d", raw.Machine3SettleAmountNum)
	raw.Machine4CashCount = pJapanese.Sprintf("%d", raw.Machine4CashCountNum)
	raw.Machine4CashAmount = pJapanese.Sprintf("%d", raw.Machine4CashAmountNum)
	raw.Machine4SettleCount = pJapanese.Sprintf("%d", raw.Machine4SettleCountNum)
	raw.Machine4SettleAmount = pJapanese.Sprintf("%d", raw.Machine4SettleAmountNum)
	raw.Machine5CashCount = pJapanese.Sprintf("%d", raw.Machine5CashCountNum)
	raw.Machine5CashAmount = pJapanese.Sprintf("%d", raw.Machine5CashAmountNum)
	raw.Machine5SettleCount = pJapanese.Sprintf("%d", raw.Machine5SettleCountNum)
	raw.Machine5SettleAmount = pJapanese.Sprintf("%d", raw.Machine5SettleAmountNum)
	raw.Machine1UnsettledCount = pJapanese.Sprintf("%d", raw.Machine1UnsettledCountNum)
	raw.Machine1UnsettledAmount = pJapanese.Sprintf("%d", raw.Machine1UnsettledAmountNum)
	raw.Machine2UnsettledCount = pJapanese.Sprintf("%d", raw.Machine2UnsettledCountNum)
	raw.Machine2UnsettledAmount = pJapanese.Sprintf("%d", raw.Machine2UnsettledAmountNum)
	raw.Machine3UnsettledCount = pJapanese.Sprintf("%d", raw.Machine3UnsettledCountNum)
	raw.Machine3UnsettledAmount = pJapanese.Sprintf("%d", raw.Machine3UnsettledAmountNum)
	raw.Machine4UnsettledCount = pJapanese.Sprintf("%d", raw.Machine4UnsettledCountNum)
	raw.Machine4UnsettledAmount = pJapanese.Sprintf("%d", raw.Machine4UnsettledAmountNum)
	raw.Machine5UnsettledCount = pJapanese.Sprintf("%d", raw.Machine5UnsettledCountNum)
	raw.Machine5UnsettledAmount = pJapanese.Sprintf("%d", raw.Machine5UnsettledAmountNum)
	raw.Machine1QrCount = pJapanese.Sprintf("%d", raw.Machine1QrCountNum)
	raw.Machine1QrAmount = pJapanese.Sprintf("%d", raw.Machine1QrAmountNum)
	raw.Machine2QrCount = pJapanese.Sprintf("%d", raw.Machine2QrCountNum)
	raw.Machine2QrAmount = pJapanese.Sprintf("%d", raw.Machine2QrAmountNum)
	raw.Machine3QrCount = pJapanese.Sprintf("%d", raw.Machine3QrCountNum)
	raw.Machine3QrAmount = pJapanese.Sprintf("%d", raw.Machine3QrAmountNum)
	raw.Machine4QrCount = pJapanese.Sprintf("%d", raw.Machine4QrCountNum)
	raw.Machine4QrAmount = pJapanese.Sprintf("%d", raw.Machine4QrAmountNum)
	raw.Machine5QrCount = pJapanese.Sprintf("%d", raw.Machine5QrCountNum)
	raw.Machine5QrAmount = pJapanese.Sprintf("%d", raw.Machine5QrAmountNum)
	raw.Machine1QrSettleCount = pJapanese.Sprintf("%d", raw.Machine1QrSettleCountNum)
	raw.Machine1QrSettleAmount = pJapanese.Sprintf("%d", raw.Machine1QrSettleAmountNum)
	raw.Machine2QrSettleCount = pJapanese.Sprintf("%d", raw.Machine2QrSettleCountNum)
	raw.Machine2QrSettleAmount = pJapanese.Sprintf("%d", raw.Machine2QrSettleAmountNum)
	raw.Machine3QrSettleCount = pJapanese.Sprintf("%d", raw.Machine3QrSettleCountNum)
	raw.Machine3QrSettleAmount = pJapanese.Sprintf("%d", raw.Machine3QrSettleAmountNum)
	raw.Machine4QrSettleCount = pJapanese.Sprintf("%d", raw.Machine4QrSettleCountNum)
	raw.Machine4QrSettleAmount = pJapanese.Sprintf("%d", raw.Machine4QrSettleAmountNum)
	raw.Machine5QrSettleCount = pJapanese.Sprintf("%d", raw.Machine5QrSettleCountNum)
	raw.Machine5QrSettleAmount = pJapanese.Sprintf("%d", raw.Machine5QrSettleAmountNum)
	raw.Machine1ECount = pJapanese.Sprintf("%d", raw.Machine1ECountNum)
	raw.Machine1EAmount = pJapanese.Sprintf("%d", raw.Machine1EAmountNum)
	raw.Machine2ECount = pJapanese.Sprintf("%d", raw.Machine2ECountNum)
	raw.Machine2EAmount = pJapanese.Sprintf("%d", raw.Machine2EAmountNum)
	raw.Machine3ECount = pJapanese.Sprintf("%d", raw.Machine3ECountNum)
	raw.Machine3EAmount = pJapanese.Sprintf("%d", raw.Machine3EAmountNum)
	raw.Machine4ECount = pJapanese.Sprintf("%d", raw.Machine4ECountNum)
	raw.Machine4EAmount = pJapanese.Sprintf("%d", raw.Machine4EAmountNum)
	raw.Machine5ECount = pJapanese.Sprintf("%d", raw.Machine5ECountNum)
	raw.Machine5EAmount = pJapanese.Sprintf("%d", raw.Machine5EAmountNum)
	raw.Machine1ESettleCount = pJapanese.Sprintf("%d", raw.Machine1ESettleCountNum)
	raw.Machine1ESettleAmount = pJapanese.Sprintf("%d", raw.Machine1ESettleAmountNum)
	raw.Machine2ESettleCount = pJapanese.Sprintf("%d", raw.Machine2ESettleCountNum)
	raw.Machine2ESettleAmount = pJapanese.Sprintf("%d", raw.Machine2ESettleAmountNum)
	raw.Machine3ESettleCount = pJapanese.Sprintf("%d", raw.Machine3ESettleCountNum)
	raw.Machine3ESettleAmount = pJapanese.Sprintf("%d", raw.Machine3ESettleAmountNum)
	raw.Machine4ESettleCount = pJapanese.Sprintf("%d", raw.Machine4ESettleCountNum)
	raw.Machine4ESettleAmount = pJapanese.Sprintf("%d", raw.Machine4ESettleAmountNum)
	raw.Machine5ESettleCount = pJapanese.Sprintf("%d", raw.Machine5ESettleCountNum)
	raw.Machine5ESettleAmount = pJapanese.Sprintf("%d", raw.Machine5ESettleAmountNum)
	raw.Machine1CCount = pJapanese.Sprintf("%d", raw.Machine1CCountNum)
	raw.Machine1CAmount = pJapanese.Sprintf("%d", raw.Machine1CAmountNum)
	raw.Machine2CCount = pJapanese.Sprintf("%d", raw.Machine2CCountNum)
	raw.Machine2CAmount = pJapanese.Sprintf("%d", raw.Machine2CAmountNum)
	raw.Machine3CCount = pJapanese.Sprintf("%d", raw.Machine3CCountNum)
	raw.Machine3CAmount = pJapanese.Sprintf("%d", raw.Machine3CAmountNum)
	raw.Machine4CCount = pJapanese.Sprintf("%d", raw.Machine4CCountNum)
	raw.Machine4CAmount = pJapanese.Sprintf("%d", raw.Machine4CAmountNum)
	raw.Machine5CCount = pJapanese.Sprintf("%d", raw.Machine5CCountNum)
	raw.Machine5CAmount = pJapanese.Sprintf("%d", raw.Machine5CAmountNum)
	raw.Machine1CSettleCount = pJapanese.Sprintf("%d", raw.Machine1CSettleCountNum)
	raw.Machine1CSettleAmount = pJapanese.Sprintf("%d", raw.Machine1CSettleAmountNum)
	raw.Machine2CSettleCount = pJapanese.Sprintf("%d", raw.Machine2CSettleCountNum)
	raw.Machine2CSettleAmount = pJapanese.Sprintf("%d", raw.Machine2CSettleAmountNum)
	raw.Machine3CSettleCount = pJapanese.Sprintf("%d", raw.Machine3CSettleCountNum)
	raw.Machine3CSettleAmount = pJapanese.Sprintf("%d", raw.Machine3CSettleAmountNum)
	raw.Machine4CSettleCount = pJapanese.Sprintf("%d", raw.Machine4CSettleCountNum)
	raw.Machine4CSettleAmount = pJapanese.Sprintf("%d", raw.Machine4CSettleAmountNum)
	raw.Machine5CSettleCount = pJapanese.Sprintf("%d", raw.Machine5CSettleCountNum)
	raw.Machine5CSettleAmount = pJapanese.Sprintf("%d", raw.Machine5CSettleAmountNum)

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
