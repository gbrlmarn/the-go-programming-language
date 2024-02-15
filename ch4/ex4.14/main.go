// Exercise 4.14: Create a web server that queries GitHub once and then allows navigation of the list of bug repors, milestones, and users.
package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type UserSearchResult struct {
	Users []User
}

type User struct {
	Login  string `json:"login"`
	Id     int    `json:"id"`
	NodeId string `json:"node_id"`
	Url    string `json:"html_url"`
	Avatar string `json:"avatar_url"`
}

const (
	usersAPI = "https://api.github.com/users"
	usage    = `
    go run main.go since=<int> per_page=<int>`
	templ = `
<h1>users</h1>      
<table>
<tr style='text-align: left'>
    <th>#</th>
    <th>User</th>
    <th>NodeID</th>
</tr>
{{range .Users}}
<tr>
    <td><a href='{{.Url}}'>{{.Id}}</a></td>'
    <td><a href='{{.Avatar}}'>{{.Login}}</a></td>'
    <td>{{.NodeId}}</td>'
</tr>
{{end}}
</table>
    `
)

var usersList = template.Must(template.New("userslist").Parse(templ))

func main() {
	if len(os.Args) < 2 {
		log.Fatal(usage)
	}
	result, err := SearchUsers(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := usersList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func SearchUsers(terms []string) (*UserSearchResult, error) {
	q := strings.Join(terms, "&")
	resp, err := http.Get(usersAPI + "?" + q)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var result UserSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result.Users); err != nil {
		log.Fatal(err)
	}
	return &result, nil
}
