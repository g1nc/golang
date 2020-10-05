package main

func main() {
	cards := newDeck()
	cards.shuffle()
	// cards.saveToFile("temp.file")

	// cards := newDeckFromFile("temp.file")
	cards.print()
}
