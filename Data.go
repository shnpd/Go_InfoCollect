package main

//
//import (
//	"github.com/shirou/gopsutil/cpu"
//	"github.com/shirou/gopsutil/disk"
//	"github.com/shirou/gopsutil/mem"
//	"github.com/shirou/gopsutil/process"
//	"net"
//)
//
//type Application struct {
//	DebInstall string `json:"debInstall"`
//	rpmInstall string `json:"rpmInstall"`
//	yumInstall string `json:"yumInstall"`
//}
//
//type Cpu struct {
//	Info          []cpu.InfoStat  `json:"info"`
//	LogicalCount  int             `json:"logicalCount"`
//	PhysicalCount int             `json:"physicalCount"`
//	Usage         []float64       `json:"usage"`
//	Time          []cpu.TimesStat `json:"time"`
//}
//
//type DataAmount struct {
//	DataAmount string `json:"dataAmount"`
//}
//
//type Device struct {
//	Name         string `json:"name"`
//	Manufacturer string `json:"manufacturer"`
//	SerialNumber string `json:"serialNumber"`
//	Version      string `json:"version"`
//}
//
//type Disk struct {
//	Partitions   []disk.PartitionStat             `json:"partitions"`
//	Usage        []*disk.UsageStat                `json:"usage"`
//	SerialNumber []string                         `json:"serialNumber"`
//	Label        []string                         `json:"label"`
//	IO           []map[string]disk.IOCountersStat `json:"io"`
//}
//
//type FireWall struct {
//	Iptables_status string `json:"iptablesStatus"`
//	Iptables_rules  string `json:"iptablesRules"`
//	Firewall_status string `json:"firewallStatus"`
//	Firewall_rules  string `json:"firewallRules"`
//}
//
////Host
//type InfoStat struct {
//	Hostname             string `json:"hostname"`
//	Uptime               uint64 `json:"uptime"`
//	BootTime             uint64 `json:"bootTime"`
//	Procs                uint64 `json:"procs"`           // number of processes
//	OS                   string `json:"os"`              // ex: freebsd, linux
//	Platform             string `json:"platform"`        // ex: ubuntu, linuxmint
//	PlatformFamily       string `json:"platformFamily"`  // ex: debian, rhel
//	PlatformVersion      string `json:"platformVersion"` // version of the complete OS
//	KernelVersion        string `json:"kernelVersion"`   // version of the OS kernel (if available)
//	KernelArch           string `json:"kernelArch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
//	VirtualizationSystem string `json:"virtualizationSystem"`
//	VirtualizationRole   string `json:"virtualizationRole"` // guest or host
//	HostID               string `json:"hostid"`             // ex: uuid
//}
//
//type Log struct {
//	SystemLog string `json:"systemLog"`
//	SecureLog string `json:"secureLog"`
//}
//
//type Memory struct {
//	SwapMemory      *mem.SwapMemoryStat      `json:"swapMemory"`
//	VirtualMemory   *mem.VirtualMemoryStat   `json:"virtualMemory"`
//	VirtualMemoryEx *mem.VirtualMemoryExStat `json:"virtualMemoryEx"`
//}
//
//type Network struct {
//	IP      map[string]string `json:"ip"`
//	MAC     map[string]string `json:"mac"`
//	GateWay []GateWay         `json:"gateway"`
//}
//
//type NetInterface struct {
//	Index        int       `json:"index"`
//	MTU          int       `json:"mtu"`
//	Name         string    `json:"name"`
//	HardwareAddr string    `json:"hardwareAddr"`
//	Flags        net.Flags `json:"flags"`
//}
//
//type Process struct {
//	Pid       int32                     `json:"pid"`
//	Name      string                    `json:"name"`
//	MemoryUse *process.MemoryInfoExStat `json:"memoryUse"`
//	CpuUse    float64                   `json:"cpuUse"`
//	Uids      []int32                   `json:"uids"`
//}
//
//type Service struct {
//	RunningService string `json:"runningService"`
//	AllService     string `json:"allService"`
//}
