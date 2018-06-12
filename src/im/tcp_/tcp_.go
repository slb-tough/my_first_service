package tcp_

import (
	"net"
	"fmt"
)

func Init_tcp() {
	net_listen, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("listen err!")
	}
	fmt.Println("Waiting for Client ...")
	for {
		conn, err := net_listen.Accept()
		if err != nil {
			fmt.Println("accept err!")
		}
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("conn read err!")
		}
		data := buffer[:n]
		fmt.Println("accept data:", data)
		fmt.Println(conn.RemoteAddr().String(), "connect success!")
	}
}
