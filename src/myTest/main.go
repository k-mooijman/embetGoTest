package main

import (
	"log"

	"mooijman.info/myTest/embed"
	_ "mooijman.info/myTest/src/myTest/lib"
)

func main() {
	var testNr = 12
	log.Printf("\n\nThis is test %v of me trying Go\n", testNr)

	embed.Start()
}
