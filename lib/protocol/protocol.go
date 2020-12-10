package protocol

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"net"
	"strings"

	"github.com/mangenotwork/servers-online-manage/lib/global"
	"github.com/mangenotwork/servers-online-manage/lib/structs"
)

//通过协议 编码数据
func EnPackSendData(sendBytes []byte) []byte {
	packetLength := len(sendBytes) + 8
	result := make([]byte, packetLength)
	result[0] = 0xFF
	result[1] = 0xFF
	result[2] = byte(uint16(len(sendBytes)) >> 8)
	result[3] = byte(uint16(len(sendBytes)) & 0xFF)
	copy(result[4:], sendBytes)
	sendCrc := crc32.ChecksumIEEE(sendBytes)
	result[packetLength-4] = byte(sendCrc >> 24)
	result[packetLength-3] = byte(sendCrc >> 16 & 0xFF)
	result[packetLength-2] = 0xFF
	result[packetLength-1] = 0xFE
	return result
}

//拿到数据包 通过协议解码数据
//根据数据包来做解析
//数据包的格式为|0xFF|0xFF|len(高)|len(低)|Data|CRC高16位|0xFF|0xFE
//其中len为data的长度，实际长度为len(高)*256+len(低)
//CRC为32位CRC，取了最高16位共2Bytes
//0xFF|0xFF和0xFF|0xFE类似于前导码
func DePackSendDataMater(conn *structs.Cli, f func(*structs.Cli, *structs.Packet)) (err error) {
	err = nil
	// 状态机状态
	state := 0x00
	// 数据包长度
	length := uint16(0)
	// crc校验和
	crc16 := uint16(0)
	var recvBuffer []byte
	// 游标
	cursor := uint16(0)
	bufferReader := bufio.NewReader(conn.Conn)
	//状态机处理数据
	for {
		recvByte, err := bufferReader.ReadByte()

		if err != nil {
			log.Println("err or close = ", err)
			//这里因为做了心跳，所以就没有加deadline时间，如果客户端断开连接
			//这里ReadByte方法返回一个io.EOF的错误，具体可考虑文档
			if err == io.EOF {
				log.Println(" %s is close!\n", conn.Conn.RemoteAddr().String())

			}
			key := strings.Split(conn.Conn.RemoteAddr().String(), ":")[0]
			global.DelSlve(key)
			//生成消息
			//在这里直接退出goroutine，关闭由defer操作完成
			err = errors.New(fmt.Sprintf("IP: %s 的Slve断开了连接", key))
			return err
		}

		//log.Println("状态机状态 = ", state)

		//进入状态机，根据不同的状态来处理
		switch state {
		case 0x00:
			if recvByte == 0xFF {
				state = 0x01
				//初始化状态机
				recvBuffer = nil
				length = 0
				crc16 = 0
			} else {
				state = 0x00
			}
			break
		case 0x01:
			if recvByte == 0xFF {
				state = 0x02
			} else {
				state = 0x00
			}
			break
		case 0x02:
			length += uint16(recvByte) * 256
			state = 0x03
			break
		case 0x03:
			length += uint16(recvByte)
			// 一次申请缓存，初始化游标，准备读数据
			recvBuffer = make([]byte, length)
			cursor = 0
			state = 0x04
			break
		case 0x04:
			//不断地在这个状态下读数据，直到满足长度为止
			recvBuffer[cursor] = recvByte
			cursor++
			if cursor == length {
				state = 0x05
			}
			break
		case 0x05:
			crc16 += uint16(recvByte) * 256
			state = 0x06
			break
		case 0x06:
			crc16 += uint16(recvByte)
			state = 0x07
			break
		case 0x07:
			if recvByte == 0xFF {
				state = 0x08
			} else {
				state = 0x00
			}
		case 0x08:
			if recvByte == 0xFE {
				//执行数据包校验
				if (crc32.ChecksumIEEE(recvBuffer)>>16)&0xFFFF == uint32(crc16) {
					var packet structs.Packet
					//把拿到的数据反序列化出来
					json.Unmarshal(recvBuffer, &packet)
					//新开协程处理数据
					go f(conn, &packet)
				} else {
					fmt.Println("丢弃数据!")
				}
			}
			//状态机归位,接收下一个包
			state = 0x00
		}
	}
}

func DePackSendData(conn net.Conn, f func(net.Conn, *structs.Packet)) {
	// 状态机状态
	state := 0x00
	// 数据包长度
	length := uint16(0)
	// crc校验和
	crc16 := uint16(0)
	var recvBuffer []byte
	// 游标
	cursor := uint16(0)
	bufferReader := bufio.NewReader(conn)
	//状态机处理数据
	for {
		recvByte, err := bufferReader.ReadByte()
		if err != nil {
			//这里因为做了心跳，所以就没有加deadline时间，如果客户端断开连接
			//这里ReadByte方法返回一个io.EOF的错误，具体可考虑文档
			if err == io.EOF {
				fmt.Printf(" %s is close!\n", conn.RemoteAddr().String())
			}
			//在这里直接退出goroutine，关闭由defer操作完成
			global.RConn <- false
			return
		}

		//log.Println("状态机状态 = ", recvByte)

		//进入状态机，根据不同的状态来处理
		switch state {
		case 0x00:
			if recvByte == 0xFF {
				state = 0x01
				//初始化状态机
				recvBuffer = nil
				length = 0
				crc16 = 0
			} else {
				state = 0x00
			}
			break
		case 0x01:
			if recvByte == 0xFF {
				state = 0x02
			} else {
				state = 0x00
			}
			break
		case 0x02:
			length += uint16(recvByte) * 256
			state = 0x03
			break
		case 0x03:
			length += uint16(recvByte)
			// 一次申请缓存，初始化游标，准备读数据
			recvBuffer = make([]byte, length)
			cursor = 0
			state = 0x04
			break
		case 0x04:
			//不断地在这个状态下读数据，直到满足长度为止
			recvBuffer[cursor] = recvByte
			cursor++
			if cursor == length {
				state = 0x05
			}
			break
		case 0x05:
			crc16 += uint16(recvByte) * 256
			state = 0x06
			break
		case 0x06:
			crc16 += uint16(recvByte)
			state = 0x07
			break
		case 0x07:
			if recvByte == 0xFF {
				state = 0x08
			} else {
				state = 0x00
			}
		case 0x08:
			if recvByte == 0xFE {
				//执行数据包校验
				if (crc32.ChecksumIEEE(recvBuffer)>>16)&0xFFFF == uint32(crc16) {
					var packet structs.Packet
					//把拿到的数据反序列化出来
					json.Unmarshal(recvBuffer, &packet)
					//新开协程处理数据
					go f(conn, &packet)
				} else {
					fmt.Println("丢弃数据!")
				}
			}
			//状态机归位,接收下一个包
			state = 0x00
		}
	}
}
