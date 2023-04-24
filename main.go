package main

import (
	_ "embed"
	"opensourcegeeks-newsletter/listmonk"
	"opensourcegeeks-newsletter/loadenv"
	"opensourcegeeks-newsletter/rss"
)

func main() {

	loadenv.LoadEnv()
	rss.Rss()
	listmonk.Listmonk()

	// fmt.Println(os.Getenv("TEST"))

}
