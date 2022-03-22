package main

import (
	"github.com/mmcdole/gofeed"
	"time"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"fmt"
)

func main () {
	refresh_int, _ := strconv.Atoi(os.Getenv("REFRESH_INTERVAL"))
	delta_int, _ := strconv.Atoi(os.Getenv("CHECK_DELTA"))
	refresh_time := time.Duration(refresh_int)
	delta_time := time.Duration(delta_int)
	for {
		checktime := time.Now().Add(time.Minute * -(refresh_time + delta_time))
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL(os.Getenv("FEED_URL"))
		timeT, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", feed.Updated)
		fmt.Println("latest found: ",feed.Items[0].Title)

		if checktime.Before(timeT) {
			data := url.Values{
				"token": {os.Getenv("PUSHOVER_TOKEN")},
				"user": {os.Getenv("PUSHOVER_USER")},
				"message": {feed.Items[0].Title},
				"url": {feed.Items[0].Link},
			}
			_, _ = http.PostForm("https://api.pushover.net/1/messages.json", data)
		}
		time.Sleep(refresh_time * time.Minute)
	}
}