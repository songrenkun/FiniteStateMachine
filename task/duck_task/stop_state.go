package duck_task

import (
	"finiteStateMachine/task"
	"fmt"
	"time"
)

type StopState struct {
	Type int
	Task *Task
}

func NewStopState(task *Task) task.State {
	return &StopState{
		Type: 2,
		Task: task,
	}
}

//初始化状态,start命令无效
func (s *StopState) Start() error {
	s.Task.ChangeState(NewEatingState(s.Task))
	return nil
}

func (s *StopState) Stop() error {
	return fmt.Errorf("当前状态为停止发呆，无法响应Stop请求")
}

func (s *StopState) Cancel() error {
	return fmt.Errorf("当前状态为停止发呆，无法响应Cancel请求")
}

func (s *StopState) Reset() error {
	return fmt.Errorf("当前状态为停止发呆，无法响应Reset请求")
}
func (s *StopState) Run() {
	fmt.Println("停止,进行发呆")
	time.Sleep(time.Second * 5)
}
