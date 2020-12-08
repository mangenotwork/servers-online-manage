// 获取计算机的系统信息的实现
// Linux
package linux

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
	"time"

	"github.com/mangenotwork/servers-online-manage/lib/cmd"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
	"github.com/mangenotwork/servers-online-manage/lib/utils"
)


//获取system-uuid
func GetSystemUUID() string {
	return cmd.LinuxSendCommand("sudo dmidecode -s system-uuid")
}

// 通过df 采样磁盘的基本
func GetSystemDF() (diskinfos []*structs.DiskInfo, allTotal int){
	diskinfos = make([]*structs.DiskInfo,0)
	rStr := cmd.LinuxSendCommand("df -m")
	if rStr == "" {
		return
	}
	rStrList := strings.Split(rStr,"\n")
	if len(rStrList) <2 {
		return
	}
	for _, v := range rStrList[1:len(rStrList)] {
		if v == "" {
			continue
		}
		log.Println(v)
		vList := strings.Split(v," ")
		nList := []string{}
		for _,n := range vList{
			if n == ""{
				continue
			}
			nList = append(nList,n)
		}
		log.Println(nList, len(nList))
		if len(nList) > 5 {
			diskinfo := &structs.DiskInfo{
				DiskName: nList[0],
				DistType: "",
				DistTotalMB : nList[1],
			}
			total := utils.Num2Int(nList[0])
			allTotal += total
			diskinfo.DistUse = &structs.DiskUseInfo{
				Total: total,
				Free: utils.Num2Int(nList[3]),
				Rate: float32(utils.Str2Int64(nList[4])),
			}
			diskinfos = append(diskinfos, diskinfo)
		}
	}
	log.Println(rStr)
	return
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
				key := utils.DeletePreAndSufSpace(d[0])
				vlue := utils.DeletePreAndSufSpace(d[1])
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
	if rStr == "" {
		return
	}
	data := make(map[string]string, 0)
	rStrList := strings.Split(rStr, "\n")
	for _, v := range rStrList {
		d := strings.Split(v, ":")
		if len(d) == 2 {
			key := utils.DeletePreAndSufSpace(d[0])
			vlue := utils.DeletePreAndSufSpace(d[1])
			data[key] = vlue
		}
	}

	memTotal := utils.Str2Int64(data["MemTotal"])
	memFree := utils.Str2Int64(data["MemFree"])
	buffers := utils.Str2Int64(data["Buffers"])
	cached := utils.Str2Int64(data["Cached"])
	mem = &structs.ProcMemInfo{
		MemTotal:   memTotal,
		MemUsed:    memTotal - memFree - buffers - cached,
		MemFree:    memFree + buffers + cached,
		MemBuffers: buffers,
		MemCached:  cached,
	}
	return
}

// /proc/stat 获取cpu 信息
func GetProcStat() (datas []*structs.ProcStatCPUData) {
	datas = make([]*structs.ProcStatCPUData, 0)
	rStr := cmd.LinuxSendCommand("cat /proc/stat")
	if rStr == "" {
		return
	}
	rStrList := strings.Split(rStr, "\n")
	for _, v := range rStrList {
		if len(v) > 3 {
			if v[:3] == "cpu" {
				vStr := strings.Split(v, " ")
				vList := make([]string, 0)
				for _, s := range vStr {
					s = utils.DeletePreAndSufSpace(s)
					if s != "" {
						vList = append(vList, s)
					}
				}
				//cpu名
				name := vList[0]
				//user 从系统启动开始累计到当前时刻，处于用户态的运行时间，不包含 nice值为负的进程。
				user := utils.Num2Int64(vList[1])

				//system 从系统启动开始累计到当前时刻，处于核心态的运行时间。
				system := utils.Num2Int64(vList[2])

				//nice 从系统启动开始累计到当前时刻，nice值为负的进程所占用的CPU时间。
				nice := utils.Num2Int64(vList[3])

				//idle 从系统启动开始累计到当前时刻，除IO等待时间以外的其它等待时间。
				idle := utils.Num2Int64(vList[4])

				//iowait 从系统启动开始累计到当前时刻，IO等待时间(since 2.5.41)。
				iowait := utils.Num2Int64(vList[5])

				//irq 从系统启动开始累计到当前时刻，硬中断时间(since 2.6.0-test4)。
				irq := utils.Num2Int64(vList[6])

				//softirq  从系统启动开始累计到当前时刻，软中断时间(since2.6.0-test4)。
				softirq := utils.Num2Int64(vList[7])

				//stealstolen  which is the time spent in otheroperating systems
				//when running in a virtualized environment(since 2.6.11)
				stealstolen := utils.Num2Int64(vList[8])

				//guest whichis the time spent running a virtual CPU  for  guest
				//operating systems under the control ofthe Linux kernel(since 2.6.24)。
				guest := utils.Num2Int64(vList[9])

				//log.Println(name, user, system, nice, idle, iowait, irq, softirq, stealstolen, guest)

				//总的cpu时间totalCpuTime = user + nice + system + idle + iowait + irq + softirq + stealstolen  +  guest
				totalCpuTime := user + nice + system + idle + iowait + irq + softirq + stealstolen + guest

				//user+nice+system+irq+softirq
				userCpuTime := user + nice + system + irq + softirq

				cpudata := &structs.ProcStatCPUData{
					Name:  name,
					Total: totalCpuTime,
					Used:  userCpuTime,
					Idle:  idle,
				}
				datas = append(datas, cpudata)
			}
		}
	}
	return
}

