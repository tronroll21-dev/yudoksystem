package models

import (
	"fmt"
	"math"
	"os"
)

type UriageNikkeiDailyData struct {
	Hiduke                             []uint8 `db:"日付"`
	Hiduke_str                         string
	Nyuukinyoteibi                     []uint8 `db:"入金予定日"`
	Nyuukinyoteibi_str                 string
	Tounyuugoukei                      int `db:"投入合計"`
	Kinsen_frontfukumu                 int `db:"釣銭（ﾌﾛﾝﾄ含む）"`
	Nyuukinyoteikingaku                int `db:"入金予定金額"`
	Kenbaikikei                        int `db:"券売機計"`
	Yuutaikenhanbai_kingaku            int `db:"優待券販売（金額）"`
	Miseisankingakugoukei              int `db:"未清算金額合計"`
	Kabusoku                           int `db:"過不足"`
	Denwaryoutounyu                    int `db:"電話料投入"`
	Honjitsumitounyuukingaku_mikakutei int `db:"本日未投入金額（未確定）"`
	Honjitsumitounyuukingaku_kakutei   int `db:"本日未投入金額（確定）"`
	Zenjitsumitounyuukingaku           int `db:"前日未投入金額"`
	Uriagekingaku                      int `db:"売上金額"`
	Uriagekingaku2                     int `db:"売上金額2"`
	Uriagerearejigoukei                int `db:"売上リアレジ合計"`
	Teuri_kingaku                      int `db:"手売（金額）"`
	Couponkaishuu                      int `db:"ｸｰﾎﾟﾝ回収"`
	Rearejikingaku                     int `db:"リアレジ金額"`
	QRkingakugoukei                    int `db:"QR金額合計"`
	Denshimoneykingakugoukei           int `db:"電子マネ金額合計"`
	Creditkingakugoukei                int `db:"クレジット金額合計"`
	Genkincashlessgoukei               int `db:"現金+キャッシュレス合計"`
	Genkincashlessgoukei_reareji       int `db:"現金+キャッシュレス合計＋リアレジ"`
	Nyuuyokushagoukei                  int `db:"入浴者合計"`
	Denshimoneyseisankingakugoukei     int `db:"電子マネ精算金額合計"`
}

type UriageNikkeiData struct {
	UriageNikkeiDailyData []UriageNikkeiDailyData
	Uriagekingakugoukei   int
	EigyouNissuu          int
	Ichinichiatariuriage  int
	Hitoriatariuriage     int
}

func GetUriageNikkeiDataByDate(startDate, endDate string) (UriageNikkeiData, error) {
	queryBytes, err := os.ReadFile("sql/uriagenikkei.sql")
	if err != nil {
		return UriageNikkeiData{}, fmt.Errorf("failed to read SQL file: %w", err)
	}
	query := string(queryBytes)

	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		return UriageNikkeiData{}, err
	}
	defer rows.Close()

	var results []UriageNikkeiDailyData
	var eigyouNissuu, uriagekingakugoukei, nyuuyokushagoukeigoukei int
	for rows.Next() {
		var rec UriageNikkeiDailyData
		err := rows.Scan(
			&rec.Hiduke,
			&rec.Kenbaikikei,
			&rec.Yuutaikenhanbai_kingaku,
			&rec.Miseisankingakugoukei,
			&rec.Kabusoku,
			&rec.Denwaryoutounyu,
			&rec.Kinsen_frontfukumu,
			&rec.Honjitsumitounyuukingaku_mikakutei,
			&rec.Honjitsumitounyuukingaku_kakutei,
			&rec.Zenjitsumitounyuukingaku,
			&rec.Nyuukinyoteibi,
			&rec.Teuri_kingaku,
			&rec.Couponkaishuu,
			&rec.Rearejikingaku,
			&rec.QRkingakugoukei,
			&rec.Denshimoneykingakugoukei,
			&rec.Creditkingakugoukei,
			&rec.Denshimoneyseisankingakugoukei,
			&rec.Nyuuyokushagoukei,
			&rec.Tounyuugoukei,
			&rec.Uriagekingaku,
			&rec.Uriagekingaku2,
			&rec.Genkincashlessgoukei,
			&rec.Nyuukinyoteikingaku,
			&rec.Uriagerearejigoukei,
			&rec.Genkincashlessgoukei_reareji,
		)
		if err != nil {
			return UriageNikkeiData{}, err
		}

		uriagekingakugoukei += rec.Uriagekingaku
		nyuuyokushagoukeigoukei += rec.Nyuuyokushagoukei

		if rec.Nyuuyokushagoukei > 0 {
			eigyouNissuu++
		}

		rec.Hiduke_str = string(rec.Hiduke)
		rec.Nyuukinyoteibi_str = string(rec.Nyuukinyoteibi)

		results = append(results, rec)
	}

	ichinichiatariuriage := int(math.Round(float64(uriagekingakugoukei) / float64(eigyouNissuu)))
	hitoriatariuriage := int(math.Round(float64(uriagekingakugoukei) / float64(nyuuyokushagoukeigoukei)))

	if err := rows.Err(); err != nil {
		return UriageNikkeiData{}, err
	}

	uriageNikkeiData := UriageNikkeiData{
		UriageNikkeiDailyData: results,
		Uriagekingakugoukei:   uriagekingakugoukei,
		Ichinichiatariuriage:  ichinichiatariuriage,
		EigyouNissuu:          eigyouNissuu,
		Hitoriatariuriage:     hitoriatariuriage,
	}

	return uriageNikkeiData, nil
}

type UriageRuikeiData struct {
	Uriagekingakuruikei        int `db:"売上金額累計"`
	Nyuujousharuikei           int `db:"入場者累計"`
	EigyouNissuuruikei         int `db:"営業日数累計"`
	Ichinichiatariuriageruikei int
	Hitoriatariuriageruikei    int
}

func GetUriageRuikeiDataByDate(startDate, endDate string) (UriageRuikeiData, error) {
	queryBytes, err := os.ReadFile("sql/uriageruikei.sql")
	if err != nil {
		return UriageRuikeiData{}, fmt.Errorf("failed to read SQL file: %w", err)
	}
	query := string(queryBytes)

	row := db.QueryRow(query)

	var rec UriageRuikeiData
	err = row.Scan(
		&rec.Uriagekingakuruikei,
		&rec.Nyuujousharuikei,
		&rec.EigyouNissuuruikei,
	)
	if err != nil {
		return UriageRuikeiData{}, err
	}

	rec.Ichinichiatariuriageruikei = int(math.Round(float64(rec.Uriagekingakuruikei) / float64(rec.EigyouNissuuruikei)))
	rec.Hitoriatariuriageruikei = int(math.Round(float64(rec.Uriagekingakuruikei) / float64(rec.Nyuujousharuikei)))

	return rec, nil
}
