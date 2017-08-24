package models

import (
	db "ace/database"
	_ "fmt"
	"strconv"
)

type (
	team struct {
		id       int
		name     string
		leagueID int
		stars    float64
	}
)

func GetTeam(name string, leagueID int, stars float64) team {
	return team{}.firstOrCreate(name, leagueID, stars)
}
func GetTeamByID(teamID int) (string, string, float64) {
	return findTeamByID(teamID)
}

func (t team) firstOrCreate(name string, leagueID int, stars float64) team {
	teamID := findTeam(name, leagueID, stars)
	if teamID == 0 {
		teamID = createTeam(name, leagueID, stars)
	}
	return team{id: teamID, name: name, leagueID: leagueID, stars: stars}
}

func findTeam(name string, leagueID int, stars float64) int {
	id := 0
	err := db.Mgr.QueryRow("SELECT id FROM team WHERE NAME = '" + name + "' AND league_id=" + string(leagueID)).Scan(&id)
	if err != nil {
		// fmt.Println(err.Error())
	}
	return id
}

func findTeamByID(id int) (string, string, float64) {
	var name, league string
	var stars float64
	tmp := strconv.Itoa(id)
	q := "SELECT teams.name as name, leagues.name as league, stars FROM teams INNER JOIN leagues on teams.league_id = leagues.id WHERE teams.id = " + tmp
	err := db.Mgr.QueryRow(q).Scan(&name, &league, &stars)
	if err != nil {
		// fmt.Println(err.Error())
	}
	return name, league, stars
}

func createTeam(name string, leagueID int, stars float64) int {
	s := strconv.FormatFloat(stars, 'g', -1, 64)
	l := strconv.FormatInt(int64(leagueID), 10)
	return db.Mgr.Insert("INSERT INTO teams (name, league_id, stars) VALUES(?,?,?)", []string{name, l, s})
}
