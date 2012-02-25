package dispatcher

import (
	"fmt"
	"gortmpd/webm"
)

type DispatchPacket struct {
	Id		uint64		// EBML ID of the packet
	Length	uint64		// Length flags for the packet (can be different from data)
	Data    []byte		// Packet data
}

func (dp DispatchPacket) GetByteRepresentation() []byte {
	id_bytes := webm.BuildVintFromNumber(dp.Id)
	length_bytes := webm.BuildVintFromNumber(dp.Length)
	var bytes = make([]byte, len(id_bytes) + len(length_bytes) + len(dp.Data))
	copy(bytes, id_bytes)
	copy(bytes[len(id_bytes):], length_bytes)
	copy(bytes[len(id_bytes)+len(length_bytes):], dp.Data)
	return bytes
}

func DispatchPackets(incoming_channel <-chan DispatchPacket) {

	for {
		packet := <- incoming_channel
		bytes := packet.GetByteRepresentation()
		fmt.Printf("Dispatching packet len %d\n", len(bytes))
	}
}
