package models

import (
	"fmt"
	"os"

	"tronroll21-dev/yudoksystem/models/helpers"
)

type NyuuyokushasuuDailyData struct {
	Hiduke                []uint8 `db:"日付"`
	Hiduke_str            string
	JapaneseWeekday       string
	AdultTicketCount      int    `db:"大人入浴券枚数"`
	AdultSetTicketCount   int    `db:"大人入浴セット券枚数"`
	ChildTicketCount      int    `db:"小人入浴券枚数"`
	InfantTicketCount     int    `db:"感謝祭招待券回収"`
	MaleTicketCount       int    `db:"男回数券枚数"`
	FemaleTicketCount     int    `db:"女回数券枚数"`
	SixMaleTicketCount    int    `db:"6男回数券枚数"`
	SixFemaleTicketCount  int    `db:"6女回数券枚数"`
	ShoutaikenTicketCount int    `db:"招待券回収"`
	YuutaikenTicketCount  int    `db:"優待券回収"`
	WeatherCondition      string `db:"天気状況"`
	PointCardAdultCount   int    `db:"ﾎﾟｲﾝﾄｶｰﾄﾞ大人回収"`
	PointCardChildCount   int    `db:"ﾎﾟｲﾝﾄｶｰﾄﾞﾞ小人回収"`
	TicketCount           int
	SixTicketCount        int
	AllTicketCount        int
	TicketSalesCount      int `db:"回数券回収"`
	SixTicketSalesCount   int `db:"6回数券回収"`
	AllTicketSalesCount   int
	SCutMale              int `db:"ｴｽｶｯﾄ男性"`
	SCutFemale            int `db:"ｴｽｶｯﾄ女性"`
	SCutChild             int `db:"ｴｽｶｯﾄ子供"`
	CouponCount           int `db:"ｸｰﾎﾟﾝ回収"`
	OldTicketCount        int `db:"過去回数券回収"`
	NyuuyokushasuuSoukei  int
}

type NyuuyokushasuuShuukeiData struct {
	NyuuyokushasuuDailyData []NyuuyokushasuuDailyData
	AdultTicketGoukei       int
	AdultSetTicketGoukei    int
	ChildTicketGoukei       int
	InfantTicketGoukei      int
	MaleTicketGoukei        int
	FemaleTicketGoukei      int
	SixMaleTicketGoukei     int
	SixFemaleTicketGoukei   int
	OtokoTicketCountRitsu   int
	OnnaTicketCountRitsu    int
	TicketCountGoukei       int
	SixTicketCountGoukei    int
	AllTicketCountGoukei    int

	ShoutaikenTicketGoukei    int
	YuutaikenTicketGoukei     int
	PointCardAdultGoukei      int
	PointCardChildGoukei      int
	NyuuyokushasuuGoukei      int
	TicketSalesGoukei         int
	SixTicketSalesGoukei      int
	AllTicketSalesGoukei      int
	SCutMaleGoukei            int
	SCutFemaleGoukei          int
	SCutChildGoukei           int
	CouponGoukei              int
	OldTicketGoukei           int
	ShoutaiYuutaiKaishuuCount int
}

