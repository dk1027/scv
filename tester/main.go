package main

import (
	"github.com/dk1027/scv/shared"
	"log"
)

func main() {
	fn := shared.TriggerLoader("../watchdir/watchdir.so")
	fn(nil, "hello")
	log.Println("finished")
}
