package main

import (
	"fmt"
	"sync"
	"time"
)

// ===== 并发基础：Goroutine =====
// Java: Thread t = new Thread(() -> { /* 代码 */ }); t.start();

func helloWorld() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, World!", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func goroutineDemo() {
	// 单个 goroutine
	// Java: new Thread(this::helloWorld).start();
	go helloWorld()

	// 主线程需要等待 goroutine 完成
	time.Sleep(600 * time.Millisecond)
	fmt.Println("Main finished")
}

// ===== 通道（Channel）=====
// Java: BlockingQueue<Integer> queue = new LinkedBlockingQueue<>();

func channelDemo() {
	// 创建一个整数通道
	// Java: 无直接对应
	messages := make(chan string)

	// 在 goroutine 中发送数据
	go func() {
		messages <- "Hello from goroutine"
		time.Sleep(100 * time.Millisecond)
		messages <- "Second message"
	}()

	// 接收数据
	msg1 := <-messages
	fmt.Println("Received:", msg1)

	msg2 := <-messages
	fmt.Println("Received:", msg2)
}

// ===== 缓冲通道 =====
// Java: BlockingQueue<Integer> queue = new LinkedBlockingQueue<>(5);

func bufferedChannelDemo() {
	// 创建容量为 2 的缓冲通道
	messages := make(chan string, 2)

	// 可以在没有接收者的情况下发送 2 个消息
	messages <- "Message 1"
	messages <- "Message 2"

	// 现在接收
	fmt.Println("Received:", <-messages)
	fmt.Println("Received:", <-messages)
}

// ===== 通道关闭和检查 =====
// Java: queue.put(null)  或 使用特殊标记

func sendMessages(messages chan<- string) {
	// 单向通道：只能发送（send-only）
	messages <- "Message 1"
	messages <- "Message 2"
	close(messages)  // 关闭通道
}

func receiveMessages(messages <-chan string) {
	// 单向通道：只能接收（receive-only）
	for msg := range messages {
		fmt.Println("Received:", msg)
	}
}

func channelClosingDemo() {
	messages := make(chan string)
	go sendMessages(messages)
	receiveMessages(messages)

	// 检查通道是否还开放
	messages2 := make(chan int, 1)
	messages2 <- 42
	close(messages2)

	// Java: 无对应
	val, ok := <-messages2
	fmt.Println("Value:", val, "Channel open:", ok)

	// 从关闭的通道接收会得到零值
	val2, ok := <-messages2
	fmt.Println("Value:", val2, "Channel open:", ok)
}

// ===== Select 语句 =====
// Java: 类似于监听多个 BlockingQueue（虽然没有直接等价）

func selectDemo() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "from channel1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "from channel2"
	}()

	// 等待第一个响应
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("Got from channel1:", msg1)
		case msg2 := <-channel2:
			fmt.Println("Got from channel2:", msg2)
		}
	}
}

// ===== 超时处理 =====
// Java: future.get(timeout, unit);

func timeoutDemo() {
	result := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		result <- "Done"
	}()

	// 设置超时
	select {
	case res := <-result:
		fmt.Println("Result:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
	}
}

// ===== Worker 池模式 =====
// Java: ExecutorService executor = Executors.newFixedThreadPool(3);

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2
	}
}

func workerPoolDemo() {
	numWorkers := 3
	numJobs := 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 创建 worker pool
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// 发送任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	for r := 0; r < numJobs; r++ {
		fmt.Println("Result:", <-results)
	}
}

// ===== 同步原语：互斥锁（Mutex）=====
// Java: synchronized  或  Lock

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func mutexDemo() {
	counter := &Counter{}

	// 并发增加计数器
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()  // 等待所有 goroutine 完成
	fmt.Println("Final count:", counter.Value())
}

// ===== 同步原语：WaitGroup =====
// Java: CountDownLatch latch = new CountDownLatch(3);

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()  // 任务完成，计数器减 1
	fmt.Printf("Task %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Task %d done\n", id)
}

func waitGroupDemo() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)  // 增加计数器
		go task(i, &wg)
	}

	wg.Wait()  // 等待计数器归零
	fmt.Println("All tasks completed")
}

// ===== 同步原语：Once =====
// Java: 使用 synchronized 或 double-checked locking

var (
	instance string
	once     sync.Once
)

func getInstance() string {
	once.Do(func() {  // 确保代码块只执行一次
		instance = "Singleton instance"
	})
	return instance
}

func onceDemo() {
	for i := 0; i < 5; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d: %s\n", id, getInstance())
		}(i)
	}
	time.Sleep(time.Second)
}

// ===== 同步原语：RWMutex（读写锁）=====
// Java: ReadWriteLock

type DataStore struct {
	mu   sync.RWMutex
	data map[string]string
}

func (ds *DataStore) Read(key string) string {
	ds.mu.RLock()      // 获取读锁
	defer ds.mu.RUnlock()
	return ds.data[key]
}

func (ds *DataStore) Write(key, value string) {
	ds.mu.Lock()       // 获取写锁
	defer ds.mu.Unlock()
	ds.data[key] = value
}

func rwMutexDemo() {
	ds := &DataStore{data: make(map[string]string)}

	// 并发写入
	for i := 0; i < 3; i++ {
		go func(id int) {
			ds.Write(fmt.Sprintf("key%d", id), fmt.Sprintf("value%d", id))
		}(i)
	}

	time.Sleep(100 * time.Millisecond)

	// 并发读取
	for i := 0; i < 3; i++ {
		go func(id int) {
			val := ds.Read(fmt.Sprintf("key%d", id))
			fmt.Printf("Read: %s\n", val)
		}(i)
	}

	time.Sleep(time.Second)
}

// ===== 通道中的 select 与 default =====
func selectWithDefaultDemo() {
	ch := make(chan string, 1)
	ch <- "Hello"

	// 非阻塞接收
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message available")
	}

	// 再尝试接收（通道已空）
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("Channel is empty")
	}
}

// ===== 管道模式（Pipeline）=====
// Java: 使用多个 BlockingQueue 或 CompletableFuture chain

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func pipelineDemo() {
	// 1 -> 2 -> 3 平方后的结果：1, 4, 9
	for result := range square(generate(1, 2, 3)) {
		fmt.Println("Pipeline result:", result)
	}
}

func main() {
	fmt.Println("=== Goroutine 基础 ===")
	goroutineDemo()

	fmt.Println("\n=== 通道基础 ===")
	channelDemo()

	fmt.Println("\n=== 缓冲通道 ===")
	bufferedChannelDemo()

	fmt.Println("\n=== 通道关闭 ===")
	channelClosingDemo()

	fmt.Println("\n=== Select 语句 ===")
	selectDemo()

	fmt.Println("\n=== 超时处理 ===")
	timeoutDemo()

	fmt.Println("\n=== Worker 池模式 ===")
	workerPoolDemo()

	fmt.Println("\n=== Mutex ===")
	mutexDemo()

	fmt.Println("\n=== WaitGroup ===")
	waitGroupDemo()

	fmt.Println("\n=== Once ===")
	onceDemo()

	fmt.Println("\n=== RWMutex ===")
	rwMutexDemo()

	fmt.Println("\n=== Select with Default ===")
	selectWithDefaultDemo()

	fmt.Println("\n=== Pipeline 模式 ===")
	pipelineDemo()
}