func GetNyuuyokushasuuShuukeiDataByDate(startDate, endDate string) (NyuuyokushasuuShuukeiData, error) {
	queryBytes, err := os.ReadFile("sql/nyuuyokushasuu.sql")
	if err != nil {
		return NyuuyokushasuuShuukeiData{}, fmt.Errorf("failed to read SQL file: %w", err)
	}
	query := string(queryBytes)

	rows, err := db.Query(query) //, startDate, endDate)
	if err != nil {
		return NyuuyokushasuuShuukeiData{}, err
	}
	defer rows.Close()

	var results []NyuuyokushasuuDailyData
	var AdultTicketGoukei, AdultSetTicketGoukei, ChildTicketGoukei, InfantTicketGoukei,
		MaleTicketGoukei, FemaleTicketGoukei, SixMaleTicketGoukei, SixFemaleTicketGoukei,
		ShoutaikenTicketGoukei, YuutaikenTicketGoukei, PointCardAdultGoukei, PointCardChildGoukei,
		NyuuyokushasuuSougoukei, TicketSalesGoukei, SixTicketSalesGoukei, AllTicketSalesGoukei,
		SCutMaleGoukei, SCutFemaleGoukei, SCutChildGoukei,
		CouponGoukei, OldTicketGoukei, TicketCountGoukei, SixTicketCountGoukei,
		AllTicketCountGoukei, ShoutaiYuutaiKaishuuCount int

	for rows.Next() {
		var rec NyuuyokushasuuDailyData
		err := rows.Scan(
			&rec.Hiduke,
			&rec.AdultTicketCount,
			&rec.AdultSetTicketCount,
			&rec.ChildTicketCount,
			&rec.InfantTicketCount,
			&rec.MaleTicketCount,
			&rec.FemaleTicketCount,
			&rec.SixMaleTicketCount,
			&rec.SixFemaleTicketCount,
			&rec.ShoutaikenTicketCount,
			&rec.YuutaikenTicketCount,
			&rec.WeatherCondition,
			&rec.PointCardAdultCount,
			&rec.PointCardChildCount,
			&rec.TicketSalesCount,
			&rec.SixTicketSalesCount,
			&rec.SCutMale,
			&rec.SCutFemale,
			&rec.SCutChild,
			&rec.CouponCount,
			&rec.OldTicketCount,
		)
		if err != nil {
			return NyuuyokushasuuShuukeiData{}, err
		}

		// uriagekingakugoukei += rec.Uriagekingaku
		// nyuuyokushagoukeigoukei += rec.Nyuuyokushagoukei

		// if rec.Nyuuyokushagoukei > 0 {
		// 	eigyouNissuu++
		// }

		rec.Hiduke_str = string(rec.Hiduke)
		rec.JapaneseWeekday, err = helpers.GetJapaneseWeekdayKanji(rec.Hiduke)
		if err != nil {
			return NyuuyokushasuuShuukeiData{}, err
		}

		rec.TicketCount = rec.MaleTicketCount + rec.FemaleTicketCount
		rec.SixTicketCount = rec.SixMaleTicketCount + rec.SixFemaleTicketCount
		rec.AllTicketCount = rec.TicketCount + rec.SixTicketCount

		rec.AllTicketSalesCount = rec.TicketSalesCount + rec.SixTicketSalesCount

		rec.NyuuyokushasuuSoukei = rec.AdultTicketCount +
			rec.AdultSetTicketCount +
			rec.ChildTicketCount +
			rec.InfantTicketCount +
			rec.TicketCount +
			rec.SixTicketCount +
			rec.ShoutaikenTicketCount +
			rec.YuutaikenTicketCount +
			rec.PointCardAdultCount +
			rec.PointCardChildCount +
			rec.CouponCount +
			rec.OldTicketCount

		AdultTicketGoukei += rec.AdultTicketCount
		AdultSetTicketGoukei += rec.AdultSetTicketCount
		ChildTicketGoukei += rec.ChildTicketCount
		InfantTicketGoukei += rec.InfantTicketCount
		MaleTicketGoukei += rec.MaleTicketCount
		FemaleTicketGoukei += rec.FemaleTicketCount
		SixMaleTicketGoukei += rec.SixMaleTicketCount
		SixFemaleTicketGoukei += rec.SixFemaleTicketCount
		ShoutaikenTicketGoukei += rec.ShoutaikenTicketCount
		YuutaikenTicketGoukei += rec.YuutaikenTicketCount
		PointCardAdultGoukei += rec.PointCardAdultCount
		PointCardChildGoukei += rec.PointCardChildCount
		NyuuyokushasuuSougoukei += rec.NyuuyokushasuuSoukei
		TicketSalesGoukei += rec.TicketSalesCount
		SixTicketSalesGoukei += rec.SixTicketSalesCount
		AllTicketSalesGoukei += rec.AllTicketSalesCount
		SCutMaleGoukei += rec.SCutMale
		SCutFemaleGoukei += rec.SCutFemale
		SCutChildGoukei += rec.SCutChild
		CouponGoukei += rec.CouponCount
		OldTicketGoukei += rec.OldTicketCount
		ShoutaiYuutaiKaishuuCount += rec.PointCardAdultCount + rec.PointCardChildCount + rec.ShoutaikenTicketCount + rec.YuutaikenTicketCount
		results = append(results, rec)
	}

	// ichinichiatariuriage := int(math.Round(float64(uriagekingakugoukei) / float64(eigyouNissuu)))
	// hitoriatariuriage := int(math.Round(float64(uriagekingakugoukei) / float64(nyuuyokushagoukeigoukei)))

	if err := rows.Err(); err != nil {
		return NyuuyokushasuuShuukeiData{}, err
	}

	TicketCountGoukei = MaleTicketGoukei + FemaleTicketGoukei
	SixTicketCountGoukei = SixMaleTicketGoukei + SixFemaleTicketGoukei
	AllTicketCountGoukei = TicketCountGoukei + SixTicketCountGoukei

	nyuuyokushasuuShuukeiData := NyuuyokushasuuShuukeiData{
		NyuuyokushasuuDailyData:   results,
		AdultTicketGoukei:         AdultTicketGoukei,
		AdultSetTicketGoukei:      AdultSetTicketGoukei,
		ChildTicketGoukei:         ChildTicketGoukei,
		InfantTicketGoukei:        InfantTicketGoukei,
		MaleTicketGoukei:          MaleTicketGoukei,
		FemaleTicketGoukei:        FemaleTicketGoukei,
		SixMaleTicketGoukei:       SixMaleTicketGoukei,
		SixFemaleTicketGoukei:     SixFemaleTicketGoukei,
		OtokoTicketCountRitsu:     helpers.CalculateRitsu(MaleTicketGoukei, TicketCountGoukei),
		OnnaTicketCountRitsu:      helpers.CalculateRitsu(FemaleTicketGoukei, TicketCountGoukei),
		TicketCountGoukei:         TicketCountGoukei,
		SixTicketCountGoukei:      SixTicketCountGoukei,
		AllTicketCountGoukei:      AllTicketCountGoukei,
		ShoutaikenTicketGoukei:    ShoutaikenTicketGoukei,
		YuutaikenTicketGoukei:     YuutaikenTicketGoukei,
		PointCardAdultGoukei:      PointCardAdultGoukei,
		PointCardChildGoukei:      PointCardChildGoukei,
		ShoutaiYuutaiKaishuuCount: ShoutaiYuutaiKaishuuCount,
		NyuuyokushasuuGoukei:      NyuuyokushasuuSougoukei,
		TicketSalesGoukei:         TicketSalesGoukei,
		SixTicketSalesGoukei:      SixTicketSalesGoukei,
		AllTicketSalesGoukei:      AllTicketSalesGoukei,
		SCutMaleGoukei:            SCutMaleGoukei,
		SCutFemaleGoukei:          SCutFemaleGoukei,
		SCutChildGoukei:           SCutChildGoukei,
		CouponGoukei:              CouponGoukei,
		OldTicketGoukei:           OldTicketGoukei,
	}

	return nyuuyokushasuuShuukeiData, nil
}

