package challenge

import (
	"math"

	"github.com/tscolari/tshirts/colors"
)

type Solver struct {
	inks     Inks
	scenario Scenario
}

func NewSolver(inks Inks, scenario Scenario) *Solver {
	return &Solver{
		inks:     inks,
		scenario: scenario,
	}
}

func (s *Solver) Solve() Solution {
	answers := []Answer{}

	for _, question := range s.scenario.Questions {
		answers = append(answers, s.solveQuestion(question))
	}

	return Solution{
		ScenarioID: s.scenario.ID,
		Answers:    answers,
	}
}

func (s *Solver) solveQuestion(question Question) (answer Answer) {
	inks := []string{}

	for _, layer := range question.Layers {
		var bestInk Ink
		smallerDist := math.MaxFloat64

		for _, ink := range s.inks {
			dist := colors.CalcDistanceHex(layer.Color, ink.Color)

			if ink.Cost < bestInk.Cost && dist < 20 {
				bestInk = ink
			} else if dist < smallerDist {
				smallerDist = dist
				bestInk = ink
			}
		}

		inks = append(inks, bestInk.ID)
	}

	return Answer{
		Inks: inks,
	}
}
