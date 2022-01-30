package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func executive_DeviceOrder(attr string) []byte {
	//name 设备名
	//manufacturer 设备厂商
	//serial_number 设备编码
	//version 设备型号
	var cmd *exec.Cmd
	if attr == "name" {
		cmd = exec.Command("/bin/bash", "-c", "sudo dmidecode -s system-product-name")
	} else
	if attr == "manufacturer" {
		cmd = exec.Command("/bin/bash", "-c", "sudo dmidecode -s system-manufacturer")
	} else
	if attr == "serial_number" {
		cmd = exec.Command("/bin/bash", "-c", "sudo dmidecode -s system-serial-number")
	} else
	if attr == "version" {
		cmd = exec.Command("/bin/bash", "-c", "sudo dmidecode -s system-version")
	} else {
		fmt.Println("输入有误！")
		return nil
	}

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return nil
	}
	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err", err)
		return nil
	}
	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return nil
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return nil
	}
	return bytes
	//fmt.Printf("stdout:\n\n%s", bytes)
}

type Device struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	SerialNumber string `json:"serial_number"`
	Version      string `json:"version"`
}

func Getdevice() {

	//name 设备名
	//manufacturer 设备厂商
	//serial_number 设备编码
	//version 设备型号
	name := string(executive_DeviceOrder("name"))
	manu := string(executive_DeviceOrder("manufacturer"))
	numb := string(executive_DeviceOrder("serial_number"))
	version := string(executive_DeviceOrder("version"))

	//去掉末尾换行符
	name = strings.TrimRight(name, "\n")
	manu = strings.TrimRight(manu, "\n")
	numb = strings.TrimRight(numb, "\n")
	version = strings.TrimRight(version, "\n")

	device := Device{
		Name:         name,
		Manufacturer: manu,
		SerialNumber: numb,
		Version:      version,
	}
	deviceJson, err := json.Marshal(device)
	if err != nil {
		panic(err)
	}
	WriteFile("device.json", deviceJson)
}
