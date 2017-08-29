package main

import (
	"log"
)

func main() {

	if true {
		log.Print("qwerty")
	}

	goto catch

catch:
	log.Fatal("error has happend")
	return
}
