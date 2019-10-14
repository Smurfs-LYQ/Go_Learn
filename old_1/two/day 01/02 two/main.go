package main

import "fmt"

func main() {
	// 定义局部常量
	const freezingF, boillingF = 32.0, 212.0

	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boillingF, fToC(boillingF)) // "212°F = 100°C"
}

func fToC(f float64) float64 {
	return (f - 32 ) * 5 / 9
}