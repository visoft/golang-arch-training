package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	First string
}

func main() {
	// p1 := Person{
	// 	First: "Damien",
	// }
	// p2 := Person{
	// 	First: "Dave",
	// }
	// xp := []Person{p1, p2}
	// bs, err := json.Marshal(xp)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println("Print JSON", string(bs))

	// xp2 := []Person{}
	// err = json.Unmarshal(bs, &xp2)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("Back into Go data structure", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := Person{
		First: "Damien",
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {

}