type NyuuyokushasuuKaisuukenHanbaiNikaishuugoukei struct {
	KaisuukenhannbaiGoukei       int `db:"回数券売数の合計"`
	KaisuukenhannbaiMaiGoukei    int
	KaisuukenkaishuuGoukei       int `db:"回数券回収の合計"`
	KaisuukenmikaishuuGoukei     int
	SixKaisuukenhannbaiGoukei    int `db:"6回数券売数の合計"`
	SixKaisuukenhannbaiMaiGoukei int
	SixKaisuukenkaishuuGoukei    int `db:"6回数券回収の合計"`
	SixKaisuukenmikaishuuGoukei  int
	YuutaikenhanbaiGoukei        int `db:"優待券販売（枚数）の合計"`
	YuutaikenkaishuuGoukei       int `db:"優待券回収の合計"`
	YuutaikenmikaishuuGoukei     int
}

func GetNyuuyokushasuuRuikeiDataByEndDate(endDate string) (NyuuyokushasuuKaisuukenHanbaiNikaishuugoukei, error) {
	queryBytes, err := os.ReadFile("sql/kaisuukenhanbainmikaishuugoukei.sql")
	if err != nil {
		return NyuuyokushasuuKaisuukenHanbaiNikaishuugoukei{}, fmt.Errorf("failed to read SQL file: %w", err)
	}
	query := string(queryBytes)

	row := db.QueryRow(query, endDate)

	var rec NyuuyokushasuuKaisuukenHanbaiNikaishuugoukei
	err = row.Scan(
		&rec.KaisuukenhannbaiGoukei,
		&rec.KaisuukenkaishuuGoukei,
		&rec.SixKaisuukenhannbaiGoukei,
		&rec.SixKaisuukenkaishuuGoukei,
		&rec.YuutaikenhanbaiGoukei,
		&rec.YuutaikenkaishuuGoukei,
	)
	if err != nil {
		return NyuuyokushasuuKaisuukenHanbaiNikaishuugoukei{}, err
	}

	rec.YuutaikenmikaishuuGoukei = rec.YuutaikenhanbaiGoukei - rec.YuutaikenkaishuuGoukei
	rec.KaisuukenhannbaiMaiGoukei = rec.KaisuukenhannbaiGoukei * 11
	rec.KaisuukenmikaishuuGoukei = rec.KaisuukenhannbaiMaiGoukei - rec.KaisuukenkaishuuGoukei
	rec.SixKaisuukenhannbaiMaiGoukei = rec.SixKaisuukenhannbaiGoukei * 6
	rec.SixKaisuukenmikaishuuGoukei = rec.SixKaisuukenhannbaiMaiGoukei - rec.SixKaisuukenkaishuuGoukei

	return rec, nil
}
