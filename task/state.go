package task


//State 任务状态
type State interface {
	Start() error
	Stop() error
	Cancel()error
	Run()
}