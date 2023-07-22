#!/bin/bash
#Adds test result to 'main.go' file
go test -bench=. the-go-programming-language/ch2/ex2.5 | awk 'PREPEND="// " {print PREPEND $0}' >> main.go
