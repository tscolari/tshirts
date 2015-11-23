package main

import (
	"fmt"

	"github.com/tscolari/tshirts/challenge"
	"github.com/tscolari/tshirts/client"
)

func main() {
	client := client.New("http://challenge.teespring.com", "1ab087b1-0518-403c-a328-d6350db4d94f")
	question, err := client.FetchQuestion()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Q -> %#v\n", question)

	inks, err := client.FetchInks()
	if err != nil {
		panic(err)
	}

	solver := challenge.NewSolver(inks, question)
	solution := solver.Solve()

	fmt.Printf("S -> %#v\n", solution)

	resp, err := client.PostAnswer(solution)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response to solution:")
	fmt.Println(resp)
}
