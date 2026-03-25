package models

import (
	"database/sql"
)

// PriceRange はデータベースのt_メニューカテゴリー範囲テーブルの構造を表します。
// JSONタグは、Ginがリクエスト/レスポンスでJSONを扱うために使用されます。
type Shiiremeisai struct {
	Nengetsu        int    `db:"nengetsu" json:"year_month"`
	TorihikisakiID  int    `db:"torihikisakiID" json:"contractor_id"`
	HimokuID        int    `db:"himokuID" json:"expense_category_id"`
	Torihikisakimei string `db:"torihikisakimei" json:"contractor_name"`
	Tekiyou         string `db:"tekiyou" json:"contractor_tekiyou"`
	Himokumei       string `db:"himokumei" json:"expense_category"`
	Shouhizei       string `db:"tekiyou" json:"shouhizei"`
	Amount          int    `db:"amount" json:"amount"`
}

type AvailableContractor struct {
	TorihikisakiID  int    `json:"contractor_id"`
	Torihikisakimei string `json:"contractor_name"`
	Tekiyou         string `json:"tekiyou"`
}

// SaveShiiremeisai performs an upsert on the shiiremeisai table based on IDs.
func SaveShiiremeisai(s Shiiremeisai) error {
	var count int
	queryCheck := `SELECT COUNT(*) FROM shiiremeisai WHERE nengetsu = ? AND torihikisakiID = ? AND himokuID = ?`
	err := db.QueryRow(queryCheck, s.Nengetsu, s.TorihikisakiID, s.HimokuID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		// Update existing record
		queryUpdate := `UPDATE shiiremeisai SET amount = ? WHERE nengetsu = ? AND torihikisakiID = ? AND himokuID = ?`
		_, err = db.Exec(queryUpdate, s.Amount, s.Nengetsu, s.TorihikisakiID, s.HimokuID)
	} else {
		// Insert new record
		queryInsert := `INSERT INTO shiiremeisai (nengetsu, torihikisakiID, himokuID, amount) VALUES (?, ?, ?, ?)`
		_, err = db.Exec(queryInsert, s.Nengetsu, s.TorihikisakiID, s.HimokuID, s.Amount)
	}

	return err
}

func DeleteShiiremeisai(s Shiiremeisai) error {

	// 既存レコードの削除
	queryInsert := `DELETE FROM shiiremeisai WHERE nengetsu = ? AND torihikisakiID = ? AND himokuID = ?`
	_, err := db.Exec(queryInsert, s.Nengetsu, s.TorihikisakiID, s.HimokuID)

	return err
}

// SaveShiiretekiyou performs an upsert on the shiiretekiyou table.
func SaveShiireshouhizei(nengetsu int, torihikisakiID int, shouhizei string) error {
	var count int
	queryCheck := `SELECT COUNT(*) FROM shiireshouhizei WHERE nengetsu = ? AND torihikisakiID = ?`
	err := db.QueryRow(queryCheck, nengetsu, torihikisakiID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		queryUpdate := `UPDATE shiireshouhizei SET shouhizei = ? WHERE nengetsu = ? AND torihikisakiID = ?`
		_, err = db.Exec(queryUpdate, shouhizei, nengetsu, torihikisakiID)
	} else {
		queryInsert := `INSERT INTO shiireshouhizei (nengetsu, torihikisakiID, shouhizei) VALUES (?, ?, ?)`
		_, err = db.Exec(queryInsert, nengetsu, torihikisakiID, shouhizei)
	}

	return err
}

// GetAvailableContractors returns a list of contractors not in the specified month, with their last used remarks.
func GetAvailableContractors(yearMonth int) ([]AvailableContractor, error) {
	query := `
		SELECT tm.torihikisakiID, tm.torihikisakimei, tm.tekiyou
		FROM torihikisaki_master tm
		WHERE tm.torihikisakiID NOT IN (
			SELECT torihikisakiID FROM shiiremeisai WHERE nengetsu = ?
		)
		ORDER BY tm.torihikisakiID
	`
	rows, err := db.Query(query, yearMonth)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contractors []AvailableContractor
	for rows.Next() {
		var c AvailableContractor
		if err := rows.Scan(&c.TorihikisakiID, &c.Torihikisakimei, &c.Tekiyou); err != nil {
			return nil, err
		}
		contractors = append(contractors, c)
	}
	return contractors, nil
}

// AddContractorToMonth initializes a contractor in a month with its last used category and remark.
func AddContractorToMonth(yearMonth int, contractorID int, tekiyou string) error {
	// Find last used himokuID for this contractor and tekiyou combination
	var himokuID int
	queryHimoku := `
		SELECT sm.himokuID
		FROM shiiremeisai sm
		WHERE sm.torihikisakiID = ?
		ORDER BY sm.nengetsu DESC
		LIMIT 1
	`
	err := db.QueryRow(queryHimoku, contractorID).Scan(&himokuID)
	if err != nil {
		// Fallback: last used category regardless of remark
		queryFallback := `SELECT himokuID FROM shiiremeisai WHERE torihikisakiID = ? ORDER BY nengetsu DESC LIMIT 1`
		err = db.QueryRow(queryFallback, contractorID).Scan(&himokuID)
		if err != nil {
			himokuID = 1
		}
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insert into shiiremeisai
	_, err = tx.Exec("INSERT INTO shiiremeisai (nengetsu, torihikisakiID, himokuID, amount) VALUES (?, ?, ?, 0)", yearMonth, contractorID, himokuID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// InitializeShiiremeisai finds the latest month with data and copies its structure to the newYearMonth.
// It copies records to shiiremeisai (with amount 0) and remarks to shiiretekiyou.
func InitializeShiiremeisai(newYearMonth int) error {
	var latestNengetsu int
	// Find the latest year_month that has data
	queryLatest := `SELECT MAX(nengetsu) FROM shiiremeisai WHERE nengetsu < ?`
	err := db.QueryRow(queryLatest, newYearMonth).Scan(&latestNengetsu)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	// If no previous data found, we might want to return an error or do nothing
	if latestNengetsu == 0 {
		return nil
	}

	// Use a transaction to ensure both tables are updated correctly
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 1. Copy records from latestNengetsu to newYearMonth with amount = 0
	queryCopyMeisai := `
		INSERT INTO shiiremeisai (nengetsu, torihikisakiID, himokuID, amount)
		SELECT ?, torihikisakiID, himokuID, 0
		FROM shiiremeisai
		WHERE nengetsu = ?
	`
	if _, err := tx.Exec(queryCopyMeisai, newYearMonth, latestNengetsu); err != nil {
		return err
	}

	return tx.Commit()
}

// FindAllShiiremeisai は全てのt_メニューカテゴリー範囲をデータベースから取得します。
func FindAllShiiremeisai(year_month string) ([]Shiiremeisai, error) {
	shiiremeisaiData := []Shiiremeisai{}

	// SQL SELECT文を実行
	// Changed to LEFT JOIN for shiiretekiyou to ensure initialized records (without tekiyou) are returned.
	query := `SELECT sm.nengetsu, sm.torihikisakiID, sm.himokuID, tm.torihikisakimei, tm.tekiyou, hm.himokumei, COALESCE(st.shouhizei, 0), sm.amount
	FROM shiiremeisai sm LEFT JOIN shiireshouhizei st
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
		if err := rows.Scan(&r.Nengetsu, &r.TorihikisakiID, &r.HimokuID, &r.Torihikisakimei, &r.Tekiyou, &r.Himokumei, &r.Shouhizei, &r.Amount); err != nil {
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
