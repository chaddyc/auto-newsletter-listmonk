package newsletter

import (
	"log"
	"os"

	"github.com/mmcdole/gofeed"
)

func RSS() {
	var rss = os.Getenv("RSS_FEED")
	var header = os.Getenv("HEADER_MESSAGE")
	var footer = os.Getenv("FOOTER_MESSAGE")

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(rss)

	f, err := os.Create("./newsletter.html")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.WriteString("<h2>" + header + "</h2><br> ")

	for i := 0; i < 5; i++ {
		f.WriteString("<p><a href='" + feed.Items[i].Link + "'>" + feed.Items[i].Title + "</a> - " + feed.Items[i].Description + "</p> \n")

	}

	// f.WriteString("Thanks for being part of the Opensource Geeks community💻🐧")
	f.WriteString(footer)
}
