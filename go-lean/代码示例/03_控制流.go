package main

import "fmt"

// ===== 控制流 =====

func ifElseDemo() {
	// ===== if-else =====
	// Java: if (x > 5) { } else if (x > 2) { } else { }
	x := 10
	
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x > 2 {
		fmt.Println("x is between 2 and 5")
	} else {
		fmt.Println("x is 2 or less")
	}

	// 在 if 条件中初始化变量
	// Java: 无对应写法
	if y := 20; y > 15 {
		fmt.Println("y is greater than 15")
	}
	// fmt.Println(y)  // Error: y 超出作用域
}

func switchDemo() {
	// ===== switch 语句 =====
	// Java: switch(value) { case 1: break; default: }
	day := 3
	
	// 基本 switch
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}

	// 多值 case
	// Java: case 1: case 2: case 3:  (需要手动处理)
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	}

	// 没有表达式的 switch（类似 if-else 链）
	// Java: 无对应写法
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade A")
	case score >= 80:
		fmt.Println("Grade B")
	case score >= 70:
		fmt.Println("Grade C")
	default:
		fmt.Println("Grade F")
	}

	// switch 中初始化变量
	// Java: 无对应
	switch x := 5; x {
	case 5:
		fmt.Println("x is 5")
	default:
		fmt.Println("x is not 5")
	}
}

func forLoopsDemo() {
	// ===== for 循环 =====
	// Go 只有 for 循环（没有 while）

	// 1. C 风格的 for 循环
	// Java: for (int i = 0; i < 5; i++)
	for i := 0; i < 5; i++ {
		fmt.Println("C-style for, i =", i)
	}

	// 2. while 风格的 for 循环
	// Java: while (x < 10) { x++ }
	x := 0
	for x < 5 {
		fmt.Println("While-style for, x =", x)
		x++
	}

	// 3. 无限循环
	// Java: for (;;)  或  while (true)
	count := 0
	for {
		if count >= 3 {
			break
		}
		fmt.Println("Infinite loop, count =", count)
		count++
	}

	// 4. range 循环（遍历数组、切片、映射、字符串）
	// Java: for (int item : array)  或  for (String key : map.keySet())
	
	// 遍历数组/切片
	slice := []int{10, 20, 30, 40}
	for i, value := range slice {
		fmt.Println("Index:", i, "Value:", value)
	}

	// 只获取索引
	for i := range slice {
		fmt.Println("Index:", i)
	}

	// 只获取值
	for _, value := range slice {
		fmt.Println("Value:", value)
	}

	// 遍历映射
	// Java: for (String key : map.keySet())
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range map1 {
		fmt.Println("Key:", key, "Value:", value)
	}

	// 遍历字符串
	// Java: char[] chars = str.toCharArray();  for (char c : chars)
	str := "Hello"
	for i, ch := range str {
		fmt.Println("Position:", i, "Character:", string(ch))
	}

	// continue 和 break
	// Java: 相同
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue  // 跳过这次迭代
		}
		if i == 4 {
			break     // 退出循环
		}
		fmt.Println("Loop with continue/break:", i)
	}

	// 标签和跳转（Go 特有）
	// Java: 无直接对应
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break OuterLoop  // 跳出外层循环
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}

// ===== 延迟执行（defer）=====
// Java: finally 块  或  try-with-resources
func deferDemo() {
	defer fmt.Println("This is deferred - executed last")
	fmt.Println("This is executed first")

	// 多个 defer 按后进先出（LIFO）执行
	for i := 1; i <= 3; i++ {
		defer fmt.Println("Deferred:", i)
	}
	fmt.Println("Main logic")
}

// ===== 条件语句和错误处理 =====
func errorHandlingDemo() {
	// Java: try-catch-finally  或  throws

	// 模拟一个返回错误的函数
	result, err := divideNumbers(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// 处理可能为 nil 的指针
	var ptr *int
	if ptr == nil {
		fmt.Println("Pointer is nil")
	}
}

func divideNumbers(a, b int) (int, error) {
	// Java: if (b == 0) throw new ArithmeticException("Division by zero");
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// ===== goto 语句 =====
// Java: 无 (已被舍弃)
// Go: 虽然有，但不推荐使用
func gotoDemo() {
	i := 0
	
Loop:
	fmt.Println("i =", i)
	i++
	
	if i < 3 {
		goto Loop  // 不推荐
	}
	
	fmt.Println("Loop finished")
}

func main() {
	fmt.Println("=== if-else ===")
	ifElseDemo()

	fmt.Println("\n=== switch ===")
	switchDemo()

	fmt.Println("\n=== for loops ===")
	forLoopsDemo()

	fmt.Println("\n=== defer ===")
	deferDemo()

	fmt.Println("\n=== error handling ===")
	errorHandlingDemo()

	fmt.Println("\n=== goto ===")
	gotoDemo()
}
