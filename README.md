# Go 语言 23 种设计模式 Demo

本仓库用 Go 语言实现了《设计模式：可复用面向对象软件的基础》（GoF）中的全部 **23 种经典设计模式**，
每种模式都包含：

- 一个**可独立运行**的 Go demo（`*.go`，`package main`，带 `main()` 演示）；
- 一份**详细讲解**的 `README.md`（意图、结构、角色、适用场景、优缺点、Go 惯用法、对比）。

## 如何运行

每个模式目录都是一个独立的 `package main`，可直接运行：

```bash
# 例如运行单例模式
go run ./creational/singleton

# 运行观察者模式
go run ./behavioral/observer
```

一次性编译检查全部：

```bash
go build ./...
go vet ./...
```

## 目录结构

设计模式按 GoF 分类分为三大类：

### 创建型模式（Creational, 5 种）—— 关注对象的创建

| 模式 | 目录 | 一句话意图 |
|------|------|-----------|
| 单例 Singleton | [`creational/singleton`](creational/singleton) | 保证一个类只有一个实例，并提供全局访问点 |
| 工厂方法 Factory Method | [`creational/factory_method`](creational/factory_method) | 定义创建对象的接口，让子类决定实例化哪个类 |
| 抽象工厂 Abstract Factory | [`creational/abstract_factory`](creational/abstract_factory) | 创建一系列相关或相互依赖的对象，而无需指定具体类 |
| 建造者 Builder | [`creational/builder`](creational/builder) | 将复杂对象的构建与表示分离，同样构建过程可创建不同表示 |
| 原型 Prototype | [`creational/prototype`](creational/prototype) | 用原型实例指定创建对象的种类，通过拷贝创建新对象 |

### 结构型模式（Structural, 7 种）—— 关注类与对象的组合

| 模式 | 目录 | 一句话意图 |
|------|------|-----------|
| 适配器 Adapter | [`structural/adapter`](structural/adapter) | 将一个类的接口转换成客户希望的另一个接口 |
| 桥接 Bridge | [`structural/bridge`](structural/bridge) | 将抽象部分与实现部分分离，使它们都可以独立变化 |
| 组合 Composite | [`structural/composite`](structural/composite) | 将对象组合成树形结构，使单个对象和组合对象使用一致 |
| 装饰器 Decorator | [`structural/decorator`](structural/decorator) | 动态地给对象添加职责，比子类更灵活 |
| 外观 Facade | [`structural/facade`](structural/facade) | 为子系统的一组接口提供一个统一的高层接口 |
| 享元 Flyweight | [`structural/flyweight`](structural/flyweight) | 运用共享技术有效支持大量细粒度对象 |
| 代理 Proxy | [`structural/proxy`](structural/proxy) | 为其他对象提供一种代理以控制对这个对象的访问 |

### 行为型模式（Behavioral, 11 种）—— 关注对象间的职责分配与通信

| 模式 | 目录 | 一句话意图 |
|------|------|-----------|
| 责任链 Chain of Responsibility | [`behavioral/chain_of_responsibility`](behavioral/chain_of_responsibility) | 让多个对象都有机会处理请求，沿链传递直到被处理 |
| 命令 Command | [`behavioral/command`](behavioral/command) | 将请求封装成对象，支持参数化、排队、撤销 |
| 解释器 Interpreter | [`behavioral/interpreter`](behavioral/interpreter) | 给定语言，定义文法表示并构建解释器 |
| 迭代器 Iterator | [`behavioral/iterator`](behavioral/iterator) | 顺序访问聚合对象元素，而不暴露内部表示 |
| 中介者 Mediator | [`behavioral/mediator`](behavioral/mediator) | 用中介对象封装一系列对象的交互，降低耦合 |
| 备忘录 Memento | [`behavioral/memento`](behavioral/memento) | 在不破坏封装的前提下捕获并恢复对象内部状态 |
| 观察者 Observer | [`behavioral/observer`](behavioral/observer) | 定义一对多依赖，对象状态改变时通知所有观察者 |
| 状态 State | [`behavioral/state`](behavioral/state) | 允许对象在内部状态改变时改变其行为 |
| 策略 Strategy | [`behavioral/strategy`](behavioral/strategy) | 定义一系列算法并封装，使它们可以互相替换 |
| 模板方法 Template Method | [`behavioral/template_method`](behavioral/template_method) | 定义算法骨架，将某些步骤延迟到子类实现 |
| 访问者 Visitor | [`behavioral/visitor`](behavioral/visitor) | 在不改变元素类的前提下定义作用于元素的新操作 |

## 阅读建议

设计模式不是银弹，Go 语言由于其**接口隐式实现、组合优于继承、函数是一等公民、并发原语内置**等特性，
很多模式的写法与经典 Java/C++ 实现有所不同。每个模式的 README 都专门有一节 **「Go 惯用写法」**
说明如何写得更「Go」，避免生搬硬套面向对象的教条。
