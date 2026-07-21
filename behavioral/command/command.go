package main

import "fmt"

// 命令模式：把「请求」封装成对象，从而支持参数化、排队、记录、撤销。
// 本例：遥控器(Invoker) 通过命令对象操作灯(Receiver)，并支持撤销。

// Command 是命令接口：执行与撤销。
type Command interface {
	Execute() string
	Undo() string
}

// Light 是接收者(Receiver)：真正执行操作的对象。
type Light struct {
	name string
	on   bool
}

func (l *Light) TurnOn() string {
	l.on = true
	return l.name + " 灯已打开"
}
func (l *Light) TurnOff() string {
	l.on = false
	return l.name + " 灯已关闭"
}

// LightOnCommand 具体命令：封装「开灯」这一请求。
type LightOnCommand struct{ light *Light }

func (c *LightOnCommand) Execute() string { return c.light.TurnOn() }
func (c *LightOnCommand) Undo() string    { return c.light.TurnOff() }

// LightOffCommand 具体命令：封装「关灯」。
type LightOffCommand struct{ light *Light }

func (c *LightOffCommand) Execute() string { return c.light.TurnOff() }
func (c *LightOffCommand) Undo() string    { return c.light.TurnOn() }

// RemoteControl 是调用者(Invoker)：持有命令、触发执行，并记录历史以支持撤销。
type RemoteControl struct {
	history []Command
}

func (r *RemoteControl) Press(c Command) {
	fmt.Println("按下按钮 ->", c.Execute())
	r.history = append(r.history, c)
}

func (r *RemoteControl) PressUndo() {
	if len(r.history) == 0 {
		fmt.Println("没有可撤销的操作")
		return
	}
	last := r.history[len(r.history)-1]
	r.history = r.history[:len(r.history)-1]
	fmt.Println("撤销 ->", last.Undo())
}

func main() {
	livingRoom := &Light{name: "客厅"}
	remote := &RemoteControl{}

	remote.Press(&LightOnCommand{livingRoom})
	remote.Press(&LightOffCommand{livingRoom})
	remote.PressUndo() // 撤销关灯 -> 又打开
	remote.PressUndo() // 撤销开灯 -> 又关闭
}
