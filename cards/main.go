package main

func main() {
	cards := newDeck()
	// cards.saveToFile("PMR.txt")
	// cards := newDeckFromFile("PMR..2txt")
	cards.shuffle()
	cards.print()
	//fmt.Println(cards.toString())
	//	cards = append(cards, "Six of Pikes")

	// hand, remaining := deal(cards, 2)
	// hand.print()
	// remaining.print()
	// cards.print()
	// fmt.Println(cards)

	// for _, card := range cards {
	// 	fmt.Println(card)
	// }

	// for i, ch := range "日本語" {
	// 	fmt.Printf("%#U starts at byte position %d\n", ch, i)
	// }

	// const s = "日本語"
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println( s[i])
	// }

	//buildAnimals()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
}
