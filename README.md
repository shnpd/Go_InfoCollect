# 项目介绍

本项目主要通过gopsutil库和直接调用linux命令两种方法搜集linux下的各种信息，每隔1s以json文件的形式输出

gopsutil库是python中的psutil库在Golang上的移植版，主要用于收集主机的各种信息，包括网络信息，进程信息，硬件信息等
具体的使用方法可以参考官方文档，这里不再赘述

直接调用linux命令即通过go语言调用执行linux的命令，然后将命令结果通过管道回显，通过此方法有些命令需要密码验证才能执行

---

**目前所能采集到的数据**：

**通过gopsutil采集**

CPU：基本信息、逻辑数量、物理数量、使用情况、时间信息

磁盘：分区、使用情况、序列号、标签、I/O信息

主机信息：主机名、运行时间、操作系统、系统版本、系统内核

内存：交换内存、虚拟内存

进程：进程ID、名称、内存占用、CPU占用

---

**通过linux命令采集**

应用：安装的应用软件、版本、描述

数据量

防火墙：iptables的状态和规则、firewall的状态和规则

日志：系统日志、安全日志

服务：正在运行的系统服务、全部系统服务

设备信息：设备名、设备厂商、设备编码、设备型号

---

**通过go内置包采集**

网络信息：ip地址、MAC地址、网关

网络接口：接口名称、最大传送单元、接口标志、接口地址

---


# 项目使用
1. 首先将本项目拉取到gopath/src目录下，注意必须是gopath目录，否则在编译的时候会出错
>可以通过`go env`查看自己的gopath目录
2. 讲项目拉取到本地后进入项目目录中通过`go build`命令编译生成可执行文件
>命令后加项目的具体路径，例如：`go build /opt/gopath/src/Go_InfoCollect`
3. 这时在项目目录下会生成一个collect文件，可以直接通过`./collect`的方式运行，也可以将其挂在后台运行
>挂在后台运行的方法：
>1. 修改文件执行权限 `chmod 777 collect`
>2. 程序后台运行 `nohup ./collect &`
>
>关闭进程的方法：
>1. 查询进程号 `ps aux|grep collect`
>2. 关闭进程 `kill 进程编号`

