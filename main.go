package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/deeper-x/gcyph/ctable"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	log.Println(ctable.New())
}
