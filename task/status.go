package task

import (
	"sync"
)

//task 信息
type Info struct {
	ID         string      //task id
	Type       int         //task 类型
	ConfigInfo interface{} //task 基础配置信息
	StatusInfo interface{} //task 状态
}

var statusMapInstance *InfoMap
var statusMapOnce sync.Once

//InfoMap 任务信息Map
type InfoMap struct {
	TaskInfoMap map[string]*Info
	lock        sync.RWMutex
}

//GetTaskStatusInstance 获取所有任务状态单例
func GetTaskStatusInstance() *InfoMap {
	statusMapOnce.Do(func() {
		statusMapInstance = &InfoMap{
			TaskInfoMap: make(map[string]*Info),
		}
	})
	return statusMapInstance
}

// 更改(含新增)任务信息map
func (m *InfoMap) ChangeInfo(taskID string, info *Info) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.TaskInfoMap[taskID] = info
}

//删除信息
func (m *InfoMap) DeleteInfo(taskID string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.TaskInfoMap, taskID)
}
