package models

// PriceRange はデータベースのt_メニューカテゴリー範囲テーブルの構造を表します。
// JSONタグは、Ginがリクエスト/レスポンスでJSONを扱うために使用されます。
type RearegiDetail struct {
	Hiduke      string `db:"日付" json:"hiduke"`
	Rank        int    `db:"ランク" json:"rank"`
	ShohinCode  int    `db:"商品区分コード" json:"shohinCode"`
	ShohinName  string `db:"商品区分名" json:"shohinName"`
	Kingaku     int    `db:"金額" json:"kingaku"`
	Suryo       int    `db:"数量" json:"suryo"`
	HeikinTanka int    `db:"平均単価" json:"heikinTanka"`
}

// SaveRearegiDetail performs an upsert on the rearegi_detail table based on IDs.
func SaveRearegiDetail(rd []RearegiDetail, hiduke string) error {

	queryDelete := `DELETE FROM T_リアレジ明細 WHERE 日付 = ?`

	_, err := db.Exec(queryDelete, hiduke)

	if err != nil {
		return err
	}

	queryInsert := `INSERT INTO T_リアレジ明細 (日付, ランク, 商品区分コード, 商品区分名, 金額, 数量, 平均単価) VALUES (?, ?, ?, ?, ?, ?, ?)`
	for _, r := range rd {
		_, err := db.Exec(queryInsert, hiduke, r.Rank, r.ShohinCode, r.ShohinName, r.Kingaku, r.Suryo, r.HeikinTanka)
		if err != nil {
			return err
		}
	}
	// }

	return nil
}

func FindRearegiDetails(start_date string, end_date string) ([]RearegiDetail, error) {
	query := `SELECT 日付, ランク, 商品区分コード, 商品区分名, 金額, 数量, 平均単価 FROM T_リアレジ明細 WHERE 日付 BETWEEN ? AND ? ORDER BY 商品区分コード`
	rows, err := db.Query(query, start_date, end_date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []RearegiDetail
	for rows.Next() {
		var d RearegiDetail
		if err := rows.Scan(&d.Hiduke, &d.Rank, &d.ShohinCode, &d.ShohinName, &d.Kingaku, &d.Suryo, &d.HeikinTanka); err != nil {
			return nil, err
		}
		details = append(details, d)
	}

	return details, nil
}
