package processor

import (
    "fmt"
)

type EBMLType int

const (
    Unknown = iota
)

type EBMLPacket struct {
    packet_type    EBMLType
    data           []byte
}

func getEBMLFromChannel(channel <-chan byte) EBMLPacket {
    num := getVintFromChannel(channel) 
    fmt.Println("Got", num, "from channel.")

    var packet EBMLPacket
    return packet
}

// Retrieves vint from channel and returns it's value
func getVintFromChannel(channel <-chan byte)(val int) {
    head := <-channel

    switch {
        case (head & 0x80) > 0:
            val = int(head & 0x7F)
        case (head & 0x40) > 0:
            val = (int(head & 0x3F) << 8) | int(<-channel)
        case (head & 0x20) > 0:
            val = (int(head & 0x1F) << 16) | (int(<-channel)) << 8 | int(<-channel)
        case (head & 0x10) > 0:
            val = (int(head & 0x0F) << 24) | (int(<-channel)) << 16 | int(<-channel) << 8 | int(<-channel)
    }

    return val
}
