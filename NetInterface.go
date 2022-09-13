package main

import (
	"net"
)

//Index			序号
//MTU			最大传输单元
//Name			名称
//HardwareAddr	硬件地址
//Flags			标志
type NetInterface struct {
	Index        int       `json:"index"`
	MTU          int       `json:"mtu"`
	Name         string    `json:"name"`
	HardwareAddr string    `json:"hardwareAddr"`
	Flags        net.Flags `json:"flags"`
}

//网络接口信息采集
func GetNetInterface() []NetInterface {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	var netInterfaces []NetInterface
	for _, n := range interfaces {
		temp := NetInterface{
			Index:        n.Index,
			MTU:          n.MTU,
			Name:         n.Name,
			HardwareAddr: n.HardwareAddr.String(),
			Flags:        n.Flags,
		}
		netInterfaces = append(netInterfaces, temp)
	}
	return netInterfaces
}

//func main() {
//	interfaces:=GetNetInterfaceInfo()
//	fmt.Println(interfaces)
//	for _, inter := range interfaces {
//		fmt.Println("--------------------")
//		fmt.Println("接口名称:",inter.Name)
//		fmt.Println("最大传送单元:",inter.MTU)
//		fmt.Println("接口标志:",inter.Flags)
//		fmt.Println("接口地址:",inter.HardwareAddr)
//	}
//}
