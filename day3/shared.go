package day3

func prioForItem(item string) int {
	r := item[0] // Convert to rune so we can do math with its byte representation

	if r > 96 {
		return int(r - 96)
	} else {
		return int(r - 38)
	}
}
