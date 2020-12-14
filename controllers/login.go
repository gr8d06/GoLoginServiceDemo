package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gr8d06/gologin/models"
)

type loginController struct {
}

func newLoginController() *loginController {
	return &loginController{}
}

func (li loginController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if r.URL.Path == "/login" {
		switch r.Method {
		case http.MethodPost:
			li.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (li *loginController) post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit the post controller")
	c, err := li.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse login request object"))
		return
	}

	err = models.ValidateCredentials(c)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid Credentials"))
	}

	u, err := models.GetUserByUserName(c.UserName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (li *loginController) parseRequest(r *http.Request) (models.Credentials, error) {
	dec := json.NewDecoder(r.Body)
	var c models.Credentials
	err := dec.Decode(&c)
	if err != nil {
		return models.Credentials{}, err
	}

	return c, nil
}
