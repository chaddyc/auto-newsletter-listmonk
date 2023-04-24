package listmonk

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

func Listmonk() {

	time.Sleep(30 * time.Second)

	fc, _ := fs.ReadFile(os.DirFS("."), "newsletter.html")
	newsletterName := "newsletter-" + time.Now().Format("01-02-2006")

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://opensourcegeeks.net/rss/")
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
		Lists:       []int{3},
		FromEmail:   "Chad at Opensource Geeks <chad@opensourcegeeks.net>",
		ContentType: "html",
		Body:        string(fc),
		Messenger:   "email",
		Type:        "regular",
		TemplateID:  3,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error occured. Err: %s", err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://newsletter.opensourcegeeks.net/api/campaigns", body)
	if err != nil {
		log.Fatalf("Error occured. Err: %s", err)
	}

	var username = os.Getenv("USER")
	var password = os.Getenv("PASSWORD")

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error occured. Err: %s", err)
	}
	defer resp.Body.Close()

	// fmt.Println(os.Getenv("TEST"))
}
