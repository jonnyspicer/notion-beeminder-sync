package notion

import (
	"encoding/json"
	"fmt"
	"github.com/jonnyspicer/go-notion"
)

func ExtractPageProperties(res notion.DatabaseQueryResponse) ([]notion.DatabasePageProperties, error) {

	err := error(nil)
	goalsInNotion := []notion.DatabasePageProperties{}
	for _, page := range res.Results {
		propsJSON, err := json.Marshal(page.Properties)
		if err != nil {
			fmt.Println(err)
		}
		var dbProps notion.DatabasePageProperties
		if err := json.Unmarshal(propsJSON, &dbProps); err != nil {
			fmt.Println(err)
		}
		goalsInNotion = append(goalsInNotion, dbProps)
	}
	return goalsInNotion, err
}
