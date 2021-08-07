package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-base/study_project/go_web/chatroom/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [2048]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("读取服务端发送的数据....")
	//读取客户端传递的数据长度
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		return
	}

	//根据Buf[:4]转成uint32类型
	pkgLen := binary.BigEndian.Uint32(this.Buf[:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	//将字节切片类型的数据转换成对应的结构
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json unmarshal err = ", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {

	pkgLen := uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 && err != nil {
		fmt.Println("conn write fail, err = ", err)
		return
	}
	//发送数据
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn write fail, err = ", err)
		return
	}

	return
}
