package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
)

type Service struct {
	RunningService string `json:"running_service"`
	AllService     string `json:"all_service"`
}

func executive_ServiceOrder(attr string) string {
	//running:获取正在运行的系统服务
	//all:获取全部的系统服务
	var cmd *exec.Cmd
	if attr == "running" {
		cmd = exec.Command("/bin/bash", "-c", "service --status-all | grep +")
	} else if attr == "all" {
		cmd = exec.Command("/bin/bash", "-c", "service --status-all")
	} else {
		fmt.Println("输入有误！")
		return ""
	}
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
	return string(bytes)
}

func GetService() {
	//running:获取正在运行的系统服务
	//all:获取全部的系统服务
	runningService := executive_ServiceOrder("running")
	allService := executive_ServiceOrder("all")
	service := Service{
		RunningService: runningService,
		AllService:     allService,
	}
	serviceJson, err := json.Marshal(service)
	if err != nil {
		panic(err)
	}
	WriteFile("Service.json", serviceJson)
}
