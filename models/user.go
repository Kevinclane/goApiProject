package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
	Password  string
}

var (
	users      []*User
	nextUserID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include an ID or it must be set to 0")
	}
	if u.Username == "" {
		return User{}, errors.New("Must provide a username")
	}
	if u.Password == "" {
		return User{}, errors.New("Must provide a password")
	}
	u.ID = nextUserID

	nextUserID++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with id '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with id '%v' not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with id '%v' not found", id)
}
