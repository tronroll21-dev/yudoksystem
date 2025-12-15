package models

import (
	"database/sql"
	"log"
)

// PriceRange はデータベースのt_メニューカテゴリー範囲テーブルの構造を表します。
// JSONタグは、Ginがリクエスト/レスポンスでJSONを扱うために使用されます。
type PriceRange struct {
	ID            int    `db:"ID" json:"id"`
	Bumon_ID      string `db:"部門ID" json:"bumon_id"`
	Han_i_kaishi  int    `db:"範囲開始" json:"han_i_kaishi"`
	Han_i_shuryo  int    `db:"範囲終了" json:"han_i_shuryo"`
	Category_name string `db:"カテゴリー名" json:"category_name"`
}

// InsertRange は新しいt_メニューカテゴリー範囲をデータベースに挿入します。
func InsertRange(r PriceRange) (PriceRange, error) {
	// SQL INSERT文を実行
	query := `INSERT INTO t_メニューカテゴリー範囲 (部門ID, 範囲開始, 範囲終了, カテゴリー名) VALUES (?, ?, ?, ?)`
	result, err := db.Exec(query, r.Bumon_ID, r.Han_i_kaishi, r.Han_i_shuryo, r.Category_name)
	if err != nil {
		return PriceRange{}, err
	}

	// 挿入された行のIDを取得
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("挿入されたIDの取得に失敗しました: ", err)
	}
	r.ID = int(id)
	return r, nil
}

// FindAllRanges は全てのt_メニューカテゴリー範囲をデータベースから取得します。
func FindAllRanges() ([]PriceRange, error) {
	ranges := []PriceRange{}

	// SQL SELECT文を実行
	query := `SELECT ID, 部門ID, 範囲開始, 範囲終了, カテゴリー名 FROM t_メニューカテゴリー範囲 ORDER BY 部門ID, 範囲開始`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 結果をループしてスキャン
	for rows.Next() {
		var r PriceRange
		if err := rows.Scan(&r.ID, &r.Bumon_ID, &r.Han_i_kaishi, &r.Han_i_shuryo, &r.Category_name); err != nil {
			return nil, err
		}
		ranges = append(ranges, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ranges, nil
}

// FindRangeByID は指定されたIDのt_メニューカテゴリー範囲をデータベースから取得します。
func FindRangeByID(id int) (PriceRange, error) {
	var r PriceRange
	// SQL SELECT文を実行
	query := `SELECT ID, 部門ID, 範囲開始, 範囲終了, カテゴリー名 FROM t_メニューカテゴリー範囲 WHERE ID = ?`
	row := db.QueryRow(query, id)

	if err := row.Scan(&r.ID, &r.Bumon_ID, &r.Han_i_kaishi, &r.Han_i_shuryo, &r.Category_name); err != nil {
		return PriceRange{}, err
	}

	return r, nil
}

// UpdateRange は指定されたIDのt_メニューカテゴリー範囲をデータベースで更新します。
func UpdateRange(id int, r PriceRange) (PriceRange, error) {
	// SQL UPDATE文を実行
	query := `UPDATE t_メニューカテゴリー範囲 SET 部門ID = ?, 範囲開始 = ?, 範囲終了 = ?, カテゴリー名 = ? WHERE ID = ?`
	result, err := db.Exec(query, r.Bumon_ID, r.Han_i_kaishi, r.Han_i_shuryo, r.Category_name, id)
	if err != nil {
		return PriceRange{}, err
	}

	// 更新された行数を確認
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return PriceRange{}, err
	}
	if rowsAffected == 0 {
		return PriceRange{}, sql.ErrNoRows // 更新対象なし
	}

	r.ID = id
	return r, nil
}

// DeleteRange は指定されたIDのt_メニューカテゴリー範囲をデータベースから削除します。
func DeleteRange(id int) error {
	// SQL DELETE文を実行
	query := `DELETE FROM t_メニューカテゴリー範囲 WHERE ID = ?`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	// 削除された行数を確認
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows // 削除対象なし
	}

	return nil
}
