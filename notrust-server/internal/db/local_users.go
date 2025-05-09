package db

import (
	"context"
)

type LocalUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type NewLocalUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetLocalUsers() ([]LocalUser, error) {
	rows, err := Pool.Query(context.Background(), "SELECT id, email FROM local_users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []LocalUser
	for rows.Next() {
		var u LocalUser
		if err := rows.Scan(&u.ID, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func CreateLocalUser(user NewLocalUser) error {
	_, err := Pool.Exec(context.Background(), "INSERT INTO local_users (email, password_hash) VALUES ($1, $2)", user.Email, user.Password)
	return err
}

func DeleteLocalUser(id string) error {
	_, err := Pool.Exec(context.Background(), "DELETE FROM local_users WHERE id=$1", id)
	return err
}

