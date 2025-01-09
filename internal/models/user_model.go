package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID           int64    `json:"id"`
	Email        string   `json:"email"`
	FirstName    string   `json:"first_name"`
	LastName     string   `json:"last_name"`
	Username     string   `json:"username"`
	Password     password `json:"-"`
	ProfileImage string   `json:"profile_image"`
	Version      int64    `json:"-"`
	Role         string   `json:"role"`
	// Role         Role     `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type password struct {
	text *string
	Hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.Hash = hash
	p.text = &text

	return nil
}
