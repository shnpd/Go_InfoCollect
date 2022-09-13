package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

//SystemLog：操作系统日志
//SecureLog：安全日志
type Log struct {
	SystemLog string `json:"systemLog"`
	SecureLog string `json:"secureLog"`
}

//通过Linux命令行获取系统日志信息，然后经过管道获取回显，放到bytes内，返回。
func executive_LogOrder(attr string) []byte {
	//System：系统日志
	//Secure：安全日志
	var cmd *exec.Cmd
	if attr == "System" {
		cmd = exec.Command("/bin/bash", "-c", "cat /var/log/syslog")
	} else if attr == "Secure" {
		cmd = exec.Command("/bin/bash", "-c", "cat /var/log/auth.log")
	} else {
		fmt.Println("输入有误！")
		return nil
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
	//返回结果
	return bytes
}

func GetLog() Log {
	//System：系统日志
	//Secure：安全日志
	systemLog := string(executive_LogOrder("System"))
	secureLog := string(executive_LogOrder("Secure"))
	log := Log{
		SystemLog: systemLog,
		SecureLog: secureLog,
	}
	return log
}
