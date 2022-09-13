package main

import (
	"io/ioutil"
	"os/exec"
)

//DebInstall	deb所安装的应用
//rpmInstall	rpm所安装的应用
//yumInstall	yum所安装的应用
type Application struct {
	DebInstall string `json:"debInstall"`
	rpmInstall string `json:"rpmInstall"`
	yumInstall string `json:"yumInstall"`
}

//GetAppInfo()函数通过linux命令采集已安装的应用软件、版本、描述
func GetAppInfo() Application {
	//dpkg -l 命令会列出系统中所有已安装的软件包信息
	cmd := exec.Command("/bin/bash", "-c", "dpkg -l")
	//rpm -qa 列出所有被安装的rpm package
	cmd2 := exec.Command("/bin/bash", "-c", "rpm -qa")
	//yum list installed列出已安装的安装包
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

	//结构化
	appInfo := Application{
		DebInstall: string(bytes),
		rpmInstall: string(bytes2),
		yumInstall: string(bytes3),
	}
	//返回数据
	return appInfo
}
