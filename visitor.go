package main

import (
	"fmt"
	"net"
	"time"
)

type Visitor struct {
	Id   int
	Conn net.Conn
}

func NewVisitor(id int, conn net.Conn) *Visitor {
	v := &Visitor{
		Id:   id,
		Conn: conn,
	}
	v.GoHandleREQ()
	return v
}

func (v *Visitor) BrodeCast() {
	BrodeCast <- "#BrodeCast!;"
}

func (v *Visitor) GoHandleREQ() {

	fmt.Println(string(v.Id) + "번째 새로운 접속자!")

	for {
		buf := make([]byte, 256)
		len, err := v.Conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}
		fmt.Println(string(v.Id) + "가 보냄 : " + string(buf[:len])) // 데이터 출력

		BrodeCast <- string(buf[:len])

		time.Sleep(0)
	}
}
