package main

import (
	"github.com/shirou/gopsutil/disk"
)

type Disk struct {
	Partitions   []disk.PartitionStat             `json:"partitions"`
	Usage        []*disk.UsageStat                `json:"usage"`
	SerialNumber []string                         `json:"serialNumber"`
	Label        []string                         `json:"label"`
	IO           []map[string]disk.IOCountersStat `json:"io"`
}

//获取磁盘使用情况
func GetUsage(path string) *disk.UsageStat {
	usage, err := disk.Usage(path)
	if err == nil {
		return usage
	} else {
		return nil
	}
}

//得到磁盘序列号
func GetSerialNumber(name string) string {
	//获取路径为name的磁盘的序列号
	sn := disk.GetDiskSerialNumber(name)
	return sn
}

//得到磁盘标签
func GetLabel(name string) string {
	label := disk.GetLabel(name)
	return label
}

//得到磁盘IO信息
func GetIO(name string) map[string]disk.IOCountersStat {
	IO, err := disk.IOCounters(name)
	if err == nil {
		return IO
	} else {
		return nil
	}
}

func GetDisk() Disk {
	//获取磁盘分区
	//如果all为false，只返回物理设备(如:硬盘、cd-rom驱动器、USB keys)，忽略其他所有设备(如:内存分区，如/dev/shm)
	DiskParti, err := disk.Partitions(false)
	if err != nil {
		panic(err)
	}
	//获取分区路径
	var paths []string
	for _, n := range DiskParti {
		paths = append(paths, n.Device)
	}
	//获取磁盘使用情况
	var diskUsage []*disk.UsageStat
	for _, path := range paths {
		diskUsage = append(diskUsage, GetUsage(path))
	}
	//得到磁盘序列号
	var diskSerialNumber []string
	for _, path := range paths {
		diskSerialNumber = append(diskSerialNumber, GetSerialNumber(path))
	}
	//得到磁盘标签
	var diskLabel []string
	for _, path := range paths {
		diskLabel = append(diskLabel, GetLabel(path))
	}
	//得到磁盘IO信息
	var diskIO []map[string]disk.IOCountersStat
	for _, path := range paths {
		diskIO = append(diskIO, GetIO(path))
	}
	disk := Disk{
		Partitions:   DiskParti,
		Usage:        diskUsage,
		SerialNumber: diskSerialNumber,
		Label:        diskLabel,
		IO:           diskIO,
	}
	return disk
}

//fmt.Println(GetPartitions())
// /dev/sda1
// /dev/sr0
//	fmt.Println(GetUsage("/dev/sr0"))
//fmt.Println(GetIO("/dev/sr0"))
