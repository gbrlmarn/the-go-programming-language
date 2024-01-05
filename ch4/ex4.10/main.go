// Exercise 4.10: Modify issues to report the results in age categories, say less than a month old, less than a year old, and more than a year old.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl/ch4/examples/github"
)

const (
	day   = 24         // a day in hours
	month = day * 31   // a month in hours
	year  = month * 12 // a year in hours
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var lmonth, lyear, myear []*github.Issue
	for _, item := range result.Items {
		if age(item.CreateAt) < month {
			lmonth = append(lmonth, item)
		} else if age(item.CreateAt) < year {
			lyear = append(lyear, item)
		} else {
			myear = append(myear, item)
		}
	}

	fmt.Printf("\n%d issues less than a month old:\n", len(lmonth))
    printIssues(lmonth)

	fmt.Printf("\n%d issues less than a year old:\n", len(lyear))
    printIssues(lyear)

	fmt.Printf("\n%d issues more than a year old:\n", len(myear))
    printIssues(myear)
}

func age(t time.Time) int {
	return int(time.Since(t).Hours())
}

func printIssues(issues []*github.Issue) {
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
