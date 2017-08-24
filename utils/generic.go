package generic

import (
	// db "ace/database"
	"ace/models"
	_ "fmt"
)

func Save(t, l string, s float64) {
	leagueID := models.GetLeague(l).GetID()
	models.GetTeam(t, leagueID, s)
}
