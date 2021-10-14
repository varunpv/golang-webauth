package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type humanReal struct {
	PhysicalBody string
	PranicBody   []string
	MindBody     string
}
type humanPerceived struct {
	PhysicalBody string
	MindBody     string `json:"mind"`
}

func main() {
	h1 := humanReal{
		PhysicalBody: "ATGTATGC",
		PranicBody:   []string{"?", "?", "?"},
		MindBody:     "state1",
	}
	h2 := humanPerceived{
		PhysicalBody: "TACG",
		MindBody:     "state2",
	}

	bh1, err := json.Marshal(h1)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(bh1))

	bh2, err := json.Marshal(h2)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(bh2))

	var rebornh1 humanReal

	err = json.Unmarshal(bh1, &rebornh1)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(rebornh1)

}
