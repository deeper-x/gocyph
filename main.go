package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/deeper-x/gocyph/gcipher"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	log.Println(gcipher.NewTable())
}
