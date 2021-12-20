package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func cript(msg string) string {
	sh := sha1.New()
	sh.Write([]byte(msg))
	sha1 := hex.EncodeToString(sh.Sum(nil))
	return sha1
}

type User struct {
	Cpf      string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Location string `json:"location"`
	Phone    string `json:"phone"`
}

type UserParams struct {
	ApiKey    string `json:"platform_api_key"`
	Signature string `json:"sha1_signature"`
	Name      string `json:"network_name"`
	Encoding  string `json:"encoding"`
	IdType    string `json:"id_type"`
	User      string `json:"user_data"`
}

type Event struct {
	Action string `json:"action"`
}

type EventParams struct {
	ApiKey    string `json:"platform_api_key"`
	Signature string `json:"sha1_signature"`
	Name      string `json:"network_name"`
	Encoding  string `json:"encoding"`
	IdType    string `json:"id_type"`
	Event     string `json:"event"`
}

func create_user(sha1, id string) {
	url_user := "https://login.plataformasocial.com.br/users/portal/" + id + "/signup"
	user := User{
		Name:     "teste-onboarding",
		Email:    "teste-onboarding@dito.com.br",
		Gender:   "female",
		Birthday: "1995-01-18",
		Location: "Manaus",
	}

	userJson, _ := json.Marshal(user)
	params := UserParams{
		ApiKey:    os.Getenv("API_KEY"),
		Signature: sha1,
		Name:      "pt",
		Encoding:  "base64",
		IdType:    "id",
		User:      string(userJson),
	}

	req, err := json.Marshal(params)

	if err != nil {
		fmt.Println("deu erro")
	}

	resp_event, err := http.Post(url_user, "application/json", bytes.NewBuffer(req))
	if err != nil {
		fmt.Println(err)
	} else {
		respp, _ := ioutil.ReadAll(resp_event.Body)
		str := string(respp[:])
		fmt.Println(str)
	}
}

func create_event(sha1, id string) {

	url_event := "http://events.plataformasocial.com.br/users/" + id

	event := Event{Action: "evento-1"}
	eventJson, _ := json.Marshal(event)
	params := EventParams{
		ApiKey:    os.Getenv("API_KEY"),
		Signature: sha1,
		Name:      "pt",
		Encoding:  "base64",
		IdType:    "id",
		Event:     string(eventJson),
	}

	req, err := json.Marshal(params)

	if err != nil {
		fmt.Println("deu erro")
	}

	resp_event, err := http.Post(url_event, "application/json", bytes.NewBuffer(req))
	if err != nil {
		fmt.Println(err)
	} else {
		respp, _ := ioutil.ReadAll(resp_event.Body)
		str := string(respp[:])
		fmt.Println(str)
	}
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sha1 := cript(os.Getenv("API_SECRET"))
	id := "92256834212"

	create_user(sha1, id)
	create_event(sha1, id)
}
