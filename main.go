package main

import "fmt"

func main() {

	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")
	cards := newDeck()

	//hand, remainingCards := deal(cards, 5)

	//fmt.Println("=======")
	//cards.print()
	//fmt.Println("=======")
	//hand.print()
	//fmt.Println("=======")
	//remainingCards.print()
	//fmt.Println(cards)
	fmt.Println("=======")
	cards = newDeck()
	fmt.Println(cards.toString())
	//cards.saveToFile("MyCards")

	//cards1 := newDeckFromFile("MyCards")
	//cards1.print()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
