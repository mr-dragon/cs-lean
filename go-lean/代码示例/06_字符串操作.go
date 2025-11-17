package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ===== 字符串操作 =====

func stringBasics() {
	// ===== 字符串声明 =====
	// Java: String str = "Hello, Go!";
	str := "Hello, Go!"
	multiline := `Line 1
Line 2
Line 3`

	fmt.Println("String:", str)
	fmt.Println("Multiline:", multiline)

	// ===== 字符串长度 =====
	// Java: str.length()
	fmt.Println("Length:", len(str))

	// ===== 字符串索引和切片 =====
	// Java: str.charAt(0);  str.substring(0, 5);
	fmt.Println("First character (byte):", str[0])
	fmt.Println("Substring [0:5]:", str[0:5])
	fmt.Println("Substring [7:]:", str[7:])

	// ===== 字符串是不可变的 =====
	// Java: 相同
	// str[0] = 'G'  // Error: cannot assign to str[0]

	// ===== Unicode 和 Rune =====
	// Java: str.charAt(0)  返回 char
	// Go: 字符串是 UTF-8 编码，需要使用 rune 处理 Unicode
	unicodeStr := "Hello 世界 🌍"
	fmt.Println("String:", unicodeStr)
	fmt.Println("Byte length:", len(unicodeStr))  // UTF-8 字节数

	// 遍历 rune（Unicode 字符）
	fmt.Print("Runes: ")
	for _, r := range unicodeStr {
		fmt.Printf("%c ", r)
	}
	fmt.Println()

	// ===== 字节和 Rune 转换 =====
	// Java: str.getBytes();  new String(bytes);
	bytes := []byte("Hello")
	fmt.Println("Bytes:", bytes)

	runes := []rune("Go")
	fmt.Println("Runes:", runes)
}

// ===== 字符串与其他类型转换 =====
func stringConversions() {
	// Java: Integer.toString(42);  String.valueOf(3.14);
	
	// 整数转字符串
	intToStr := strconv.Itoa(42)
	fmt.Println("Int to String:", intToStr)

	// 字符串转整数
	// Java: Integer.parseInt("42");
	strToInt, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("String to Int:", strToInt)

	// 使用 ParseInt 更灵活
	// Java: Integer.parseInt("42", 10);
	parsed, _ := strconv.ParseInt("42", 10, 64)
	fmt.Println("Parsed Int64:", parsed)

	// 浮点数转字符串
	// Java: Double.toString(3.14);
	floatToStr := strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Println("Float to String:", floatToStr)

	// 字符串转浮点数
	// Java: Double.parseDouble("3.14");
	strToFloat, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println("String to Float:", strToFloat)

	// 布尔值转字符串
	// Java: Boolean.toString(true);
	boolToStr := strconv.FormatBool(true)
	fmt.Println("Bool to String:", boolToStr)

	// 字符串转布尔值
	// Java: Boolean.parseBoolean("true");
	strToBool, _ := strconv.ParseBool("true")
	fmt.Println("String to Bool:", strToBool)
}

// ===== 字符串方法（使用 strings 包）=====
func stringMethods() {
	// Java: str.contains("Go");
	str := "Hello, Go!"
	
	// 包含
	contains := strings.Contains(str, "Go")
	fmt.Println("Contains 'Go':", contains)

	// 前缀和后缀
	// Java: str.startsWith("Hello");  str.endsWith("!");
	fmt.Println("Starts with 'Hello':", strings.HasPrefix(str, "Hello"))
	fmt.Println("Ends with '!':", strings.HasSuffix(str, "!"))

	// 索引
	// Java: str.indexOf("Go");
	index := strings.Index(str, "Go")
	fmt.Println("Index of 'Go':", index)

	// 替换
	// Java: str.replace("Go", "Golang");
	replaced := strings.ReplaceAll(str, "Go", "Golang")
	fmt.Println("Replaced:", replaced)

	// 大小写转换
	// Java: str.toUpperCase();  str.toLowerCase();
	fmt.Println("Upper:", strings.ToUpper(str))
	fmt.Println("Lower:", strings.ToLower(str))

	// 去除空白
	// Java: str.trim();
	spacedStr := "  Hello, Go!  "
	fmt.Println("Trimmed:", strings.TrimSpace(spacedStr))

	// 分割
	// Java: str.split(",");
	parts := strings.Split(str, ",")
	fmt.Println("Split by ',':", parts)

	// 连接
	// Java: String.join(",", array);
	joined := strings.Join([]string{"Hello", "Go", "World"}, " ")
	fmt.Println("Joined:", joined)

	// 重复
	// Java: 无直接对应
	repeated := strings.Repeat("Go ", 3)
	fmt.Println("Repeated:", repeated)

	// 计数
	// Java: 无直接对应
	count := strings.Count("banana", "a")
	fmt.Println("Count of 'a' in 'banana':", count)

	// 字段分割（按空白分割）
	// Java: str.split("\\s+");
	fields := strings.Fields("Hello   Go   World")
	fmt.Println("Fields:", fields)
}

