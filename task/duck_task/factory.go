package duck_task

import (
	"finiteStateMachine/task"
	"fmt"
)

type Factory struct{}

func (f *Factory) CreateTask(taskID string, configInfoData interface{}) (task.Task, error) {
	configInfo, ok := configInfoData.(ConfigInfo)
	if !ok {
		return nil, fmt.Errorf("config err")
	}

	baseTask := task.NewBaseTask(taskID, DuckTaskType, configInfo, &StatusInfo{})
	task := &Task{
		baseTask,
	}
	task.SetState(NewInitialState(task))
	return task, nil
}
