package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
)

type DataAmount struct {
	DataAmount string `json:"data_amount"`
}

func GetDataAmount() {
	cmd1 := exec.Command("/bin/bash", "-c", "vnstat -u")
	cmd := exec.Command("/bin/bash", "-c", "vnstat")

	//分别执行更新命令和读取数据量命令
	stdout, err := cmd.StdoutPipe()
	stdout2, err2 := cmd1.StdoutPipe()
	stdout2 = stdout2
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}
	if err2 != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err2)
		return
	}

	//执行命令
	if err2 := cmd1.Start(); err2 != nil {
		fmt.Println("Error:The command is err,", err2)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}
	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
	//返回结果
	dataAmount := DataAmount{DataAmount: string(bytes)}
	dataAmountJson, err := json.Marshal(dataAmount)
	if err != nil {
		panic(err)
	}
	WriteFile("DataAmount.json", dataAmountJson)
}
