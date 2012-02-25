package processor

import (
    "fmt"
    "gortmpd/webm"
)

type ProcessorStates int
const (
    SearchingForHeader = iota
    SearchingForSegment
    SearchingForSegmentInfo
    ProcessingBlocks
)

func ProcessData(channel <-chan byte) {
    state := SearchingForHeader

    // Loop forever
    for {
        switch state {
            case SearchingForHeader:
                ebmlId, _ := webm.GetVintFromChannel(channel)

                if ebmlId != 0xA45DFA3 {
                    fmt.Println("ERROR - file does not begin with an EBML!")
                    return
                }

                ebmlLength, _ := webm.GetVintFromChannel(channel)
                webm.InputStream.SetEBMLHeader(getBytes(channel, ebmlLength))
                state = SearchingForSegment

            case SearchingForSegment:
                id, length, _ := webm.GetEBMLHeaderFromChannel(channel)
                // Segment header
                if id == 0x8538067 {
                    state = SearchingForSegmentInfo
                    fmt.Printf("[ProcessData] Found segment, length %d bytes.\n", length)
                } else {
                    skipBytes(channel, length)
                }
            case SearchingForSegmentInfo:
                id, length, _ := webm.GetEBMLHeaderFromChannel(channel)
                if id == 0x549A966 {
                    state = ProcessingBlocks
                    fmt.Println("[ProcessData] Found segment info!")
                    webm.InputStream.SetStreamInfo(getSegmentInfo(channel, length))
                } else {
                    skipBytes(channel, length)
                }

            case ProcessingBlocks:
                id,length, _ := webm.GetEBMLHeaderFromChannel(channel)
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
            webm.InputStream.SetTrackInfo(getBytes(channel, length))
            fmt.Printf("[Block] Found track info, size %dB.\n", length)
        default:
            skipBytes(channel, length)
    }

}
