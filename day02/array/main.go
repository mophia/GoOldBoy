package main

import "fmt"

// 数组 array：存放元素的容器，
// 必须指定存放元素的类型和容量（长度），
// 元素的个数（数组的长度）是数组类型的一部分

func main() {
	var a1 [3]bool // [true false true]
	var a2 [4]bool // [true true false false]

	fmt.Printf("a1:%T a2:%T", a1, a2)

	// 数组的初始化
	// 如果不初始化：默认元素都是零值（布尔值：false，整型和浮点型：0，字符串：""）。
	fmt.Println(a1, a2)

	// 1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 2.初始化方式2
	a10 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a10)
	// 3.初始化方式3：根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	// 数组的遍历
	cities := [...]string{"北京", "上海", "深圳"}
	// 1. 根据索引遍历
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	// 2. for range 遍历
	for index, value := range cities {
		fmt.Println(index, value)
	}

	// 多维数组
	// [ [1 2] [3 4] [5 6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a11)

	// 多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型（不是引用类型）
	b1 := [3]int{1, 2, 3} // [1 2 3]
	b2 := b1              // [1 2 3] Ctrl+C Ctrl+V => 把word文档从文件夹A拷贝到文件夹B
	b2[0] = 100           // b2: [100 2 3]
	fmt.Println(b1, b2)   // b1: [1 2 3]

}
