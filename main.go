package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	// h1 := humanReal{
	// 	PhysicalBody: "ATGTATGC",
	// 	PranicBody:   []string{"?", "?", "?"},
	// 	MindBody:     "state1",
	// }
	// h2 := humanPerceived{
	// 	PhysicalBody: "TACG",
	// 	MindBody:     "state2",
	// }

	// bh1, err := json.Marshal(h1)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(string(bh1))

	// bh2, err := json.Marshal(h2)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(string(bh2))

	// var rebornh1 humanReal

	// err = json.Unmarshal(bh1, &rebornh1)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(rebornh1)

	//encode and decode

	http.HandleFunc("/", encoderhandler)
	http.HandleFunc("/decode", decodehandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func encoderhandler(w http.ResponseWriter, r *http.Request) {
	h2 := []humanReal{{
		PhysicalBody: "ATGCTAGCC",
		MindBody:     "state2",
		PranicBody:   []string{"pranavayu", "udana vayu"},
	},
		{PhysicalBody: "TAC", MindBody: "state2", PranicBody: []string{"sama vayu"}},
	}
	jencode := json.NewEncoder(w)
	jencode.Encode(h2)
}

func decodehandler(w http.ResponseWriter, r *http.Request) {
	var h2 []humanPerceived

	b := bytes.Buffer{}
	b.ReadFrom(r.Body)
	fmt.Println()
	json.NewDecoder(r.Body).Decode(h2)
	fmt.Println(h2)

}
