package processor

import (
    "fmt"
)

func getEBMLHeaderFromChannel(channel <-chan byte) (id uint64, length uint64) {
    id = getVintFromChannel(channel)
    length = getVintFromChannel(channel)
    return
}

// Retrieves vint from channel and returns it's value
func getVintFromChannel(channel <-chan byte)(val uint64) {
    head := <-channel

    switch {
        case (head & 0x80) > 0:
            val = uint64(head & 0x7F)
        case (head & 0x40) > 0:
            val = (uint64(head & 0x3F) << 8) | uint64(<-channel)
        case (head & 0x20) > 0:
            val = (uint64(head & 0x1F) << 16) | (uint64(<-channel)) << 8 | uint64(<-channel)
        case (head & 0x10) > 0:
            val = (uint64(head & 0x0F) << 24) | (uint64(<-channel)) << 16 | uint64(<-channel) << 8 | uint64(<-channel)
        case (head & 0x08) > 0:
            val = (uint64(head & 0x07) << 32) | (uint64(<-channel)) << 24 | uint64(<-channel) << 16 | uint64(<-channel) << 8 | uint64(<-channel)
        case (head & 0x04) > 0:
             val = (uint64(head & 0x03) << 40) | (uint64(<-channel)) << 32 | (uint64(<-channel)) << 24 | uint64(<-channel) << 16 | uint64(<-channel) << 8 | uint64(<-channel)
        case (head & 0x02) > 0:
             val = (uint64(head & 0x01) << 48) | (uint64(<-channel)) << 40 | (uint64(<-channel)) << 32 | (uint64(<-channel)) << 24 | uint64(<-channel) << 16 | uint64(<-channel) << 8 | uint64(<-channel)
        case (head & 0x01) > 0:
             val = uint64(<-channel) << 48 | (uint64(<-channel)) << 40 | (uint64(<-channel)) << 32 | (uint64(<-channel)) << 24 | uint64(<-channel) << 16 | uint64(<-channel) << 8 | uint64(<-channel)
        default:
            fmt.Println("ERROR - vint too big!")
    }

    return val
}
