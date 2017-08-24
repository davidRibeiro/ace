package models

import (
	db "ace/database"
	_ "fmt"
)

type (
	league struct {
		id   int
		name string
	}
)

func GetLeague(name string) league {
	return league{}.firstOrCreate(name)
}

func (l league) GetID() int {
	return l.id
}

func (l league) GetName() string {
	return l.name
}

func (l league) firstOrCreate(name string) league {
	leagueID := findLeague(name)
	if leagueID == 0 {
		leagueID = createLeague(name)
	}
	// fmt.Println(league{id: leagueID, name: name})
	return league{id: leagueID, name: name}
}

func findLeague(name string) int {
	id := 0
	err := db.Mgr.QueryRow("SELECT id FROM leagues WHERE NAME = '" + name + "'").Scan(&id)
	if err != nil {
		// fmt.Println(err.Error())
	}
	return id
}

func createLeague(name string) int {
	return db.Mgr.Insert("INSERT INTO leagues (name) VALUES(?)", []string{name})
}
