# 项目介绍

本项目主要通过gopsutil库和直接调用linux命令两种方法搜集linux下的各种信息，每隔1s以json文件的形式输出

gopsutil库是python中的psutil库在Golang上的移植版，主要用于收集主机的各种信息，包括网络信息，进程信息，硬件信息等
具体的使用方法可以参考官方文档，这里不再赘述

直接调用linux命令即通过go语言调用执行linux的命令，然后将命令结果通过管道回显，通过此方法有些命令需要密码验证才能执行

# 项目使用
1. 首先将本项目拉取到gopath/src目录下，注意必须是gopath目录，否则在编译的时候会出错
>可以通过`go env`查看自己的gopath目录
2. 讲项目拉取到本地后进入项目目录中通过`go build`命令编译生成可执行文件
>命令后加项目的具体路径，例如：`go build /opt/gopath/src/Go_InfoCollect`
3. 这时在项目目录下会生成一个collect文件，可以直接通过`./collect`的方式运行，也可以将其挂在后台运行
>挂在后台运行的方法：
>1. 修改文件执行权限 `chmod 777 collect`
>2. 程序后台运行 `nohup ./collect &`
>关闭进程的方法：
>1. 查询进程号 `ps aux|grep collect`
>2. 关闭进程 `kill 进程编号`

