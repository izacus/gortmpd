package dispatcher

type DispatchPacket struct {
	id		uint64		// EBML ID of the packet
	length	uint64		// Length flags for the packet (can be different from data)
	data    []byte		// Packet data
}

func (dp DispatchPacket) getByteRepresentation() []byte {
	return nil
}

func DispatchPackets() {

}
