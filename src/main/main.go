package main

import (
	util "../util"
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
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	http.HandleFunc("/todd", util.ToddFunc)
	http.HandleFunc("/ming", util.MingFunc)
	http.HandleFunc("/rio", util.RioFunc)
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body><a href="/login">Google Log In</a></body></html>`
	fmt.Fprint(w, html)
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

	if err != nil {
		fmt.Printf("could nor parse response: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	/*Useful for debugging
	fmt.Fprint(w, `<html><body>`)
	fmt.Fprintf(w, "Response: %s", content)
	home := `<a href=/todd>Home</a>`
	fmt.Fprint(w, home)
	fmt.Fprint(w, `</body></html>`)
	*/
	//Uses home.html as the home page

	data, err := ioutil.ReadFile("templates/home.html")
	if err == nil {
		w.Write(data)
	} else {
		log.Println(err.Error())
	}
}
