package task

//Task 任务接口
type Task interface {
	GetID() string
	GetType() int
	GetInfo() *Info
	Start() error
	Stop() error
	Cancel() error
	Run()
}

//BaseTask 任务
type BaseTask struct {
	Info         *Info
	currentState State
}

func NewBaseTask(taskID string, taskType int, configInfo interface{}, statusInfo interface{}) *BaseTask {
	return &BaseTask{
		Info: &Info{
			ID:         taskID,
			Type:       taskType,
			ConfigInfo: configInfo,
			StatusInfo: statusInfo,
		},
	}
}

func (t *BaseTask) SetState(state State) {
	t.currentState = state
}

// GetID 获取任务ID
func (t *BaseTask) GetID() string {
	return t.Info.ID
}

//GetType 任务类型
func (t *BaseTask) GetType() int {
	return t.Info.Type
}

//GetStatus 获取状态
func (t *BaseTask) GetInfo() *Info {
	return t.Info
}

//Start 任务开启
func (t *BaseTask) Start() error {
	return t.currentState.Start()
}

//Stop 任务停止
func (t *BaseTask) Stop() error {
	return t.currentState.Stop()
}

//Cancel 任务取消
func (t *BaseTask) Cancel() error {
	return t.currentState.Cancel()
}

func (t *BaseTask) Run() {
	t.currentState.Run()
}

//ChangeState 更改状态
func (t *BaseTask) ChangeState(newState State) {
	t.currentState = newState
	go func(state State) {
		t.currentState.Run()
	}(t.currentState)
}
