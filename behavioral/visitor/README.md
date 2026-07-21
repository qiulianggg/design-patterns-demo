# 访问者模式（Visitor）

## 意图

> 表示一个作用于某对象结构中各元素的操作。它使你可以在**不改变各元素类**的前提下，
> 定义作用于这些元素的**新操作**。

## 解决什么问题

假设你有一组稳定的元素类型（圆、矩形、三角形），需要对它们施加**很多种不同的操作**
（算面积、算周长、导出 XML、渲染、序列化……）。

如果把每个操作都写成元素类的方法，那么每加一个操作，就要改动**所有元素类**——它们会越来越臃肿，且违反单一职责。

访问者模式把「操作」抽出来，集中到一个**访问者对象**里。新增操作 = 新增一个访问者类，元素类**一行都不用改**。

## 双分派（Double Dispatch）——核心机制

难点在于：调用哪个逻辑要**同时取决于两件事**——元素的具体类型 + 访问者的具体类型。
访问者用「两次分派」解决：

1. `shape.Accept(visitor)` —— 由 **shape 的实际类型**决定进入 `Circle.Accept` 还是 `Rectangle.Accept`；
2. 在 `Accept` 内部调用 `visitor.VisitCircle(c)` —— 由 **visitor 的实际类型**决定具体操作。

两次动态分派组合，精确定位到「(某元素类型, 某操作)」的实现。

## 角色

| 角色 | 说明 | 本例对应 |
|------|------|---------|
| Visitor | 为每种元素声明 `VisitXxx` | `Visitor` |
| ConcreteVisitor | 具体操作 | `AreaVisitor`、`XMLExportVisitor` |
| Element | 声明 `Accept(Visitor)` | `Shape` |
| ConcreteElement | 实现 `Accept`，回调对应 `VisitXxx` | `Circle`、`Rectangle` |

## 结构

```
shape.Accept(v)                v.VisitCircle(c)
  由元素类型分派 ──第1次分派──► 由访问者类型分派 ──第2次分派──► 具体逻辑
Circle.Accept ─► v.VisitCircle
Rectangle.Accept ─► v.VisitRectangle
```

## Go 惯用写法

Go 没有方法重载，所以访问者接口必须为每种元素显式声明不同名字的方法（`VisitCircle`、`VisitRectangle`），
这正是 GoF 访问者的标准形态。元素的 `Accept` 里回调对应方法完成双分派。

**Go 的实用替代**：很多时候用 **type switch** 就能达到「集中定义操作」的效果，且更简单：

```go
func Area(s Shape) float64 {
    switch t := s.(type) {
    case *Circle:    return 3.14159 * t.Radius * t.Radius
    case *Rectangle: return t.Width * t.Height
    }
    return 0
}
```

区别：type switch 在**新增元素类型**时要改每个操作函数；而访问者在新增元素类型时要改所有访问者接口/实现。
两者都对「加操作」友好、对「加元素类型」不友好（见下）。元素类型很稳定、操作频繁增加时，才值得上完整访问者模式。

## 「加操作 easy，加元素 hard」

- **加操作**：写一个新访问者即可，不碰元素 —— 容易；
- **加元素类型**：要给 `Visitor` 接口加一个 `VisitXxx`，所有现有访问者都得实现它 —— 困难。

所以访问者适用于**元素结构稳定、操作经常扩展**的场景（如编译器 AST：节点类型固定，遍历/优化/生成等操作不断增加）。

## 适用场景

- 对象结构稳定，但要经常对其定义新操作（编译器 AST、文档对象、报表）；
- 需要对一个复杂对象结构做多种不相关的操作，且不想污染元素类。

## 优点

- 新增操作方便，符合开闭原则（对操作而言）；
- 相关操作集中在访问者里，元素类保持干净；
- 可在访问者中累积跨元素的状态（如本例 `Total`）。

## 缺点

- 新增元素类型代价大（要改所有访问者）；
- 访问者需要访问元素内部，可能破坏封装；
- 结构较重、理解成本高——Go 里常可用 type switch 替代。

## 运行

```bash
go run ./behavioral/visitor
```