// ===== 字符串拼接 =====
func stringConcatenation() {
	// Java: str1 + str2  或  StringBuilder

	// 1. 直接 + 拼接（简单情况）
	str1 := "Hello"
	str2 := "Go"
	result := str1 + ", " + str2
	fmt.Println("Direct concatenation:", result)

	// 2. fmt.Sprintf（用于复杂格式化）
	// Java: String.format()
	name := "Alice"
	age := 30
	formatted := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println("Formatted:", formatted)

	// 3. strings.Builder（高性能拼接）
	// Java: StringBuilder sb = new StringBuilder();
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("Go")
	builder.WriteString("!")
	fmt.Println("Builder result:", builder.String())

	// 4. strings.Join（拼接切片）
	// Java: String.join(",", list);
	parts := []string{"Apple", "Banana", "Cherry"}
	joined := strings.Join(parts, ", ")
	fmt.Println("Joined:", joined)
}

// ===== 正则表达式 =====
func regularExpressions() {
	// Java: Pattern p = Pattern.compile("\\d+");  Matcher m = p.matcher(str);
	
	// Go 使用 regexp 包（需要导入）
	// 这里仅展示使用模式

	// import "regexp"
	// re := regexp.MustCompile(`\d+`)
	// matches := re.FindAllString("abc123def456", -1)  // 查找所有数字
	// matched := re.MatchString("test123")              // 匹配检查
	// result := re.ReplaceAllString("hello123world", "")// 替换

	fmt.Println("正则表达式示例：")
	fmt.Println("- 导入 regexp 包")
	fmt.Println("- 使用 regexp.MustCompile() 创建正则表达式")
	fmt.Println("- 使用 FindAllString(), MatchString(), ReplaceAllString() 等方法")
}

// ===== 字符串和字节数组相互转换 =====
func bytesConversion() {
	// ===== 字符串转字节数组 =====
	// Java: str.getBytes();
	str := "Hello, Go!"
	bytes := []byte(str)
	fmt.Println("String to bytes:", bytes)
	fmt.Println("As string:", string(bytes))

	// ===== 字节数组转字符串 =====
	// Java: new String(bytes);
	byteArray := []byte{72, 101, 108, 108, 111}  // "Hello"
	strFromBytes := string(byteArray)
	fmt.Println("Bytes to string:", strFromBytes)

	// ===== Rune 数组 =====
	// Java: char[] chars = str.toCharArray();
	runeArray := []rune("Hello")
	fmt.Println("Rune array:", runeArray)
	fmt.Println("Back to string:", string(runeArray))
}

// ===== 字符串比较 =====
func stringComparison() {
	str1 := "Hello"
	str2 := "Hello"
	str3 := "World"

	// ===== 相等比较 =====
	// Java: str1.equals(str2);
	fmt.Println("str1 == str2:", str1 == str2)
	fmt.Println("str1 == str3:", str1 == str3)

	// ===== 字典序比较 =====
	// Java: str1.compareTo(str2);
	cmp := strings.Compare(str1, str3)
	fmt.Println("Compare 'Hello' with 'World':", cmp)  // -1 (小于)

	// ===== 不区分大小写比较 =====
	// Java: str1.equalsIgnoreCase(str2);
	fmt.Println("Equal (ignore case):", strings.EqualFold("Hello", "hello"))
}

// ===== 字符串检查 =====
func stringChecks() {
	// ===== 是否为空 =====
	// Java: str == null || str.isEmpty();
	empty := ""
	fmt.Println("Is empty:", empty == "")

	// ===== 是否仅包含空白 =====
	// Java: str.trim().isEmpty();
	spaced := "   "
	fmt.Println("Is blank:", strings.TrimSpace(spaced) == "")

	// ===== 是否为字母数字 =====
	// Java: 无直接对应，需要正则表达式
	fmt.Println("String checks:")
	fmt.Println("- 使用 regexp 包进行更复杂的检查")
	fmt.Println("- 使用 unicode 包检查字符属性")
}

// ===== 高级示例：字符串解析 =====
func advancedStringParsing() {
	// Java: String[] parts = csv.split(",");  parseKey(parts[0]);
	
	csv := "name,age,city"
	fields := strings.Split(csv, ",")
	fmt.Println("CSV fields:", fields)

	// 解析 key=value 格式
	kvPair := "host=localhost:8080"
	parts := strings.Split(kvPair, "=")
	if len(parts) == 2 {
		key, value := parts[0], parts[1]
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// 构建 query string
	// Java: URLEncoder.encode()
	queryString := fmt.Sprintf("name=%s&age=%d&city=%s", "Alice", 30, "New York")
	fmt.Println("Query string:", queryString)
}

func main() {
	fmt.Println("=== 字符串基础 ===")
	stringBasics()

	fmt.Println("\n=== 字符串与其他类型转换 ===")
	stringConversions()

	fmt.Println("\n=== 字符串方法 ===")
	stringMethods()

	fmt.Println("\n=== 字符串拼接 ===")
	stringConcatenation()

	fmt.Println("\n=== 正则表达式 ===")
	regularExpressions()

	fmt.Println("\n=== 字节数组转换 ===")
	bytesConversion()

	fmt.Println("\n=== 字符串比较 ===")
	stringComparison()

	fmt.Println("\n=== 字符串检查 ===")
	stringChecks()

	fmt.Println("\n=== 高级字符串解析 ===")
	advancedStringParsing()
}
