package main

import (
	"html/template"
	"log"
	"os"

	"gopl/ch4/examples/github"
)

const templ = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
    <th>#</th>
    <th>State</th>
    <th>User</th>
    <th>Title</th>
</tr>
{{range .Items}}
<tr>
    <td><a href='.HTMLURL'>{{.Number}}</a></td>
    <td>{{.State}}</td>
    <td><a href='.User.HTMLURL'>{{.User.Login}}</a></td>
    <td><a href='.HTMLRUL'>{{.Title | printf "%.64s"}}</a></td>
</tr>
{{end}}
</table>
`

var report = template.Must(template.New("issuelist").
    Parse(templ))

func main() {
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
    if err := report.Execute(os.Stdout, result); err != nil {
        log.Fatal(err)
    }
}