//总cpu使用率的计算：
// 1、采样两个足够短的时间间隔的Cpu快照，分别记作t1,t2，其中t1、t2的结构均为
// 2、计算总的Cpu时间片totalCpuTime
// a)   把第一次的所有cpu使用情况求和，得到s1;
// b)   把第二次的所有cpu使用情况求和，得到s2;
// c)   s2 - s1得到这个时间间隔内的所有时间片，即totalCpuTime = j2 - j1 ;
// 3、计算空闲时间idle
// idle对应第四列的数据，用第二次的第四列 - 第一次的第四列即可
// idle=第二次的第四列 - 第一次的第四列
// 4、计算cpu使用率
// pcpu =100* (total-idle)/total
func ProcStat() {
	var pcpu float64

	procStat1 := GetProcStat()
	//睡眠延时500ms
	time.Sleep(500 * time.Millisecond)
	procStat2 := GetProcStat()
	for _, t1 := range procStat1 {
		for _, t2 := range procStat2 {
			if t1.Name == t2.Name {
				total := t2.Total - t1.Total
				idle := t2.Idle - t1.Idle
				pcpu = 100 * float64((total - idle)) / float64(total)
				log.Println("cpu Name = ", t1.Name)
				log.Println("pcpu = ", pcpu)
				log.Println("__________________________")
			}
		}
	}
}

