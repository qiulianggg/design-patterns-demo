package main

import "fmt"

// 享元模式：通过共享，避免大量相似对象重复占用内存。
// 把对象状态拆成：
//   - 内部状态(intrinsic)：可共享、不随环境变化 —— 本例的「字符字形/字体」
//   - 外部状态(extrinsic)：不可共享、由客户端传入 —— 本例的「位置」
// 相同内部状态的对象只创建一份，反复复用。

// Glyph 是享元：只保存可共享的内部状态（字符 + 字体）。
type Glyph struct {
	char rune
	font string
}

// Render 接收外部状态（坐标），不把它存进对象里。
func (g *Glyph) Render(x, y int) string {
	return fmt.Sprintf("在(%d,%d)绘制字符'%c' [字体:%s]", x, y, g.char, g.font)
}

// GlyphFactory 是享元工厂：维护一个池，保证相同 (char,font) 只有一个 Glyph 实例。
type GlyphFactory struct {
	pool map[string]*Glyph
}

func NewGlyphFactory() *GlyphFactory {
	return &GlyphFactory{pool: make(map[string]*Glyph)}
}

func (f *GlyphFactory) Get(char rune, font string) *Glyph {
	key := fmt.Sprintf("%c-%s", char, font)
	if g, ok := f.pool[key]; ok {
		return g // 命中缓存，复用
	}
	g := &Glyph{char: char, font: font}
	f.pool[key] = g
	fmt.Printf("  [新建享元] %s\n", key)
	return g
}

func (f *GlyphFactory) Count() int { return len(f.pool) }

func main() {
	factory := NewGlyphFactory()

	// 渲染字符串 "hello"，其中 'l' 出现两次、共享同一个享元
	text := "hello"
	font := "Arial"
	for i, ch := range text {
		g := factory.Get(ch, font)
		fmt.Println(g.Render(i*10, 0)) // 位置是外部状态，运行时传入
	}

	fmt.Printf("共渲染 %d 个字符，但只创建了 %d 个享元对象\n", len(text), factory.Count())
}
