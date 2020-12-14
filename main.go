package main

import (
	"net/http"

	"github.com/gr8d06/gologin/controllers"
	"github.com/gr8d06/gologin/models"
)

func main() {
	//initialize with the seed data for the test.
	models.AddUser(models.User{
		ID:        0,
		UserName:  "c137@onecause.com",
		FirstName: "Tim",
		LastName:  "Sublette",
	})

	models.AddCreds("c137@onecause.com", "#th@nH@rm#y#r!$100%D0p#")

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
