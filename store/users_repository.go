package store

import (
	"context"
	"database/sql"
	"log"
	"userStore/model"
)

type (
	UsersRepository struct {
		store *Store
	}
)

func NewUsersRepository(store *Store) *UsersRepository {
	return &UsersRepository{
		store: store,
	}
}

func (r *UsersRepository) GetAll(ctx context.Context) ([]model.User, error) {
	results := []model.User{}

	rows, err := r.store.db.QueryContext(ctx, "SELECT u.id, u.username, up.first_name, up.last_name, up.city, ud.school FROM user u LEFT JOIN user_profile up ON u.id = up.user_id LEFT JOIN user_data ud ON u.id = ud.user_id")
	if err != nil {
		log.Println("Getall Query error: ", err)
		return nil, model.ErrInternalServerError
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.City, &user.School)
		if err != nil {
			return nil, model.ErrInternalServerError
		}
		results = append(results, user)
	}

	if err := rows.Err(); err != nil {
		return nil, model.ErrInternalServerError
	}

	return results, nil
}

func (r *UsersRepository) Find(username string, ctx context.Context) (model.User, error) {
	var user model.User
	err := r.store.db.QueryRowContext(ctx, "SELECT u.id, u.username, up.first_name, up.last_name, up.city, ud.school FROM user u LEFT JOIN user_profile up ON u.id = up.user_id LEFT JOIN user_data ud ON u.id = ud.user_id WHERE u.username=?", username).Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.City, &user.School)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Find QueryRow ErrNoRows error: ", err)
			return model.User{}, model.ErrInternalServerError
		}
		log.Println("Find QueryRow error: ", err)
		return model.User{}, model.ErrInternalServerError
	}

	return user, nil
}
