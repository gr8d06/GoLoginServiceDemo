package models

import (
	"errors"
	"fmt"
	"time"
)

//TwoFactor : an object for the longin holding the username and a one time key (sent to email, auth app, etc..)
type TwoFactor struct {
	UserName string
	OneKey   string
}

// would typically generate a "life" that the token would have too.
var oneTimeTokens = make(map[string]string)

//GenerateOneTimeKey : upon successfull username and password verification, a one time 4 digit key will be created and sent to the user email.
func GenerateOneTimeKey(userName string) string {
	t := time.Now()
	hrs := fmt.Sprintf("%02d", t.Hour())
	mins := fmt.Sprintf("%02d", t.Minute())
	strOneKey := hrs + mins
	//strconv.Itoa(t.Hour()) + strconv.Itoa(t.Minute())

	//for debugging purposes
	fmt.Println(strOneKey)

	oneTimeTokens[userName] = strOneKey
	return strOneKey
}

//ValidateOneTimeKey : checks to see if the submitted number is equal to the one in storage for the user.
func ValidateOneTimeKey(tf TwoFactor) error {

	if tf.OneKey != oneTimeTokens[tf.UserName] {
		fmt.Printf("Input: %s   Stored: %s", tf.OneKey, oneTimeTokens[tf.UserName])
		return errors.New("invalid two factor key")
	}

	return nil
}
