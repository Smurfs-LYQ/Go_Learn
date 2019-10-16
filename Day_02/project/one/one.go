package one

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

// One 把字符串的IP地址"192.168.19.200"转换成整数
func One() {
	ip := "192.168.19.200"
	ipInt := InetAtoN(ip)

	fmt.Printf("convert string ip [%s] to int: %d\n", ip, ipInt)
	fmt.Printf("convert int ip [%d] to string: %s\n", ipInt, InetNtoA(ipInt))
}
