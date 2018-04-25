package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
	Num      int    `gorm:"AUTO_INCREMENT"` // 自增
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/databasetest?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer db.Close()

	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	db.NewRecord(user) // => 主键为空返回`true`

	db.Create(&user)

}