//指定进程的 /proc/*/stat 获取cpu 信息
/*
每个参数意思为：
参数                                                       解释
		0  pid=6873                                              进程(包括轻量级进程，即线程)号
		1  comm=a.out                                          应用程序或命令的名字
		2  task_state=R                                        任务的状态，R:runnign, S:sleeping (TASK_INTERRUPTIBLE), D:disk sleep (TASK_UNINTERRUPTIBLE), T: stopped, T:tracing stop,Z:zombie, X:dead
		3  ppid=6723                                            父进程ID
		4  pgid=6873                                            线程组号
		5  sid=6723                                              该任务所在的会话组ID
		6  tty_nr=34819(pts/3)                            该任务的tty终端的设备号，INT（34817/256）=主设备号，（34817-主设备号）=次设备号
		7  tty_pgrp=6873                                     终端的进程组号，当前运行在该任务所在终端的前台任务(包括shell 应用程序)的PID。
		8  task->flags=8388608                           进程标志位，查看该任务的特性
		9  min_flt=77                                            该任务不需要从硬盘拷数据而发生的缺页（次缺页）的次数
		10  cmin_flt=0                                            累计的该任务的所有的waited-for进程曾经发生的次缺页的次数目
		11  maj_flt=0                                              该任务需要从硬盘拷数据而发生的缺页（主缺页）的次数
		12  cmaj_flt=0                                            累计的该任务的所有的waited-for进程曾经发生的主缺页的次数目
		13  utime=1587                                          该任务在用户态运行的时间，单位为jiffies
		14  stime=1                                                该任务在核心态运行的时间，单位为jiffies
		15  cutime=0                                              累计的该任务的所有的waited-for进程曾经在用户态运行的时间，单位为jiffies
		16  cstime=0                                              累计的该任务的所有的waited-for进程曾经在核心态运行的时间，单位为jiffies
		17  priority=25                                           任务的动态优先级
		18  nice=0                                                  任务的静态优先级
		19  num_threads=3                                    该任务所在的线程组里线程的个数
		20  it_real_value=0                                     由于计时间隔导致的下一个 SIGALRM 发送进程的时延，以 jiffy 为单位.
		21  start_time=5882654                             该任务启动的时间，单位为jiffies
		22  vsize=1409024（page）                       该任务的虚拟地址空间大小
		23  rss=56(page)                                        该任务当前驻留物理地址空间的大小
			 Number of pages the process has in real memory,minu 3 for administrative purpose.
              这些页可能用于代码，数据和栈。
		24  rlim=4294967295（bytes）                  该任务能驻留物理地址空间的最大值
		25  start_code=134512640                        该任务在虚拟地址空间的代码段的起始地址
		26  end_code=134513720                         该任务在虚拟地址空间的代码段的结束地址
		27  start_stack=3215579040                     该任务在虚拟地址空间的栈的结束地址
		28  kstkesp=0                                            esp(32 位堆栈指针) 的当前值, 与在进程的内核堆栈页得到的一致.
		29  kstkeip=2097798                                 指向将要执行的指令的指针, EIP(32 位指令指针)的当前值.
		30  pendingsig=0                                       待处理信号的位图，记录发送给进程的普通信号
		31  block_sig=0                                          阻塞信号的位图
		32  sigign=0                                               忽略的信号的位图
		33  sigcatch=082985                                  被俘获的信号的位图
		34  wchan=0                                               如果该进程是睡眠状态，该值给出调度的调用点
		35  nswap                                                   被swapped的页数，当前没用
		36  cnswap                                                 所有子进程被swapped的页数的和，当前没用
		37  exit_signal=17                                      该进程结束时，向父进程所发送的信号
		38  task_cpu(task)=0                                  运行在哪个CPU上
		39  task_rt_priority=0                                 实时进程的相对优先级别
		40  task_policy=0                                        进程的调度策略，0=非实时进程，1=FIFO实时进程；2=RR实时进程
*/
func GetProcessProcStat(pid string) (data *structs.ProcessProcStatCPUData) {
	data = &structs.ProcessProcStatCPUData{}
	rStr := cmd.LinuxSendCommand(fmt.Sprintf("cat /proc/%s/stat", pid))
	if rStr == "" {
		return
	}
	rStrList := strings.Split(rStr, " ")
	if len(rStrList) > 30 && rStrList[0] == pid {
		//comm 应用程序或命令的名字
		comm := rStrList[1]
		data.Name = comm
		//log.Println("comm 应用程序或命令的名字 = ", comm)

		//task_state  任务的状态，R:runnign, S:sleeping (TASK_INTERRUPTIBLE),
		//D:disk sleep (TASK_UNINTERRUPTIBLE), T: stopped, T:tracing stop,Z:zombie, X:dead
		taskState := rStrList[2]
		data.TaskState = taskState
		//log.Println("task_state  任务的状态 = ", taskState)

		//ppid 父进程ID
		ppid := rStrList[3]
		data.Ppid = ppid
		//log.Println("ppid 父进程ID = ", ppid)

		//pgid 线程组号
		pgid := rStrList[4]
		data.Pgid = pgid
		//log.Println("pgid 线程组号 = ", pgid)

		//sid  该任务所在的会话组ID
		sid := rStrList[5]
		data.Sid = sid
		//log.Println("sid  该任务所在的会话组ID = ", sid)

		//utime 该任务在用户态运行的时间，单位为jiffies
		utime := utils.Num2Int64(rStrList[13])
		//log.Println("utime 该任务在用户态运行的时间 = ", utime)

		//stime 该任务在核心态运行的时间，单位为jiffies
		stime := utils.Num2Int64(rStrList[14])
		//log.Println("stime 该任务在核心态运行的时间 = ", stime)

		//cutime 累计的该任务的所有的waited-for进程曾经在用户态运行的时间，单位为jiffies
		cutime := utils.Num2Int64(rStrList[15])
		//log.Println("cutime 累计的该任务的所有的waited-for进程曾经在用户态运行的时间 = ", cutime)

		//cstime 累计的该任务的所有的waited-for进程曾经在核心态运行的时间，单位为jiffies
		cstime := utils.Num2Int64(rStrList[16])
		//log.Println("cstime 累计的该任务的所有的waited-for进程曾经在核心态运行的时间 = ", cstime)

		//num_threads 该任务所在的线程组里线程的个数
		numThreads := rStrList[19]
		data.NumThreads = numThreads
		//log.Println("num_threads 该任务所在的线程组里线程的个数 = ", numThreads)

		//task_cpu(task) 运行在哪个CPU上
		//taskCpu := rStrList[38]
		//log.Println("task_cpu(task) 运行在哪个CPU上 = ", taskCpu)

		//进程的总Cpu时间processCpuTime = utime + stime + cutime + cstime，该值包括其所有线程的cpu时间。
		processCpuTime := utime + stime + cutime + cstime
		data.ProcessCpuTime = processCpuTime
		//log.Println("进程的总Cpu时间processCpuTime = ", processCpuTime)
	}
	return
}

