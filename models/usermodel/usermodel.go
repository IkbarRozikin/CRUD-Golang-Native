package usermodel

import (
	"crud-go-native/config"
	"crud-go-native/entities"
	"fmt"
)

func CreatUser(user entities.User) (int64, error) {
	result, err := config.DB.Exec(`
	INSERT INTO users (username, email, password)
	VALUES (?, ?, ?)`,
		user.Username, user.Email, user.Password,
	)

	if err != nil {
		panic(err)
	}

	LasInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return LasInsertId, nil
}

// func GetUsername(username string) (*entities.User, error) {

// 	var user entities.User

// 	query := fmt.Sprintf("SELECT id, username, password FROM users WHERE username = '%s'", username)

// 	err := config.DB.QueryRow(query).Scan(&user.Id, &user.Username, &user.Password)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &user, nil
// }

func GetUsername(username string) (*entities.User, bool) {

	var user entities.User

	query := fmt.Sprintf("SELECT id, username, password FROM users WHERE username = '%s'", username)

	err := config.DB.QueryRow(query).Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		return &user, false

	}
	return &user, true
}
