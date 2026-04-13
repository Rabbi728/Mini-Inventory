package user

import (
	"basic-inventory-app/config"
	"time"
)

type UserService struct{}

func (s *UserService) GetAllUsers() ([]User, error) {
	var users []User
	query := `SELECT id, name, username, email, password, created_at, updated_at FROM users`
	err := config.DB.Select(&users, query)
	return users, err
}

func (s *UserService) CreateUser(user *User) error {
	query := `INSERT INTO users (name, username, email, password, created_at, updated_at) 
			  VALUES (:name, :username, :email, :password, :created_at, :updated_at) RETURNING id`
	
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	
	nstmt, err := config.DB.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer nstmt.Close()
	
	err = nstmt.Get(&user.ID, user)
	return err
}

func (s *UserService) GetUserByID(id string) (User, error) {
	var user User
	query := `SELECT id, name, username, email, password, created_at, updated_at FROM users WHERE id = $1`
	err := config.DB.Get(&user, query, id)
	return user, err
}

func (s *UserService) UpdateUser(user *User) error {
	query := `UPDATE users SET name = :name, username = :username, email = :email, 
			  password = :password, updated_at = :updated_at WHERE id = :id`
	
	user.UpdatedAt = time.Now()
	_, err := config.DB.NamedExec(query, user)
	return err
}

func (s *UserService) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := config.DB.Exec(query, id)
	return err
}