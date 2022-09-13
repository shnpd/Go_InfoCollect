package main

import (
	"github.com/shirou/gopsutil/mem"
)

//SwapMemory	  交换内存
//VirtualMemory	  虚拟内存
//VirtualMemoryEx 虚拟交换内存
type Memory struct {
	SwapMemory      *mem.SwapMemoryStat      `json:"swapMemory"`
	VirtualMemory   *mem.VirtualMemoryStat   `json:"virtualMemory"`
	VirtualMemoryEx *mem.VirtualMemoryExStat `json:"virtualMemoryEx"`
}

//获得内存有关信息
func GetVirtualMemInfo() *mem.VirtualMemoryStat {
	memInfo, _ := mem.VirtualMemory()
	return memInfo
}

func GetVirtualMemExInfo() *mem.VirtualMemoryExStat {
	memInfo, _ := mem.VirtualMemoryEx()
	return memInfo
}

func GetMemory() Memory {
	swapInfo, err := mem.SwapMemory()
	if err != nil {
		panic(err)
	}
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}
	memExinfo, err := mem.VirtualMemoryEx()
	if err != nil {
		panic(err)
	}

	memory := Memory{
		SwapMemory:      swapInfo,
		VirtualMemory:   memInfo,
		VirtualMemoryEx: memExinfo,
	}
	return memory
}

//
//func main(){
//	fmt.Println(GetVirtualMemExInfo())
//	fmt.Println(GetVirtualMemInfo())
//}
