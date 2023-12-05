package daythree

import (
	"bufio"
	"fmt"
	"os"
)

type Solution struct{}

func (sl Solution) Name() string {
	return "Day Three"
}

type val string

func (v val) number() bool {
	return true
}

type schematic struct {
	ubX, ubY int
	rows     []val
}

func newSchematic(rows []string) schematic {
	return schematic{
		ubX:  len(rows[0]) - 1,
		ubY:  len(rows) - 1,
		rows: rows,
	}
}

func (s schematic) neighbours(x, y int) (int, error) {
	v, err := s.at(x, y)
	if err != nil {
		return 0, fmt.Errorf("getting pos: %w", err)
	}

	return 0, nil
}

func (s schematic) at(x, y int) (string, error) {
	if x > s.ubX {
		return "", fmt.Errorf(fmt.Sprintf("upper x is %q", s.ubX))
	}
	if y > s.ubY {
		return "", fmt.Errorf(fmt.Sprintf("upper y is %q", s.ubY))
	}
	return string(s.rows[y][x]), nil
}

func (sl Solution) Solve() (string, error) {
	f, err := os.Open("internal/day_two/input.txt")
	if err != nil {
		return "", fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	rows := []string{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}

	newSchematic(rows)

	return "", nil
}
