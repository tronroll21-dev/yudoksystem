package models

// "database/sql"
// "log"

// PriceRange はデータベースのt_メニューカテゴリー範囲テーブルの構造を表します。
// JSONタグは、Ginがリクエスト/レスポンスでJSONを扱うために使用されます。
type Shiiremeisai struct {
	Nengetsu        int    `db:"nengetsu" json:"year_month"`
	Torihikisakimei string `db:"torihikisakimei" json:"contractor_id"`
	Himokumei       string `db:"himokumei" json:"expense_category"`
	Tekiyou         string `db:"tekiyou" json:"tekiyou"`
	Amount          int    `db:"amount" json:"amount"`
}

// InsertRange は新しいt_メニューカテゴリー範囲をデータベースに挿入します。
// func InsertRange(r PriceRange) (PriceRange, error) {
// 	// SQL INSERT文を実行
// 	query := `INSERT INTO t_メニューカテゴリー範囲 (部門ID, 範囲開始, 範囲終了, カテゴリー名) VALUES (?, ?, ?, ?)`
// 	result, err := db.Exec(query, r.Bumon_ID, r.Han_i_kaishi, r.Han_i_shuryo, r.Category_name)
// 	if err != nil {
// 		return PriceRange{}, err
// 	}

// 	// 挿入された行のIDを取得
// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		log.Println("挿入されたIDの取得に失敗しました: ", err)
// 	}
// 	r.ID = int(id)
// 	return r, nil
// }

// FindAllShiiremeisai は全てのt_メニューカテゴリー範囲をデータベースから取得します。
func FindAllShiiremeisai(year_month string) ([]Shiiremeisai, error) {
	shiiremeisaiData := []Shiiremeisai{}

	// SQL SELECT文を実行
	query := `SELECT sm.nengetsu, tm.torihikisakimei, hm.himokumei, st.tekiyou, sm.amount
	FROM shiiremeisai sm INNER JOIN shiiretekiyou st
	ON sm.nengetsu = st.nengetsu AND sm.torihikisakiID = st.torihikisakiID INNER JOIN
	himokumaster hm ON sm.himokuID = hm.himokuID INNER JOIN
	torihikisaki_master tm ON sm.torihikisakiID = tm.torihikisakiID
	WHERE sm.nengetsu = ?;`

	rows, err := db.Query(query, year_month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 結果をループしてスキャン
	for rows.Next() {
		var r Shiiremeisai
		if err := rows.Scan(&r.Nengetsu, &r.Torihikisakimei, &r.Himokumei, &r.Tekiyou, &r.Amount); err != nil {
			return nil, err
		}
		shiiremeisaiData = append(shiiremeisaiData, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return shiiremeisaiData, nil
}

// FindRangeByID は指定されたIDのt_メニューカテゴリー範囲をデータベースから取得します。
// func FindRangeByID(id int) (PriceRange, error) {
// 	var r PriceRange
// 	// SQL SELECT文を実行
// 	query := `SELECT ID, 部門ID, 範囲開始, 範囲終了, カテゴリー名 FROM t_メニューカテゴリー範囲 WHERE ID = ?`
// 	row := db.QueryRow(query, id)

// 	if err := row.Scan(&r.ID, &r.Bumon_ID, &r.Han_i_kaishi, &r.Han_i_shuryo, &r.Category_name); err != nil {
// 		return PriceRange{}, err
// 	}

// 	return r, nil
// }

// UpdateRange は指定されたIDのt_メニューカテゴリー範囲をデータベースで更新します。
// func UpdateRange(id int, r PriceRange) (PriceRange, error) {
// 	// SQL UPDATE文を実行
// 	query := `UPDATE t_メニューカテゴリー範囲 SET 部門ID = ?, 範囲開始 = ?, 範囲終了 = ?, カテゴリー名 = ? WHERE ID = ?`
// 	result, err := db.Exec(query, r.Bumon_ID, r.Han_i_kaishi, r.Han_i_shuryo, r.Category_name, id)
// 	if err != nil {
// 		return PriceRange{}, err
// 	}

// 	// 更新された行数を確認
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return PriceRange{}, err
// 	}
// 	if rowsAffected == 0 {
// 		return PriceRange{}, sql.ErrNoRows // 更新対象なし
// 	}

// 	r.ID = id
// 	return r, nil
// }

// DeleteRange は指定されたIDのt_メニューカテゴリー範囲をデータベースから削除します。
// func DeleteRange(id int) error {
// 	// SQL DELETE文を実行
// 	query := `DELETE FROM t_メニューカテゴリー範囲 WHERE ID = ?`
// 	result, err := db.Exec(query, id)
// 	if err != nil {
// 		return err
// 	}

// 	// 削除された行数を確認
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if rowsAffected == 0 {
// 		return sql.ErrNoRows // 削除対象なし
// 	}

// 	return nil
// }
