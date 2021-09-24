package duck_task

import "finiteStateMachine/task"

/*

假设一个鸭子吃东西的任务
状态：
0.初始状态
1.吃
2.停止吃,发呆

3.吃饱了


状态转换

*/

const DuckTaskType = -1

type Task struct {
	*task.BaseTask //提供baseTask相关方法
}

//创建任务的一些必要信息(每种任务根据情况定制化)
type ConfigInfo struct {
	Name string
}

func (t *Task) SetStatusInfo(status *StatusInfo) {
	t.Info.StatusInfo = status
}

func (t *Task) GetStatusInfo() *StatusInfo {
	return t.Info.StatusInfo.(*StatusInfo)
}

//表示任务状态的信息(每种任务定制化)
type StatusInfo struct {
	AccomplishmentRate float64 //完成百分比
	Err                string  // 发生错误时的错误信息
}
