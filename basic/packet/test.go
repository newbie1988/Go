package main
import (
	"fmt"
	"./pkt"
	)

func main() {
	fmt.Println("Packet Name: ", pkt.PacketName);
	fmt.Println("Packet Version: ", pkt.Version);
	fmt.Println("Packet Description: ");
	fmt.Println(pkt.Whoami());
}
