package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// 获取　年月日　时分秒
	year := now.Year()
	month := now.Month()
	day := now.Day()

	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second)

	// 格式化日期，输出字符串
	layout := "2006-01-02 15:04:05"
	timeStr := now.Format(layout)
	fmt.Println(timeStr)

}
