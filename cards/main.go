package main

//import "fmt"

func main() {
	//var cards string = "Ace of spades"
	//cards := "Ace of spades"
	//cards := newCard()

	//cards := []string{newCard(), "Ace of spades"}
	//cards := deck{newCard(), "Ace of spades"}

	//cards = append(cards, "Six of spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	//cards := newDeck()
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
	//fmt.Println(cards.toString())
	//	cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
