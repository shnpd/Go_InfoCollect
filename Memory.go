package main

import (
	"encoding/json"
	"github.com/shirou/gopsutil/mem"
)

type Memory struct {
	SwapMemory      *mem.SwapMemoryStat      `json:"swap_memory"`
	VirtualMemory   *mem.VirtualMemoryStat   `json:"virtual_memory"`
	VirtualMemoryEx *mem.VirtualMemoryExStat `json:"virtual_memory_ex"`
}

func getMemory() {
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

	memoryJson, err := json.Marshal(memory)
	if err != nil {
		panic(err)
	}
	WriteFile("Memory.json", memoryJson)
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

//
//func main(){
//	fmt.Println(GetVirtualMemExInfo())
//	fmt.Println(GetVirtualMemInfo())
//}
