package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
)

type State struct {
	Full_text string `json:"full_text"`
	Color     string `json:"color"`
	Name      string `json:"name"`
	MinWidth  int    `json:"min_width"`
}

func main() {
	flag.Parse()
	args := flag.Args()
	i := 0
	states := []State{}

	for i < len(args) {
		mw, _ := strconv.Atoi(args[i+3])
		s := State{
			Full_text: args[i],
			Name:      args[i+1],
			Color:     args[i+2],
			MinWidth:  mw,
		}
		states = append(states, s)
		i = i + 4
	}
	a, _ := json.Marshal(states)
	fmt.Printf("%s\n", a)
}
