package main

import (
	htmpTemplate "html/template"
	"os"
	"text/template"
)

func main() {
	HTMLTemplate()
}

type user struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Privilege int    `json:"privilege"`
	URL       string `json:"url"`
}

var users = []user{
	user{ID: "id-1", Name: "user-1<lt>", Privilege: 100, URL: "http://google.com"},
	user{ID: "id-2", Name: "user-2", Privilege: 200, URL: "http://google.com"},
	user{ID: "id-3", Name: "user-3", Privilege: 300, URL: "http://google.com"},
}

func AdjustPriv(priv int) int {
	return 500 - priv
}

func HTMLTemplate() {
	templ := `<h1>Users:</h1>
	<table style='text-align: center'>
	<tr>
	<th>ID</th>
	<th>Name</th>
	<th>Privilege</th>
	<th>URL</th>
	</tr>
	{{range .}}
	<tr>
	<td>{{.ID}}</td>
	<td>{{.Name | printf "Mr/Mrs %s"}}</td>
	<td>{{.Privilege | adjustPriv}}</td>
	<td><a href={{.URL}}>{{.URL}}</a></td>
	</tr>
	{{end}}
	</table>`

	report := htmpTemplate.
		Must(htmpTemplate.
			New("users").
			Funcs(htmpTemplate.FuncMap{"adjustPriv": AdjustPriv}).
			Parse(templ))

	file, err := os.Create("user.html")
	if err != nil {
		panic(err)
	}

	err = report.Execute(file, users)
	if err != nil {
		panic(err)
	}
}

func TextTemplate() {
	templ := `Users:
	{{range .}}---------------------------
	ID: {{.ID}}
	Name: {{.Name | printf "Mr/Mrs %s"}}
	Privilege:{{.Privilege | adjustPriv}}
	URL: {{.URL}}
	{{end}}`
	report := template.Must(template.New("users").Funcs(template.FuncMap{"adjustPriv": AdjustPriv}).Parse(templ))

	err := report.Execute(os.Stdout, users)
	if err != nil {
		panic(err)
	}
}
