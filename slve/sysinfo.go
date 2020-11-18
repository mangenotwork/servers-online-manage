//获取计算机的系统信息的实现
package slve

import (
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/mangenotwork/csdemo/lib/cmd"
	"github.com/mangenotwork/csdemo/structs"
)

//获取本机ip
func GetMyIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Println("[Error] :", err)
		return ""
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String())
	// ip := strings.Split(localAddr.String(), ":")[0]
	return localAddr.String()

}

func GetSysType() string {
	return runtime.GOOS
}

func GetHostName() string {
	name, err := os.Hostname()
	if err != nil {
		name = "null"
	}
	return name
}

//获取系统信息
func SysInfo() {
	log.Println(`系统类型：`, runtime.GOOS)
	log.Println(`系统架构：`, runtime.GOARCH)
	log.Println(`CPU 核数：`, runtime.GOMAXPROCS(0))
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	log.Println(`电脑名称：`, name)
}

func GetHostInfo() *structs.HostInfo {
	return &structs.HostInfo{
		HostName:      GetHostName(),
		SysType:       runtime.GOOS,
		SysArch:       runtime.GOARCH,
		CpuCoreNumber: fmt.Sprintf("cpu 核心数: %d", runtime.GOMAXPROCS(0)),
	}
}

//获取system-uuid
func GetSystemUUID() string {
	return cmd.LinuxSendCommand("sudo dmidecode -s system-uuid")
}

// 从 /proc/cpuinfo 获取cpu相关信息
func ProcCpuinfo() (cpuinfos []map[string]string) {

	cpuinfos = make([]map[string]string, 0)

	rStr := cmd.LinuxSendCommand("cat /proc/cpuinfo")
	if rStr == "" {
		return
	}

	rStrList := strings.Split(rStr, "processor")
	for _, v := range rStrList {
		v = "processor" + v
		vList := strings.Split(v, "\n")
		data := make(map[string]string, 0)
		for _, i := range vList {
			d := strings.Split(i, ":")
			if len(d) == 2 {
				key := DeletePreAndSufSpace(d[0])
				vlue := DeletePreAndSufSpace(d[1])
				data[key] = vlue
			}
		}
		cpuinfos = append(cpuinfos, data)
	}
	return
}

// 从 /proc/meminfo 中读取内存
func ProcMeminfo() (mem *structs.ProcMemInfo) {
	rStr := cmd.LinuxSendCommand("cat /proc/meminfo")
	log.Println(rStr)

	data := make(map[string]string, 0)
	rStrList := strings.Split(rStr, "\n")
	for _, v := range rStrList {
		d := strings.Split(v, ":")
		if len(d) == 2 {
			key := DeletePreAndSufSpace(d[0])
			vlue := DeletePreAndSufSpace(d[1])
			data[key] = vlue
		}
	}
	log.Println(data)

	//MemTotal: 所有可用RAM大小 （即物理内存减去一些预留位和内核的二进制代码大小）
	memTotal := Str2Int64(data["MemTotal"])

	//MemFree: LowFree与HighFree的总和，被系统留着未使用的内存
	memFree := Str2Int64(data["MemFree"])

	memUsed := memTotal - memFree

	//Buffers: 用来给文件做缓冲大小
	buffers := Str2Int64(data["Buffers"])

	//Cached: 被高速缓冲存储器（cache memory）用的内存的大小（等于diskcache minus SwapCache ）.
	cached := Str2Int64(data["Cached"])

	log.Println(memTotal, memUsed, memFree, buffers, cached)
	log.Println("Used = ", memTotal-memFree)
	log.Println("-buffers/cache反映的是被程序实实在在吃掉的内存 : ", memUsed-buffers-cached)
	log.Println("+buffers/cache反映的是可以挪用的内存总数 : ", memFree+buffers+cached)

	mem = &structs.ProcMemInfo{
		MemTotal:   memTotal,
		MemUsed:    memUsed - buffers - cached,
		MemFree:    memFree + buffers + cached,
		MemBuffers: buffers,
		MemCached:  cached,
	}

	return
}

