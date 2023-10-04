package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	NickName  string    `json:"nickName,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare() error {
	err := user.validate()
	if err != nil {
		return err
	}
	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("field name is required")
	}
	if user.NickName == "" {
		return errors.New("field nickname is required")
	}
	if user.Email == "" {
		return errors.New("field email is required")
	}
	if user.Password == "" {
		return errors.New("field password is required")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Email = strings.TrimSpace(user.Email)

}
