package main

import (
	"github.com/shirou/gopsutil/process"
)


//Pid 		进程ID
//Name 		进程名
//MemoryUse 内存占用
//CpuUse 	CPU占用
//Uids 		用户ID
type Process struct {
	Pid       int32                     `json:"pid"`
	Name      string                    `json:"name"`
	MemoryUse *process.MemoryInfoExStat `json:"memoryUse"`
	CpuUse    float64                   `json:"cpuUse"`
	Uids      []int32                   `json:"uids"`
}

//ProcessId()函数通过调用gopsutil的process模块，循环获取进程id
func ProcessId() (pid []int32) {
	//进程ID获取
	pids, _ := process.Pids()

	//循环获取所有进程ID
	for _, p := range pids {
		pid = append(pid, p)
	}

	//返回数据
	return pid
}

//ProcessName()函数通过调用gopsutil的process模块，循环获取每个进程的进程名
func ProcessName() (pname []string) {
	//获取进程ID
	pids, _ := process.Pids()

	//依据进程ID循环获取进程名
	for _, pid := range pids {
		pn, _ := process.NewProcess(pid)
		pName, _ := pn.Name()
		pname = append(pname, pName)
	}

	//返回数据
	return pname
}

//ProcessMemory()函数通过调用gopsutil的process模块，循环获取每个进程的内存占用
func ProcessMemory() (pmemory []string) {
	//获取进程ID
	pids, _ := process.Pids()

	//依据进程ID循环获取内存占用
	for _, pid := range pids {
		pn, _ := process.NewProcess(pid)
		pmry, _ := pn.MemoryInfoEx()
		pMemory := pmry.String()
		pmemory = append(pmemory, pMemory)
	}

	//返回数据
	return pmemory
}

//ProcessCpu()函数通过调用gopsutil的process模块，循环获取每个进程的CPU占用
func ProcessCpu() (pcpu []float64) {
	//获取进程ID
	pids, _ := process.Pids()

	//依据进程ID循环获取CPU占用
	for _, pid := range pids {
		pn, _ := process.NewProcess(pid)
		pCpu, _ := pn.CPUPercent()
		pcpu = append(pcpu, pCpu)
	}

	//返回数据
	return pcpu
}

//GetProcess()函数通过调用gopsutil的process模块，收集所有进程的进程ID、名称、内存占用、CPU占用等信息
func GetProcess() []Process {
	//tempName 进程名数据组
	//tempMemoryUse 内存占用数据组
	//tempCpuUse CPU占用数据组
	//tempUids 用户ID数据组

	//获取进程ID
	pids, err := process.Pids()
	if err != nil {
		panic(err)
	}
	var Processes []Process

	//依据进程ID循环获取、名称、内存占用、CPU占用、用户ID等信息
	for _, pid := range pids {
		tempProcess, err := process.NewProcess(pid)
		if err != nil {
			panic(err)
		}

		//获取进程名
		tempName, err := tempProcess.Name()
		if err != nil {
			panic(err)
		}

		//获取内存占用
		tempMemoryUse, err := tempProcess.MemoryInfoEx()
		if err != nil {
			panic(err)
		}

		//获取CPU占用
		tempCpuUse, err := tempProcess.CPUPercent()
		if err != nil {
			panic(err)
		}

		//获取用户ID
		tempUids, err := tempProcess.Uids()
		if err != nil {
			panic(err)
		}

		//结构化
		temp := Process{
			Pid:       pid,
			Name:      tempName,
			MemoryUse: tempMemoryUse,
			CpuUse:    tempCpuUse,
			Uids:      tempUids,
		}
		Processes = append(Processes, temp)
	}

	//返回数据
	return Processes
}
