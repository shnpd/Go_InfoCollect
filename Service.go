package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

//RunningService	 运行中的服务
//AllService 		所有服务
type Service struct {
	RunningService string `json:"runningService"`
	AllService     string `json:"allService"`
}

//executive_ServiceOrder()函数服务于GetService()函数，通过调用Linux命令行收集对应attr服务信息
func executive_ServiceOrder(attr string) string {
	//running:获取正在运行的系统服务
	//all:获取全部的系统服务
	var cmd *exec.Cmd
	if attr == "running" {
		//获取运行中的服务
		cmd = exec.Command("/bin/bash", "-c", "service --status-all | grep +")
	} else if attr == "all" {
		//获取所有服务
		cmd = exec.Command("/bin/bash", "-c", "service --status-all")
	} else {
		fmt.Println("输入有误！")
		return ""
	}
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	//执行命令
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		panic(err)
	}
	if err := cmd.Wait(); err != nil {
		panic(err)
	}

	//返回数据
	return string(bytes)
}
//GetService()函数用于收集正在运行中的系统服务和全部的系统服务
func GetService() Service {
	//running:获取正在运行的系统服务
	//all:获取全部的系统服务

	//通过调用executive_ServiceOrder()函数完成对两种服务的收集工作
	runningService := executive_ServiceOrder("running")
	allService := executive_ServiceOrder("all")

	//结构化
	service := Service{
		RunningService: runningService,
		AllService:     allService,
	}

	//返回数据
	return service
}
