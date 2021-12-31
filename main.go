package main

import (
	"fmt"

	"github.com/stackfull/days-since/src"
)

// title contains the name of the project
const title = "days-since"

/*
ProjectName returns the value of `title` string
*/
func ProjectName() string {
	return title
}

func main() {
	fmt.Printf("Running project: %s\n", src.ProjectName())
}
