package helper

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TimeToString(t *time.Time) *string {
	if t == nil {
		return nil
	}
	formattedTime := t.Format("2006-01-02 15:04:05")
	return &formattedTime
}

func DateToString(t *time.Time) *string {
	if t == nil {
		return nil
	}
	formattedDate := t.Format("2006-01-02")
	return &formattedDate
}

func StringToTime(s string) (*time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func StringToDate(s string) (*time.Time, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}