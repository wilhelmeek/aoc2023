package dayone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var nums = []string{
	"🌭",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

type Solution struct{}

func (d Solution) Name() string {
	return "Day One"
}

func (d Solution) Solve() (string, error) {
	f, err := os.Open("internal/day_one/input.txt")
	if err != nil {
		return "", fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	words := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text())
	}

	var sum int
	for _, w := range words {
		eIdx := -1
		lIdx := -1
		var e, l int
		for n, nStr := range nums {
			if o := strings.Index(w, nStr); o != -1 {
				if o <= eIdx || eIdx == -1 {
					eIdx = o
					e = n
				}
			}
			if o := strings.LastIndex(w, nStr); o != -1 {
				if o > lIdx {
					lIdx = o
					l = n
				}
			}
		}

		if eIdx != -1 {
			// Adjust for overlapped numbers (1 rune)
			olIdx := eIdx + len(nums[e]) - 1
			if olIdx == lIdx {
				w = w[:olIdx] + string(w[olIdx]) + w[olIdx:]
			}

			w = strings.Replace(w, nums[e], strconv.Itoa(e), 1)
			w = strings.ReplaceAll(w, nums[l], strconv.Itoa(l))
		}

		var fr, lr rune
		for _, r := range w {
			if !unicode.IsDigit(r) {
				continue
			}
			lr = r
			if fr == 0 {
				fr = r
			}
		}

		tup, err := strconv.Atoi(string(fr) + string(lr))
		if err != nil {
			return "", fmt.Errorf("opening file: %w", err)
		}

		sum += tup
	}

	return strconv.Itoa(sum), nil
}
