# 工厂方法模式（Factory Method）

## 意图

> 定义一个创建对象的接口，但让**子类（具体工厂）决定**实例化哪一个具体产品类。
> 工厂方法使一个类的实例化延迟到其子类。

## 解决什么问题

当一段业务逻辑需要创建某种产品，但**具体创建哪个产品应当可扩展、可替换**时，
如果直接写 `new ConcreteProduct{}`，业务代码就和具体类型强耦合了，新增产品类型就要改动业务代码。

工厂方法把「创建产品」这一步抽象成一个方法，业务逻辑只依赖抽象产品接口，
新增产品只需新增一个工厂，符合**开闭原则**。

## 角色

| 角色 | 说明 | 本例对应 |
|------|------|---------|
| Product（抽象产品） | 产品的公共接口 | `Transport` |
| ConcreteProduct（具体产品） | 接口的具体实现 | `Truck`、`Ship` |
| Creator（抽象工厂/创建者） | 声明工厂方法，含使用产品的业务逻辑 | `Logistics` / `baseLogistics` |
| ConcreteCreator（具体工厂） | 重写工厂方法返回具体产品 | `RoadLogistics`、`SeaLogistics` |

## 结构

```
Logistics(接口)                Transport(接口)
  ├─ CreateTransport() ───────► ┌───────────┐
  └─ PlanDelivery()  复用工厂方法 └───────────┘
        ▲                          ▲     ▲
 RoadLogistics──►Truck        Truck    Ship
 SeaLogistics ──►Ship
```

## 与「简单工厂」的区别

- **简单工厂**：一个函数根据参数 `switch` 返回不同产品——不是 GoF 模式，但 Go 里很常用。
- **工厂方法**：把「创建」变成可被覆盖的方法/接口，靠**多态**而非 `switch` 扩展，新增产品不改老代码。

## Go 惯用写法

Go 没有类继承，用**组合 + 接口**实现：`baseLogistics` 内嵌到具体工厂中复用 `PlanDelivery`，
并把工厂方法以函数字段 `factory` 注入。更「Go」的轻量做法其实是直接用**构造函数**：

```go
// 最常见：一个返回接口的函数就是工厂
func NewTransport(kind string) Transport {
    switch kind {
    case "road": return Truck{}
    case "sea":  return Ship{}
    }
    return nil
}
```

也可以用 **函数类型** 当工厂：`type Factory func() Transport`。是否需要完整的接口式工厂方法，
取决于「创建逻辑是否要与一套复用的业务方法绑定」。

## 适用场景

- 一个框架需要标准化对象的创建，但把具体类型交给使用方扩展（如 `database/sql` 的 driver 注册）；
- 运行时才能决定创建哪种对象；
- 想把产品的创建集中管理、便于替换与测试。

## 优点

- 业务代码只依赖抽象产品，符合开闭原则、依赖倒置；
- 创建逻辑集中，便于维护。

## 缺点

- 每加一种产品往往要加一个工厂，类数量增多；
- 对简单场景是过度设计——Go 里优先考虑简单工厂函数。

## 运行

```bash
go run ./creational/factory_method
```
