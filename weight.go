package gUtils

import (
	"errors"
	"fmt"
	"math/rand"
)

type Weight struct {
	Value  interface{}
	Weight int32
}

func CalcWeight(weights []*Weight) (*Weight, error) {
	var totalWeight int32
	for _, v := range weights {
		totalWeight += v.Weight
	}
	if totalWeight == 0 {
		return nil, errors.New("zero weight")
	}
	n := rand.Int31n(totalWeight)

	for _, v := range weights {
		if n < v.Weight {
			return v, nil
		}
		n -= v.Weight
	}

	return nil, errors.New("null")
}

func TestFormatWeight(n int) {
	weights := []*Weight{
		{Value: "1", Weight: 1},
		{Value: "2", Weight: 1},
		{Value: "3", Weight: 3},
		{Value: "4", Weight: 4},
	}

	m := make(map[interface{}]int)

	for i := 0; i < n; i++ {
		w, err := CalcWeight(weights)
		if err == nil {
			if v, ok := m[w.Value]; ok {
				m[w.Value] = v + 1
			} else {
				m[w.Value] = 1
			}
		}
	}
	fmt.Println("FormatWeight:", m)
}
