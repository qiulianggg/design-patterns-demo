# 抽象工厂模式（Abstract Factory）

## 意图

> 提供一个接口，用于创建**一系列相关或相互依赖的对象（产品族）**，
> 而无需指定它们的具体类。

## 解决什么问题

当系统需要创建**成套、需要互相搭配**的对象时（比如同一套 UI 风格里的按钮、复选框、菜单必须风格一致），
抽象工厂保证「一次选定一个工厂，拿到的整族产品都是配套的」，避免出现「Windows 按钮 + macOS 复选框」这种错配。

## 工厂方法 vs 抽象工厂

- **工厂方法**：创建**一个**产品（一个工厂方法）。
- **抽象工厂**：创建**一族**产品（多个工厂方法组成一个工厂接口）。抽象工厂通常由多个工厂方法组成。

## 角色

| 角色 | 说明 | 本例对应 |
|------|------|---------|
| AbstractFactory | 声明创建各类产品的接口 | `GUIFactory` |
| ConcreteFactory | 生产同一产品族的具体产品 | `WinFactory`、`MacFactory` |
| AbstractProduct | 每类产品的接口 | `Button`、`Checkbox` |
| ConcreteProduct | 具体产品 | `WinButton`、`MacCheckbox` … |
| Client | 只依赖抽象工厂与抽象产品 | `renderUI` |

## 结构

```
        GUIFactory(接口)
      ┌───────┴────────┐
  WinFactory        MacFactory
   │    │             │    │
CreateButton      CreateButton
CreateCheckbox    CreateCheckbox
   ▼    ▼             ▼    ▼
WinButton WinCheckbox MacButton MacCheckbox
   （Windows 产品族）   （macOS 产品族）
```

## Go 惯用写法

Go 里用**接口**天然表达抽象工厂：具体工厂是空结构体（无状态），实现 `GUIFactory` 即可。
选择工厂的逻辑常放在一个构造函数里：

```go
func NewFactory(os string) GUIFactory {
    switch os {
    case "windows": return WinFactory{}
    case "mac":     return MacFactory{}
    }
    return nil
}
```

客户端 `renderUI(f GUIFactory)` 面向接口编程，替换整个产品族只需换一个工厂。

## 适用场景

- 系统要独立于产品的创建、组合与表示（换肤、跨平台 UI）；
- 需要强约束「一族产品必须配套使用」；
- 数据库/云厂商适配：一套接口，多套实现族（如 AWS 族、GCP 族的存储+队列+计算）。

## 优点

- 保证产品族的一致性；
- 隔离具体类，切换产品族只改一处；
- 符合开闭原则（新增产品族）。

## 缺点

- **新增一类产品**（给接口加方法）会牵连所有工厂——对「产品种类」不友好，对「产品族」友好；
- 类/接口数量多，结构较重。

## 运行

```bash
go run ./creational/abstract_factory
```
