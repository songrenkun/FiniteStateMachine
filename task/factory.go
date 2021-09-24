package task


//Factory 任务工厂
type Factory interface {
	CreateTask(taskID string, configInfoData interface{}) (Task, error)
}
