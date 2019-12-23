package main

import "fmt"

func main() {
	int1 := []int{4, 2, 1, 3, 5}

	for i := 0; i < len(int1); i++ { // 因为每个元素都需要进行冒泡, 所以做一次全面的循环, 这样就可以调用到所有的元素了
		for j := i + 1; j < len(int1); j++ { // 拿到此次要冒泡的元素之后, 要拿剩下所有的元素, 所以需要再循环一次, i := i+1 也保证了拿的要冒泡的数字后面的所有数字, 并且跳过之前已经冒泡完毕的数字
			if int1[i] > int1[j] { // 进行两个元素的比对
				int1[i], int1[j] = int1[j], int1[i] // 如果要冒泡的数字比比对的数字大, 那么他们两个进行换位操作
			}
			fmt.Println(int1)

		}
		fmt.Println("############")
	}
	fmt.Println(int1)
}
