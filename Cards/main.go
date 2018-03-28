package main

import (
	"fmt"
)

func main() {
	fmt.Println("新建卡组")
	cards := newDeck()
	cards.print()

	fmt.Println("重分卡组")
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	fmt.Println("将卡组转换成一个字符串")
	fmt.Println(cards.toString())

	fmt.Println("保存卡组字符串到文件my_cards")
	cards.saveToFile("my_cards")

	fmt.Println("错误读取卡组文件")
	cardsErr := newDeckFromFile("my_cards_ERR")
	cardsErr.print()

	fmt.Println("读取卡组文件")
	cardsNew := newDeckFromFile("my_cards")
	cardsNew.print()

	fmt.Println("洗牌")
	cards.shuffle()
	cards.print()
}
