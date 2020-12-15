package models

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//Credentials : an object for the login process
type Credentials struct {
	UserName string
	PassWord string
	OneKey   string
}

// a map of the userId and the encrypted password
var storedCreds = make(map[string][]byte)

//ValidateCredentials : takes a Credential and User object. Uses bCrypt to compare the stored hashed password, with the hashed version of the password that was received.
func ValidateCredentials(cred Credentials) error {
	err := bcrypt.CompareHashAndPassword([]byte(storedCreds[cred.UserName]), []byte(cred.PassWord))

	if err != nil {
		return err
	}
	GenerateOneTimeKey(cred.UserName)
	//would typically send the one time key to a secondary account(email, text, authapp), or have one generated.
	tf := TwoFactor{cred.UserName, cred.OneKey}

	err = (ValidateOneTimeKey(tf))
	fmt.Println(err)
	return err
}

//AddCreds : add an entry to the storedCreds map.
func AddCreds(userName string, strPass string) error {
	if len(strPass) < 8 {
		return errors.New("password must be longer than 8 characters")
	}

	passBytes, err := bcrypt.GenerateFromPassword([]byte(strPass), 8) //https://xkcd.com/936/
	if err != nil {
		return errors.New("could not generate an encrypted password")
	}
	storedCreds[userName] = passBytes

	return err
}
