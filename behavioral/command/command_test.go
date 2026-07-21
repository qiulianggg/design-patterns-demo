package main

import "testing"

// 执行开灯命令后灯应为开；撤销后应关闭。
func TestCommandExecuteAndUndo(t *testing.T) {
	light := &Light{name: "客厅"}
	remote := &RemoteControl{}

	remote.Press(&LightOnCommand{light})
	if !light.on {
		t.Fatal("开灯命令执行后灯应为开")
	}

	remote.PressUndo()
	if light.on {
		t.Fatal("撤销开灯后灯应为关")
	}
}

// 关灯命令的撤销应重新开灯。
func TestOffCommandUndo(t *testing.T) {
	light := &Light{name: "卧室", on: true}
	remote := &RemoteControl{}
	remote.Press(&LightOffCommand{light})
	if light.on {
		t.Fatal("关灯命令后灯应为关")
	}
	remote.PressUndo()
	if !light.on {
		t.Fatal("撤销关灯后灯应为开")
	}
}

// 无历史时撤销应安全（不 panic）。
func TestUndoEmpty(t *testing.T) {
	(&RemoteControl{}).PressUndo()
}
