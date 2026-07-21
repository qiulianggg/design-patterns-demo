# 命令模式（Command）

## 意图

> 将一个**请求封装成一个对象**，从而让你可以用不同的请求对客户进行参数化，支持请求的**排队、记录日志、撤销**等操作。

## 解决什么问题

「调用某个方法」这个动作本身，通常是写死在代码里的、一次性的。但有时我们希望把「做某件事」当作**一等对象**来传递、存储、延迟执行、排队、记录、甚至撤销。

命令模式把「谁来做（Receiver）」「做什么（方法）」「参数」一起打包进一个命令对象，
调用者（Invoker）只管触发命令，不关心细节。这样就能实现撤销/重做、宏命令、任务队列等。

## 角色

| 角色 | 说明 | 本例对应 |
|------|------|---------|
| Command | 命令接口，声明 `Execute`（可含 `Undo`） | `Command` |
| ConcreteCommand | 绑定接收者与动作 | `LightOnCommand`、`LightOffCommand` |
| Receiver | 真正执行操作的对象 | `Light` |
| Invoker | 触发命令、可保存历史 | `RemoteControl` |
| Client | 创建命令并配置接收者 | `main` |

## 结构

```
Client 创建 ► LightOnCommand{light}
Invoker(RemoteControl).Press(cmd)
      │ cmd.Execute() ──► Receiver(Light).TurnOn()
      └ 记录 history 以支持 Undo()
```

## Go 惯用写法

Go 里命令可以是实现接口的结构体（本例，便于携带撤销、状态）；
若不需要撤销等能力，直接用**闭包/函数值**就是最轻量的「命令」：

```go
type Command func()
queue := []Command{
    func() { light.TurnOn() },
    func() { light.TurnOff() },
}
for _, cmd := range queue { cmd() }
```

任务队列、`time.AfterFunc`、goroutine 里跑的 `func()`，都是命令思想的体现。
需要撤销/重做、序列化、日志重放时，才升级为完整的命令对象。

## 适用场景

- 需要撤销/重做（编辑器、绘图工具）；
- 请求需要排队、调度、异步执行（任务队列、job）；
- 需要记录操作日志以便崩溃后重放；
- 宏命令：把多个命令组合成一个。

## 优点

- 解耦「触发操作的对象」与「执行操作的对象」；
- 命令是一等对象，可组合、排队、记录、撤销；
- 新增命令符合开闭原则。

## 缺点

- 每个操作一个命令类，数量可能膨胀（Go 用闭包可缓解）。

## 运行

```bash
go run ./behavioral/command
```
