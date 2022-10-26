package main

import (
	"fmt"
	code "lem-in/general"
	"os"
)

func main() {
	example := os.Args[1:]
	Info := code.All{}

	if err := Info.IsValid(example); err != nil {
		fmt.Println("ERROR: invalid data format, ", err)
		return
	}

	graph, err := Info.ConstructGraph()
	if err != nil {
		fmt.Println("ERROR: invalid data format, ", err)
		return
	}

	if err = graph.Paths(&Info); err != nil {
		fmt.Println("ERROR: invalid data format, ", err)
		return
	}

	Info.WhatPaths()

	fmt.Println(Info.Map)

	if Info.StepsBfs <= Info.StepsDisjoint {
		fmt.Print(Info.Bfsres)
	} else {
		fmt.Print(Info.DisRes)
	}
}
