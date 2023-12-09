package daythree

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Solution struct{}

func (sl Solution) Name() string {
	return "Day Three"
}

type row []rune
type val rune

func (v val) number() bool {
	return unicode.IsNumber(rune(v))
}

func (v val) symbol() bool {
	if v == 0 {
		return false
	}
	if v == '.' {
		return false
	}
	return !unicode.IsDigit(rune(v))
}

type schematic struct {
	lbX, lbY, ubX, ubY int
	rows               []row
}

func newSchematic(rows []row) schematic {
	return schematic{
		lbX:  0,
		lbY:  0,
		ubX:  len(rows[0]) - 1,
		ubY:  len(rows) - 1,
		rows: rows,
	}
}

func (s schematic) symbolAdj(x, y int) bool {
	return false ||
		s.at(x-1, y).symbol() ||
		s.at(x+1, y).symbol() ||
		s.at(x, y-1).symbol() ||
		s.at(x, y+1).symbol() ||
		s.at(x+1, y+1).symbol() ||
		s.at(x+1, y-1).symbol() ||
		s.at(x-1, y+1).symbol() ||
		s.at(x-1, y-1).symbol()
}

func (s schematic) parts() []string {
	var parts []string

	var part []rune
	var symAdj bool
	for yIdx, r := range s.rows {
		for xIdx, v := range r {
			if val(v).number() {
				part = append(part, v)
				if s.symbolAdj(xIdx, yIdx) {
					symAdj = true
				}
			}
			if len(part) != 0 {
				// Moving out of number
				if !val(v).number() || xIdx == s.ubX {
					if symAdj {
						parts = append(parts, string(part))
					}
					part = nil
					symAdj = false
				}
			}
		}
	}

	return parts
}

func (s schematic) at(x, y int) (v val) {
	if x > s.ubX {
		return
	}
	if x < s.lbX {
		return
	}
	if y > s.ubY {
		return
	}
	if y < s.lbY {
		return
	}
	return val(s.rows[y][x])
}

func (sl Solution) Solve() (string, error) {
	f, err := os.Open("internal/day_three/input.txt")
	if err != nil {
		return "", fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	rows := []row{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		rows = append(rows, row(sc.Text()))
	}

	s := newSchematic(rows)

	var sum int
	for _, p := range s.parts() {
		v, err := strconv.Atoi(p)
		if err != nil {
			return "", fmt.Errorf("converting part: %w", err)
		}
		sum += v
	}

	return fmt.Sprintf("Sum of parts: %v", sum), nil
}
