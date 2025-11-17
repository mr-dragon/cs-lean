package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ===== 文件操作 =====

func createAndWriteFile() {
	// ===== 创建文件 =====
	// Java: FileWriter fw = new FileWriter("file.txt");
	
	filename := "output.txt"
	content := "Hello, Go!\nThis is a test file.\n"

	// 写入文件（覆盖）
	// Java: Files.write(path, content.getBytes());
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File created:", filename)
}

func readFile() {
	// ===== 读取整个文件 =====
	// Java: new String(Files.readAllBytes(path));
	
	filename := "output.txt"
	
	// 读取整个文件
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	
	fmt.Println("File content:")
	fmt.Println(string(content))
}

func appendToFile() {
	// ===== 追加到文件 =====
	// Java: new FileWriter(filename, true);
	
	filename := "output.txt"
	appendContent := "Appended line.\n"

	// 打开文件追加模式
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(appendContent)
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
	fmt.Println("Content appended")
}

func readFileLineByLine() {
	// ===== 逐行读取文件 =====
	// Java: BufferedReader reader = new BufferedReader(new FileReader("file.txt"));
	
	filename := "output.txt"
	
	// 方法 1: 读取所有内容并按行分割
	content, _ := ioutil.ReadFile(filename)
	lines := string(content)
	
	fmt.Println("Reading file line by line:")
	for _, line := range lines[:len(lines)-1] {  // 避免最后的换行符
		if line == '\n' {
			fmt.Println()
		} else {
			fmt.Print(string(line))
		}
	}

	// 更好的方法：使用 bufio.Scanner
	// import "bufio"
	// file, _ := os.Open(filename)
	// defer file.Close()
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	//     line := scanner.Text()
	//     fmt.Println(line)
	// }
}

func fileInfo() {
	// ===== 获取文件信息 =====
	// Java: File file = new File("file.txt");  file.exists();  file.length();
	
	filename := "output.txt"
	
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("File name: %s\n", info.Name())
	fmt.Printf("Size: %d bytes\n", info.Size())
	fmt.Printf("Mode: %v\n", info.Mode())
	fmt.Printf("Modified: %v\n", info.ModTime())
	fmt.Printf("Is directory: %v\n", info.IsDir())
}

func deleteFile() {
	// ===== 删除文件 =====
	// Java: new File("file.txt").delete();
	
	filename := "temp.txt"
	
	// 创建临时文件
	ioutil.WriteFile(filename, []byte("temp"), 0644)
	
	// 删除文件
	err := os.Remove(filename)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File deleted:", filename)
}

func directoryOperations() {
	// ===== 目录操作 =====
	// Java: new File("mydir").mkdirs();  File.listFiles();
	
	dirName := "mydir"
	
	// 创建目录
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	fmt.Println("Directory created:", dirName)

	// 列出目录内容
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println("Directory contents:")
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("[DIR]", file.Name())
		} else {
			fmt.Println("[FILE]", file.Name())
		}
	}

	// 删除目录
	os.RemoveAll(dirName)
	fmt.Println("Directory deleted")
}

func filePathOperations() {
	// ===== 路径操作 =====
	// Java: new File(dir, filename);  file.getAbsolutePath();
	
	dir := "data"
	filename := "test.txt"

	// 连接路径
	// Java: Paths.get(dir, filename);
	fullPath := filepath.Join(dir, filename)
	fmt.Println("Full path:", fullPath)

	// 获取目录
	directory := filepath.Dir(fullPath)
	fmt.Println("Directory:", directory)

	// 获取文件名
	name := filepath.Base(fullPath)
	fmt.Println("Filename:", name)

	// 获取文件扩展名
	ext := filepath.Ext(fullPath)
	fmt.Println("Extension:", ext)

	// 获取绝对路径
	// Java: file.getAbsolutePath();
	absPath, _ := filepath.Abs(fullPath)
	fmt.Println("Absolute path:", absPath)

	// 获取相对路径
	// Java: 无直接对应
	rel, _ := filepath.Rel("/home/user", "/home/user/documents/file.txt")
	fmt.Println("Relative path:", rel)
}

// ===== JSON 操作 =====

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func jsonMarshaling() {
	// ===== 对象转 JSON 字符串 =====
	// Java: new ObjectMapper().writeValueAsString(person);
	
	person := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}

	// 转换为 JSON
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}

	fmt.Println("JSON:", string(jsonBytes))

	// 格式化 JSON（缩进）
	// Java: mapper.writerWithDefaultPrettyPrinter().writeValueAsString(person);
	prettyJSON, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}

	fmt.Println("Formatted JSON:")
	fmt.Println(string(prettyJSON))
}

