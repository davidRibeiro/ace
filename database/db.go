package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

type singleton struct {
	driver   string
	file     string
	database *sql.DB
}

var instance *singleton
var once sync.Once
var migrations = []string{
	"CREATE TABLE IF NOT EXISTS `teams` (`id`	INTEGER NOT NULL UNIQUE, `name`	TEXT NOT NULL, `league_id`	INTEGER NOT NULL, `stars` NUMERIC NOT NULL, PRIMARY KEY(`id`))",
	"CREATE TABLE IF NOT EXISTS `leagues` (`id` INTEGER NOT NULL UNIQUE, `name` TEXT NOT NULL, PRIMARY KEY(`id`))"}

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{driver: "sqlite3", file: "cache/ace.sqlite3"}
		instance.init()
	})
	return instance
}

func (single singleton) init() (bool, error) {

	db, err := sql.Open(single.driver, single.file)
	single.database = db
	if err != nil {
		return false, err
	}

	defer single.database.Close()
	for _, query := range migrations {
		err = single.runMigration(query)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (single singleton) runMigration(query string) error {
	_, err := single.database.Exec(query)
	return err
}
