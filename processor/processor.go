package processor

import (
    "fmt"
    "gortmpd/ebml"
)

type ProcessorStates int
const (
    SearchingForHeader = iota
    SearchingForSegment
    SearchingForSegmentInfo
    ProcessingBlocks
)

type InputStreamContext struct {
    ebmlHeader      []byte      // EBML header
    streamInfo      []byte      // Stored incoming stream info
    trackInfo       []byte      // Stored incoming track info
}

var inputStream InputStreamContext

func ProcessData(channel <-chan byte) {
    state := SearchingForHeader

    // Loop forever
    for {
        switch state {
            case SearchingForHeader:
                ebmlId, _ := ebml.GetVintFromChannel(channel)

                if ebmlId != 0xA45DFA3 {
                    fmt.Println("ERROR - file does not begin with an EBML!")
                    return
                }

                ebmlLength, _ := ebml.GetVintFromChannel(channel)
                inputStream.ebmlHeader = getBytes(channel, ebmlLength)
                state = SearchingForSegment

            case SearchingForSegment:
                id, length, _ := ebml.GetEBMLHeaderFromChannel(channel)
                // Segment header
                if id == 0x8538067 {
                    state = SearchingForSegmentInfo
                    fmt.Printf("[ProcessData] Found segment, length %d bytes.\n", length)
                } else {
                    skipBytes(channel, length)
                }
            case SearchingForSegmentInfo:
                id, length, _ := ebml.GetEBMLHeaderFromChannel(channel)
                if id == 0x549A966 {
                    state = ProcessingBlocks
                    fmt.Println("[ProcessData] Found segment info!")
                    inputStream.streamInfo = getSegmentInfo(channel, length)
                } else {
                    skipBytes(channel, length)
                }

            case ProcessingBlocks:
                id,length, _ := ebml.GetEBMLHeaderFromChannel(channel)
                processBlock(channel, id, length)
        }
    }
}

func getSegmentInfo(channel <-chan byte, size uint64) []byte {
    data := getBytes(channel, size)
    fmt.Printf("[SegmentInfo] Stored %d bytes of head.\n", size)
    return data
}

func processBlock(channel <-chan byte, id uint64, length uint64) {
    fmt.Printf("[Block] Block ID %X size %d.\n", id, length)

    switch id {
        case 0x654AE6B:             // Track info
            inputStream.trackInfo = getBytes(channel, length)
            fmt.Printf("[Block] Found track info, size %dB.\n", length)
        default:
            skipBytes(channel, length)
    }

}
