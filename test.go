package main

import "time"

func main() {

	for {
		GetAppInfo()
		getCpu()
		GetDataAmount()
		Getdevice()
		GetDisk()
		GetFireWall()
		GetHostInfo()
		GetLog()
		getMemory()
		GetNetInfo()
		GetNetInterface()
		GetProcess()
		GetService()
		time.Sleep(time.Second * 1)
	}
}
