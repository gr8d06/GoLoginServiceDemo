package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

//RegisterControllers : does the front end routing to the proper controllers.
func RegisterControllers() {

	//uc := newUserController()
	//http.Handle("/users", *uc)
	//http.Handle("/users/", *uc)

	lc := newLoginController()
	http.Handle("/login", *lc)

	tf := newTwoFactorController()
	http.Handle("/twofactor", tf)

}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
