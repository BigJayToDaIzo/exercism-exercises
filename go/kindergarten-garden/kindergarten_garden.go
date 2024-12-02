package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden struct {
	r1, r2     string
	children   []string
	seedLookup map[byte]string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	dupeCheck := make(map[string]bool)
	for _, child := range children {
		if dupeCheck[child] {
			return nil, errors.New("duplicate child name")
		}
		dupeCheck[child] = true
	}
	rows := strings.Split(diagram, "\n")
	if len(rows) != 3 {
		return nil, errors.New("diagram must have 3 rows")
	}
	if len(rows[1]) != len(rows[2]) {
		return nil, errors.New("rows must have the same length")
	}
	if len(rows[1])%2 != 0 {
		return nil, errors.New("rows must have an even number of plants")
	}
	orderedChildren := make([]string, len(children))
	copy(orderedChildren, children)
	sort.Strings(orderedChildren)
	seedLookup := map[byte]string{
		'G': "grass",
		'C': "clover",
		'R': "radishes",
		'V': "violets",
	}
	for i := 1; i < 3; i++ {
		for _, cup := range rows[i] {
			if _, ok := seedLookup[byte(cup)]; !ok {
				return nil, errors.New("invalid plant")
			}
		}
	}

	return &Garden{rows[1], rows[2], orderedChildren, seedLookup}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	childIndex := -1
	for i, childName := range g.children {
		if childName == child {
			childIndex = i
			break
		}
	}
	if childIndex == -1 {
		return nil, false
	}
	plantsIndex := childIndex * 2
	return []string{
		g.seedLookup[g.r1[plantsIndex]],
		g.seedLookup[g.r1[plantsIndex+1]],
		g.seedLookup[g.r2[plantsIndex]],
		g.seedLookup[g.r2[plantsIndex+1]],
	}, true
}
