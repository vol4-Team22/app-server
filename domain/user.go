package domain

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserID int64

type User struct {
	UserID   UserID    `json:"user_id" db:"user_id"`
	UserName string    `json:"user_name" db:"user_name"`
	Password string    `json:"password" db:"password"`
	Role     string    `json:"role" db:"role"`
	Created  time.Time `json:"created_at" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
}

func (u *User) ComparePassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
}
