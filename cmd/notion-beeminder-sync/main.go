package main

import (
	"fmt"

	"github.com/jonnyspicer/notion-beeminder-sync/pkg/beeminder"
)

func main() {
	// fmt.Println("Hello, World!")
	// create new instance of beeminder client
	// have client retrieve user info
	// print user info
	key := ""
	client := beeminder.NewClient("https://www.beeminder.com/api/v1", key)
	var diff int64 = 1721303299
	params := beeminder.GetUserParams{
		Username:     "jjspicer",
		DiffSince:    &diff,
		Associations: true,
	}
	res, err := client.GetUser(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", res)

	// query notion DB
	// retrieve each page
	// find data block
	// print data block
}
