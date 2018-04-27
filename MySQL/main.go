package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB  *gorm.DB
	err error
)

// Scanner
type NullValue struct {
	Id      int64
	Name    sql.NullString  `sql:"not null"`
	Gender  *sql.NullString `sql:"not null"`
	Age     sql.NullInt64
	Male    sql.NullBool
	Height  sql.NullFloat64
	AddedAt NullTime
}

type NullTime struct {
	Time  time.Time
	Valid bool
}

type User0 struct {
	gorm.Model
	Birthday time.Time
	Age0     int
	Name     string `root:"size:255"`       // string默认长度为255, 使用这种tag重设。
	Num      int    `root:"AUTO_INCREMENT"` // 自增
	Email    string
}

type MyUser struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `root:"size:255"`       // string默认长度为255, 使用这种tag重设。
	Num      int    `root:"AUTO_INCREMENT"` // 自增
}

func main() {
	DB, err := gorm.Open("mysql", "root:root@/databasetest?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer DB.Close()
	// 检查模型`User0`表是否存在
	// if !DB.HasTable(&User0{}) {
	// 	fmt.Println("INIT: Model User0 already Exist, Delete it and INIT again")
	// 	DB.DropTable(&User0{})
	// 	user000 := User0{Name: "Losser", Age0: -10, Birthday: time.Now().Local(), Email: "6666@163.com"}
	// 	DB.Save(&user000)
	// }
	// // 检查表`users`是否存在
	// if !DB.HasTable("users") {
	// 	fmt.Println("INIT: Error, Table users not found")
	// } else {
	// 	fmt.Println("INIT: Table users create Success")
	// }

	// 自动迁移模式
	//	DB.AutoMigrate(&User0{}, &Product{})

	// 检查模型`User0`表是否存在
	if DB.HasTable(&MyUser{}) {
		fmt.Println("INIT: Model MyUser already Exist, Delete it and INIT again")
		DB.DropTable(&MyUser{})
		// 创建表`users'时将“ENGINE = InnoDB”附加到SQL语句
		DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&MyUser{})
		// 为模型`User0`创建表
		users111 := MyUser{Name: "Winner", Age: 18, Birthday: time.Now().Local()}
		DB.Save(&users111)
		users222 := MyUser{Name: "ProGame", Age: 88, Birthday: time.Now().UTC()}
		DB.Save(&users222)
		user1 := MyUser{Name: "RowsUser1", Age: 1, Birthday: time.Now().UTC()}
		user2 := MyUser{Name: "RowsUser2", Age: 10, Birthday: time.Now().UTC()}
		user3 := MyUser{Name: "RowsUser3", Age: 20, Birthday: time.Now().UTC()}
		DB.Save(&user1).Save(&user2).Save(&user3)
		// 获取所有记录
		fmt.Println(DB.Find(&users111))
		// fmt.Println(DB.First(&MyUser))
		// 检查表`users`是否存在
		if DB.HasTable(users111) && DB.HasTable(users222) {
			fmt.Println("LAST: OK, Table users Created")
		}
	}
	// // 检查表`users`是否存在
	// if !DB.HasTable("users") {
	// 	fmt.Println("INIT: Error, Table users not found")
	// } else {
	// 	fmt.Println("INIT: Table users create Success")
	// }

	// // 创建
	// DB.Create(&Product{Code: "L1212", Price: 1000})

	// // 读取
	// var product Product
	// DB.First(&product, 1)                   // 查询id为1的product
	// DB.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	// // 更新 - 更新product的price为2000
	// DB.Model(&product).Update("Price", 2000)

	// 删除模型`User0`的表
	// DB.DropTable(&User0{})
	// DB.DropTable(&MyUser{})
	// // // 删除表`users`
	// // DB.DropTable("users")

	// // // 删除模型`User0`的表和表`products`
	// // DB.DropTableIfExists(&User0{}, "products")

	// // 检查模型`User0`表是否存在
	// if !DB.HasTable(&MyUser{}) {
	// 	fmt.Println("LAST: OK, Model User0 is deleted")
	// }
	// inRows()
}

func inRows() {

	// timeLayout := "2006-01-02 15:04:05"
	// loc, _ := time.LoadLocation("Local")
	// toBeCharge := "2015-01-01 00:00:00"
	user1 := User0{Name: "RowsUser1", Age0: 1, Birthday: time.Now()}
	user2 := User0{Name: "RowsUser2", Age0: 10, Birthday: time.Now()}
	user3 := User0{Name: "RowsUser3", Age0: 20, Birthday: time.Now()}
	DB.Save(&user1).Save(&user2).Save(&user3)

	// rows, err := DB.Table("users").Where("name = ? or name = ?", user2.Name, user3.Name).Select("name, age").Rows()
	// if err != nil {
	// 	fmt.Println("Not error should happen, got ", err)
	// }

	// count := 0
	// for rows.Next() {
	// 	var name string
	// 	var age int64
	// 	rows.Scan(&name, &age)
	// 	count++
	// }

	// if count != 2 {
	// 	fmt.Println("Should found two records")
	// }
}

// func parseTime(str string) *time.Time {
// 	t := now.New(time.Now().UTC()).MustParse(str)
// 	return &t
// }
