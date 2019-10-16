package Width

import "fmt"

// One 宽度标识符
func One() {
	T1 := 12.34
	// %f	 默认宽度，默认精度
	fmt.Printf("%f\n", T1)

	// %10f  宽度10，默认精度
	fmt.Printf("%10f\n", T1)

	// %.2f  默认宽度，精度2
	fmt.Printf("%.2f\n", T1)

	// %10.3f 宽度10，精度3
	fmt.Printf("%10.3f\n", T1)

	// %10.f  宽度10，精度0
	fmt.Printf("%10.f\n", T1)
}
