package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var (
	Visitors  = make(map[int]Visitor)
	BrodeCast = make(chan string)
)

func main() {
	fmt.Println("게임서버 0.1 시작...")

	port := "8080"
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("서버리슨 오류. 종료")
		os.Exit(1)
	}
	fmt.Println("서버 리슨시작 포트번호 : " + port)
	count := 0
	for {
		conn, errAc := listener.Accept()
		if errAc != nil {
			//어셉트한게 오류있으면 이건 짜르고 다시반복
			fmt.Println("Connection accepting failed.")
			conn.Close()
			time.Sleep(100 * time.Millisecond)
			continue
		}

		count++
		Visitors[count] = NewVisitor(count, conn)

		time.Sleep(0) //100 * time.Millisecond)
	}
}
