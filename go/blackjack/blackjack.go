package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "king" || card == "queen" || card == "jack" || card == "ten":
		return 10
	case card == "nine":
		return 9
	case card == "eight":
		return 8
	case card == "seven":
		return 7
	case card == "six":
		return 6
	case card == "five":
		return 5
	case card == "four":
		return 4
	case card == "three":
		return 3
	case card == "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardSum := ParseCard(card1) + ParseCard(card2)
	dcFloat := ParseCard(dealerCard)
	if cardSum == 22 {
		return "P"
	} else if cardSum == 21 {
		if dcFloat < 10 {
			return "W"
		}
		return "S"
	} else if cardSum >= 17 && cardSum <= 20 {
		return "S"
	} else if cardSum >= 12 && cardSum <= 16 {
		if dcFloat >= 7 {
			return "H"
		}
		return "S"
	}
	return "H"
}
