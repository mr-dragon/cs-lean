package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ===== HTTP 客户端 =====
// Java: HttpClient client = HttpClient.newBuilder().build();

func basicHttpGet() {
	// ===== 简单 GET 请求 =====
	// Java: HttpRequest request = HttpRequest.newBuilder().uri(new URI(url)).GET().build();

	url := "https://httpbin.org/get"
	
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应体
	// Java: response.body().asString()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	fmt.Println("Status:", response.Status)
	fmt.Println("Body:", string(body[:100]))  // 只显示前 100 个字符
}

func basicHttpPost() {
	// ===== POST 请求 =====
	// Java: HttpRequest request = HttpRequest.newBuilder()
	//           .uri(new URI(url))
	//           .POST(HttpRequest.BodyPublishers.ofString(body))
	//           .build();

	url := "https://httpbin.org/post"
	data := "name=Alice&age=30"
	contentType := "application/x-www-form-urlencoded"

	response, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("POST Response:", string(body[:100]))
}

func httpWithCustomHeaders() {
	// ===== 自定义请求头 =====
	// Java: request.header("Authorization", "Bearer token");

	url := "https://httpbin.org/headers"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 添加请求头
	// Java: request.headers("User-Agent", "...");
	req.Header.Add("Authorization", "Bearer my-token")
	req.Header.Add("User-Agent", "Go-HTTP-Client/1.0")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response with headers:", string(body[:100]))
}

func httpWithTimeout() {
	// ===== 设置超时 =====
	// Java: HttpClient.newBuilder().connectTimeout(Duration.ofSeconds(10)).build();

	url := "https://httpbin.org/delay/5"

	// 创建客户端并设置超时
	// Java: HttpClient client = HttpClient.newBuilder()
	//           .connectTimeout(Duration.ofSeconds(3))
	//           .build();
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	response, err := client.Get(url)
	if err != nil {
		fmt.Println("Error (expected timeout):", err)
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response:", string(body[:50]))
}

func httpFormData() {
	// ===== 表单数据 =====
	// Java: HttpRequest.BodyPublishers.ofString(encode form data);

	url := "https://httpbin.org/post"

	// 创建表单数据
	// Java: URLEncoder.encode() 或 form = new HashMap<>()
	data := url.Values{}
	data.Set("username", "alice")
	data.Set("password", "secret123")
	data.Set("remember", "true")

	response, err := http.PostForm(url, data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Form POST Response:", string(body[:100]))
}

func httpJsonRequest() {
	// ===== JSON 请求 =====
	// Java: HttpRequest.BodyPublishers.ofString(json);

	url := "https://httpbin.org/post"
	jsonData := `{"name":"Bob","age":25,"email":"bob@example.com"}`

	// 创建请求
	// Java: HttpRequest request = HttpRequest.newBuilder()
	//           .uri(new URI(url))
	//           .header("Content-Type", "application/json")
	//           .POST(HttpRequest.BodyPublishers.ofString(jsonData))
	//           .build();
	req, _ := http.NewRequest("POST", url, strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, _ := client.Do(req)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("JSON POST Response:", string(body[:100]))
}

func httpErrorHandling() {
	// ===== 错误处理 =====
	// Java: try-catch, HttpException

	url := "https://httpbin.org/status/404"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer response.Body.Close()

	// 检查 HTTP 状态码
	// Java: response.statusCode()
	if response.StatusCode != http.StatusOK {
		fmt.Printf("HTTP Error: %s (Status Code: %d)\n", http.StatusText(response.StatusCode), response.StatusCode)
	}
}

// ===== HTTP 服务器 =====
// Java: HttpServer server = HttpServer.create(new InetSocketAddress(8080), 0);

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// ===== 处理请求 =====
	// Java: public void handle(HttpExchange exchange)

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	// ===== 返回 JSON =====
	// Java: response.setContentType("application/json");

	w.Header().Set("Content-Type", "application/json")
	json := `{"message":"Hello, World!","status":"ok"}`
	fmt.Fprint(w, json)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// ===== 处理 POST 请求 =====
	// Java: if ("POST".equals(request.getMethod()))

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed")
		return
	}

	// 解析表单数据
	// Java: request.getParameter("name");
	r.ParseForm()
	name := r.FormValue("name")
	age := r.FormValue("age")

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Name: %s, Age: %s", name, age)
}

func middlewareExample(next http.Handler) http.Handler {
	// ===== 中间件 =====
	// Java: Filter, Spring Interceptor

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 前置处理
		fmt.Println("Request received:", r.Method, r.URL.Path)

		// 调用下一个处理器
		next.ServeHTTP(w, r)

		// 后置处理
		fmt.Println("Response sent")
	})
}

