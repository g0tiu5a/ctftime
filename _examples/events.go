package main

import (
	"fmt"
	"log"

	"github.com/g0tiu5a/ctftime"
)

func GetEvents() {
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

func GetSpecifiedEvent() {
	ctx := ctftime.APIContext{
		"event_id": 1,
	}

	url, err := ctftime.GetUrl("events", ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[==>] Requesting %s ...\n", url)

	obj, err := ctftime.GetAPIData("events", ctx)
	if err != nil {
		log.Fatal(err)
	}

	event, ok := obj.(ctftime.Event)
	if !ok {
		log.Fatal("event")
	}

	fmt.Printf("%#v\n", event)
}

func main() {
	fmt.Println("[*] Trying All events api ...")
	GetEvents()
	fmt.Println("[*] Trying Specified event api ...")
	GetSpecifiedEvent()
}
