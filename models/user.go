package models

import (
	"errors"
	"fmt"
)

//User : a struct to hold general information about a user.
type User struct {
	ID        int
	UserName  string
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

//GetUsers : returns a slice of pointers containingt User structs
func GetUsers() []*User {
	return users
}

//GetUserByUserName : Kinda self explanitory?
func GetUserByUserName(uName string) (User, error) {
	for _, u := range users {
		if u.UserName == uName {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("UserName '%v' not found", uName)
}

//GetUserByID : Kinda self explanitory?
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

//AddUser : add an address of a user object to the users var
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set to Zero")
	}

	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

//UpdateUser : Kinda self explanitory?
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

//RemoveUserByID : Kinda self explanitory?
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:1], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
