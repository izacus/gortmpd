package webm

import (
    "fmt"
)

func GetEBMLHeaderFromChannel(channel <-chan byte) (id uint64, length uint64, read uint64) {
    id, readId := GetVintFromChannel(channel)
    length, readLen := GetVintFromChannel(channel)
    read = readId + readLen
    return
}

// Retrieves vint from channel and returns it's value and bytes read
func GetVintFromChannel(channel <-chan byte)(val uint64, read uint64) {
    head := <-channel

    switch {
        case (head & 0x80) > 0:
            val = uint64(head & 0x7F)
            read = 1
        case (head & 0x40) > 0:
            val = (uint64(head & 0x3F) << 8) | uint64(<-channel)
            read = 2
        case (head & 0x20) > 0:
            val = (uint64(head & 0x1F) << 16) | (uint64(<-channel)) << 8 | uint64(<-channel)
            read = 3
        case (head & 0x10) > 0:
            val = (uint64(head & 0x0F) << 24) | (uint64(<-channel)) << 16 | uint64(<-channel) << 8 | uint64(<-channel)
            read = 4
        case (head & 0x08) > 0:
            val = (uint64(head & 0x07) << 32) | (uint64(<-channel)) << 24 | uint64(<-channel) << 16 | uint64(<-channel) << 8 | uint64(<-channel)
            read = 5
        case (head & 0x04) > 0:
             val = (uint64(head & 0x03) << 40) | (uint64(<-channel)) << 32 | (uint64(<-channel)) << 24 | uint64(<-channel) << 16 | uint64(<-channel) << 8 | uint64(<-channel)
             read = 6
        case (head & 0x02) > 0:
             val = (uint64(head & 0x01) << 48) | (uint64(<-channel)) << 40 | (uint64(<-channel)) << 32 | (uint64(<-channel)) << 24 | uint64(<-channel) << 16 | uint64(<-channel) << 8 | uint64(<-channel)
             read = 7
        case (head & 0x01) > 0:
             val = uint64(<-channel) << 48 | (uint64(<-channel)) << 40 | (uint64(<-channel)) << 32 | (uint64(<-channel)) << 24 | uint64(<-channel) << 16 | uint64(<-channel) << 8 | uint64(<-channel)
             read = 8
        default:
            fmt.Println("ERROR - vint too big!")
    }

    return val, read
}

func getNumber(channel <-chan byte, length uint64) uint64 {
    num := uint64(0)

    for i := uint64(0); i < length; i++ {
        num = (num << 8) | uint64(<-channel)
    }
    
    return num
}
