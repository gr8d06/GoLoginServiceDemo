package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gr8d06/gologin/models"
)

type twoFactorController struct{}

func newTwoFactorController() *twoFactorController {
	return &twoFactorController{}
}

func (tf twoFactorController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/twofactor" {
		switch r.Method {
		case http.MethodPost:
			tf.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (tf *twoFactorController) post(w http.ResponseWriter, r *http.Request) {
	tfRequest, err := tf.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse two factor request object"))
		return
	}

	err = models.ValidateOneTimeKey(tfRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//it would be typical to return a session token, but this will due for a demo.
	encodeResponseAsJSON("Success", w)
}

func (tf twoFactorController) parseRequest(r *http.Request) (models.TwoFactor, error) {
	dec := json.NewDecoder(r.Body)
	var t models.TwoFactor
	err := dec.Decode(&t)
	if err != nil {
		return models.TwoFactor{}, err
	}

	return t, nil
}