func jsonUnmarshaling() {
	// ===== JSON 字符串转对象 =====
	// Java: new ObjectMapper().readValue(json, Person.class);
	
	jsonStr := `{"name":"Bob","age":25,"email":"bob@example.com"}`

	var person Person
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	fmt.Printf("Person: Name=%s, Age=%d, Email=%s\n", person.Name, person.Age, person.Email)
}

func jsonWithDynamicData() {
	// ===== 动态 JSON 处理 =====
	// Java: JSONObject json = new JSONObject(jsonStr);

	// 使用 map[string]interface{} 处理动态 JSON
	jsonStr := `{"name":"Charlie","age":35,"city":"New York"}`

	var data map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &data)

	fmt.Println("Dynamic JSON:")
	for key, value := range data {
		fmt.Printf("%s: %v\n", key, value)
	}

	// 创建动态 JSON
	newData := map[string]interface{}{
		"name": "David",
		"age":  28,
		"city": "San Francisco",
	}

	jsonBytes, _ := json.Marshal(newData)
	fmt.Println("Created JSON:", string(jsonBytes))
}

func jsonArrayHandling() {
	// ===== 处理 JSON 数组 =====
	// Java: JSONArray array = json.getJSONArray("items");

	type Product struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Price float64 `json:"price"`
	}

	products := []Product{
		{ID: 1, Name: "Laptop", Price: 999.99},
		{ID: 2, Name: "Mouse", Price: 29.99},
		{ID: 3, Name: "Keyboard", Price: 79.99},
	}

	// 数组转 JSON
	jsonBytes, _ := json.MarshalIndent(products, "", "  ")
	fmt.Println("JSON Array:")
	fmt.Println(string(jsonBytes))

	// JSON 转数组
	jsonStr := `[{"id":1,"name":"Item1","price":10.5},{"id":2,"name":"Item2","price":20.5}]`
	var items []Product
	json.Unmarshal([]byte(jsonStr), &items)

	fmt.Println("Parsed items:")
	for _, item := range items {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", item.ID, item.Name, item.Price)
	}
}

func jsonFileOperations() {
	// ===== JSON 文件操作 =====
	// Java: mapper.writeValue(new File("data.json"), person);
	
	filename := "data.json"

	// 写入 JSON 文件
	person := Person{
		Name:  "Eve",
		Age:   28,
		Email: "eve@example.com",
	}

	jsonBytes, _ := json.MarshalIndent(person, "", "  ")
	ioutil.WriteFile(filename, jsonBytes, 0644)
	fmt.Println("JSON written to file:", filename)

	// 读取 JSON 文件
	content, _ := ioutil.ReadFile(filename)
	
	var loadedPerson Person
	json.Unmarshal(content, &loadedPerson)
	
	fmt.Printf("Loaded: Name=%s, Age=%d, Email=%s\n", 
		loadedPerson.Name, loadedPerson.Age, loadedPerson.Email)

	// 清理
	os.Remove(filename)
}

func main() {
	fmt.Println("=== 创建和写入文件 ===")
	createAndWriteFile()

	fmt.Println("\n=== 读取文件 ===")
	readFile()

	fmt.Println("\n=== 追加到文件 ===")
	appendToFile()

	fmt.Println("\n=== 逐行读取 ===")
	readFileLineByLine()

	fmt.Println("\n=== 文件信息 ===")
	fileInfo()

	fmt.Println("\n=== 文件操作 ===")
	deleteFile()

	fmt.Println("\n=== 目录操作 ===")
	directoryOperations()

	fmt.Println("\n=== 路径操作 ===")
	filePathOperations()

	fmt.Println("\n=== JSON 序列化 ===")
	jsonMarshaling()

	fmt.Println("\n=== JSON 反序列化 ===")
	jsonUnmarshaling()

	fmt.Println("\n=== 动态 JSON ===")
	jsonWithDynamicData()

	fmt.Println("\n=== JSON 数组 ===")
	jsonArrayHandling()

	fmt.Println("\n=== JSON 文件操作 ===")
	jsonFileOperations()

	// 清理测试文件
	os.Remove("output.txt")
}
