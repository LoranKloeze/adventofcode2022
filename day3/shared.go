package day3

func prioForItem(item rune) int {

	if item > 96 {
		return int(item - 96)
	} else {
		return int(item - 38)
	}
}
