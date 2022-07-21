package main

import (
	"io/ioutil"
	"os/exec"
)

type DataAmount struct {
	DataAmount string `json:"dataAmount"`
}

func GetDataAmount() DataAmount {
	cmd1 := exec.Command("/bin/bash", "-c", "vnstat -u")
	cmd := exec.Command("/bin/bash", "-c", "vnstat")

	//分别执行更新命令和读取数据量命令
	stdout, err := cmd.StdoutPipe()
	stdout2, err2 := cmd1.StdoutPipe()
	stdout2 = stdout2
	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}

	//执行命令
	if err2 := cmd1.Start(); err2 != nil {
		panic(err2)
	}

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
	//返回结果
	dataAmount := DataAmount{DataAmount: string(bytes)}
	return dataAmount
}
