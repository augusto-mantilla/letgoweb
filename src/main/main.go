package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//	"encoding/json"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	util "../util"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
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
	util.InitDB("letme:hello@tcp(localhost:3306)/letmefix")
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)

	http.HandleFunc("/newRequest", util.NewRequest)

	http.HandleFunc("/professional", util.Professional)
	http.HandleFunc("/commercial", util.Commercial)
	http.HandleFunc("/private", util.Private)

	http.HandleFunc("/commercial/profile", util.CommercialProfile)
	http.HandleFunc("/commercial/showRequests", util.CommercialShowRequests)

	http.HandleFunc("/professional/profile", util.ProfessionalProfile)
	http.HandleFunc("/professional/proposal", util.Proposal)
	http.HandleFunc("/professional/showProposals", util.ShowProposals)

	http.HandleFunc("/private/profile", util.PrivateProfile)
	http.HandleFunc("/private/showRequests", util.PrivateShowRequests)

	http.HandleFunc("/request/public", util.PublicRequest)
	http.HandleFunc("/request/direct", util.DirectRequest)
	http.HandleFunc("/request/execution", util.ExecutionRequest)

	http.HandleFunc("/proposals/showPublicProposals", util.ShowPublicProposals)

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
