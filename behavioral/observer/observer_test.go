package main

import "testing"

// 测试用观察者：记录收到的温度。
type mockObserver struct {
	name string
	last float64
	hits int
}

func (m *mockObserver) Name() string { return m.name }
func (m *mockObserver) Update(temp float64) {
	m.last = temp
	m.hits++
}

// 所有订阅者都应收到通知。
func TestObserverNotify(t *testing.T) {
	station := &WeatherStation{}
	o1 := &mockObserver{name: "o1"}
	o2 := &mockObserver{name: "o2"}
	station.Subscribe(o1)
	station.Subscribe(o2)

	station.SetTemperature(25.5)
	if o1.last != 25.5 || o2.last != 25.5 {
		t.Fatalf("观察者未收到正确温度: o1=%v o2=%v", o1.last, o2.last)
	}
}

// 退订后不应再收到通知。
func TestObserverUnsubscribe(t *testing.T) {
	station := &WeatherStation{}
	o := &mockObserver{name: "o"}
	station.Subscribe(o)
	station.SetTemperature(10)
	station.Unsubscribe(o)
	station.SetTemperature(20)

	if o.hits != 1 {
		t.Fatalf("退订后应只收到 1 次通知, 实际 %d", o.hits)
	}
}
