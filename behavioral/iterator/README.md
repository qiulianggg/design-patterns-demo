# 迭代器模式（Iterator）

## 意图

> 提供一种方法**顺序访问**一个聚合对象中的各个元素，而**不暴露**该对象的内部表示。

## 解决什么问题

不同的集合有不同的内部结构（数组、链表、树、哈希、环形缓冲）。如果每种集合都让使用方直接按其内部结构去遍历，代码就和具体结构耦合了。

迭代器把「如何遍历」封装成一个独立对象，对外只暴露 `HasNext()` / `Next()`。
使用方用统一方式遍历任何集合，且集合的内部实现可以随意改变。

## 角色

| 角色 | 说明 | 本例对应 |
|------|------|---------|
| Iterator | 遍历接口 | `Iterator` |
| ConcreteIterator | 具体遍历逻辑、维护游标 | `ringIterator` |
| Aggregate | 聚合接口，能产出迭代器 | `Aggregate` |
| ConcreteAggregate | 具体集合 | `RingBuffer` |

## 结构

```
RingBuffer.CreateIterator() ──► ringIterator
Client: for it.HasNext() { it.Next() }
           （不知道内部是环形+取模）
```

本例 `RingBuffer` 从逻辑起点开始、用取模环绕遍历，客户端完全看不到这些细节。

## Go 惯用写法

这是**最需要提醒「不要生搬硬套」**的模式。Go 已内建强大的遍历机制：

1. **`range`** 直接遍历 slice/map/channel/string——绝大多数情况用它就够了；
2. **channel + goroutine** 可作为惰性/无限序列的迭代器：
   ```go
   func gen() <-chan int { ch := make(chan int); go func(){ ... ; close(ch) }(); return ch }
   for v := range gen() { ... }
   ```
3. **Go 1.23+ 的 range-over-func 迭代器**（`iter.Seq`）是官方版迭代器模式，最地道：
   ```go
   func (r *RingBuffer) All() iter.Seq[int] {
       return func(yield func(int) bool) {
           for i := range r.data {
               idx := (r.start + i) % len(r.data)
               if !yield(r.data[idx]) { return }
           }
       }
   }
   // 用法：for v := range rb.All() { ... }
   ```

本 demo 用经典的 `HasNext/Next` 接口是为了展示 GoF 原型；实际项目请优先用 `range` 或 `iter.Seq`。

## 适用场景

- 需要遍历复杂/自定义的聚合结构而不暴露其内部；
- 需要为同一集合提供多种遍历方式（正序、逆序、过滤）；
- 需要统一不同集合的遍历接口。

## 优点

- 遍历逻辑与集合分离，单一职责；
- 支持多种、并行的遍历；隐藏内部结构。

## 缺点

- 对简单集合是多余抽象——Go 里尤其明显；
- 经典 `HasNext/Next` 接口不如 `range`/`iter.Seq` 简洁安全。

## 运行

```bash
go run ./behavioral/iterator
```