/*
每个独立进程: /proc/1/*
cmdline — 启动当前进程的完整命令，但僵尸进程目录中的此文件不包含任何信息；
cwd — 指向当前进程运行目录的一个符号链接；
environ — 当前进程的环境变量列表，彼此间用空字符（NULL）隔开；变量用大写字母表示，其值用小写字母表示；
exe — 指向启动当前进程的可执行文件（完整路径）的符号链接，通过/proc/N/exe可以启动当前进程的一个拷贝；
fd — 这是个目录，包含当前进程打开的每一个文件的文件描述符（file descriptor），这些文件描述符是指向实际文件的一个符号链接；
limits — 当前进程所使用的每一个受限资源的软限制、硬限制和管理单元；此文件仅可由实际启动当前进程的UID用户读取；（2.6.24以后的内核版本支持此功能）；
maps — 当前进程关联到的每个可执行文件和库文件在内存中的映射区域及其访问权限所组成的列表；
mem — 当前进程所占用的内存空间，由open、read和lseek等系统调用使用，不能被用户读取； （用户读不到）
root — 指向当前进程运行根目录的符号链接；在Unix和Linux系统上，通常采用chroot命令使每个进程运行于独立的根目录；
stat — 当前进程的状态信息，包含一系统格式化后的数据列，可读性差，通常由ps命令使用；
statm — 当前进程占用内存的状态信息，通常以“页面”（page）表示；
status — 与stat所提供信息类似，但可读性较好，如下所示，每行表示一个属性信息；其详细介绍请参见 proc的man手册页；

全局: /proc/*
/proc/buddyinfo  --  用于诊断内存碎片问题的相关信息文件；
/proc/cmdline  --  在启动时传递至内核的相关参数信息，这些信息通常由lilo或grub等启动管理工具进行传递；
/proc/cpuinfo  -- 处理器的相关信息的文件；
/proc/crypto  -- 系统上已安装的内核使用的密码算法及每个算法的详细信息列表；
/proc/devices  -- 系统已经加载的所有块设备和字符设备的信息，包含主设备号和设备组（与主设备号对应的设备类型）名；
/proc/diskstats  -- 每块磁盘设备的磁盘I/O统计信息列表；（内核2.5.69以后的版本支持此功能）
/proc/dma  -- 每个正在使用且注册的ISA DMA通道的信息列表；
/proc/execdomains  --  内核当前支持的执行域（每种操作系统独特“个性”）信息列表；
/proc/fb  -- 帧缓冲设备列表文件，包含帧缓冲设备的设备号和相关驱动信息；
/proc/filesystems  -- 当前被内核支持的文件系统类型列表文件，被标示为nodev的文件系统表示不需要块设备的支持；通常mount一个设备时，如果没有指定文件系统类型将通过此文件来决定其所需文件系统的类型；
/proc/interrupts  -- X86或X86_64体系架构系统上每个IRQ相关的中断号列表；多路处理器平台上每个CPU对于每个I/O设备均有自己的中断号；
/proc/iomem  -- 每个物理设备上的记忆体（RAM或者ROM）在系统内存中的映射信息；
/proc/ioports  -- 当前正在使用且已经注册过的与物理设备进行通讯的输入-输出端口范围信息列表；如下面所示，第一列表示注册的I/O端口范围，其后表示相关的设备；
/proc/kallsyms   -- 模块管理工具用来动态链接或绑定可装载模块的符号定义，由内核输出；（内核2.5.71以后的版本支持此功能）；通常这个文件中的信息量相当大；
/proc/kcore  -- 系统使用的物理内存，以ELF核心文件（core file）格式存储，其文件大小为已使用的物理内存（RAM）加上4KB；这个文件用来检查内核数据结构的当前状态，
				因此，通常由GBD通常调试工具使用，但不能使用文件查看命令打开此文件；
/proc/kmsg   -- 此文件用来保存由内核输出的信息，通常由/sbin/klogd或/bin/dmsg等程序使用，不要试图使用查看命令打开此文件；
/proc/loadavg  -- 保存关于CPU和磁盘I/O的负载平均值，其前三列分别表示每1秒钟、每5秒钟及每15秒的负载平均值，类似于uptime命令输出的相关信息；第四列是由斜线隔开的两个数值，
				前者表示当前正由内核调度的实体（进程和线程）的数目，后者表示系统当前存活的内核调度实体的数目；第五列表示此文件被查看前最近一个由内核创建的进程的PID；
/proc/locks  -- 保存当前由内核锁定的文件的相关信息，包含内核内部的调试数据；每个锁定占据一行，且具有一个惟一的编号；如下输出信息中每行的第二列表示当前锁定使用的锁定类别，
				POSIX表示目前较新类型的文件锁，由lockf系统调用产生，FLOCK是传统的UNIX文件锁，由flock系统调用产生；第三列也通常由两种类型，ADVISORY表示不允许其他用户锁定此文件，
				但允许读取，MANDATORY表示此文件锁定期间不允许其他用户任何形式的访问；
/proc/mdstat  -- 保存RAID相关的多块磁盘的当前状态信息，在没有使用RAID机器上，其显示为如下状态：
/proc/meminfo  -- 系统中关于当前内存的利用状况等的信息，常由free命令使用；可以使用文件查看命令直接读取此文件，其内容显示为两列，前者为统计属性，后者为对应的值；
/proc/modules  -- 当前装入内核的所有模块名称列表，可以由lsmod命令使用，也可以直接查看；如下所示，其中第一列表示模块名，第二列表示此模块占用内存空间大小，
				第三列表示此模块有多少实例被装入，第四列表示此模块依赖于其它哪些模块，第五列表示此模块的装载状态（Live：已经装入；Loading：正在装入；Unloading：正在卸载），
				第六列表示此模块在内核内存（kernel memory）中的偏移量；
/proc/partitions  -- 块设备每个分区的主设备号（major）和次设备号（minor）等信息，同时包括每个分区所包含的块（block）数目（如下面输出中第三列所示）；
/proc/pci  -- 内核初始化时发现的所有PCI设备及其配置信息列表，其配置信息多为某PCI设备相关IRQ信息，可读性不高，可以用“/sbin/lspci –vb”命令获得较易理解的相关信息；
				在2.6内核以后，此文件已为/proc/bus/pci目录及其下的文件代替；
/proc/slabinfo  -- 在内核中频繁使用的对象（如inode、dentry等）都有自己的cache，即slab pool，而/proc/slabinfo文件列出了这些对象相关slap的信息；详情可以参见内核文档中slapinfo的手册页；

/proc/stat
实时追踪自系统上次启动以来的多种统计信息；如下所示，其中，
“cpu”行后的八个值分别表示以1/100（jiffies）秒为单位的统计值（包括系统运行于用户模式、低优先级用户模式，运系统模式、空闲模式、I/O等待模式的时间等）；
“intr”行给出中断的信息，第一个为自系统启动以来，发生的所有的中断的次数；然后每个数对应一个特定的中断自系统启动以来所发生的次数；
“ctxt”给出了自系统启动以来CPU发生的上下文交换的次数。
“btime”给出了从系统启动到现在为止的时间，单位为秒；
“processes (total_forks) 自系统启动以来所创建的任务的个数目；
“procs_running”：当前运行队列的任务的数目；
“procs_blocked”：当前被阻塞的任务的数目；

/proc/swaps
当前系统上的交换分区及其空间利用信息，如果有多个交换分区的话，则会每个交换分区的信息分别存储于/proc/swap目录中的单独文件中，而其优先级数字越低，被使用到的可能性越大；下面是作者系统中只有一个交换分区时的输出信息；

/proc/uptime   --  系统上次启动以来的运行时间，如下所示，其第一个数字表示系统运行时间，第二个数字表示系统空闲时间，单位是秒；

/proc/version
当前系统运行的内核版本号，在作者的RHEL5.3上还会显示系统安装的gcc版本，如下所示；

/proc/vmstat
当前系统虚拟内存的多种统计数据，信息量可能会比较大，这因系统而有所不同，可读性较好；下面为作者机器上输出信息的一个片段；（2.6以后的内核支持此文件）

/proc/zoneinfo
内存区域（zone）的详细信息列表，信息量较大，


设备唯一id
sudo dmidecode -s system-uuid


*/

//删除字符串前后两端的所有空格
func DeletePreAndSufSpace(str string) string {
	strList := []byte(str)
	spaceCount, count := 0, len(strList)
	for i := 0; i <= len(strList)-1; i++ {
		if strList[i] == 32 {
			spaceCount++
		} else {
			break
		}
	}

	strList = strList[spaceCount:]
	spaceCount, count = 0, len(strList)
	for i := count - 1; i >= 0; i-- {
		if strList[i] == 32 {
			spaceCount++
		} else {
			break
		}
	}

	return string(strList[:count-spaceCount])
}

//字符串转flot64
func Str2Flot64(s string) float64 {
	floatnum, err := strconv.ParseFloat(s, 64)
	if err != nil {
		floatnum = 0
	}
	return floatnum
}

//字符串转int64
func Str2Int64(s string) int64 {
	reg := regexp.MustCompile(`[0-9]+`)
	sList := reg.FindAllString(s, -1)
	log.Println(sList)
	if len(sList) == 0 {
		return 0
	}

	int64num, err := strconv.ParseInt(sList[0], 10, 64)
	if err != nil {
		return 0
	}
	return int64num
}
