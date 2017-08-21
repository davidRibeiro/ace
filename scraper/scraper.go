package scraper

import (
	_ "ace/database"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strconv"
	"strings"
)

const URLBase = "https://www.fifaindex.com/pt/teams/"

func Scrap() {
	index := 1
	for {
		url := URLBase + strconv.Itoa(index)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if resp.StatusCode == http.StatusNotFound {
			return
		}
		scrapPage(resp)
		index++

	}
}

func scrapPage(resp *http.Response) {
	b := resp.Body
	defer b.Close()
	doc := html.NewTokenizer(b)
	tbodyHasStarted := false
	for tokenType := doc.Next(); tokenType != html.ErrorToken; tokenType = doc.Next() {
		token := doc.Token()
		if token.Data == "tbody" && token.Type == html.StartTagToken {
			tbodyHasStarted = true
		}
		if token.Data == "tbody" && token.Type == html.EndTagToken {
			tbodyHasStarted = false
			return
		}
		if token.Data == "tr" && token.Type == html.StartTagToken && tbodyHasStarted {
			doc = scrapRow(doc)
		}
	}
}

func scrapRow(doc *html.Tokenizer) *html.Tokenizer {
	team, league := "", ""
	stars := 0.0
	for tokenType := doc.Next(); tokenType != html.ErrorToken; tokenType = doc.Next() {
		token := doc.Token()
		if token.Data == "tr" && token.Type == html.EndTagToken {
			fmt.Println(team, league, stars)
			return doc
		}
		if token.Type == html.TextToken {
			if team == "" {
				team = token.String()
			} else if league == "" {
				league = token.String()
			}
		}
		if token.Data == "i" && token.Type == html.StartTagToken {
			for _, i := range token.Attr {
				if strings.Contains(i.Val, "fa-star-o") {
					continue
				}
				if strings.Contains(i.Val, "fa-star-half-o") {
					stars += 0.5
				} else if strings.Contains(i.Val, "fa-star") {
					stars += 1.0
				}
			}
		}
	}
	return doc
}
