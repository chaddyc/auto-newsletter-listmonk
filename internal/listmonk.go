package newsletter

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mmcdole/gofeed"
)

func Listmonk() {

	// env vars
	var api = os.Getenv("LISTMONK_API")
	var username = os.Getenv("USER")
	var password = os.Getenv("PASSWORD")
	var rss = os.Getenv("RSS_FEED")
	var email = os.Getenv("FROM_EMAIL")
	var t = os.Getenv("TEMPLATE_ID")
	template, err := strconv.Atoi(t)
	var l = os.Getenv("LISTS")
	list, err := strconv.Atoi(l)

	time.Sleep(30 * time.Second)

	fc, _ := fs.ReadFile(os.DirFS("."), "newsletter.html")
	newsletterName := "newsletter-" + time.Now().Format("01-02-2006")

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(rss)
	newsletterSubject := feed.Items[0].Title

	type Payload struct {
		Name        string `json:"name"`
		Subject     string `json:"subject"`
		Lists       []int  `json:"lists"`
		FromEmail   string `json:"from_email"`
		ContentType string `json:"content_type"`
		Body        string `json:"body"`
		Messenger   string `json:"messenger"`
		Type        string `json:"type"`
		TemplateID  int    `json:"template_id"`
	}

	data := Payload{
		Name:        newsletterName,
		Subject:     newsletterSubject,
		Lists:       []int{list},
		FromEmail:   email,
		ContentType: "html",
		Body:        string(fc),
		Messenger:   "email",
		Type:        "regular",
		TemplateID:  template,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error occured. Err: %s", err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", api, body)
	if err != nil {
		log.Fatalf("Error occured. Err: %s", err)
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error occured. Err: %s", err)
	}
	defer resp.Body.Close()

}
