package main

import (
	util "../util"
	//	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL: "http://localhost:8080/callback",
		ClientID:    os.Getenv("GOOGLE_CLIENT_ID"),
		//"608621652614-2qd9a4gbf7gq4eu9god80o8bnh9pm2qq.apps.googleusercontent.com",
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		//"4-LZrLdA26_P5KkIdjky1HUI",
		Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}

	//TODO: randomize
	randomState = "random"
	admin       = User{"I don't", "Care", "aug.ornelas@gmail.com"}
)

type User struct {
	name     string
	lastname string
	email    string
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	http.HandleFunc("/professional", util.Professional)
	http.HandleFunc("/commercial", util.Commercial)
	http.HandleFunc("/professional/profile", util.ProfessionalProfile)
	http.HandleFunc("/professional/proposal", util.Proposal)
	http.HandleFunc("/private", util.Private)
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	path := "templates/login.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != randomState {
		fmt.Println("state is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))

	if err != nil {
		fmt.Printf("could not get token: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("could not create get request: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()
	//	content, err := ioutil.ReadAll(resp.Body)

	/*	c := make(map[string]string)
		//	e := json.Unmarshal(content, &c)

		if e != nil {
			fmt.Printf("could not parse response: %s\n", e.Error())
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}

		user := User{email: c["email"]}*/
	if err != nil {
		fmt.Printf("could nor parse response: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	//Useful for debugging
	/*	fmt.Fprint(w, `<html><body>`)
		fmt.Fprintf(w, "Response: %s", content)
		home := `<a href=/professional>Home</a>`
		fmt.Fprintf(w, home)

		fmt.Fprint(w, `</body></html>`)

		//Uses home.html as the home page
		//	fmt.Fprint(w, user.email)*/
	data, err := ioutil.ReadFile("templates/register.html")
	if err == nil {
		w.Write(data)
	} else {
		log.Println(err.Error())
	}
}
