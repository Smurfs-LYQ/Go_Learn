package Project_2

/*
Project2
打印出100-999中所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字立方等于该数本身。例如：153是一个“水仙花数”，因为153=1的三次方+5的三次方+3的三次方
*/

import (
	"fmt"
	"strconv"
)

func Project2() {
	var str string
	fmt.Scanf("%s", &str) // 这个就是等待用户的输入, 拿用户输入的一个值, 类似于python中的input

	var res = 0
	for i := 0; i < len(str); i++ {
		num := int(str[i] - '0') // test := '0' 单引号声明的都是byte类型
		res += (num * num * num) //求该数字的三次方
	}
	/*
		这里利用的是
			通过下标获取数字字符num 比如num=3
			返回的是num的ASCII值, 而3的ASCII值是51 0的ASCII值是48, 51-48=3, 任意数字的ASCII值 减 0的ASCII值 等于 其数字的十进制格式
	*/

	number, err := strconv.Atoi(str) //Atoi是做字符串转int类型用的一个函数, 返回值(转换之后的内容, 错误提示), 如果转换成功 错误提示会返回一个指针
	if err != nil {
		fmt.Printf("不能将%s转换为int类型\n", str)
		return
	}

	if res == number {
		fmt.Printf("%d是莲花函数\n", number)
	} else {
		fmt.Printf("%d不是莲花函数\n", number)
	}
}
