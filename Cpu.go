package main

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
)

//Info			CPU基本信息
//LogicalCount	逻辑CPU数量
//PhysicalCount	物理CPU数量
//Usage			CPU使用情况
//Time			CPU有关时间
type Cpu struct {
	Info          []cpu.InfoStat  `json:"info"`
	LogicalCount  int             `json:"logicalCount"`
	PhysicalCount int             `json:"physicalCount"`
	Usage         []float64       `json:"usage"`
	Time          []cpu.TimesStat `json:"time"`
}

func GetCpu() Cpu {
	//cpu基本信息
	cpuInfos, err := cpu.Info()
	if err != nil {
		panic(err)
	}
	//cpu数量;true逻辑核心数量，false物理核心数量
	cpuLogicalCount, err := cpu.Counts(true)
	if err != nil {
		panic(err)
	}
	cpuPhysicalCount, err := cpu.Counts(false)
	if err != nil {
		panic(err)
	}
	//cpu利用率;true为每个cpu，false为总的cpu
	cpuUsage, err := cpu.Percent(time.Second, true)
	if err != nil {
		panic(err)
	}
	//cpu有关时间信息;true为每个cpu，false为总的cpu
	cpuTime, err := cpu.Times(true)
	cpu := Cpu{
		Info:          cpuInfos,
		LogicalCount:  cpuLogicalCount,
		PhysicalCount: cpuPhysicalCount,
		Usage:         cpuUsage,
		Time:          cpuTime,
	}

	return cpu
}

/*
cpu基本信息输出各项的含义：
processor　：系统中逻辑处理核的编号。对于单核处理器，则课认为是其CPU编号，对于多核处理器则可以是物理核、或者使用超线程技术虚拟的逻辑核
vendor_id　：CPU制造商
cpu family　：CPU产品系列代号
model　　　：CPU属于其系列中的哪一代的代号
model name：CPU属于的名字及其编号、标称主频
stepping　 ：CPU属于制作更新版本
cpu MHz　 ：CPU的实际使用主频
cache size ：CPU二级缓存大小
physical id ：单个CPU的标号
siblings ：单个CPU逻辑物理核数
core id ：当前物理核在其所处CPU中的编号，这个编号不一定连续
cpu cores ：该逻辑核所处CPU的物理核数
apicid ：用来区分不同逻辑核的编号，系统中每个逻辑核的此编号必然不同，此编号不一定连续
fpu ：是否具有浮点运算单元（Floating Point Unit）
fpu_exception ：是否支持浮点计算异常
cpuid level ：执行cpuid指令前，eax寄存器中的值，根据不同的值cpuid指令会返回不同的内容
wp ：表明当前CPU是否在内核态支持对用户空间的写保护（Write Protection）
flags ：当前CPU支持的功能
bogomips ：在系统内核启动时粗略测算的CPU速度（Million Instructions Per Second）
clflush size ：每次刷新缓存的大小单位
cache_alignment ：缓存地址对齐单位
address sizes ：可访问地址空间位数
power management ：对能源管理的支持，有以下几个可选支持功能：
　ts：　　temperature sensor
　fid：　 frequency id control
　vid：　 voltage id control
　ttp：　 thermal trip
　tm：
　stc：
　100mhzsteps：
　hwpstate：
CPU信息中flags各项含义：
fpu： Onboard (x87) Floating Point Unit
vme： Virtual Mode Extension
de： Debugging Extensions
pse： Page Size Extensions
tsc： Time Stamp Counter: support for RDTSC and WRTSC instructions
msr： Model-Specific Registers
pae： Physical Address Extensions: ability to access 64GB of memory; only 4GB can be accessed at a time though
mce： Machine Check Architecture
cx8： CMPXCHG8 instruction
apic： Onboard Advanced Programmable Interrupt Controller
sep： Sysenter/Sysexit Instructions; SYSENTER is used for jumps to kernel memory during system calls, and SYSEXIT is used for jumps： back to the user code
mtrr： Memory Type Range Registers
pge： Page Global Enable
mca： Machine Check Architecture
cmov： CMOV instruction
pat： Page Attribute Table
pse36： 36-bit Page Size Extensions: allows to map 4 MB pages into the first 64GB RAM, used with PSE.
pn： Processor Serial-Number; only available on Pentium 3
clflush： CLFLUSH instruction
dtes： Debug Trace Store
acpi： ACPI via MSR
mmx： MultiMedia Extension
fxsr： FXSAVE and FXSTOR instructions
sse： Streaming SIMD Extensions. Single instruction multiple data. Lets you do a bunch of the same operation on different pieces of input： in a single clock tick.
sse2： Streaming SIMD Extensions-2. More of the same.
selfsnoop： CPU self snoop
acc： Automatic Clock Control
IA64： IA-64 processor Itanium.
ht： HyperThreading. Introduces an imaginary second processor that doesn’t do much but lets you run threads in the same process a bit quicker.
nx： No Execute bit. Prevents arbitrary code running via buffer overflows.
pni： Prescott New Instructions aka. SSE3
vmx： Intel Vanderpool hardware virtualization technology
svm： AMD “Pacifica” hardware virtualization technology
lm： “Long Mode,” which means the chip supports the AMD64 instruction set
tm： “Thermal Monitor” Thermal throttling with IDLE instructions. Usually hardware controlled in response to CPU temperature.
tm2： “Thermal Monitor 2″ Decrease speed by reducing multipler and vcore.
est： “Enhanced SpeedStep”
*/
/*
cpu有关时间输出各项的含义：
user:用户态的CPU时间
nice：低优先级程序所占用的用户态的cpu时间。
system：系统态的CPU时间
idle：CPU空闲的时间（不包含IO等待）
iowait：等待IO响应的时间
irq：处理硬件中断的时间
softirq：处理软中断的时间
steal:其他系统所花的时间（个人理解是针对虚拟机）
guest：运行时间为客户操作系统下的虚拟CPU控制（个人理解是访客控制CPU的时间）
guest_nice：低优先级程序所占用的用户态的cpu时间。（访客的）
*/