// 某一进程Cpu使用率的计算
// 计算方法：
// 1．采样两个足够短的时间间隔的cpu快照与进程快照，
// a)  每一个cpu快照均为(user、nice、system、idle、iowait、irq、softirq、stealstolen、guest)的9元组;
// b)  每一个进程快照均为 (utime、stime、cutime、cstime)的4元组；
// 2．分别计算出两个时刻的总的cpu时间与进程的cpu时间，分别记作：totalCpuTime1、totalCpuTime2、processCpuTime1、processCpuTime2
// 3．计算该进程的cpu使用率pcpu = 100*( processCpuTime2 – processCpuTime1) / (totalCpuTime2 – totalCpuTime1)
// (按100%计算，如果是多核情况下还需乘以cpu的个数);
func ProcessProcStat(pid string) {
	var pcpu float64

	processCpuTime1 := GetProcessProcStat(pid)
	procStat1 := GetProcStat()
	//睡眠延时500ms
	time.Sleep(500 * time.Millisecond)
	processCpuTime2 := GetProcessProcStat(pid)
	procStat2 := GetProcStat()
	if len(procStat1) == 0 || len(procStat2) == 0 {
		return
	}
	pcpu = 100 * float64((processCpuTime2.ProcessCpuTime - processCpuTime1.ProcessCpuTime)) / float64((procStat2[0].Total - procStat1[0].Total))
	log.Println("pcpu = ", pcpu)
}

// 从 /proc/diskstats  -- 每块磁盘设备的磁盘I/O统计信息列表；（内核2.5.69以后的版本支持此功能）
func GetProcDiskstats() (datas []*structs.ProcDiskstatsData) {
	datas = make([]*structs.ProcDiskstatsData, 0)
	//（内核2.5.69以后的版本支持此功能）
	_, version := ProcVersion()
	versionList := strings.Split(version, ".")
	if len(versionList) < 2 {
		return
	}
	if utils.Num2Int64(versionList[0]) < 2 {
		if utils.Num2Int64(versionList[1]) < 5 {
			return
		}
		return
	}

	rStr := cmd.LinuxSendCommand("cat /proc/diskstats")
	if rStr == "" {
		return
	}

	//1:设备号  2:编号  3:设备   4:读完成次数  5:合并完成次数   6:读扇区次数   7:读操作花费毫秒数   8:写完成次数
	//9:合并写完成次数    10:写扇区次数    11:写操作花费的毫秒数    12:正在处理的输入/输出请求数    13:输入/输出操作花费的毫秒数
	//14:输入/输出操作花费的加权毫秒数。
	rStrList := strings.Split(rStr, "\n")
	for _, v := range rStrList {
		vList := strings.Split(v, " ")
		diskListData := []string{}
		for _, d := range vList {
			if d != "" {
				diskListData = append(diskListData, d)
			}
		}
		if len(diskListData) < 13 {
			continue
		}

		datas = append(datas, &structs.ProcDiskstatsData{
			DiskName:     diskListData[2],
			IOTime:       utils.Num2Int64(diskListData[12]),
			ReadRequest:  utils.Num2Int64(diskListData[3]),
			WriteRequest: utils.Num2Int64(diskListData[7]),
			MsecRead:     utils.Num2Int64(diskListData[6]),
			MsecWrite:    utils.Num2Int64(diskListData[10]),
		})
	}
	return
}

