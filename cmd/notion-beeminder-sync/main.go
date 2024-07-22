package main

import (
	"context"
	"fmt"

	"github.com/jonnyspicer/go-notion"
	"github.com/jonnyspicer/notion-beeminder-sync/pkg/beeminder"
	nhelper "github.com/jonnyspicer/notion-beeminder-sync/pkg/notion"
	"github.com/jonnyspicer/notion-beeminder-sync/pkg/unified"
	"github.com/jonnyspicer/notion-beeminder-sync/utils"
	"github.com/spf13/viper"
)

func beeApiToIntermediary(goal beeminder.Goal, user string) unified.Goal {
	// convert beeminder goal to intermediary goal
	return unified.Goal{
		Goal:      goal.Slug,
		Deadline:  goal.Losedate,
		Delta:     goal.Delta,
		GoalUnits: goal.Gunits,
		SafeDays:  goal.SafeBuf,
		Pledge:    goal.Pledge,
		Link:      fmt.Sprintf("https://www.beeminder.com/%s/%s", user, goal.Slug),
		Autodata:  goal.Autodata,
		GoalType:  unified.GoalType(goal.GoalType),
		RateUnit:  unified.RateUnit(goal.Runits),
	}
}

func notionApiToIntermediary(goal notion.DatabasePageProperties) unified.Goal {
	// convert notion goal to intermediary goal
	return unified.Goal{
		Goal:      goal["Goal"].Value().(string),
		Deadline:  goal["Deadline"].Value().(int64), // this is wrong, it's going to be a date coming out of notion, needs converting to epoch time
		Delta:     goal["Delta"].Value().(float64),
		GoalUnits: goal["Goal Units"].Value().(string),
		SafeDays:  goal["Safe Days"].Value().(int64),
		Pledge:    goal["Pledge"].Value().(float64),
		Link:      goal["Link"].Value().(string),
		Autodata:  goal["Autodata"].Value().(string),
		GoalType:  unified.GoalType(goal["Goal Type"].Value().(string)),
		RateUnit:  unified.RateUnit(goal["Rate Unit"].Value().(string)),
	}
}

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

	intGoals := []unified.Goal{}
	for _, goal := range user.Goals {
		intGoals = append(intGoals, beeApiToIntermediary(goal, user.Username))
	}

	for _, goal := range intGoals {
		fmt.Println(goal.Goal)
	}

	// query notion DB
	notionKey := viper.GetString("NOTION_API_KEY")
	notionClient := notion.NewClient(notionKey)
	ctx := context.Background()
	databaseId := viper.GetString("NOTION_DATABASE_ID")
	databaseQuery := notion.DatabaseQuery{}
	db, err := notionClient.QueryDatabase(ctx, databaseId, &databaseQuery)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println()

	goalsInNotion, err := nhelper.ExtractPageProperties(db)
	for _, property := range goalsInNotion {
		for k := range property {
			fmt.Println(k)
		}
	}

	// retrieve each page
	// check if all slugs have matches
	// if not, create new page
	// find data block
	// print data block

}
