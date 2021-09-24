package task

import (
	"fmt"
	"strings"
	"sync"
)

var controllerInstance *controller
var controllerOnce sync.Once

//controller 任务控制器
type controller struct {
	taskFactoryMap map[int]Factory
	taskMap        map[string]Task
	taskMapRWLock  sync.RWMutex
}

func InitTaskController(factoryMap map[int]Factory) {
	controllerOnce.Do(func() {
		controllerInstance = &controller{
			taskFactoryMap: factoryMap,
			taskMap:        make(map[string]Task),
		}
	})
}

//GetControllerInstance 获取任务控制器单例
func GetControllerInstance() *controller {
	return controllerInstance
}

//Create 创建任务
func (c *controller) Create(taskID string, taskType int, taskConfigInfo interface{}) error {

	if strings.EqualFold(taskID, "") {
		return fmt.Errorf("task id is nil")
	}
	c.taskMapRWLock.Lock()
	defer c.taskMapRWLock.Unlock()
	if c.taskMap[taskID] != nil {
		return fmt.Errorf("task not exit in taskMap")
	}
	taskFactory, ok := controllerInstance.taskFactoryMap[taskType]
	if !ok {
		return fmt.Errorf("task type not registered in task factory")
	}

	task, err := taskFactory.CreateTask(taskID, taskConfigInfo)
	if err != nil {
		return err
	}
	controllerInstance.taskMap[taskID] = task
	go func() {
		task.Run()
	}()
	return nil
}

//Start 开启任务
func (c *controller) Start(taskID string) error {
	c.taskMapRWLock.Lock()
	task := c.taskMap[taskID]
	if task == nil {
		c.taskMapRWLock.Unlock()
		return fmt.Errorf("task id = %s not exit", taskID)

	}
	c.taskMapRWLock.Unlock()
	return task.Start()
}

//Stop 暂停任务
func (c *controller) Stop(taskID string) error {
	c.taskMapRWLock.Lock()
	task := c.taskMap[taskID]
	if task == nil {
		c.taskMapRWLock.Unlock()
		return fmt.Errorf("task id = %s not exit", taskID)
	}
	c.taskMapRWLock.Unlock()
	return task.Stop()
}

//Cancel 取消任务
func (c *controller) Cancel(taskID string) error {
	c.taskMapRWLock.Lock()
	task := c.taskMap[taskID]
	if task == nil {
		c.taskMapRWLock.Unlock()
		return fmt.Errorf("task id = %s not exit", taskID)
	}
	c.taskMapRWLock.Unlock()
	return task.Cancel()
}

func (c *controller) GetStatus(taskID string) interface{} {
	c.taskMapRWLock.Lock()
	task := c.taskMap[taskID]
	if task == nil {
		c.taskMapRWLock.Unlock()
		return fmt.Errorf("task id = %s not exit", taskID)
	}
	c.taskMapRWLock.Unlock()
	return task.GetInfo()
}
