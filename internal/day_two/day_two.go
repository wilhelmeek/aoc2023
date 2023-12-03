package daytwo

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type DayTwo struct{}

func (d DayTwo) Name() string {
	return "Day Two"
}

type set struct{ red, green, blue int }

func (s set) satisfiable(r, g, b int) bool {
	return s.red <= r && s.green <= g && s.blue <= b
}

func (s set) pow() int {
	return s.red * s.green * s.blue
}

type round []set

func (rd round) min() set {
	var r, g, b float64
	for _, s := range rd {
		r = math.Max(float64(s.red), r)
		g = math.Max(float64(s.green), g)
		b = math.Max(float64(s.blue), b)
	}
	return set{red: int(r), green: int(g), blue: int(b)}
}

func (rd round) satisfiable(r, g, b int) bool {
	for _, s := range rd {
		if !s.satisfiable(r, g, b) {
			return false
		}
	}
	return true
}

func (d DayTwo) Solve() error {
	f, err := os.Open("internal/day_two/input.txt")
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	rows := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		rows = append(rows, s.Text())
	}

	expr, err := regexp.Compile("Game (\\d*): (.*)")
	if err != nil {
		return fmt.Errorf("compiling regexp: %w", err)
	}

	var satSum, minSum int
	for _, r := range rows {
		rd := round{}

		sm := expr.FindStringSubmatch(r)
		if len(sm) != 3 {
			return fmt.Errorf("getting submatch: %w", err)
		}

		id, err := strconv.Atoi(sm[1])
		if err != nil {
			return fmt.Errorf("parsing id: %w", err)
		}

		for _, r := range strings.Split(sm[2], "; ") {
			s := set{}
			for _, t := range strings.Split(r, ", ") {
				tup := strings.Split(t, " ")
				if len(tup) != 2 {
					return fmt.Errorf("round must have count and color")
				}
				col := tup[1]
				ct, err := strconv.Atoi(tup[0])
				if err != nil {
					return fmt.Errorf("parsing to int: %w", err)
				}
				switch col {
				case "red":
					s.red += ct
				case "green":
					s.green += ct
				case "blue":
					s.blue += ct
				}
			}
			rd = append(rd, s)
		}

		minSum += rd.min().pow()
		if rd.satisfiable(12, 13, 14) {
			satSum += id
		}
	}

	fmt.Println(fmt.Sprintf("Sum of satisfiable sets: %v", satSum))
	fmt.Println(fmt.Sprintf("Sum of min-set round powers %v", minSum))

	return nil
}
