package repository

import (
	"database/sql"
	"fmt"
	"roar-dashboard-backend/database"
	"roar-dashboard-backend/models"
)

type Repo struct {
	DB *sql.DB
}

func (repo *Repo) GetUserIdFromToken(token string) (int, error) {
	// query := "SELECT id from auth_utoken where key=(?)"
	// ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancelfunc()
	var userId int
	fmt.Println(repo.DB, token)
	if err := repo.DB.QueryRow("SELECT * FROM auth_token",
		token).Scan(&userId); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("0000000000")
			return 0, err
		}
		fmt.Println("///////////////", err.Error())
		return 0, err
	}
	fmt.Println("//////////////", userId)
	return userId, nil
}

func GetUserDetailsFromId(userId int) (models.User, error) {
	var user models.User
	if err := NewRepo().DB.QueryRow("SELECT id, username, is_staff, is_active from auth_user where user_id = ?",
		userId).Scan(&user); err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, err
		}
		return models.User{}, err
	}
	return user, nil
}

func NewRepo() Repo {
	return Repo{
		DB: database.InitDb(),
	}
}
