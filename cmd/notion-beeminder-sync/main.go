package main

import (
	"context"
	"fmt"
	"github.com/jonnyspicer/go-notion"
	"github.com/jonnyspicer/notion-beeminder-sync/pkg/beeminder"
	"github.com/jonnyspicer/notion-beeminder-sync/utils"
	"github.com/spf13/viper"
)

func main() {
	utils.LoadEnv()
	beeKey := viper.GetString("BEEMINDER_API_KEY")
	beeClient := beeminder.NewClient("https://www.beeminder.com/api/v1", beeKey)
	var diff int64 = 1721303299
	params := beeminder.GetUserParams{
		Username:     "jjspicer",
		DiffSince:    &diff,
		Associations: true,
	}
	user, err := beeClient.GetUser(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", len(user.Goals))
	fmt.Printf("%+v", user.Goals[0].Slug)

	// query notion DB
	notionKey := viper.GetString("NOTION_API_KEY")
	notionClient := notion.NewClient(notionKey)
	ctx := context.Background()
	db, err := notionClient.FindDatabaseByID(ctx, viper.GetString("NOTION_DATABASE_ID"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Printf("%+v", db)
	// retrieve each page
	// find data block
	// print data block
}
