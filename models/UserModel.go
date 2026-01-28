package models

import "log"

type User struct {
	ID       uint   `db:"担当ｺｰﾄﾞ"`
	Username string `db:"担当"`
	Password string `db:"password"`
}

func UpdatePassword(user *User) (*User, error) {

	updateRes, err := db.Exec("UPDATE 担当ﾏｽﾀ SET password = ? WHERE 担当ｺｰﾄﾞ = ?;", user.Password, user.ID)

	if err != nil {
		log.Printf("Error updating password for user ID %d: %v", user.ID, err)
		return &User{}, err
	}

	rows, err := updateRes.RowsAffected()

	if rows == 0 {
		log.Printf("UpdatePassword: No rows affected for user ID %d", user.ID)
		return &User{ID: 0, Username: "", Password: ""}, err
	}

	var result User

	err = db.QueryRow("SELECT * FROM 担当ﾏｽﾀ WHERE 担当ｺｰﾄﾞ = ?;", user.ID).Scan(&result.ID, &result.Username, &result.Password)

	if err != nil {
		log.Printf("Error fetching user after password update for ID %d: %v", user.ID, err)
		return nil, err
	}

	return &User{ID: result.ID, Username: result.Username, Password: result.Password}, nil
}

func GetUserById(id uint) (*User, error) {

	var result User

	log.Printf("Attempting to get user by ID: %d", id)
	err := db.QueryRow("SELECT * FROM 担当ﾏｽﾀ WHERE 担当ｺｰﾄﾞ = ?;", id).Scan(&result.ID, &result.Username, &result.Password)
	if err != nil {
		log.Printf("Error getting user by ID %d: %v", id, err)
		return nil, err
	}

	return &result, nil
}
