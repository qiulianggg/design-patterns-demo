# 观察者模式（Observer）

## 意图

> 定义对象间的一种**一对多**依赖关系，当一个对象（主题）的状态发生改变时，
> 所有依赖于它的对象（观察者）都会得到通知并**自动更新**。

## 解决什么问题

当「一个对象的变化需要联动更新其他若干对象」，且你**不想让主题写死它有哪些依赖者**时，
观察者提供了一种松耦合的发布-订阅机制：主题只维护一个观察者列表，状态变化时逐个通知，
完全不关心观察者具体是谁、要做什么。这就是「发布-订阅」的基础。

## 角色

| 角色 | 说明 | 本例对应 |
|------|------|---------|
| Subject | 主题接口：订阅/退订/通知 | `Subject` |
| ConcreteSubject | 持有状态，状态变则通知 | `WeatherStation` |
| Observer | 观察者接口：`Update` | `Observer` |
| ConcreteObserver | 收到通知后更新自己 | `PhoneDisplay`、`WindowDisplay` |

## 结构

```
WeatherStation(Subject)  observers: [Phone, Window]
    SetTemperature() ──► Notify() ──► 逐个 o.Update(temp)
                                   ├─► PhoneDisplay
                                   └─► WindowDisplay
```

## Go 惯用写法

两种主流实现：

### 1. 接口列表（本例）

主题持有 `[]Observer`，`Notify` 遍历调用 `Update`。经典、直观。

### 2. channel（更 Go）

Go 里「发布-订阅」常用 **channel** 表达：主题给每个订阅者一个 channel，发送事件即通知。
结合 goroutine 可实现异步、并发的通知：

```go
type Broker struct{ subs []chan float64 }
func (b *Broker) Subscribe() <-chan float64 { ch := make(chan float64, 1); b.subs = append(b.subs, ch); return ch }
func (b *Broker) Publish(v float64) { for _, ch := range b.subs { ch <- v } }
```

**并发注意**：多 goroutine 下增删观察者、通知需要加锁（`sync.RWMutex`）；同步通知时，某个观察者阻塞会拖慢整体，考虑异步/带缓冲 channel。

## 观察者 vs 相近模式

- **观察者 vs 中介者**：观察者是单向的一对多广播；中介者是集中协调多对多交互。
- **观察者 vs Pub/Sub**：Pub/Sub 通常多一层「消息代理/主题」解耦发布者与订阅者，观察者中主题直接持有观察者引用。

## 适用场景

- 一个变化需要通知/联动多个对象，且对象数量、种类动态变化；
- 事件系统、GUI 数据绑定、MVC 中 Model 通知 View；
- 消息通知、监控告警。

## 优点

- 主题与观察者松耦合，符合开闭原则；
- 运行时动态订阅/退订；支持广播通信。

## 缺点

- 通知顺序不确定；观察者过多时通知开销大；
- 若观察者又反过来改主题，可能引发循环通知；
- 同步通知下异常/阻塞会影响主题。

## 运行

```bash
go run ./behavioral/observer
```
