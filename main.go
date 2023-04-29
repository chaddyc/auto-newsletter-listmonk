package main

import (
	_ "embed"
	newsletter "opensourcegeeks-newsletter/internal"
)

func main() {

	newsletter.LoadEnv()
	newsletter.RSS()
	newsletter.Listmonk()

	// fmt.Println(os.Getenv("TEST"))

}
