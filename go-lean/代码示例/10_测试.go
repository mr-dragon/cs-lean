package main

import (
	"fmt"
	"log"
	"testing"
)

// ===== 单元测试 =====
// Java: @Test public void testAdd() { assertEquals(4, add(2, 2)); }

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// ===== 基本测试函数 =====
// 注意：Go 测试函数必须以 Test 开头，接收 *testing.T 参数
// 文件必须以 _test.go 结尾

// Java: @Test
//       assertEquals(4, add(2, 2));
func TestAdd(t *testing.T) {
	result := add(2, 2)
	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(5, 3)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestMultiply(t *testing.T) {
	result := multiply(3, 4)
	if result != 12 {
		t.Errorf("Expected 12, got %d", result)
	}
}

// ===== 表驱动测试 =====
// Java: @ParameterizedTest @ValueSource(ints = {1, 2, 3})
func TestAddMultiple(t *testing.T) {
	// Java: @ParameterizedTest
	//       @CsvSource({"2, 2, 4", "3, 5, 8", "0, 0, 0"})
	cases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"2+2=4", 2, 2, 4},
		{"3+5=8", 3, 5, 8},
		{"0+0=0", 0, 0, 0},
		{"-1+1=0", -1, 1, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := add(c.a, c.b)
			if result != c.expected {
				t.Errorf("add(%d, %d) = %d, want %d", c.a, c.b, result, c.expected)
			}
		})
	}
}

// ===== 错误测试 =====
// Java: assertThrows(ArithmeticException.class, () -> { divide(10, 0); });
func TestDivideByZero(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Errorf("Expected error for division by zero, got nil")
	}
}

func TestDivideValid(t *testing.T) {
	result, err := divide(10, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

// ===== 基准测试 =====
// Java: JMH (Java Microbenchmark Harness)
// 运行: go test -bench=.

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(2, 2)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiply(3, 4)
	}
}

// ===== 子基准测试 =====
func BenchmarkMultipleOperations(b *testing.B) {
	b.Run("Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			add(5, 3)
		}
	})

	b.Run("Multiply", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			multiply(5, 3)
		}
	})
}

// ===== 模糊测试（Fuzz Testing）=====
// Go 1.18+ 支持，Java 没有直接对应

func FuzzAdd(f *testing.F) {
	f.Add(1, 2)
	f.Add(-1, 1)
	f.Add(0, 0)

	f.Fuzz(func(t *testing.T, a, b int) {
		result := add(a, b)
		// 属性：a + b = result
		if result != a+b {
			t.Errorf("add(%d, %d) = %d, want %d", a, b, result, a+b)
		}
	})
}

// ===== 测试设置和清理 =====
// Java: @Before, @After 或 setUp(), tearDown()

func setupTest(t *testing.T) func() {
	// Java: @Before
	fmt.Println("Test setup")

	return func() {
		// Java: @After
		fmt.Println("Test cleanup")
	}
}

func TestWithSetupCleanup(t *testing.T) {
	cleanup := setupTest(t)
	defer cleanup()

	result := add(1, 2)
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

// ===== 测试工具函数 =====
// 辅助函数用于共通的测试逻辑

func assertInt(t *testing.T, got, want int) {
	// Java: assertEquals(want, got);
	t.Helper()  // 标记为辅助函数，错误行号指向调用者
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestWithHelper(t *testing.T) {
	assertInt(t, add(2, 3), 5)
	assertInt(t, multiply(4, 5), 20)
}

// ===== 跳过测试 =====
// Java: @Ignore 或使用条件测试

func TestSkipped(t *testing.T) {
	if testing.Short() {
		// Java: @Ignore("Skip in short mode")
		t.Skip("Skipping long-running test in short mode")
	}
	fmt.Println("This test runs in long mode")
}

// ===== 并发测试 =====
// Java: 使用多线程进行并发测试

func counter(results chan int) {
	for i := 0; i < 10; i++ {
		results <- i
	}
	close(results)
}

func TestConcurrency(t *testing.T) {
	results := make(chan int)
	go counter(results)

	count := 0
	for range results {
		count++
	}

	if count != 10 {
		t.Errorf("Expected 10 results, got %d", count)
	}
}

func main() {
	fmt.Println("=== Go 测试框架 ===")
	fmt.Println(`
Go 的测试框架使用标准库 testing 包，非常简洁高效。

文件命名：xxx_test.go
测试函数命名：TestXxx(t *testing.T)
基准测试命名：BenchmarkXxx(b *testing.B)

运行测试：
  go test              # 运行当前包的所有测试
  go test -v           # 详细输出
  go test -run TestAdd # 运行特定测试
  go test -bench=.     # 运行基准测试
  go test -cover       # 显示覆盖率
  go test -coverprofile=coverage.out  # 生成覆盖率文件
  go tool cover -html=coverage.out    # 查看覆盖率报告

基本测试函数签名：
  func TestXxx(t *testing.T) {
      if condition {
          t.Errorf("error message: %v", value)
          // 或 t.Fatalf("fatal error") - 停止执行
          // 或 t.Fail() - 仅标记失败
          // 或 t.FailNow() - 立即停止
      }
  }

基准测试函数签名：
  func BenchmarkXxx(b *testing.B) {
      for i := 0; i < b.N; i++ {
          // 要测试的代码
      }
  }

与 Java 对比：
┌──────────────────────┬──────────────────────────┬──────────────────────┐
│       特性           │        Java (JUnit)       │    Go (testing)       │
├──────────────────────┼──────────────────────────┼──────────────────────┤
│ 测试注解             │ @Test                    │ 函数命名 Test*       │
│ 断言                 │ assertEquals()           │ if + t.Errorf()      │
│ 跳过测试             │ @Ignore                  │ t.Skip()             │
│ 设置/清理            │ @Before/@After           │ setupTest() + defer  │
│ 参数化测试           │ @ParameterizedTest       │ 表驱动测试           │
│ 基准测试             │ JMH                      │ testing.B (内置)     │
│ 测试框架             │ JUnit, TestNG            │ testing (标准库)     │
│ 模拟对象             │ Mockito, EasyMock        │ 无，手写 Mock 实现   │
│ 并发测试             │ 手写线程                 │ goroutine + channel  │
└──────────────────────┴──────────────────────────┴──────────────────────┘

表驱动测试示例：
  func TestAddMultiple(t *testing.T) {
      tests := []struct {
          name     string
          a, b     int
          expected int
      }{
          {"case1", 1, 1, 2},
          {"case2", 2, 3, 5},
      }

      for _, tt := range tests {
          t.Run(tt.name, func(t *testing.T) {
              got := add(tt.a, tt.b)
              if got != tt.expected {
                  t.Errorf("got %d, want %d", got, tt.expected)
              }
          })
      }
  }

基准测试示例：
  func BenchmarkAdd(b *testing.B) {
      for i := 0; i < b.N; i++ {
          add(2, 3)
      }
  }
  
  运行：go test -bench=BenchmarkAdd -benchmem

模糊测试示例（Go 1.18+）：
  func FuzzAdd(f *testing.F) {
      f.Add(1, 2)
      f.Fuzz(func(t *testing.T, a, b int) {
          result := add(a, b)
          if result != a+b {
              t.Errorf("property violated")
          }
      })
  }

常见断言：
  if condition != expected {
      t.Errorf("got %v, want %v", got, expected)
  }
  
  if err != nil {
      t.Fatalf("unexpected error: %v", err)
  }
  
  if !bytes.Equal(got, want) {
      t.Errorf("output mismatch")
  }
`)

	log.Println("详见本文件中的具体测试函数实现")
}
