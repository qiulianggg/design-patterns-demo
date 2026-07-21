package main

import "fmt"

// 代理模式：为另一个对象提供替身，以控制对它的访问。
// 本例是「虚拟代理 + 保护代理」的结合：
//   - 延迟加载（真正需要时才创建昂贵的真实对象）
//   - 访问控制（校验权限、记录日志、缓存）

// Image 是真实主题与代理共同实现的接口。
type Image interface {
	Display() string
}

// RealImage 是真实主题：创建成本高（模拟从磁盘加载大图）。
type RealImage struct {
	filename string
}

func newRealImage(filename string) *RealImage {
	fmt.Printf("  [加载] 从磁盘读取大图 %s（耗时操作）\n", filename)
	return &RealImage{filename: filename}
}

func (r *RealImage) Display() string {
	return "显示图片 " + r.filename
}

// ProxyImage 是代理：持有对真实对象的引用，按需创建并附加额外控制。
type ProxyImage struct {
	filename string
	real     *RealImage // 延迟初始化
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename: filename}
}

func (p *ProxyImage) Display() string {
	// 虚拟代理：第一次 Display 时才真正加载
	if p.real == nil {
		p.real = newRealImage(p.filename)
	}
	// 这里还可以加：权限校验、访问日志、结果缓存等
	fmt.Println("  [代理] 记录一次访问日志")
	return p.real.Display()
}

func main() {
	// 创建代理时不会加载大图（省资源）
	img := NewProxyImage("photo.png")
	fmt.Println("代理已创建，此时尚未加载真实图片")

	// 第一次使用 -> 触发加载
	fmt.Println(img.Display())
	// 第二次使用 -> 复用已加载的真实对象，不再加载
	fmt.Println(img.Display())
}
