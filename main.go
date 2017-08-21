package main

import (
	_ "ace/database"
	"ace/scraper"
	_ "fmt"
)

func main() {
	scraper.Scrap()
}
