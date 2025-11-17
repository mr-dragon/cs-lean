package main

import "fmt"

// ===== 基本数据类型 =====

func basicTypes() {
	// ===== 整数类型 =====
	// Java: byte, short, int, long
	var b byte = 255           // 8 位无符号
	var sh int16 = 32767       // 16 位有符号
	var i int32 = 2147483647   // 32 位有符号
	var l int64 = 9223372036854775807 // 64 位有符号
	var ui uint = 100          // 无符号，大小依赖平台
	var ui32 uint32 = 200      // 32 位无符号
	var ui64 uint64 = 300      // 64 位无符号
	var num int = 42           // 默认 int 类型（平台相关）
	
	fmt.Println("byte:", b, "int16:", sh, "int32:", i, "int64:", l)
	fmt.Println("uint:", ui, "uint32:", ui32, "uint64:", ui64, "int:", num)

	// ===== 浮点类型 =====
	// Java: float (32 位), double (64 位)
	var f32 float32 = 3.14     // 32 位浮点
	var f64 float64 = 2.71828  // 64 位浮点 (默认)
	fval := 9.99               // 默认为 float64
	
	fmt.Println("float32:", f32, "float64:", f64, "default:", fval)

	// ===== 布尔类型 =====
	// Java: boolean
	var isActive bool = true
	isEnabled := false
	
	fmt.Println("isActive:", isActive, "isEnabled:", isEnabled)

	// ===== 字符和字符串 =====
	// Java: char (单个 Unicode 字符), String (字符串)
	var ch rune = 'A'          // rune 是 int32 的别名，表示一个 Unicode 字符
	var str string = "Hello, Go!"
	multilineStr := `This is a
multiline string
in Go`
	
	fmt.Println("rune:", ch, "string:", str)
	fmt.Println(multilineStr)
}

// ===== 复合数据类型 =====

func arrayAndSlice() {
	// ===== 数组 =====
	// Java: int[] arr = {1, 2, 3};  或  int[] arr = new int[5];
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	arr2 := [...]int{10, 20, 30} // 长度自动推导
	arr3 := [5]int{}             // 默认初始化为零
	
	fmt.Println("Array:", arr)
	fmt.Println("Array2:", arr2)
	fmt.Println("Array3:", arr3)
	fmt.Println("Array length:", len(arr))
	fmt.Println("Array capacity:", cap(arr))

	// ===== 切片（Slice）=====
	// Java: List<Integer> list = new ArrayList<>(Arrays.asList(1, 2, 3));
	slice := []int{1, 2, 3, 4, 5}  // 不指定长度则为切片
	slice2 := make([]int, 5)        // 创建长度为 5 的切片
	slice3 := make([]int, 3, 10)    // 长度 3，容量 10
	
	fmt.Println("Slice:", slice)
	fmt.Println("Slice2:", slice2)
	fmt.Println("Slice3:", slice3)
	fmt.Println("Slice length:", len(slice))
	fmt.Println("Slice capacity:", cap(slice))

	// 切片操作
	slice = append(slice, 6, 7)                // 追加元素
	subSlice := slice[1:4]                    // 切片切割 [2, 3, 4]
	fmt.Println("After append:", slice)
	fmt.Println("Sub-slice [1:4]:", subSlice)

	// ===== 映射（Map）=====
	// Java: Map<String, Integer> map = new HashMap<>();
	map1 := make(map[string]int)
	map1["apple"] = 5
	map1["banana"] = 3
	
	// 字面量形式
	map2 := map[string]string{
		"name": "Gopher",
		"lang": "Go",
	}
	
	fmt.Println("Map1:", map1)
	fmt.Println("Map2:", map2)
	fmt.Println("Value of 'apple':", map1["apple"])

	// 检查键是否存在
	// Java: if (map.containsKey("apple"))
	value, exists := map1["apple"]
	fmt.Println("'apple' exists:", exists, "value:", value)

	// 删除元素
	// Java: map.remove("apple");
	delete(map1, "apple")
	fmt.Println("After delete:", map1)

	// ===== 结构体（Struct）=====
	// Java: class Person { String name; int age; }
	type Person struct {
		Name string
		Age  int
		Email string
	}
	
	// 创建结构体实例
	person1 := Person{"Alice", 30, "alice@example.com"}
	person2 := Person{Name: "Bob", Age: 25}  // 命名字段
	
	fmt.Println("Person1:", person1)
	fmt.Println("Person2:", person2)
	fmt.Println("Name:", person1.Name, "Age:", person1.Age)

	// ===== 嵌套结构体 =====
	// Java: class Company { Person ceo; }
	type Company struct {
		Name string
		CEO  Person
	}
	
	company := Company{
		Name: "TechCorp",
		CEO: Person{Name: "Charlie", Age: 45},
	}
	
	fmt.Println("Company:", company)
	fmt.Println("CEO Name:", company.CEO.Name)
}

// ===== 指针 =====
func pointersDemo() {
	// Java: 一切都是引用，使用 new 创建对象
	x := 10
	
	// & 取地址
	var ptr *int = &x
	
	// * 解引用
	fmt.Println("Value at pointer:", *ptr)
	fmt.Println("Pointer address:", ptr)

	// 修改指针指向的值
	*ptr = 20
	fmt.Println("After modification, x =", x)

	// nil 指针
	// Java: null
	var nilPtr *int
	fmt.Println("nil pointer:", nilPtr)

	// 结构体指针
	type Dog struct {
		Name string
	}
	
	dog := Dog{"Buddy"}
	dogPtr := &dog
	
	// Java: dog.name  或  dogPtr.name  (自动解引用)
	// Go: dog.Name   或  dogPtr.Name  (自动解引用)
	fmt.Println("Dog name via value:", dog.Name)
	fmt.Println("Dog name via pointer:", dogPtr.Name)
	
	// 修改
	dogPtr.Name = "Max"
	fmt.Println("After modification:", dog.Name)
}

func main() {
	fmt.Println("=== 基本数据类型 ===")
	basicTypes()
	
	fmt.Println("\n=== 数组、切片、映射、结构体 ===")
	arrayAndSlice()
	
	fmt.Println("\n=== 指针 ===")
	pointersDemo()
}
