1. 把字符串的IP地址"192.168.19.200"转换成整数
```
package main

import (
	"fmt"
	"math/big"
	"net"
)

// InetAtoN 将IP转换成整数
func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

// InetNtoA 将整数转换成IP
func InetNtoA(ip int64) string {
	fmt.Println(ip)
	fmt.Printf("%T %v\n", ip>>24, ip>>24)
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func main() {
	ip := "192.168.19.200"
	ipInt := InetAtoN(ip)

	fmt.Printf("convert string ip [%s] to int: %d\n", ip, ipInt)
	fmt.Printf("convert int ip [%d] to string: %s\n", ipInt, InetNtoA(ipInt))
}

```
2. 变量d输出什么
```
var a = []int{1,2,3,4,5,6,7}
b := a[2:5]
fmt.Println(c)		// [3, 4, 5]
fmt.Println(len(c)) // 3
fmt.Println(cap(c)) // 5

c := b[:5]
fmt.Println(c)		// [3, 4, 5, 6, 7]
/*
	输出[3, 4, 5, 6, 7]的原因是b在切割完之后长度为3，所以b只输出[3，4，5]，但是的容量为5，其中包含的其实是[3, 4, 5, 6, 7]，但是只
*/
```
3. 查看下面的代码的输出结果
```
var a = make([]int, 5, 10)
for i := 0; i < 10; i++ {
	a = append(a, i)
}
fmt.Println(a)
```
4. "hou do you do"中每个单词出现的次数