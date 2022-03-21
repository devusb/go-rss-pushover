package main

import (
	"github.com/mmcdole/gofeed"
	"time"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main () {
	refresh_int, _ := strconv.Atoi(os.Getenv("REFRESH_INTERVAL"))
	refresh_time := time.Duration(refresh_int)
	for {
		checktime := time.Now().Add(time.Minute * -(refresh_time + 1))
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
		time.Sleep(refresh_time * time.Minute)
	}
}