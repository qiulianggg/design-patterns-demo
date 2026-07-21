# 策略模式（Strategy）

## 意图

> 定义一系列算法，把它们**一个个封装**起来，并且使它们可以**相互替换**。
> 策略模式让算法的变化独立于使用它的客户端。

## 解决什么问题

同一件事有多种做法（多种折扣、多种排序、多种支付、多种压缩算法），且需要在运行时选择或切换。
如果写成 `if type == A { ... } else if type == B { ... }`，新增算法就要改这段逻辑，违反开闭原则。

策略模式把每种算法封装成独立对象，让它们实现同一接口。上下文持有一个策略引用，
把工作委托给它；换算法只需换一个策略对象。

## 角色

| 角色 | 说明 | 本例对应 |
|------|------|---------|
| Strategy | 算法接口 | `DiscountStrategy` |
| ConcreteStrategy | 各具体算法 | `NoDiscount`、`PercentageDiscount`、`ThresholdDiscount` |
| Context | 持有并调用策略 | `Cart` |

## 结构

```
Cart(Context).discount ──► DiscountStrategy(接口)
   Checkout() ─委托─► discount.Apply(amount)
                        ├─ NoDiscount
                        ├─ PercentageDiscount
                        └─ ThresholdDiscount
```

## Go 惯用写法

策略是 Go 里**最自然、最常用**的模式之一，因为它天然契合接口和一等函数：

### 1. 接口策略（本例）——适合策略有状态/多方法

### 2. 函数策略（更 Go，最常见）

如果策略只有一个方法，直接用**函数类型**，无需定义结构体：

```go
type DiscountFunc func(float64) float64

func percentage(p float64) DiscountFunc {
    return func(a float64) float64 { return a * (1 - p) }
}
// Cart 持有一个 DiscountFunc 字段即可
```

标准库里 `sort.Slice(s, less func(i,j int) bool)`、`http.HandlerFunc`、
`sort.Interface` 都是策略思想。传一个函数进去 = 注入一个策略。

## 策略 vs 相近模式

- **策略 vs 状态**：结构相同；策略由客户端选择、互不感知；状态会自我切换、构成状态机。
- **策略 vs 模板方法**：策略用**组合+委托**在运行时换整个算法；模板方法用**继承/内嵌**在编译期让子类覆盖算法的某些步骤。
- **策略 vs 桥接**：策略是一个维度换算法；桥接是两个维度做结构解耦。

## 适用场景

- 一个功能有多种可互换的算法/行为；
- 想在运行时选择算法；
- 想消除大量算法选择的条件分支。

## 优点

- 算法独立封装、可自由替换，符合开闭原则；
- 消除条件分支；算法可独立测试复用。

## 缺点

- 客户端需了解不同策略的差异才能选择；
- 策略数量多时对象/类型增多（Go 用函数可缓解）。

## 运行

```bash
go run ./behavioral/strategy
```
