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
type gear struct{ nums map[int]int }

func (g *gear) add(n int) {
	g.nums[n]++
}

func (v val) num() bool {
	return unicode.IsNumber(rune(v))
}

func (v val) sym() bool {
	return v != 0 &&
		v != '.' &&
		!unicode.IsDigit(rune(v))
}

func (v val) gear() bool {
	return v == '*'
}

type schematic struct {
	lbX, lbY, ubX, ubY int
	rows               []row
	gears              map[string]*gear
	parts              []int
}

func newSchematic(rows []row) schematic {
	s := schematic{
		lbX:   0,
		lbY:   0,
		ubX:   len(rows[0]) - 1,
		ubY:   len(rows) - 1,
		rows:  rows,
		gears: map[string]*gear{},
		parts: []int{},
	}

	var part []rune
	var partGears []*gear
	var partSymAdj bool

	for yIdx, r := range s.rows {
		for xIdx, v := range r {
			if val(v).num() {
				part = append(part, v)
				if ok, g := s.symAdj(xIdx, yIdx); ok {
					partSymAdj = true
					partGears = append(partGears, g...)
				}
			}
			if len(part) != 0 {
				if !val(v).num() || xIdx == s.ubX {
					if partSymAdj {
						p, err := strconv.Atoi(string(part))
						if err != nil {
							panic("handle me")
						}

						s.parts = append(s.parts, p)
						for _, g := range partGears {
							g.add(p)
						}
					}
					part = nil
					partGears = nil
					partSymAdj = false
				}
			}
		}
	}

	return s
}

func (s *schematic) gear(x, y int) *gear {
	k := fmt.Sprintf("%v,%v", x, y)

	if s.at(x, y).gear() {
		g := s.gears[k]
		if g == nil {
			g = &gear{nums: map[int]int{}}
		}
		s.gears[k] = g
	}

	return s.gears[k]
}

func (s schematic) symAdj(x, y int) (adj bool, gears []*gear) {
	coords := [][]int{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
		{x + 1, y + 1},
		{x + 1, y - 1},
		{x - 1, y + 1},
		{x - 1, y - 1},
	}

	for _, c := range coords {
		if v := s.at(c[0], c[1]); v.sym() {
			adj = true
			if g := s.gear(c[0], c[1]); g != nil {
				gears = append(gears, g)
			}
		}
	}

	return adj, gears
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
	for _, p := range s.parts {
		sum += p
	}

	var gears int
	for _, g := range s.gears {
		if len(g.nums) == 2 {
			vals := []int{}
			for n := range g.nums {
				vals = append(vals, n)
			}
			gears += vals[0] * vals[1]
		}
	}

	return fmt.Sprintf("Sum of parts: %v; gears: %v", sum, gears), nil
}
