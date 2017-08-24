package main

import (
	db "ace/database"
	"ace/models"
	"ace/scraper"
	"fmt"
)

func main() {

	scraper.Scrap()
	ids := db.Mgr.BuildLeagues()
	for index := 0; index < len(ids); index++ {
		name, league, stars := models.GetTeamByID(ids[index])
		fmt.Println(name, league, stars)
		if index+1 == 20 || index+1 == 20+24 || index+1 == 20+24+24 || index+1 == 20+24+24+24 {
			fmt.Println("-----------")
		}
	}
}
