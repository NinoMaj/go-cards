package main

func main() {
	cards := newCards()

	// hand, remainingCards := deal(cards, 5)
	// fmt.Println(hand, remainingCards)

	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")
}
