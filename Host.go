package main

import (
	"github.com/shirou/gopsutil/host"
)

//type InfoStat struct {
//	Hostname             主机名称
//	Uptime               开机时间
//	BootTime             boot时间
//	Procs                进程数目
//	OS                   操作系统 如freebsd, linux
//	Platform             如： ubuntu, linuxmint
//	PlatformFamily       如: debian, rhel
//	PlatformVersion      操作系统版本
//	KernelVersion        操作系统内核版本
//	KernelArch           内核架构
//	VirtualizationSystem 虚拟系统
//	VirtualizationRole   虚拟角色 guest or host
//	HostID               hostid  // ex: uuid
//}
func GetHostInfo() *host.InfoStat {
	HostInfo, err := host.Info()
	if err != nil {
		panic(err)
	}
	return HostInfo
}
