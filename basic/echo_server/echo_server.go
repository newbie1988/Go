package main
import (
	"fmt"
	"flag"
	"net"
)

func Process(con net.Conn) {
	defer con.Close();
	const BUFF_LEN = 512;
	buffer := make([]byte, BUFF_LEN);
	for {
		len, err := con.Read(buffer);
		if nil != err {
			fmt.Println("read error:", err.Error());
			break;
		}
		fmt.Println(string(buffer[:len]));
	}
}

func main() {
	flag.Parse();
	if 2 != flag.NArg() {
		panic("Usage: <ip> <port>");
	}
	
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1));
	//serverIp, _ := net.ResolveTCPAddr("tcp", hostAndPort);
	listener, _ := net.Listen("tcp", hostAndPort);
	fmt.Println("Server Start, Listern Port:", flag.Arg(1));
	for {
		con, _ := listener.Accept();
		fmt.Println("A New Connection From: ", con.RemoteAddr().String());
		go Process(con);		
	}
}
