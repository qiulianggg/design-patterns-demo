package main

import (
	"fmt"
	"strings"
)

// 解释器模式：为一种简单语言定义文法，并构建一个解释器来求值该语言的句子。
// 本例：解释布尔表达式，如 "true AND false OR true"（从左到右，无优先级，仅演示）。

// Context 携带解释所需的变量绑定（本例简单起见用不到复杂上下文）。
type Context struct {
	vars map[string]bool
}

// Expression 是抽象表达式：所有终结符/非终结符都实现它。
type Expression interface {
	Interpret(ctx *Context) bool
}

// Constant 终结符表达式：常量 true / false。
type Constant struct{ value bool }

func (c Constant) Interpret(*Context) bool { return c.value }

// Variable 终结符表达式：从上下文取变量值。
type Variable struct{ name string }

func (v Variable) Interpret(ctx *Context) bool { return ctx.vars[v.name] }

// AndExpression 非终结符：逻辑与。
type AndExpression struct{ left, right Expression }

func (e AndExpression) Interpret(ctx *Context) bool {
	return e.left.Interpret(ctx) && e.right.Interpret(ctx)
}

// OrExpression 非终结符：逻辑或。
type OrExpression struct{ left, right Expression }

func (e OrExpression) Interpret(ctx *Context) bool {
	return e.left.Interpret(ctx) || e.right.Interpret(ctx)
}

// parse 把 token 串构建成表达式树（左结合，无优先级，仅供演示）。
func parse(tokens []string, ctx *Context) Expression {
	var expr Expression = toTerminal(tokens[0], ctx)
	for i := 1; i < len(tokens); i += 2 {
		op := tokens[i]
		right := toTerminal(tokens[i+1], ctx)
		switch op {
		case "AND":
			expr = AndExpression{expr, right}
		case "OR":
			expr = OrExpression{expr, right}
		}
	}
	return expr
}

func toTerminal(tok string, ctx *Context) Expression {
	switch tok {
	case "true":
		return Constant{true}
	case "false":
		return Constant{false}
	default:
		return Variable{tok} // 视为变量
	}
}

func main() {
	ctx := &Context{vars: map[string]bool{"x": true, "y": false}}

	for _, src := range []string{
		"true AND false OR true",
		"x AND y",
		"x OR y AND true",
	} {
		tree := parse(strings.Fields(src), ctx)
		fmt.Printf("%-25s => %v\n", src, tree.Interpret(ctx))
	}
}
