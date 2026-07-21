package main

import "fmt"

// 外观模式：为一组复杂的子系统提供一个简单、统一的高层入口。
// 本例：一键「启动家庭影院」背后其实要操作投影仪、音响、播放器等多个子系统。

// ---------- 子系统（复杂、各自独立）----------

type Projector struct{}

func (Projector) On() string       { return "投影仪开机" }
func (Projector) WideScreen() string { return "投影仪切换宽屏模式" }

type Amplifier struct{}

func (Amplifier) On() string            { return "功放开机" }
func (Amplifier) SetVolume(v int) string { return fmt.Sprintf("音量设为 %d", v) }

type Player struct{}

func (Player) On() string              { return "播放器开机" }
func (Player) Play(movie string) string { return "开始播放《" + movie + "》" }

// ---------- 外观：把上面一堆操作封装成简单接口 ----------

type HomeTheaterFacade struct {
	projector Projector
	amplifier Amplifier
	player    Player
}

func NewHomeTheater() *HomeTheaterFacade { return &HomeTheaterFacade{} }

// WatchMovie 把「看电影」需要的一连串子系统调用封装成一个方法。
func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("=== 准备观影 ===")
	steps := []string{
		h.projector.On(),
		h.projector.WideScreen(),
		h.amplifier.On(),
		h.amplifier.SetVolume(8),
		h.player.On(),
		h.player.Play(movie),
	}
	for _, s := range steps {
		fmt.Println("  ->", s)
	}
}

func main() {
	// 客户端只需一个简单调用，无需了解子系统细节与调用顺序
	theater := NewHomeTheater()
	theater.WatchMovie("盗梦空间")
}
