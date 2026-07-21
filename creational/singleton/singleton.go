package main

import (
	"fmt"
	"sync"
)

// config 是要被单例化的对象：全局唯一的配置中心。
type config struct {
	settings map[string]string
}

var (
	instance *config
	once     sync.Once
)

// GetConfig 返回全局唯一的 config 实例。
//
// 使用 sync.Once 保证初始化逻辑无论被多少 goroutine 并发调用，
// 都只会执行一次，这是 Go 中实现线程安全懒加载单例的惯用方式。
func GetConfig() *config {
	once.Do(func() {
		fmt.Println("[初始化] 只会打印一次：正在加载配置...")
		instance = &config{settings: map[string]string{
			"host": "127.0.0.1",
			"port": "8080",
		}}
	})
	return instance
}

func (c *config) Get(key string) string { return c.settings[key] }
func (c *config) Set(key, val string)    { c.settings[key] = val }

func main() {
	// 并发获取，验证只初始化一次。
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = GetConfig()
		}()
	}
	wg.Wait()

	c1 := GetConfig()
	c2 := GetConfig()

	fmt.Printf("c1 == c2 ? %v (地址: %p vs %p)\n", c1 == c2, c1, c2)

	c1.Set("port", "9090")
	// 通过 c2 读取，能看到 c1 的修改，证明是同一个实例。
	fmt.Println("通过 c2 读取 port:", c2.Get("port"))
}
