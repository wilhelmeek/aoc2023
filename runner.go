package main

import (
	"fmt"
	"log"

	dayone "github.com/wilhelmeek/aoc2023/internal/day_one"
	daytwo "github.com/wilhelmeek/aoc2023/internal/day_one/day_two"
	"github.com/wilhelmeek/aoc2023/internal/solver"
)

func main() {
	solvers := []solver.Solver{
		dayone.DayOne{},
		daytwo.DayTwo{},
	}

	for _, s := range solvers {
		fmt.Println(fmt.Sprintf("Solving %q", s.Name()))
		if err := s.Solve(); err != nil {
			log.Fatalf("Failed to solve %q: %s", s.Name(), err)
		}
	}
}
