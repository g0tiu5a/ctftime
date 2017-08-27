package main

import (
	"fmt"

	"github.com/g0tiu5a/ctftime"
)

func GetTop10() {
	url := ctftime.GetUrl("top10", nil)
	fmt.Printf("[==>] Requesting %s ...\n", url)

	top10s := ctftime.GetAPIData("top10", nil)
	for key, top10 := range top10s.(ctftime.Top10) {
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
	url := ctftime.GetUrl("top10", ctx)
	fmt.Printf("[==>] Requesting %s ...\n", url)

	top10 := ctftime.GetAPIData("top10", ctx)
	for idx, team := range top10.(ctftime.Top10)["2017"] {
		fmt.Printf("	[%d] %#v\n", idx, team)
	}
}

func main() {
	fmt.Printf("[*] Trying All year's top10 api ...\n")
	GetTop10()

	fmt.Printf("[*] Trying 2017's top10 api ...\n")
	Get2017Top10()
}
