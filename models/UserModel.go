package models

type User struct {
	ID       uint   `db:"担当ｺｰﾄﾞ"`
	Username string `db:"担当"`
	Password string `db:"password"`
}

func UpdatePassword(user *User) (*User, error) {

	updateRes, err := db.Exec("UPDATE 担当ﾏｽﾀ SET password = ? WHERE 担当ｺｰﾄﾞ = ?;", user.Password, user.ID)

	if err != nil {
		return &User{}, err
	}

	rows, err := updateRes.RowsAffected()

	if rows == 0 {
		return &User{ID: 0, Username: "", Password: ""}, err
	}

	var result User

	err = db.QueryRow("SELECT * FROM 担当ﾏｽﾀ WHERE 担当ｺｰﾄﾞ = ?;", user.ID).Scan(&result.ID, &result.Username, &result.Password)

	if err != nil {
		return nil, err
	}

	return &User{ID: result.ID, Username: result.Username, Password: result.Password}, nil
}

func GetUserById(id uint) (*User, error) {

	var result User

	err := db.QueryRow("SELECT * FROM 担当ﾏｽﾀ WHERE 担当ｺｰﾄﾞ = ?;", id).Scan(&result.ID, &result.Username, &result.Password)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
