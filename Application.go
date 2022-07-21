package main

import (
	"io/ioutil"
	"os/exec"
)

type Application struct {
	DebInstall string `json:"debInstall"`
	rpmInstall string `json:"rpmInstall"`
	yumInstall string `json:"yumInstall"`
}

func GetAppInfo() Application {
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
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}
	if err3 != nil {
		panic(err3)
	}
	//执行命令
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	if err2 := cmd2.Start(); err2 != nil {
		panic(err2)
	}
	if err3 := cmd3.Start(); err3 != nil {
		panic(err3)
	}
	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	bytes2, err2 := ioutil.ReadAll(stdout2)
	bytes3, err3 := ioutil.ReadAll(stdout3)
	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}
	if err3 != nil {
		panic(err3)
	}
	if err := cmd.Wait(); err != nil {
		panic(err)
	}
	if err2 := cmd2.Wait(); err2 != nil {
		panic(err2)
	}
	if err3 := cmd3.Wait(); err3 != nil {
		panic(err3)
	}

	appInfo := Application{
		DebInstall: string(bytes),
		rpmInstall: string(bytes2),
		yumInstall: string(bytes3),
	}
	return appInfo
}
