package main

import (
	"encoding/json"
	"finiteStateMachine/task"
	"finiteStateMachine/task/duck_task"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//初始化工厂map
	factoryMap := map[int]task.Factory{
		-1: &duck_task.Factory{},
	}

	//初始化任务控制器
	task.InitTaskController(factoryMap)

	//创建任务
	err := task.GetControllerInstance().Create("test", -1, duck_task.ConfigInfo{"duck"})
	if err != nil {
		fmt.Println("创建duck 任务失败")
		return
	}
	//假设模拟对任务进行 停止操作（所有状态下停止,都进入停止状态）
	time.Sleep(time.Second * 2)
	err = task.GetControllerInstance().Stop("test")
	if err != nil {
		fmt.Println(err.Error())
	}
	time.Sleep(time.Second * 2)
	//停止状态再停止,报错误
	err = task.GetControllerInstance().Stop("test")
	if err != nil {
		fmt.Println(err.Error())
	}

	//继续开吃
	err = task.GetControllerInstance().Start("test")
	if err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(time.Second)
	//获取任务状态
	info := task.GetControllerInstance().GetStatus("test")
	b, _ := json.Marshal(info)
	fmt.Println(string(b))

	term := make(chan os.Signal, 1)

	signal.Notify(term, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	sig := <-term

	fmt.Printf("Got signal: %v, exit", sig)
}