// 从 /proc/diskstats  -- 每块磁盘设备的磁盘I/O统计信息
// 采样两个足够短的时间间隔的磁盘快照，标记为t1、t2，计算t1时间的输入/输出操作花费的毫秒数used1，
// 计算t2时间的输入/输出操作花费的毫秒数used2。
// 于是磁盘IO操作百分比为：
// 100 * （used2 - used1）/ （t2 - t1）
// BUG:  不准确
func ProcDiskstats() {

	used1List := GetProcDiskstats()
	t1 := time.Now().UnixNano()

	//睡眠延时500ms
	time.Sleep(1000 * time.Millisecond)
	used2List := GetProcDiskstats()

	t2 := time.Now().UnixNano()

	totalDuration := float64((t2 - t1) / 1000 / 1000)

	for _, u1 := range used1List {
		for _, u2 := range used2List {
			if u1.DiskName == u2.DiskName {
				log.Println(*u1, *u2, (t2-t1)/1000/1000)
				DiskIO := (float64(u2.IOTime-u1.IOTime) / totalDuration * 100)
				log.Println("DiskIO = ", DiskIO)
				read_use_io := float64(u2.MsecRead - u1.MsecRead)
				write_use_io := float64(u2.MsecWrite - u1.MsecWrite)
				read_io := float64(u2.ReadRequest - u1.ReadRequest)
				write_io := float64(u2.WriteRequest - u1.WriteRequest)
				read_write_io := float64(u2.IOTime - u1.IOTime)
				readwrite_io := read_io + write_io
				io_awit := int(read_use_io + write_use_io/readwrite_io*10000)
				io_rs := (read_io / totalDuration) * 10000
				io_ws := (write_io / totalDuration) * 10000
				io_util := (read_write_io / (totalDuration * 1000)) * 10000
				log.Println("io_awit = ", io_awit)
				log.Println("io_rs = ", io_rs)
				log.Println("io_ws = ", io_ws)
				log.Println("io_util = ", io_util)
				log.Println("_________________________________________")
			}
		}
	}
}

// 从 /proc/version 当前系统运行的内核版本号
func ProcVersion() (string, string) {
	rStr := cmd.LinuxSendCommand("cat /proc/version")
	if rStr == "" {
		return rStr, ""
	}
	version := ""
	reg := regexp.MustCompile(`Linux version(.*?)-`)
	sList := reg.FindStringSubmatch(rStr)
	if len(sList) > 1 {
		version = sList[1]
	}
	return rStr, utils.DeletePreAndSufSpace(version)
}

//从/proc/net/dev中读取  采集网卡信息
func GetProcNetDev() (datas []*structs.ProcNetDevData) {
	datas = make([]*structs.ProcNetDevData, 0)
	rStr := cmd.LinuxSendCommand("cat /proc/net/dev")
	if rStr == "" {
		return
	}

	rStrList := strings.Split(rStr, "\n")
	for _, v := range rStrList {
		vList := strings.Split(v, " ")
		dataList := make([]string, 0)
		for _, i := range vList {
			if i != "" {
				dataList = append(dataList, i)
			}
		}
		if len(dataList) == 17 {
			name := dataList[0]
			recv := utils.Num2Int64(dataList[1])
			send := utils.Num2Int64(dataList[9])
			// log.Println(name, recv, send)
			datas = append(datas, &structs.ProcNetDevData{name, recv, send})
		}
	}
	return
}

