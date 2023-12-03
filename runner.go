package main

import (
	"log"

	dayone "github.com/wilhelmeek/aoc2023/internal/day_one"
	"github.com/wilhelmeek/aoc2023/internal/solver"
)

func main() {
	solvers := []solver.Solver{
		dayone.DayOne{},
	}

	for _, s := range solvers {
		if err := s.Solve(); err != nil {
			log.Fatalf("Failed to solve %q: %s", s.Name(), err)
		}
	}
}
