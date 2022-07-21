package main

import (
	"encoding/json"
	"github.com/shirou/gopsutil/host"
)

type Linux struct {
	Host         *host.InfoStat `json:"host"`
	Cpu          Cpu            `json:"cpu"`
	Disk         Disk           `json:"disk"`
	Memory       Memory         `json:"memory"`
	Network      Network        `json:"network"`
	NetInterface []NetInterface `json:"netInterface"`
	Process      []Process      `json:"process"`
	Service      Service        `json:"service"`
	Device       Device         `json:"device"`
	Log          Log            `json:"log"`
	FireWall     FireWall       `json:"firewall"`
	DataAmount   DataAmount     `json:"dataAmount"`
	Application  Application    `json:"application"`
}

func main() {
	cpu := GetCpu()
	disk := GetDisk()
	application := GetAppInfo()
	dataamount := GetDataAmount()
	device := Getdevice()
	firewall := GetFireWall()
	host := GetHostInfo()
	log := GetLog()
	memory := GetMemory()
	network := GetNetInfo()
	netinterface := GetNetInterface()
	process := GetProcess()
	service := GetService()

	linux := Linux{
		Host:         host,
		Cpu:          cpu,
		Disk:         disk,
		Memory:       memory,
		Network:      network,
		NetInterface: netinterface,
		Process:      process,
		Service:      service,
		Device:       device,
		Log:          log,
		FireWall:     firewall,
		DataAmount:   dataamount,
		Application:  application,
	}
	Linuxes, err := json.Marshal(linux)
	if err != nil {
		panic(err)
	}
	WriteFile("Linux.json", Linuxes)
}
