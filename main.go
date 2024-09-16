package main

import (
	"fmt"
	"net"
)

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("something went wrong")
		return
	}
	fmt.Println("sir lhad l port (8080)")

	read := make([]byte, 4096)
	my_map := make(map[net.Conn]string)
	for {
		connections, err := listner.Accept()
		if err != nil {
			fmt.Fprintf(connections, "weraweng yalaa")
			continue
		}
		fmt.Fprintf(connections, "enter your name please")
		n, err := connections.Read(read)
		if err != nil {
			fmt.Fprintf(connections, "wooow hold up a bitt you can't connectto the server")
			continue
		}
		name := read[:n]
		my_map[connections] = string(name)
		fmt.Fprintf(connections, "%v", string(name))
	}
}

func HandelConection(connections net.Conn, my_map map[net.Conn]string, name string) {
}
