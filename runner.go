package main

import (
	"fmt"
	"log"

	dayone "github.com/wilhelmeek/aoc2023/internal/day_one"
	daythree "github.com/wilhelmeek/aoc2023/internal/day_three"
	daytwo "github.com/wilhelmeek/aoc2023/internal/day_two"
	"github.com/wilhelmeek/aoc2023/internal/solver"
)

func main() {
	solvers := []solver.Solver{
		dayone.Solution{},
		daytwo.Solution{},
		daythree.Solution{},
	}

	for _, s := range solvers {
		out, err := s.Solve()
		if err != nil {
			log.Fatalf("Failed to solve %q: %s", s.Name(), err)
		}

		fmt.Println(fmt.Sprintf("\nAnswer for %q:\n\n%s\n\n===", s.Name(), out))
	}
}
