package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"mini-inventory/config"
	"mini-inventory/modules/user"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (s *AuthService) Register(input user.User) (user.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user.User{}, err
	}
	input.Password = string(hashedPassword)

	query := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
	err = config.DB.QueryRow(query, input.Name, input.Email, input.Password, time.Now(), time.Now()).Scan(&input.ID, &input.CreatedAt, &input.UpdatedAt)
	if err != nil {
		return user.User{}, err
	}

	return input, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	var u user.User
	query := `SELECT * FROM users WHERE email = $1`
	err := config.DB.Get(&u, query, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", err
	}
	token := hex.EncodeToString(tokenBytes)

	hashedToken := hashToken(token)
	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	tokenQuery := `INSERT INTO personal_access_tokens (user_id, token, expires_at, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err = config.DB.Exec(tokenQuery, u.ID, hashedToken, expiresAt, time.Now(), time.Now())
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Logout(token string) error {
	hashedToken := hashToken(token)
	query := `DELETE FROM personal_access_tokens WHERE token = $1`
	_, err := config.DB.Exec(query, hashedToken)
	return err
}

func (s *AuthService) GetUserByToken(token string) (user.User, error) {
	hashedToken := hashToken(token)
	var u user.User
	query := `
		SELECT u.* FROM users u
		JOIN personal_access_tokens t ON u.id = t.user_id
		WHERE t.token = $1 AND t.expires_at > NOW()
	`
	err := config.DB.Get(&u, query, hashedToken)
	if err != nil {
		return user.User{}, errors.New("unauthorized")
	}
	return u, nil
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
