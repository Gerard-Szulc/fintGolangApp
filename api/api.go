package api

import (
	"encoding/json"
	"fintGolangApp/helpers"
	"fintGolangApp/interfaces"
	"fintGolangApp/users"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Login struct {
	Username string
	Password string
}
type ErrResponse struct {
	Message string
}

func StartApi() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	fmt.Println("App is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))

}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	fmt.Println(formattedBody.Username, formattedBody.Password)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	if login["message"] == "success.response_success" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "error.wrong_username_or_password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	// Read body
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)
	// Handle registration
	var formattedBody interfaces.Register
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	fmt.Println(formattedBody.Username, formattedBody.Password)
	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)
	// Prepare response
	if register["message"] == "success.response_success" {
		resp := register
		json.NewEncoder(w).Encode(resp)
		// Handle error in else
	} else {
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}
