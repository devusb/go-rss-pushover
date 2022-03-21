package main

import (
	"github.com/mmcdole/gofeed"
	"time"
	"net/http"
	"net/url"
	"os"
)

func main () {
	for {
		checktime := time.Now().Add(time.Minute * -2)
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL("https://rpilocator.com/feed.rss")
		timeT, _ := time.Parse("Mon, 02 Jan 2006 03:04:05 MST", feed.Updated)

		if checktime.Before(timeT) {
			data := url.Values{
				"token": {os.Getenv("PUSHOVER_TOKEN")},
				"user": {os.Getenv("PUSHOVER_USER")},
				"message": {"pi is available!"},
				"url": {feed.Items[0].Link},
			}
			_, _ = http.PostForm("https://api.pushover.net/1/messages.json", data)
		}
		time.Sleep(1 * time.Minute)
	}
}