// 从/proc/net/dev中读取信息并计算
//采样两个时间段的网卡信息 n1,n2 ,  时间t1,t2
//网络(kb/sec) = n2-n1/1024*(t2-t1)
func ProcNetDev() {
	n1 := GetProcNetDev()
	//休眠1s,误差忽略
	time.Sleep(1 * time.Second)
	n2 := GetProcNetDev()

	for _, v1 := range n1 {
		for _, v2 := range n2 {
			if v1.Name == v2.Name {
				receice_rate := (v2.Recv - v1.Recv) / 1024 * 1
				send_rate := (v2.Send - v1.Send) / 1024 * 1
				log.Println(v1.Name, "RX :", receice_rate, " | TX: ", send_rate, " |TOL: ", receice_rate+send_rate)
			}
		}
	}
}

// 从/proc/net/snmp 采集各层网络协议的收发包的情况
// tcp : CurrEstab(TCP连接数)
func ProcNetSnmp() {

	rStr := cmd.LinuxSendCommand("cat /proc/net/snmp")
	if rStr == "" {
		return
	}
	rStrList := strings.Split(rStr, "\n")

	for i := 0; i < len(rStrList)-1; i++ {
		if (i+1)%2 == 0 {
			continue
		}
		//log.Println(rStrList[i], rStrList[i+1])
		keyList := strings.Split(rStrList[i], " ")
		vlueList := strings.Split(rStrList[i+1], " ")
		//log.Println(keyList)
		//log.Println(vlueList)
		mapData := map[string]string{}
		mapData["name"] = vlueList[0]
		for i := 0; i < len(keyList); i++ {
			mapData[keyList[i]] = vlueList[i]
		}
		log.Println(mapData)
	}
}

// 从/proc/<pid>/cmdline  获取启动当前进程的完整命令，但僵尸进程目录中的此文件不包含任何信息；
func ProcPIDCmdline(pid string) {
	rStr := cmd.LinuxSendCommand(fmt.Sprintf("cat /proc/%s/cmdline", pid))
	if rStr == "" {
		return
	}
	log.Println(rStr)
}

// 从/proc/<pid>/environ 当前进程的环境变量列表，彼此间用空字符（NULL）隔开；变量用大写字母表示，其值用小写字母表示；
func ProcPIDEnviron(pid string) {
	rStr := cmd.LinuxSendCommand(fmt.Sprintf("cat /proc/%s/environ", pid))
	if rStr == "" {
		return
	}
	log.Println(rStr)
	reg := regexp.MustCompile(`(.*?)=[^A-Z]+`)
	sList := reg.FindAllString(rStr, -1)
	//log.Println(sList, len(sList))
	for k, v := range sList {
		log.Println(k, v)
	}
}

// 从 /proc/<pid>/limits —> 当前进程所使用的每一个受限资源的软限制、硬限制和管理单元；
//此文件仅可由实际启动当前进程的UID用户读取；（2.6.24以后的内核版本支持此功能）；
func ProcPIDLimits(pid string) {
	rStr := cmd.LinuxSendCommand(fmt.Sprintf("cat /proc/%s/limits", pid))
	if rStr == "" {
		return
	}
	log.Println(rStr)
}

// 从 /proc/<pid>/maps — 当前进程关联到的每个可执行文件和库文件在内存中的映射区域及其访问权限所组成的列表；
func ProcPIDMaps(pid string) {
	rStr := cmd.LinuxSendCommand(fmt.Sprintf("cat /proc/%s/maps", pid))
	if rStr == "" {
		return
	}
	log.Println(rStr)
}

// 从 /proc/<pid>/status —当前进程的状态信息，包含一系统格式化后的数据列，可读性较好
func ProcPIDStatus(pid string) {
	rStr := cmd.LinuxSendCommand(fmt.Sprintf("cat /proc/%s/status", pid))
	if rStr == "" {
		return
	}
	log.Println(rStr)
}

// 从 /proc/crypto  -- 系统上已安装的内核使用的密码算法及每个算法的详细信息列表；
func ProcCrypto() {
	rStr := cmd.LinuxSendCommand("cat /proc/crypto")
	if rStr == "" {
		return
	}
	log.Println(rStr)
}

