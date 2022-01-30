package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
)

type AppInfo struct {
	DebInstall string `json:"deb_install"`
	rpmInstall string `json:"rpm_install"`
	yumInstall string `json:"yum_install"`
}

func GetAppInfo() {
	//deb安装
	cmd := exec.Command("/bin/bash", "-c", "dpkg -l")
	//rpm安装
	cmd2 := exec.Command("/bin/bash", "-c", "rpm -qa")
	//yum安装
	cmd3 := exec.Command("/bin/bash", "-c", "yum list installed")
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	stdout2, err2 := cmd2.StdoutPipe()
	stdout3, err3 := cmd3.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}
	if err2 != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err2)
		return
	}
	if err3 != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err3)
		return
	}
	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err", err)
		return
	}
	if err2 := cmd2.Start(); err2 != nil {
		fmt.Println("Error:The command is err", err2)
		return
	}
	if err3 := cmd3.Start(); err3 != nil {
		fmt.Println("Error:The command is err", err3)
		return
	}
	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	bytes2, err2 := ioutil.ReadAll(stdout2)
	bytes3, err3 := ioutil.ReadAll(stdout3)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return
	}
	if err2 != nil {
		fmt.Println("ReadAll Stdout:", err2.Error())
		return
	}
	if err3 != nil {
		fmt.Println("ReadAll Stdout:", err3.Error())
		return
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
	if err2 := cmd2.Wait(); err2 != nil {
		fmt.Println("wait:", err2.Error())
		return
	}
	if err3 := cmd3.Wait(); err3 != nil {
		fmt.Println("wait:", err3.Error())
		return
	}

	appInfo := AppInfo{
		DebInstall: string(bytes),
		rpmInstall: string(bytes2),
		yumInstall: string(bytes3),
	}
	appInfoJson, err := json.Marshal(appInfo)
	WriteFile("Application.json", appInfoJson)
}
