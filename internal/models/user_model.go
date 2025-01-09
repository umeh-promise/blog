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
	RoleObj      Role     `json:"-"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
}

type password struct {
	text *string
	Hash []byte
}

func (p *password) HashPassword(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.Hash = hash
	p.text = &text

	return nil
}

func (p *password) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword(p.Hash, []byte(password))
}
