package main

import (
	"testing"
)

func TestGameScore(t *testing.T) {
	cases := []struct {
		game   string
		score  int
		powers int
	}{
		{
			game:   "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			score:  1,
			powers: 48,
		},
		{
			game:   "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			score:  2,
			powers: 12,
		},
		{
			game:   "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			score:  0,
			powers: 1560,
		},
		{
			game:   "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			score:  0,
			powers: 630,
		},
		{
			game:   "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			score:  5,
			powers: 36,
		},
	}

	sum := 0
	powers := 0
	for _, c := range cases {
		res, p := GameScore(c.game, map[string]int{
			"red":   12,
			"green": 13,
			"blue":  14,
		})

		sum += res
		powers += p
		if res != c.score {
			t.Errorf("Res Error, got: %d, want: %d", res, c.score)
		}
		if p != c.powers {
			t.Errorf("Res Error, got: %d, want: %d", p, c.powers)
		}
	}

	if sum != 8 {
		t.Errorf("Sum Error, got: %d, want 8", sum)
	}
}
