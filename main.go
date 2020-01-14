package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const MaxWindow = 144 // the window for the SMA
const Voters = 200000 // amount of real voters
const Target = 0.5    // the percentage of voters voting "yay"

var rng *rand.Rand

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Vote struct {
	Score uint64
	Yay   bool
}

// Winners simulates the grading process. Each voter generates a vote,
// which is assigned a random uint64 value. The results are then sorted,
// and the top 50 are returned.
func Winners(consensus float64) []Vote {
	votes := make([]Vote, Voters)
	yays := int(Voters * consensus)

	for i := range votes {
		votes[i].Score = rng.Uint64()
		votes[i].Yay = i < yays
	}

	sort.SliceStable(votes, func(i, j int) bool {
		return votes[i].Score > votes[j].Score
	})

	return votes[:50]
}

// Tally tallies up the yay votes inside of a result
func Tally(votes []Vote) int {
	var yes int
	for _, v := range votes {
		if v.Yay {
			yes++
		}
	}
	return yes
}

var results []float64

// Generate appends another voting result to the list.
// A result is the percentage of "yay" votes of the top 50.
func Generate(consensus float64) {
	w := Winners(consensus)
	yes := Tally(w)
	results = append(results, float64(yes)/50)
}

// GetAverage calculates the SMA for a given window size
func GetAverage(window int) float64 {
	percents := 0.0
	for i := 0; i < window; i++ {
		percents += results[len(results)-1-i]
	}
	return percents / float64(window)
}

func main() {
	// precalculate enough entries to have a complete window available
	for i := 0; i < MaxWindow; i++ {
		Generate(Target)
		if i%(MaxWindow/10) == 0 {
			fmt.Println(float64(i)/MaxWindow*100, "done")
		}
	}

	fmt.Printf("Target,Block")
	for i := 0; i <= MaxWindow; i += 6 {
		j := i
		if i == 0 {
			j = 1
		}
		fmt.Printf(",SMA-%d", j)
	}
	fmt.Println()
	for c := 0; c < 100; c++ {
		Generate(Target)
		fmt.Printf("%.1f,%.4f", Target, results[len(results)-1])
		for i := 0; i <= MaxWindow; i += 6 {
			j := i
			if j == 0 {
				j = 1
			}
			fmt.Printf(",%.4f", GetAverage(j))
		}
		fmt.Println()
	}
}
