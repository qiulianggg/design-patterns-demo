# 适配器模式（Adapter）

## 意图

> 将一个类的接口**转换成客户希望的另一个接口**。适配器让原本因接口不兼容而不能一起工作的类能够协作。

## 解决什么问题

你有一段代码依赖接口 A，而现成的组件（第三方库、遗留代码）提供的是接口 B，两者签名不一致又不能改。
适配器像「电源转换插头」一样，包裹住 B，对外暴露 A，把调用翻译过去。

## 角色

| 角色 | 说明 | 本例对应 |
|------|------|---------|
| Target（目标接口） | 客户端期望使用的接口 | `PaymentProcessor` |
| Adaptee（被适配者） | 已有的、接口不兼容的类 | `ThirdPartyPay` |
| Adapter（适配器） | 实现 Target，内部调用 Adaptee | `ThirdPartyAdapter` |
| Client | 只面向 Target 编程 | `checkout` |

## 结构

```
Client ──► PaymentProcessor(Target)
                 ▲
          ThirdPartyAdapter ──持有──► ThirdPartyPay(Adaptee)
             Pay(cents)  ───翻译──►    MakePayment(yuan)
```

## Go 惯用写法

Go 里适配器就是「一个实现目标接口、内部委托给被适配对象的结构体」，用**组合**完成，非常自然。
由于 Go 接口是**隐式实现**的，只要 Adapter 的方法集匹配 Target，就自动满足接口，无需显式声明。

函数适配也很常见，例如标准库的 `http.HandlerFunc`：把一个普通函数适配成 `http.Handler` 接口：

```go
type HandlerFunc func(ResponseWriter, *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) { f(w, r) }
```

这就是用**函数类型 + 方法**做的适配器。

## 对象适配器 vs 类适配器

- **对象适配器**（组合，本例）：Adapter 内部**持有** Adaptee 实例。Go 只能用这种（无继承），也更灵活。
- **类适配器**（多重继承）：Go 不支持，忽略。

## 适用场景

- 接入第三方 SDK / 遗留系统，接口对不上；
- 统一多个不同来源、相似功能但接口各异的组件；
- 想复用一个类，但它的接口不符合当前需要。

## 优点

- 复用现有类而无需修改它（符合开闭原则）；
- 隔离接口差异，客户端代码保持稳定。

## 缺点

- 增加一层间接，代码整体复杂度上升；
- 适配器过多时会让调用链变绕。

## 与相近模式区别

- **适配器**：改变接口，不改变功能（事后补救，让 A 能用 B）。
- **装饰器**：保持接口，增强功能。
- **外观**：为一组复杂子系统提供一个简化的新接口（不是为了兼容既有接口）。

## 运行

```bash
go run ./structural/adapter
```
