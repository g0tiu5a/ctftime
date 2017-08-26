package main

import (
	"fmt"

	"github.com/g0tiu5a/ctftime"
)

func main() {
	url := ctftime.GetUrl("events", nil)
	fmt.Printf("[==>] Requesting %s ...\n", url)
	events := ctftime.GetAPIData("events", nil)
	for idx, event := range events.([]ctftime.Event) {
		fmt.Printf("[event%d]\n", idx)
		fmt.Printf("%#v\n", event)
	}
}
