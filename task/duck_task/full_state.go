package duck_task

import (
	"finiteStateMachine/task"
	"fmt"
	"time"
)

type FullState struct {
	Type int
	Task *Task
}

func NewFullState(task *Task) task.State {
	return &FullState{
		Type: 3,
		Task: task,
	}
}

//初始化状态,start命令无效
func (s *FullState) Start() error {
	return fmt.Errorf("当前状态为吃饱状态，无法响应Start请求")

}

//吃的状态通过stop命令转换成 停止发呆状态
func (s *FullState) Stop() error {
	s.Task.ChangeState(NewStopState(s.Task))
	return nil
}

func (s *FullState) Cancel() error {
	return fmt.Errorf("当前状态为吃饱状态，无法响应Cancel请求")
}

func (s *FullState) Run(){
	fmt.Println("吃饱了。。")
	time.Sleep(time.Second * 1)
}
