# 责任链模式（Chain of Responsibility）

## 意图

> 使多个对象都有机会处理请求，从而避免请求发送者与接收者之间的耦合。将这些对象连成一条**链**，
> 沿着链传递请求，直到有对象处理它（或到达链尾）。

## 解决什么问题

一个请求需要经过**一系列检查/处理步骤**（认证、限流、日志、参数校验、业务处理……）。
如果写成一个大函数里一堆 `if`，既臃肿又难以增删、调序步骤。

责任链把每一步做成独立的处理器节点，串成链。每个节点只关心「我处不处理 / 要不要往下传」，
步骤的**组合与顺序**变得灵活可配，符合单一职责与开闭原则。

## 角色

| 角色 | 说明 | 本例对应 |
|------|------|---------|
| Handler | 处理器接口，含指向下一节点的能力 | `Handler` / `baseHandler` |
| ConcreteHandler | 具体处理器，处理或转发 | `AuthHandler`、`RateLimitHandler`、`BusinessHandler` |
| Client | 组装链并发起请求 | `main` |

## 结构

```
Request ──► Auth ──► RateLimit ──► Business
             │  任一节点都可「中断」并直接返回
             └── 或调用 callNext() 传给下一个
```

## Go 惯用写法

- 用 `baseHandler` 内嵌保存 `next` 并提供 `callNext`，具体处理器只写自己的判断逻辑；
- `SetNext` 返回传入的 handler，支持 `a.SetNext(b).SetNext(c)` 的链式串联。

**更「Go」的等价写法**：这其实就是 Web 框架里的**中间件链**。用函数式表达更常见：

```go
type Middleware func(http.Handler) http.Handler
// 洋葱式包裹：Auth(RateLimit(Business))
```

`net/http` 中间件、gRPC 拦截器链，本质都是责任链。写业务时优先考虑这种函数式中间件。

## 适用场景

- 请求需要多个对象按顺序处理，且处理者集合/顺序在运行时可变；
- Web 中间件、审批流、事件过滤链、日志处理管道。

## 优点

- 解耦发送者与接收者；
- 动态增删、重排处理节点，符合开闭原则；
- 每个处理器单一职责。

## 缺点

- 请求可能走到链尾都没被处理（需兜底）；
- 链过长影响性能、调试时不易追踪；
- 不保证一定被处理。

## 运行

```bash
go run ./behavioral/chain_of_responsibility
```
