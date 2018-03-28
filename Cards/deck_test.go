package main

import (
	"os"
	"testing"
)

//测试思路： 在完全新的文件环境 -> 创建卡组 -> 保存卡组到文件 -> 加载卡组文件
// -> 测试文件正确性 -> 删除测试文件

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Deck长度应为16，而不是 %v", len(d))
	}
}

func TestNewDeckAndNewDeckFromFile(t *testing.T) {
	//长文件名是为了对应上相关函数名
	os.Remove("_desktesting")
	deck := newDeck()
	deck.saveToFile("_desktesting")

	loadedDeck := newDeckFromFile("_desktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("卡组数目不正确，现在是 %v", len(loadedDeck))
	}
	os.Remove("_desktesting")
}
