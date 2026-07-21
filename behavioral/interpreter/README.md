# 解释器模式（Interpreter）

## 意图

> 给定一种语言，定义它的**文法表示**，并定义一个**解释器**，用该表示来解释语言中的句子。

## 解决什么问题

当你需要反复解释/求值某类**结构化的句子**（表达式、规则、查询、简单脚本）时，
可以为这门「小语言（DSL）」定义文法：每条文法规则对应一个类，把句子解析成**抽象语法树（AST）**，
再递归地对树求值。

## 角色

| 角色 | 说明 | 本例对应 |
|------|------|---------|
| AbstractExpression | 声明 `Interpret` | `Expression` |
| TerminalExpression | 终结符（叶子，直接求值） | `Constant`、`Variable` |
| NonterminalExpression | 非终结符（组合子表达式） | `AndExpression`、`OrExpression` |
| Context | 解释时的全局信息（变量表等） | `Context` |
| Client | 构建 AST 并触发解释 | `main` + `parse` |

## 结构

```
"x AND y OR true" 解析成 AST：
            Or
           /  \
         And   Constant(true)
        /   \
   Var(x)  Var(y)
每个节点 Interpret(ctx) 递归求值。
```

> 注意：本例的 `parse` 只是为了演示，做了极简的「左结合、无优先级」解析。解释器模式本身聚焦于
> **文法的对象表示与递归求值**，词法/语法分析（parser）通常是配套但独立的部分。

## Go 惯用写法

Go 里每个文法规则是一个实现 `Expression` 接口的结构体，`Interpret` 递归调用子表达式。
这与「组合模式」的树形结构一脉相承——AST 本身就是一棵组合树。

现实中，除非语言极其简单，很少手写完整解释器模式；更常见的是用 `text/scanner`、
`go/parser`，或 `goyacc`、参与解析库（如 `participle`）生成，再对 AST 求值。
本模式的价值更多在于**理解 AST + 递归求值**这一思想。

## 适用场景

- 有一门简单、稳定的语言/规则需要反复解释（配置规则引擎、过滤表达式、计算器）；
- 文法相对简单，且效率不是首要因素。

## 优点

- 文法规则与类一一对应，易于扩展新规则（加一个 Expression 实现）；
- 改变/扩展文法方便。

## 缺点

- 文法复杂时类数量爆炸、难以维护；
- 递归求值效率一般，不适合复杂/高性能语言——那时应上专业的解析器与虚拟机。

## 运行

```bash
go run ./behavioral/interpreter
```
