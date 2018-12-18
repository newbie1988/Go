package pkt
import "bytes" 

const PacketName = "Packet1";
const Version = "1.2"

func Whoami() string {
	var buffer bytes.Buffer;
	buffer.WriteString("Packet Name: " + PacketName);
	buffer.WriteString("\n");
	buffer.WriteString("Packet Version: " + Version);
	return buffer.String();	
}

