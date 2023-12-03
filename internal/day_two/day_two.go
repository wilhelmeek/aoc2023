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

type Solution struct{}

func (d Solution) Name() string {
	return "Day Two"
}

type set struct{ r, g, b int }

func (s set) sat(r, g, b int) bool {
	return s.r <= r && s.g <= g && s.b <= b
}

func (s set) pow() int {
	return s.r * s.g * s.b
}

type round []set

func (rd round) min() set {
	var r, g, b float64
	for _, s := range rd {
		r = math.Max(float64(s.r), r)
		g = math.Max(float64(s.g), g)
		b = math.Max(float64(s.b), b)
	}
	return set{r: int(r), g: int(g), b: int(b)}
}

func (rd round) sat(r, g, b int) bool {
	for _, s := range rd {
		if !s.sat(r, g, b) {
			return false
		}
	}
	return true
}

func (d Solution) Solve() (string, error) {
	f, err := os.Open("internal/day_two/input.txt")
	if err != nil {
		return "", fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	rows := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		rows = append(rows, s.Text())
	}

	expr, err := regexp.Compile("Game (\\d*): (.*)")
	if err != nil {
		return "", fmt.Errorf("compiling regexp: %w", err)
	}

	var satSum, minSum int
	for _, r := range rows {
		rd := round{}

		sm := expr.FindStringSubmatch(r)
		if len(sm) != 3 {
			return "", fmt.Errorf("getting submatch: %w", err)
		}

		id, err := strconv.Atoi(sm[1])
		if err != nil {
			return "", fmt.Errorf("parsing id: %w", err)
		}

		for _, r := range strings.Split(sm[2], "; ") {
			var s set
			for _, t := range strings.Split(r, ", ") {
				tup := strings.Split(t, " ")
				if len(tup) != 2 {
					return "", fmt.Errorf("round must have count and color")
				}
				col := tup[1]
				ct, err := strconv.Atoi(tup[0])
				if err != nil {
					return "", fmt.Errorf("parsing to int: %w", err)
				}
				switch col {
				case "red":
					s.r += ct
				case "green":
					s.g += ct
				case "blue":
					s.b += ct
				}
			}
			rd = append(rd, s)
		}

		minSum += rd.min().pow()
		if rd.sat(12, 13, 14) {
			satSum += id
		}
	}

	return fmt.Sprintf("Part 1: %d\nPart 2: %d", satSum, minSum), nil
}
