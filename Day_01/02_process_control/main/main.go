package main

import (
	"Go_Learn/Day_01/02_process_control/forDemo"
	"Go_Learn/Day_01/02_process_control/gotoDemo"
	"Go_Learn/Day_01/02_process_control/ifDemo"
	"Go_Learn/Day_01/02_process_control/switchDemo"
	"fmt"
)

func main() {
	fmt.Println("-----------if_Demo-----------")
	ifDemo.One()
	fmt.Println("-----------for_Demo----------")
	forDemo.One()
	fmt.Println("-----------switch_Demo----------")
	switchDemo.One()
	fmt.Println("-----------goto_Demo----------")
	gotoDemo.One()
}
