package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func handler(r net.Conn, localPort int) {
	buf := make([]byte, 1024)
	for {
		log.Println("okok")
		//先从远程读数据
		n, err := r.Read(buf)
		if err != nil {
			log.Println("先从远程读数据 err = ", err)
		}
		log.Println("先从远程读数据 = ", n)
		data := buf[:n]
		//建立与本地80服务的连接
		local, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", localPort))
		if err != nil {
			log.Println("建立与本地服务的连接 err = ", err)
		}
		log.Println("local = ", local)
		//向80服务写数据
		n, err = local.Write(data)
		if err != nil {
			log.Println("服务写数据 err = ", err)
		}
		log.Println("向80服务写数据 = ", n)
		//读取80服务返回的数据
		n, err = local.Read(buf)
		//关闭80服务，因为本地80服务是http服务，不是持久连接
		//一个请求结束，就会自动断开。所以在for循环里我们要不断Dial，然后关闭。
		local.Close()
		if err != nil {
			log.Println("读取80服务返回的数据 err = ", err)
			continue
		}
		data = buf[:n]
		//向远程写数据
		n, err = r.Write(data)
		if err != nil {
			log.Println("向远程写数据 err = ", err)
			continue
		}
	}
}

func main() {
	//参数解析
	var port int
	var host string
	flag.StringVar(&host, "host", "", "服务器地址")
	flag.IntVar(&port, "port", 80, "本地端口")
	flag.Parse()
	if flag.NFlag() != 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	//建立与服务器的连接
	//log.Println("strate")
	//"118.25.137.1"
	remote, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(remote)
	handler(remote, port)

	for {
	}
}
