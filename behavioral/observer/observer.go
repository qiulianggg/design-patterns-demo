package main

import "fmt"

// 观察者模式：定义对象间一对多的依赖，当一个对象(主题)状态改变时，
// 所有依赖它的对象(观察者)都会收到通知并自动更新。
// 本例：气象站发布温度，多个显示面板订阅并自动更新。

// Observer 是观察者接口。
type Observer interface {
	Update(temp float64)
	Name() string
}

// Subject 是主题接口：管理订阅关系并广播通知。
type Subject interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Notify()
}

// WeatherStation 是具体主题。
type WeatherStation struct {
	observers []Observer
	temp      float64
}

func (s *WeatherStation) Subscribe(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *WeatherStation) Unsubscribe(target Observer) {
	for i, o := range s.observers {
		if o == target {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}

func (s *WeatherStation) Notify() {
	for _, o := range s.observers {
		o.Update(s.temp)
	}
}

// SetTemperature 状态变化 -> 自动通知所有观察者。
func (s *WeatherStation) SetTemperature(t float64) {
	fmt.Printf("\n气象站: 温度更新为 %.1f°C\n", t)
	s.temp = t
	s.Notify()
}

// PhoneDisplay 具体观察者。
type PhoneDisplay struct{ name string }

func (p *PhoneDisplay) Name() string { return p.name }
func (p *PhoneDisplay) Update(temp float64) {
	fmt.Printf("  📱 %s 显示: 当前 %.1f°C\n", p.name, temp)
}

// WindowDisplay 具体观察者。
type WindowDisplay struct{ name string }

func (w *WindowDisplay) Name() string { return w.name }
func (w *WindowDisplay) Update(temp float64) {
	fmt.Printf("  🖥️  %s 显示: 当前 %.1f°C\n", w.name, temp)
}

func main() {
	station := &WeatherStation{}

	phone := &PhoneDisplay{name: "手机App"}
	window := &WindowDisplay{name: "桌面窗口"}
	station.Subscribe(phone)
	station.Subscribe(window)

	station.SetTemperature(25.5) // 两个观察者都会更新

	station.Unsubscribe(window) // 桌面窗口取消订阅
	station.SetTemperature(30.0) // 只有手机会更新
}
