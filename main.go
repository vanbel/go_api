package main

import (
	"log"
	"test/api"
)

func main() {
	a := api.New(":9111")
	log.Fatal(a.Start())
}
