package one

import (
	"fmt"
	"time"
)

// One 获取当前时间，格式化输出为2017/06/19 20:30:05格式
func One() {
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 03:04:05"))
}
