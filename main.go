package main

import (
	"fmt"
	"log"

	"github.com/thingful/zenroom-go"
)

func main() {
	script := `print("holass")`
	res, err := zenroom.Exec(script, "", "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
