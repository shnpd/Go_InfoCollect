package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
)

type FireWall struct {
	Iptables_status string `json:"iptables_status"`
	Iptables_rules  string `json:"iptables_rules"`
	Firewall_status string `json:"firewall_status"`
	Firewall_rules  string `json:"firewall_rules"`
}

func GetFireWall() {
	ipstatus, iprules := Getiptables()     //ip防火墙信息
	firestatus, firerules := Getfirewall() //firewall防火墙信息
	firewall := FireWall{
		Iptables_status: ipstatus,
		Iptables_rules:  iprules,
		Firewall_status: firestatus,
		Firewall_rules:  firerules,
	}
	firewallJson, err := json.Marshal(firewall)
	if err != nil {
		panic(err)
	}
	WriteFile("Firewall.json", firewallJson)
}

//iptables防火墙
func Getiptables() (string, string) {
	cmd := exec.Command("/bin/bash", "-c", "service iptables status")
	cmd2 := exec.Command("/bin/bash", "-c", "sudo iptables -L")
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	stdout2, err2 := cmd2.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return "", ""
	}
	if err2 != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err2)
		return "", ""
	}
	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err", err)
		return "", ""
	}
	if err2 := cmd2.Start(); err2 != nil {
		fmt.Println("Error:The command is err", err2)
		return "", ""
	}
	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	bytes2, err2 := ioutil.ReadAll(stdout2)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return "", ""
	}
	if err2 != nil {
		fmt.Println("ReadAll Stdout:", err2.Error())
		return "", ""
	}
	//if err := cmd.Wait(); err != nil {
	//	fmt.Println("wait:", err.Error())
	//	return
	//}
	if err2 := cmd2.Wait(); err2 != nil {
		fmt.Println("wait:", err2.Error())
		return "", ""
	}

	return string(bytes), string(bytes2)
	//fmt.Println("iptables的状态：")
	//fmt.Printf("%s\n", bytes)
	//fmt.Println("iptables的规则：")
	//fmt.Printf("%s\n", bytes2)
}

//firewall防火墙
func Getfirewall() (string, string) {
	cmd := exec.Command("/bin/bash", "-c", "firewall-cmd --state")
	cmd2 := exec.Command("/bin/bash", "-c", "firewall-cmd --list-all")
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	stdout2, err2 := cmd2.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return "", ""
	}
	if err2 != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err2)
		return "", ""
	}
	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err", err)
		return "", ""
	}
	if err2 := cmd2.Start(); err2 != nil {
		fmt.Println("Error:The command is err", err2)
		return "", ""
	}
	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	bytes2, err2 := ioutil.ReadAll(stdout2)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return "", ""
	}
	if err2 != nil {
		fmt.Println("ReadAll Stdout:", err2.Error())
		return "", ""
	}
	//if err := cmd.Wait(); err != nil {
	//	fmt.Println("wait:", err.Error())
	//	return
	//}
	if err2 := cmd2.Wait(); err2 != nil {
		fmt.Println("wait:", err2.Error())
		return "", ""
	}
	return string(bytes), string(bytes2)
	//fmt.Println("firewall的状态：")
	//fmt.Printf("%s\n", bytes)
	//fmt.Println("firewall的规则：")
	//fmt.Printf("%s\n", bytes2)
}
