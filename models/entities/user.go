package entities

import "errors"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) Exist() bool {
	return u != nil
}

func (u *User) GetID() int {
	if u.Exist() {
		return 0
	}

	return u.ID
}

func (u *User) GetName() string {
	if u.Exist() {
		return ""
	}

	return u.Name
}

func (u *User) GetEmail() string {
	if u.Exist() {
		return ""
	}

	return u.Email
}

func (u *User) SetID(id int) error {
	if u.Exist() {
		return errors.New("user is not created")
	}

	u.ID = id
	return nil
}

func (u *User) SetName(name string) error {
	if u.Exist() {
		return errors.New("user is not created")
	}

	u.Name = name
	return nil
}

func (u *User) SetEmail(email string) error {
	if u.Exist() {
		return errors.New("user is not created")
	}

	u.Email = email
	return nil
}
