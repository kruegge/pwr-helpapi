package malo

import (
	"fmt"
	"math/rand"
	"time"
)

func Create() string {
	min, max := 100000000, 999999999
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	r2 := rand.New(s1)

	part1 := r1.Intn(max-min) + min
	identifier := fmt.Sprintf("%d%d", part1, r2.Intn(9))

	return fmt.Sprintf("%s%d", identifier, generateCheckDigit(identifier))
}

func generateCheckDigit(identifier string) int {
	weightSum := 0
	weight := 1
	for _, v := range identifier {
		weightSum += weight * int(v)
		weight = weight%2 + 1
	}

	return (10 - weightSum%10) % 10
}
