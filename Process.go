package main

import (
	"github.com/shirou/gopsutil/process"
)

type Process struct {
	Pid       int32                     `json:"pid"`
	Name      string                    `json:"name"`
	MemoryUse *process.MemoryInfoExStat `json:"memoryUse"`
	CpuUse    float64                   `json:"cpuUse"`
	Uids      []int32                   `json:"uids"`
}

//获取进程id
func ProcessId() (pid []int32) {
	pids, _ := process.Pids()
	for _, p := range pids {
		pid = append(pid, p)
	}
	return pid
}

//获取进程名
func ProcessName() (pname []string) {
	pids, _ := process.Pids()
	for _, pid := range pids {
		pn, _ := process.NewProcess(pid)
		pName, _ := pn.Name()
		pname = append(pname, pName)
	}
	return pname
}

//内存占用
func ProcessMemory() (pmemory []string) {
	pids, _ := process.Pids()
	for _, pid := range pids {
		pn, _ := process.NewProcess(pid)
		pmry, _ := pn.MemoryInfoEx()
		pMemory := pmry.String()
		pmemory = append(pmemory, pMemory)
	}
	return pmemory
}

//CPU占用
func ProcessCpu() (pcpu []float64) {
	pids, _ := process.Pids()
	for _, pid := range pids {
		pn, _ := process.NewProcess(pid)
		pCpu, _ := pn.CPUPercent()
		pcpu = append(pcpu, pCpu)
	}
	return pcpu
}

func GetProcess() []Process {
	pids, err := process.Pids()
	if err != nil {
		panic(err)
	}
	var Processes []Process
	for _, pid := range pids {
		tempProcess, err := process.NewProcess(pid)
		if err != nil {
			panic(err)
		}
		tempName, err := tempProcess.Name()
		if err != nil {
			panic(err)
		}
		tempMemoryUse, err := tempProcess.MemoryInfoEx()
		if err != nil {
			panic(err)
		}
		tempCpuUse, err := tempProcess.CPUPercent()
		if err != nil {
			panic(err)
		}
		tempUids, err := tempProcess.Uids()
		if err != nil {
			panic(err)
		}
		temp := Process{
			Pid:       pid,
			Name:      tempName,
			MemoryUse: tempMemoryUse,
			CpuUse:    tempCpuUse,
			Uids:      tempUids,
		}
		Processes = append(Processes, temp)
	}
	return Processes
}
