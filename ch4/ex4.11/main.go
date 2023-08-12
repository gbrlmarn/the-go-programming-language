// Exercise 4.11: Build a tool that lets create, read, update, and close GitHub issues from the command line, invoking their preferred text editor when substantial text input is required.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"the-go-programming-language/ch4/examples/github"
)

const (
	GithubAPI = "https://api.github.com/repos"
    usage     = `usage: 
    search repo:<reponame> is:<open/close> <branch> <description>
    [read] owner repo issuenumber`
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
    cmd := os.Args[1]
    args := os.Args[2:]
    if cmd == "search" {
	    search(os.Args[1:])
        os.Exit(0)
    }

    owner, repo, number := args[0], args[1], args[2]
    switch cmd {
    case "read":
        read(owner, repo, number)
    }
}

func search(name []string) {
	res, err := github.SearchIssues(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", res.TotalCount)
	for _, item := range res.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func read(owner, repo, number string) {
    issue, err := GetIssue(owner, repo, number)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Printf("#%-5d %9.9s %.55s\n%s\n",
	    issue.Number, issue.User.Login, issue.Title, issue.Body)
}

func GetIssue(owner, repo, number string) (*github.Issue, error) {
	url := strings.Join([]string{
		GithubAPI, owner,
		repo, "issues", number}, "/")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issue github.Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

