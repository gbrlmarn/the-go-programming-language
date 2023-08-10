// Exercise 4.11: Build a tool that lets create, read, update, and close GitHub issues from the command line, invoking their preferred text editor when substantial text input is required. 
package main

import (
    "fmt"
    "log"
)

func main() {

}

func search(name []string) {
	res, err := SearchIssues(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", res.TotalCount)
	for _, item := range res.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

