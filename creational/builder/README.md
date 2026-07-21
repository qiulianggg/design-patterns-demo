# 建造者模式（Builder）

## 意图

> 将一个**复杂对象的构建过程**与它的**表示**分离，使得同样的构建过程可以创建不同的表示。

## 解决什么问题

当一个对象有**很多可选字段**、或构建过程分**多个步骤**、或需要在创建时做**校验/默认值**时，
用一个巨大的构造函数（telescoping constructor，`New(a, b, c, d, e, ...)`）既难读又易错。

建造者把「一步步设置各部件」和「最终产出对象」分开，让创建过程清晰、可读、可复用。

## 角色

| 角色 | 说明 | 本例对应 |
|------|------|---------|
| Product | 被构建的复杂对象 | `Computer` |
| Builder | 定义构建各部件的方法 | `ComputerBuilder` |
| Director（可选） | 封装固定的构建流程/顺序 | `Director` |
| Client | 使用 Builder（或 Director）得到产品 | `main` |

## 结构

```
Director ──uses──► ComputerBuilder ──Build()──► Computer
 (预设流程)          WithCPU/WithRAM/...          (最终产品)
```

## Go 惯用写法

Go 里建造者最常见的两种形态：

### 1. 链式 Builder（本例）

每个 `WithXxx` 返回 `*Builder`，支持流式调用，最后 `Build()` 收尾：

```go
pc := NewComputerBuilder().WithCPU("M3").WithRAM(24).Build()
```

### 2. Functional Options（Go 社区最推崇）

用可变参数 + 选项函数，无需单独的 builder 类型，非常适合「带很多可选项的构造函数」：

```go
type Option func(*Computer)
func WithGPU(g string) Option { return func(c *Computer){ c.GPU = g } }

func NewComputer(cpu string, opts ...Option) Computer {
    c := Computer{CPU: cpu, RAM: 8} // 默认值
    for _, o := range opts { o(&c) }
    return c
}
// 用法：NewComputer("M3", WithGPU("RTX4090"), WithRAM(64))
```

Functional Options 在标准库和知名库（gRPC、zap 等）中极为普遍，是 Go 版建造者的首选。

## 适用场景

- 对象字段多、可选项多，且很多有默认值；
- 构建过程需要分步、有顺序或需要校验；
- 想用同一套步骤产出不同配置（`Director` 的 `BuildOfficePC` / `BuildGamingPC`）。

## 优点

- 分步构建，代码可读性强；
- 可复用构建流程；可在 `Build()` 集中校验；
- 避免超长参数列表。

## 缺点

- 需要额外的 Builder 类型，增加代码量；
- 对字段少的简单对象是过度设计。

## 运行

```bash
go run ./creational/builder
```
