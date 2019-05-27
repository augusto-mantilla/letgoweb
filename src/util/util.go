package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type MyHandler struct {
}

type Context struct {
	FirstName string
	Message   string
	URL       string
	Beers     []string
	Title     string
}

type Context1 struct {
	FirstName string
	Message   string
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)
	//    path := "templates" + r.URL.Path

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Write(data)
	} else if path == "" {
		data1, err1 := ioutil.ReadFile("templates/home.html")
		if err1 == nil {
			w.Write(data1)
		} else {
			log.Println(err1.Error())
		}
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 My Friend - " + http.StatusText(404)))
	}
}

func MyHandlerFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content Type", "text/html")
	templates := template.New("template")
	templates.New("test").Parse(doc)
	templates.New("header").Parse(head)
	templates.New("footer").Parse(foot)
	context := Context{
		"Todd",
		"more beer, please",
		req.URL.Path,
		[]string{"New Belgium", "La Fin Du Monde", "The Alchemist"},
		"Favorite Beers",
	}
	templates.Lookup("test").Execute(w, context)
}

const doc = `
{{template "header" .Title}}
<body>

    <h1>{{.FirstName}} says, "{{.Message}}"</h1>

    {{if eq .URL "/nobeer"}}
        <h2>We're out of beer, {{.FirstName}}. Sorry!</h2>
    {{else}}
        <h2>Yes, grab another beer, {{.FirstName}}</h2>
        <ul>
            {{range .Beers}}
            <li>{{.}}</li>
            {{end}}
        </ul>
    {{end}}

    <hr>

    <h2>Here's all the data:</h2>
    <p>{{.}}</p>
<a href="/ming/">Ming!</a>
</body>
{{template "footer"}}`

const head = `
<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>{{.}}</title>
</head>
`

const foot = `
</html>
`

const doc1 = `<!DOCTYPE html>
<html>
  <head lang="en">
    <meta charset="UTF-8">
    <title>First Template</title>
  </head>
  <body>
    <h1>My name is {{.FirstName}}</h1>
    <p>{{.Message}}</p>
    <a href="/ming">Ming!</a>
  </body>
</html>`

const filename = "templates/test.html"

func Professional(w http.ResponseWriter, req *http.Request) {
	path := "templates/professional.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	InsertAddress("9020-123", "Caminho")
	w.Write(data)
}

func ProfessionalProfile(w http.ResponseWriter, req *http.Request) {
	path := "templates/professionalProfile.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}

func PrivateProfile(w http.ResponseWriter, req *http.Request) {
	path := "templates/privateProfile.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}

func CommercialProfile(w http.ResponseWriter, req *http.Request) {
	path := "templates/commercialProfile.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}

func Commercial(w http.ResponseWriter, req *http.Request) {
	path := "templates/commercial.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}

func Private(w http.ResponseWriter, req *http.Request) {
	path := "templates/private.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}

func Proposal(w http.ResponseWriter, req *http.Request) {
	path := "templates/proposalForm.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}

func PublicRequest(w http.ResponseWriter, req *http.Request) {
	path := "templates/publicRequest.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}

func DirectRequest(w http.ResponseWriter, req *http.Request) {
	path := "templates/directRequest.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}

func ShowPublicProposals(w http.ResponseWriter, req *http.Request) {
	path := "templates/presentPublicProposals.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}

func ExecutionRequest(w http.ResponseWriter, req *http.Request) {
	path := "templates/executionRequest.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}

func CommercialShowRequests(w http.ResponseWriter, req *http.Request) {
	path := "templates/commercialShowRequests.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}

func PrivateShowRequests(w http.ResponseWriter, req *http.Request) {
	path := "templates/privateShowRequests.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}

func NewRequest(w http.ResponseWriter, req *http.Request) {
	path := "templates/newRequest.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}

func ToddFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content Type", "text/html")
	tmpl, err := template.New("toddTemplate").Parse(doc1)
	if err == nil {
		context := Context1{"Todd", "more Go, please"}
		tmpl.Execute(w, context)
	} else {
		log.Println(err.Error())
	}
}

func Message(w http.ResponseWriter, req *http.Request) {
	path := "templates/showRequests.html"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not load page: %s\n", path)
		return
	}
	w.Write(data)
}
