package duck_task

import (
	"finiteStateMachine/task"
	"fmt"
	"time"
)

/*
0.初始化状态
*/

//  初始化状态
type InitialState struct {
	Type int
	Task *Task
}

func NewInitialState(task *Task) task.State {
	return &InitialState{
		Type: 0,
		Task: task,
	}
}

//初始化状态,start命令无效
func (s *InitialState) Start() error {
	return fmt.Errorf("当前状态为初始化，无法响应Start请求")

}

func (s *InitialState) Stop() error {
	s.Task.ChangeState(NewStopState(s.Task))
	return nil
}

func (s *InitialState) Cancel() error {
	return fmt.Errorf("当前状态为初始化，无法响应Cancel请求")
}

func (s *InitialState) Reset() error {
	return fmt.Errorf("当前状态为初始化，无法响应Reset请求")
}

//初始化状态下
func (s *InitialState) Run() {
	fmt.Println("任务开始前的准备期")
	time.Sleep(time.Second * 5)

	s.Task.ChangeState(NewEatingState(s.Task))
}
