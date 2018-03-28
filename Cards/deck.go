package main

import (
	//"os"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

//流程思路： 创建卡组-> 保存到文件-> 检查加载文件->
type deck []string //创建method

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"} //牌样
	cardValues := []string{"Ace", "Two", "Three", "Four"}          //牌点
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit) //字符串整合
		}
	}
	return cards //合并成牌组
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
	fmt.Println("This Func End.")
} //打印序号和卡牌

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
} //将cards按照handSize切片,返回切片后的两部分

func (d deck) toString() string {
	return strings.Join([]string(d), ", ") //将卡组整合成一个字符串，用“，”分割不同牌
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) //保存到文件
	//“0666” 表示此文件所有用户均可读写
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) //bs = byteString
	if err != nil {
		// Option #1 - 输出错误信息并返回
		fmt.Println("Error: ", err)

		// Option #2 - 输出错误并完全退出程序
		// os.Exit(1)
	}
	s := strings.Split(string(bs), ",") //将“，”隔开的牌组恢复成单张
	return deck(s)
}

func (d deck) shuffle() { //洗牌
	source := rand.NewSource(time.Now().UnixNano()) //依靠时间信息创建完全随机的Rand函数的种子随机数
	//UnixNano根据现有时间生成独一无二的int64作为rand函数随机种子
	r := rand.New(source)

	for i := range d {
		newPostion := r.Intn(len(d) - 1) //随机生成新的牌位顺序

		d[i], d[newPostion] = d[newPostion], d[i] //交换位置
	}
}
