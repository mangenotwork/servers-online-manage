package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

type MidServer struct {
	//客户端监听
	clientLis *net.TCPListener
	//后端服务连接
	transferLis *net.TCPListener
	//所有通道
	channels map[int]*Channel
	//当前通道ID
	curChannelId int
}

type Channel struct {
	//通道ID
	id int
	//客户端连接
	client net.Conn
	//后端服务连接
	transfer net.Conn
	//客户端接收消息
	clientRecvMsg chan []byte
	//后端服务发送消息
	transferSendMsg chan []byte
}

//创建一个服务器
func New() *MidServer {
	return &MidServer{
		channels:     make(map[int]*Channel),
		curChannelId: 0,
	}
}

//启动服务
func (m *MidServer) Start(clientPort, transferPort string) error {
	addr, err := net.ResolveTCPAddr("tcp", clientPort)
	if err != nil {
		return err
	}
	m.clientLis, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	addr, err = net.ResolveTCPAddr("tcp", transferPort)
	if err != nil {
		return err
	}
	m.transferLis, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	go m.AcceptLoop()
	return nil
}

//关闭服务
func (m *MidServer) Stop() {
	m.clientLis.Close()
	m.transferLis.Close()
	//循环关闭通道连接
	for _, v := range m.channels {
		v.client.Close()
		v.transfer.Close()
	}
}

//删除通道
func (m *MidServer) DelChannel(id int) {
	chs := m.channels
	delete(chs, id)
	m.channels = chs
}

//处理连接
func (m *MidServer) AcceptLoop() {
	transfer, err := m.transferLis.Accept()
	if err != nil {
		return
	}
	for {
		//获取连接
		client, err := m.clientLis.Accept()
		if err != nil {
			continue
		}

		//创建一个通道
		ch := &Channel{
			id:              m.curChannelId,
			client:          client,
			transfer:        transfer,
			clientRecvMsg:   make(chan []byte),
			transferSendMsg: make(chan []byte),
		}
		m.curChannelId++

		//把通道加入channels中
		chs := m.channels
		chs[ch.id] = ch
		m.channels = chs

		//启一个goroutine处理客户端消息
		go m.ClientMsgLoop(ch)
		//启一个goroutine处理后端服务消息
		go m.TransferMsgLoop(ch)
		go m.MsgLoop(ch)
	}
}

//处理客户端消息
func (m *MidServer) ClientMsgLoop(ch *Channel) {
	defer func() {
		fmt.Println("ClientMsgLoop exit")
	}()
	for {
		select {
		case data, isClose := <-ch.transferSendMsg:
			{
				//判断channel是否关闭，如果是则返回
				if !isClose {
					return
				}
				_, err := ch.client.Write(data)
				if err != nil {
					return
				}
			}
		}
	}
}

//处理后端服务消息
func (m *MidServer) TransferMsgLoop(ch *Channel) {
	defer func() {
		fmt.Println("TransferMsgLoop exit")
	}()
	for {
		select {
		case data, isClose := <-ch.clientRecvMsg:
			{
				//判断channel是否关闭，如果是则返回
				if !isClose {
					return
				}
				_, err := ch.transfer.Write(data)
				if err != nil {
					return
				}
			}
		}
	}
}

//客户端与后端服务消息处理
func (m *MidServer) MsgLoop(ch *Channel) {
	defer func() {
		//关闭channel，好让ClientMsgLoop与TransferMsgLoop退出
		close(ch.clientRecvMsg)
		close(ch.transferSendMsg)
		m.DelChannel(ch.id)
		fmt.Println("MsgLoop exit")
	}()
	buf := make([]byte, 1024)
	for {
		n, err := ch.client.Read(buf)
		if err != nil {
			return
		}
		ch.clientRecvMsg <- buf[:n]
		n, err = ch.transfer.Read(buf)
		if err != nil {
			return
		}
		ch.transferSendMsg <- buf[:n]
	}
}

func main() {
	//参数解析     -ext 127.0.0.1:18080 -in 127.0.0.1:18888
	var localPort, remotePort string
	flag.StringVar(&localPort, "ext", "", "客户端访问地址")
	flag.StringVar(&remotePort, "in", "", "服务访问地址")
	flag.Parse()
	if flag.NFlag() != 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	ms := New()
	//启动服务
	err := ms.Start(localPort, remotePort)
	log.Println(err)
	for {
	}
}
