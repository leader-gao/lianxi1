package src

import (
	"fmt"
	"log"
	"github.com/google/googleopacket"
	"github.com/google/gopacket/pcap"
)

var (
	handle *pcap.Handle
	err    error
)

func main() {
	// Open pcap file
	handle, err = pcap.OpenOffline("test.pcap")
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Process packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
