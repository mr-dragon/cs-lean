package main

import "fmt"

// ===== 函数基础 =====

// 1. 基本函数定义
// Java: public int add(int a, int b) { return a + b; }
func add(a int, b int) int {
	return a + b
}

// 2. 多参数类型相同，可以合并
// Java: public int multiply(int a, int b, int c)
func multiply(a, b, c int) int {
	return a * b * c
}

// 3. 多返回值（Java 需要返回对象或 Map）
// Java: 无对应，需要创建对象
func divideWithRemainder(a, b int) (int, int) {
	return a / b, a % b
}

// 4. 命名返回值
// Java: 无对应
func rectangleArea(length, width int) (area int, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return  // 自动返回命名返回值
}

// 5. 可变参数（... 表示可变长）
// Java: public void printNumbers(int... numbers)
func printNumbers(numbers ...int) {
	for _, num := range numbers {
		fmt.Println(num)
	}
}

// 6. 空返回值（void）
// Java: public void greet(String name)
func greet(name string) {
	fmt.Println("Hello, " + name)
}

// 7. 函数作为参数（高阶函数）
// Java: void callFunction(Consumer<Integer> func)  或使用 lambda
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// 8. 函数作为返回值
// Java: Function<Integer, Integer> getMultiplier(int factor)
func getMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// ===== 匿名函数和闭包 =====
// Java: new Thread(() -> { /* 代码 */ }).start();
func closuresDemo() {
	// 匿名函数
	func() {
		fmt.Println("This is an anonymous function")
	}()

	// 闭包（capture 外部变量）
	x := 10
	increment := func() int {
		x++  // 修改外部变量
		return x
	}
	fmt.Println("Increment:", increment())  // 11
	fmt.Println("Increment:", increment())  // 12
	fmt.Println("Final x:", x)              // 12
}

// ===== 延迟函数调用（defer）=====
// Java: finally 块或 try-with-resources
func deferExecution() {
	fmt.Println("Start")
	
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")
	
	fmt.Println("Middle")
}

// 实际使用场景：关闭文件
func readFile(filename string) {
	// file, _ := os.Open(filename)
	// defer file.Close()  // 确保文件被关闭
	fmt.Println("Reading file:", filename)
}

// ===== 值接收者和指针接收者 =====

type Counter struct {
	value int
}

// 值接收者（不修改原对象）
// Java: public int getValue()
func (c Counter) getValue() int {
	return c.value
}

// 指针接收者（修改原对象）
// Java: public void increment()
func (c *Counter) increment() {
	c.value++
}

// ===== 接口 =====
// Java: public interface Shape { double area(); }
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width float64
}

// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
	return 3.14159 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.radius
}

// Rectangle 实现 Shape 接口
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

// 接受任何实现 Shape 的对象
// Java: public void printShapeInfo(Shape shape)
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// 空接口（可以接受任何类型）
// Java: public void process(Object obj)
func process(v interface{}) {
	fmt.Println("Received:", v)
}

// ===== 错误处理 =====
// Java: throws Exception  或 try-catch
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// 自定义错误
// Java: class CustomException extends Exception
type CustomError struct {
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func validateInput(input string) error {
	if input == "" {
		return CustomError{Code: 400, Message: "input cannot be empty"}
	}
	return nil
}

// ===== 性能优化：值 vs 指针 =====

type LargeStruct struct {
	data [1000]int
}

// 值接收者（复制整个结构体，性能低）
func (ls LargeStruct) processValue() {
	fmt.Println("Processing by value")
}

// 指针接收者（只传递指针，高效）
func (ls *LargeStruct) processPointer() {
	fmt.Println("Processing by pointer")
}

// ===== 递归函数 =====
// Java: public int factorial(int n)
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 斐波那契数列（带缓存防止重复计算）
func fibonacci(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, exists := memo[n]; exists {
		return val
	}
	result := fibonacci(n-1, memo) + fibonacci(n-2, memo)
	memo[n] = result
	return result
}

func main() {
	fmt.Println("=== 基本函数 ===")
	fmt.Println("add(3, 5) =", add(3, 5))
	fmt.Println("multiply(2, 3, 4) =", multiply(2, 3, 4))
	
	q, r := divideWithRemainder(10, 3)
	fmt.Println("divideWithRemainder(10, 3): quotient =", q, "remainder =", r)
	
	area, perimeter := rectangleArea(5, 3)
	fmt.Println("Rectangle: area =", area, "perimeter =", perimeter)

	fmt.Println("\n=== 可变参数 ===")
	printNumbers(1, 2, 3, 4, 5)
	
	fmt.Println("\n=== 函数作为参数 ===")
	result := applyOperation(10, 5, add)
	fmt.Println("applyOperation(10, 5, add) =", result)

	fmt.Println("\n=== 函数作为返回值 ===")
	multiply3 := getMultiplier(3)
	fmt.Println("multiply3(5) =", multiply3(5))
	fmt.Println("multiply3(10) =", multiply3(10))

	fmt.Println("\n=== 闭包 ===")
	closuresDemo()

	fmt.Println("\n=== defer ===")
	deferExecution()

	fmt.Println("\n=== 接收者方法 ===")
	counter := Counter{value: 0}
	fmt.Println("Counter value:", counter.getValue())
	counter.increment()  // 这不会改变 counter，因为是值接收者
	fmt.Println("Counter value after increment:", counter.getValue())
	
	counterPtr := &counter
	counterPtr.increment()  // 使用指针接收者
	fmt.Println("Counter value after increment (via pointer):", counterPtr.getValue())

	fmt.Println("\n=== 接口 ===")
	circle := Circle{radius: 5}
	rect := Rectangle{length: 4, width: 6}
	
	fmt.Println("Circle:")
	printShapeInfo(circle)
	fmt.Println("Rectangle:")
	printShapeInfo(rect)

	fmt.Println("\n=== 空接口 ===")
	process(42)
	process("Hello")
	process(3.14)
	process([]int{1, 2, 3})

	fmt.Println("\n=== 错误处理 ===")
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n=== 自定义错误 ===")
	err = validateInput("")
	if err != nil {
		fmt.Println("Validation error:", err)
	}

	fmt.Println("\n=== 递归 ===")
	fmt.Println("factorial(5) =", factorial(5))
	fmt.Println("fibonacci(10) =", fibonacci(10, make(map[int]int)))

	fmt.Println("\n=== 匿名函数 ===")
	func(a, b int) {
		fmt.Printf("Sum: %d + %d = %d\n", a, b, a+b)
	}(7, 3)
}
