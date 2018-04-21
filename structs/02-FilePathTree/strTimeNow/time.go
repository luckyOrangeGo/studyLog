package strTimeNow

import (
	"fmt"
	"time"
)

//获取当前时间字符串
func StrTm() (str string) {

	//格式化为字符串,tm为Time类型
	tm := time.Unix(TimestampNow(), 0)
	str = tm.Format("2006-01-02 03:04:05 PM")
	return str
}

//获取时间戳int64
func TimestampNow() int64 {
	timestamp := time.Now().Unix()

	return timestamp
}

//将当前时间戳int64转换为字符串
func StrTmsp() string {
	ts := fmt.Sprintf("%d", TimestampNow())
	return ts
}
