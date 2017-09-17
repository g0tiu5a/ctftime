package main

import (
	"fmt"
	"log"

	"github.com/g0tiu5a/ctftime"
)

func GetTop10() {
	url, err := ctftime.GetUrl("top10", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[==>] Requesting %s ...\n", url)

	obj, err := ctftime.GetAPIData("top10", nil)
	if err != nil {
		log.Fatal(err)
	}

	top10s, ok := obj.(ctftime.Top10s)
	if !ok {
		log.Fatal("top10")
	}

	for key, top10 := range top10s {
		fmt.Printf("[%s]", key)
		for idx, team := range top10 {
			fmt.Printf("	[%d] %#v\n", idx, team)
		}
		fmt.Printf("\n")
	}
}

func Get2017Top10() {
	ctx := ctftime.APIContext{
		"year": "2017",
	}
	url, err := ctftime.GetUrl("top10", ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[==>] Requesting %s ...\n", url)

	obj, err := ctftime.GetAPIData("top10", ctx)
	if err != nil {
		log.Fatal(err)
	}

	top10, ok := obj.(ctftime.Top10)
	if !ok {
		log.Fatal("top10")
	}

	for idx, team := range top10 {
		fmt.Printf("	[%d] %#v\n", idx, team)
	}
}

func main() {
	fmt.Printf("[*] Trying All year's top10 api ...\n")
	GetTop10()

	fmt.Printf("[*] Trying 2017's top10 api ...\n")
	Get2017Top10()
}
