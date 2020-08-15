package matchers

import (
	"fmt"
	"strings"
)

const (
	MatcherOpRetain    = 0x1
	MatcherOpTransform = 0x2
	MatcherOpStrip     = 0x4
	MatcherOpNormalize = 0x8
	MatcherOpRemove    = 0x10
)

type MatcherOperation int

type Matcher struct {
	MatcherString      string
	MatcherOp          MatcherOperation
	NormalizeCityState bool
	NormalizeNumbers   bool
}

/*
AUTOPAY \digit{10}RAUTOPAY AUTO-PMT
CHIPOTLE \digit{4} \city \stateshort
COSTCO WHSE #\digit{4} \city \stateshort
*/
func FindMatchers(description string) ([]Matcher, error) {
	retainMatcher := Matcher{}

	var b strings.Builder

	inDigit := false
	digitsLen := 0

	//words := strings.Split(description, " ")

	for _, r := range description {
		if isDigit(r) {
			digitsLen++
			if !inDigit {
				inDigit = true
			}
		} else {
			if inDigit {
				b.WriteString(fmt.Sprintf("\\digit{%v}", digitsLen))

				inDigit = false
				digitsLen = 0
			}

			_, err := b.WriteRune(r)
			if err != nil {
				return nil, err
			}
		}
	}

	fmt.Printf(b.String())

	return []Matcher{
		retainMatcher,
	}, nil
}

func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}

	return false
}

func main() {
	fmt.Println("vim-go")
}
