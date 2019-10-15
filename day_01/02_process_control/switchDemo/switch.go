package switchDemo

import "fmt"

// One switch 示例
func One() {
	// 使用 switch 语句可方便的对大量的值进行条件判断
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("鸡爪？")
	}

	// 一个分支可以有多个值，多个case值之间使用英文逗号分隔
	switch i := 9; i {
	case 1, 3, 5, 7, 9:
		fmt.Println(i, "是基数")
	case 2, 4, 6, 8:
		fmt.Println(i, "是偶数")
	default:
		fmt.Println("超出测试范围了")
	}

	// 分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量
	age := 21
	switch {
	case age < 18:
		fmt.Println("虽然你未成年，但是你也不能为所欲为")
	case age >= 18 && age <= 60:
		fmt.Println("没权没势的，老实做人吧")
	case age > 60:
		fmt.Println("咱是有素质的人啊，碰瓷那种事咱不能干")
	default:
		fmt.Println("你已超出三界之外，不在无形之中")
	}

	// fallthrough 语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的
	switch s := "a"; s {
	case "a":
		fmt.Println("a, 如果打印出了a，那么下面应该也会打印出b")
		fallthrough
	case "b":
		fmt.Println("b")
	case "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