// 从 /proc/modules  -- 当前装入内核的所有模块名称列表，可以由lsmod命令使用，也可以直接查看；如下所示，其中第一列表示模块名，第二列表示此模块占用内存空间大小，
//				第三列表示此模块有多少实例被装入，第四列表示此模块依赖于其它哪些模块，第五列表示此模块的装载状态（Live：已经装入；Loading：正在装入；Unloading：正在卸载），
//				第六列表示此模块在内核内存（kernel memory）中的偏移量；
func ProcModules() {
	rStr := cmd.LinuxSendCommand("cat /proc/modules")
	if rStr == "" {
		return
	}
	log.Println(rStr)
}

// 从 /proc/uptime   --  系统上次启动以来的运行时间，如下所示，其第一个数字表示系统运行时间，第二个数字表示系统空闲时间，单位是秒；
func ProcUptime() {
	rStr := cmd.LinuxSendCommand("cat /proc/uptime")
	if rStr == "" {
		return
	}
	log.Println(rStr)
	rStrList := strings.Split(rStr, " ")
	runTime := rStrList[0]
	ldleTime := rStrList[1]
	log.Println("runTime = ", runTime, " | ldleTime = ", ldleTime)

}

//  ==================== 未使用  =========================
/*
每个独立进程: /proc/1/*
cwd — 指向当前进程运行目录的一个符号链接；
exe — 指向启动当前进程的可执行文件（完整路径）的符号链接，通过/proc/N/exe可以启动当前进程的一个拷贝；
fd — 这是个目录，包含当前进程打开的每一个文件的文件描述符（file descriptor），这些文件描述符是指向实际文件的一个符号链接；
mem — 当前进程所占用的内存空间，由open、read和lseek等系统调用使用，不能被用户读取； （用户读不到）
root — 指向当前进程运行根目录的符号链接；在Unix和Linux系统上，通常采用chroot命令使每个进程运行于独立的根目录；
stat — 当前进程的状态信息，包含一系统格式化后的数据列，可读性差，通常由ps命令使用；
statm — 当前进程占用内存的状态信息，通常以“页面”（page）表示；
status — 与stat所提供信息类似，但可读性较好，如下所示，每行表示一个属性信息；其详细介绍请参见 proc的man手册页；
*/

/*
全局: /proc/*
/proc/buddyinfo  --  用于诊断内存碎片问题的相关信息文件；
/proc/cmdline  --  在启动时传递至内核的相关参数信息，这些信息通常由lilo或grub等启动管理工具进行传递；
/proc/cpuinfo  -- 处理器的相关信息的文件；
/proc/devices  -- 系统已经加载的所有块设备和字符设备的信息，包含主设备号和设备组（与主设备号对应的设备类型）名；
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
/proc/partitions  -- 块设备每个分区的主设备号（major）和次设备号（minor）等信息，同时包括每个分区所包含的块（block）数目（如下面输出中第三列所示）；
/proc/pci  -- 内核初始化时发现的所有PCI设备及其配置信息列表，其配置信息多为某PCI设备相关IRQ信息，可读性不高，可以用“/sbin/lspci –vb”命令获得较易理解的相关信息；
				在2.6内核以后，此文件已为/proc/bus/pci目录及其下的文件代替；
/proc/slabinfo  -- 在内核中频繁使用的对象（如inode、dentry等）都有自己的cache，即slab pool，而/proc/slabinfo文件列出了这些对象相关slap的信息；详情可以参见内核文档中slapinfo的手册页；
/proc/swaps
当前系统上的交换分区及其空间利用信息，如果有多个交换分区的话，则会每个交换分区的信息分别存储于/proc/swap目录中的单独文件中，而其优先级数字越低，被使用到的可能性越大；下面是作者系统中只有一个交换分区时的输出信息；
/proc/vmstat
当前系统虚拟内存的多种统计数据，信息量可能会比较大，这因系统而有所不同，可读性较好；下面为作者机器上输出信息的一个片段；（2.6以后的内核支持此文件）
/proc/zoneinfo
内存区域（zone）的详细信息列表，信息量较大，
*/

//获取网卡Mac
func GetNotCardMAC(){
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println("Error:" + err.Error())
		return
	}
	for _, inter := range interfaces {
		log.Println(inter.Name)
		log.Println(inter.Index)
		log.Println(inter.HardwareAddr)
	}

}