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

func ProcessData(context webm.Context) {
    state := SearchingForHeader

    // Loop forever
    for {
        switch state {
            case SearchingForHeader:
                ebmlId, _ := webm.GetVintFromChannel(context.InputChannel)

                if ebmlId != 0xA45DFA3 {
                    fmt.Println("ERROR - file does not begin with an EBML!")
                    return
                }
                ebmlLength, _ := webm.GetVintFromChannel(context.InputChannel)
                context.InputStream.SetEBMLHeader(getBytes(context.InputChannel, ebmlLength))

                fmt.Println("Input stream ", context.InputStream)
                dispatchPacket(context.DispatchChannel, ebmlId, ebmlLength, context.InputStream.GetEBMLHeader())

                state = SearchingForSegment

            case SearchingForSegment:
                id, length, _ := webm.GetEBMLHeaderFromChannel(context.InputChannel)
                // Segment header
                if id == 0x8538067 {
                    state = SearchingForSegmentInfo
                    fmt.Printf("[ProcessData] Found segment, length %d bytes.\n", length)

                    dispatchPacket(context.DispatchChannel, id, length, nil)
                } else {
                    skipBytes(context.InputChannel, length)
                }
            case SearchingForSegmentInfo:
                id, length, _ := webm.GetEBMLHeaderFromChannel(context.InputChannel)
                if id == 0x549A966 {
                    state = ProcessingBlocks
                    fmt.Println("[ProcessData] Found segment info!")
                    context.InputStream.SetStreamInfo(getSegmentInfo(context.InputChannel, length))

                    dispatchPacket(context.DispatchChannel, id, length, context.InputStream.GetStreamInfo())

                } else {
                    skipBytes(context.InputChannel, length)
                }

            case ProcessingBlocks:
                id,length, _ := webm.GetEBMLHeaderFromChannel(context.InputChannel)
                processBlock(context, id, length)
        }
    }
}

func dispatchPacket(channel chan<- webm.DispatchPacket, id uint64, length uint64, data []byte) {
    var packet webm.DispatchPacket
    packet.Id = id
    packet.Length = length
    packet.Data = data
    channel <- packet
}

func getSegmentInfo(channel <-chan byte, size uint64) []byte {
    data := getBytes(channel, size)
    fmt.Printf("[SegmentInfo] Stored %d bytes of head.\n", size)
    return data
}

func processBlock(context webm.Context, id uint64, length uint64) {
    fmt.Printf("[Block] Block ID %X size %d.\n", id, length)

    switch id {
        case 0x654AE6B:             // Track info
            context.InputStream.SetTrackInfo(getBytes(context.InputChannel, length))
            dispatchPacket(context.DispatchChannel, id, length, context.InputStream.GetTrackInfo())
            fmt.Printf("[Block] Found track info, size %dB.\n", length)
        default:
            data := getBytes(context.InputChannel, length)
            dispatchPacket(context.DispatchChannel, id, length, data)
    }

}
