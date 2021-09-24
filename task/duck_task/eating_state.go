package duck_task

import (
	"finiteStateMachine/task"
	"fmt"
	"time"
)

/*
1.吃的状态
*/

type EatingState struct {
	Type int
	Task *Task
}

func NewEatingState(task *Task) task.State {
	return &EatingState{
		Type: 1,
		Task: task,
	}
}

//初始化状态,start命令无效
func (s *EatingState) Start() error {
	return fmt.Errorf("当前状态为初始化，无法响应Start请求")

}

//吃的状态通过stop命令转换成 停止发呆状态
func (s *EatingState) Stop() error {

	s.Task.ChangeState(NewStopState(s.Task))
	return nil
}

func (s *EatingState) Cancel() error {
	return fmt.Errorf("当前状态为初始化，无法响应Cancel请求")
}

func (s *EatingState) Run() {
	fmt.Println("开始吃！！")
	time.Sleep(time.Second * 1)
	s.Task.SetStatusInfo(&StatusInfo{AccomplishmentRate: 0.1})
	fmt.Println("吃1！！")
	time.Sleep(time.Second * 1)
	s.Task.SetStatusInfo(&StatusInfo{AccomplishmentRate: 0.3})
	fmt.Println("吃2！！")
	time.Sleep(time.Second * 1)
	s.Task.SetStatusInfo(&StatusInfo{AccomplishmentRate: 0.6})
	fmt.Println("吃3！！")
	time.Sleep(time.Second * 1)
	s.Task.SetStatusInfo(&StatusInfo{AccomplishmentRate: 0.7})
	fmt.Println("吃4！！")
	time.Sleep(time.Second * 1)
	s.Task.SetStatusInfo(&StatusInfo{AccomplishmentRate: 1})
	fmt.Println("吃5！！")
	s.Task.ChangeState(NewFullState(s.Task))
}
