package main

import (
	"fmt"
	"log"

	"github.com/g0tiu5a/ctftime"
)

func main() {
	url, err := ctftime.GetUrl("events", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[==>] Requesting %s ...\n", url)

	obj, err := ctftime.GetAPIData("events", nil)
	if err != nil {
		log.Fatal(err)
	}

	events, ok := obj.(ctftime.Events)
	if !ok {
		log.Fatal("events")
	}

	for idx, event := range events {
		fmt.Printf("[event%d]\n", idx)
		fmt.Printf("%#v\n", event)
	}
}
