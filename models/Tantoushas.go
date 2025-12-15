package models

type Tantousha struct {
	ID   int    `db:"担当ｺｰﾄﾞ" json:"id"`
	Name string `db:"担当" json:"name"`
}

func GetTantoushas() ([]Tantousha, error) {
	rows, err := db.Query("SELECT 担当ｺｰﾄﾞ, 担当 FROM 担当ﾏｽﾀ ORDER BY 担当ｺｰﾄﾞ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tantoushas []Tantousha
	for rows.Next() {
		var t Tantousha
		err := rows.Scan(&t.ID, &t.Name)
		if err != nil {
			return nil, err
		}
		tantoushas = append(tantoushas, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tantoushas, nil
}