func startServer() {
	// ===== 启动服务器 =====
	// Java: server.start();

	fmt.Println("\n=== HTTP 服务器启动示例 ===")
	fmt.Println(`
import "net/http"

// 注册处理器
http.HandleFunc("/hello", helloHandler)
http.HandleFunc("/json", jsonHandler)
http.HandleFunc("/post", postHandler)

// 启动服务器
// Java: HttpServer server = HttpServer.create(new InetSocketAddress(8080), 0);
//       server.start();
err := http.ListenAndServe(":8080", nil)
if err != nil {
    log.Fatal(err)
}

// 或者使用自定义服务器配置
server := &http.Server{
    Addr:         ":8080",
    Handler:      nil,
    ReadTimeout:  15 * time.Second,
    WriteTimeout: 15 * time.Second,
    IdleTimeout:  60 * time.Second,
}
server.ListenAndServe()

// 关闭服务器
server.Shutdown(context.Background())
`)
}

func restApiExample() {
	fmt.Println("\n=== 简单 REST API 示例 ===")
	fmt.Println(`
package main

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID   int    \`json:"id"\`
    Name string \`json:"name"\`
    Age  int    \`json:"age"\`
}

var users = []User{
    {1, "Alice", 30},
    {2, "Bob", 25},
}

// GET /api/users - 获取所有用户
func getUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// GET /api/users/{id} - 获取单个用户
func getUser(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    for _, user := range users {
        if fmt.Sprintf("%d", user.ID) == id {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(user)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
}

// POST /api/users - 创建用户
func createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)
    user.ID = len(users) + 1
    users = append(users, user)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func main() {
    http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            getUsers(w, r)
        } else if r.Method == http.MethodPost {
            createUser(w, r)
        }
    })
    http.ListenAndServe(":8080", nil)
}
`)
}

func webFrameworksExample() {
	fmt.Println("\n=== Web 框架对比 ===")
	fmt.Println(`
Go Web 框架（类似 Java 的 Spring Boot）：

1. Gin - 轻量级、高性能（推荐）
   import "github.com/gin-gonic/gin"
   
   r := gin.Default()
   r.GET("/hello", func(c *gin.Context) {
       c.JSON(200, gin.H{"message": "Hello"})
   })
   r.Run(":8080")

2. Echo - 简洁强大
   import "github.com/labstack/echo/v4"
   
   e := echo.New()
   e.GET("/hello", func(c echo.Context) error {
       return c.JSON(200, map[string]string{"message": "Hello"})
   })
   e.Start(":8080")

3. 标准库 net/http - 无依赖
   (见上面的示例)

┌─────────────────────┬────────────────────────┬─────────────────────┐
│      特性           │    Java (Spring)        │    Go (Gin/Echo)     │
├─────────────────────┼────────────────────────┼─────────────────────┤
│ 路由注册            │ @RequestMapping         │ router.GET/POST()   │
│ 依赖注入            │ @Autowired              │ 构造函数注入         │
│ 中间件              │ Filter/Interceptor      │ middleware           │
│ 参数绑定            │ @RequestParam/@PathVar  │ c.Param()/Query()   │
│ 请求体解析          │ @RequestBody            │ c.BindJSON()        │
│ 响应JSON            │ return object           │ c.JSON()            │
│ 异常处理            │ @ExceptionHandler       │ middleware/recover  │
│ 跨域处理            │ @CrossOrigin            │ middleware          │
└─────────────────────┴────────────────────────┴─────────────────────┘
`)
}

func main() {
	fmt.Println("=== HTTP 客户端示例 ===")
	fmt.Println("注意：以下示例需要网络连接，演示目的")

	fmt.Println("\n--- 基本 GET 请求 ---")
	fmt.Println(`
response, _ := http.Get("https://httpbin.org/get")
body, _ := ioutil.ReadAll(response.Body)
fmt.Println(string(body))
`)

	fmt.Println("\n--- 基本 POST 请求 ---")
	fmt.Println(`
data := "name=Alice&age=30"
response, _ := http.Post("https://httpbin.org/post", "application/x-www-form-urlencoded", strings.NewReader(data))
body, _ := ioutil.ReadAll(response.Body)
fmt.Println(string(body))
`)

	fmt.Println("\n--- 自定义请求头 ---")
	fmt.Println(`
req, _ := http.NewRequest("GET", url, nil)
req.Header.Add("Authorization", "Bearer token")
client := &http.Client{}
response, _ := client.Do(req)
`)

	fmt.Println("\n--- 超时设置 ---")
	fmt.Println(`
client := &http.Client{Timeout: 5 * time.Second}
response, _ := client.Get(url)
`)

	fmt.Println("\n--- 表单 POST ---")
	fmt.Println(`
data := url.Values{}
data.Set("username", "alice")
response, _ := http.PostForm(url, data)
`)

	fmt.Println("\n--- JSON POST ---")
	fmt.Println(`
jsonData := \`{"name":"Bob","age":25}\`
req, _ := http.NewRequest("POST", url, strings.NewReader(jsonData))
req.Header.Set("Content-Type", "application/json")
client := &http.Client{}
response, _ := client.Do(req)
`)

	fmt.Println("\n--- 错误处理 ---")
	fmt.Println(`
response, err := http.Get(url)
if err != nil {
    fmt.Println("Connection error:", err)
}
if response.StatusCode != http.StatusOK {
    fmt.Printf("HTTP Error: %d", response.StatusCode)
}
`)

	startServer()
	restApiExample()
	webFrameworksExample()
}
