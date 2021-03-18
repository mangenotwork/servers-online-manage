## ManGe 服务器在线监控与管理解决方案
> 用于对服务器的监控, 操控, docker操作, 部署二进制可执行文件, 常用命令操作, 文件传输,远程部署常用中间件如redis, nginx等

## Task
1. 主机监控(控制报警),信息查看
2. 进程监控
3. Docker管理与操作
1. 部署与管理服务(编译好的2进制服务)
2. 文件传输
1. 方便常见中间件集群操作
2. 内网穿透


## TODO List
- Windows 磁盘IO
- Slve 信息详情
- Slve 资源信息采集
- Slve 查看环境变量
- Slve 查看空间
- docker 开启api linux&windows
- 通过docker api&sdk 管理docker
- 通过cmd 管理docker
- 部署docker images
- 部署可执行文件
- nginx 信息获取与监控
- 内网穿透服务创建
- 内网代理服务创建
- 代理服务创建

## Complete
1. 获取Slve网卡，IP等信息
2. 文件传输 
3. 配置文件
4. 断开从连
5. Windows CPU 信息
6. Sqlist 数据库初始化与创建表
7. Windows 系统基本信息
8. Windows 磁盘&内存信息
9. Windows 网络IO - 命令采集实现 -> <iphlpapi.h>与iphlpapi.li导包问题，暂时不用C实现
10. Linux 信息获取
11. Linux CPU, 内存, 磁盘 使用率获取
12. 


## 细节:

1. slve 采集系统信息, 发送心跳包带上采集数据
master 存储, 并检查监控点, 

2. slve 第一次连接后 需要上报基础信息
master 存储

3. web 使用认证

4. 查看进程, 对进程设置监控, slve 启用改进程的数据采集,  同心跳发送


5. docker 管理, 镜像,容器  pull 部署等


## 展示
- Slve 服务列表

![slve host list ](https://file.mange.run/mange-server_manage/20201208134139.png)



## 服务器监控的目标:
CPU使用量、内存消耗、I/O、网络、磁盘使用量、进程, 服务, 系统日志, 端口