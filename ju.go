package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

type State struct {
	Full_text string `json:"full_text"`
	Color     string `json:"color"`
	Name      string `json:"name"`
}

func main() {
	flag.Parse()
	args := flag.Args()
	i := 0
	states := []State{}

	for i < len(args) {
		s := State{
			Full_text: args[i],
			Name:      args[i+1],
			Color:     args[i+2],
		}
		states = append(states, s)
		i = i + 3
	}
	a, _ := json.Marshal(states)
	fmt.Printf("%s\n", a)
}
