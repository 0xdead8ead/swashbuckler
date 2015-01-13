// Swashbuckler (connection hijacker)

package main


import (
	"code.google.com/p/gopacket/pcap"
	"code.google.com/p/gopacket"
	"fmt"
)

func handlePacket(packet gopacket.Packet){
	fmt.Println("Packet Recieved\n----------------\n")
	fmt.Println(packet)
}

func main() {
	

	handle, err := pcap.OpenLive("en0", 1600, true, pcap.BlockForever)

	if err != nil {
	  panic(err)
	}
	err = handle.SetBPFFilter("icmp")
	if err != nil {  // optional
	  panic(err)
	} 


	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	
	for packet := range packetSource.Packets() {
		
		handlePacket(packet)  // Do something with a packet here.
	}


}