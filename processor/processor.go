package processor

import ( 
    "fmt" 
)

type ProcessorStates int
const (
    SearchingForHeader = iota
    SearchingForSegment
    SearchingForSegmentInfo
    ProcessingBlocks
)

func skipBytes(channel <-chan byte, num uint64) {
    for i:=uint64(0); i < num; i++ {
        <-channel
    }
}

func ProcessData(channel <-chan byte) {
    state := SearchingForHeader

    // Loop forever
    for {
        switch state {
            case SearchingForHeader:
                ebmlId, _ := getVintFromChannel(channel)

                if ebmlId != 0xA45DFA3 {
                    fmt.Println("ERROR - file does not begin with an EBML!")
                    return
                }

                ebmlLength, _ := getVintFromChannel(channel)
                skipBytes(channel, ebmlLength)
                state = SearchingForSegment

            case SearchingForSegment:
                id, length, _ := getEBMLHeaderFromChannel(channel)
                // Segment header
                if id == 0x8538067 {
                    state = SearchingForSegmentInfo
                    fmt.Printf("[ProcessData] Found segment, length %d bytes.\n", length)
                } else {
                    skipBytes(channel, length)
                }
            case SearchingForSegmentInfo:
                id, length, _ := getEBMLHeaderFromChannel(channel)
                if id == 0x549A966 {
                    state = ProcessingBlocks
                    fmt.Println("[ProcessData] Found segment info!")
                    getSegmentInfo(channel, length)
                } else {
                    skipBytes(channel, length)
                }

            case ProcessingBlocks:
                id,length, _ := getEBMLHeaderFromChannel(channel)
                processBlock(channel, id, length)
        }
    }
}

func getSegmentInfo(channel <-chan byte, size uint64) SegmentHead {
    var head SegmentHead
    read := uint64(0)

    for read < size {
        id, length, read_pass := getEBMLHeaderFromChannel(channel)
        fmt.Printf("[SegmentInfo] ID %X len %d.\n", id, length)
        read += read_pass

        switch id {
            case 0xAD7B1:   // Timecode scale
               scale := getNumber(channel, length)
               fmt.Println("[SegmentInfo] Timecode scale: ", scale)
               head.TimecodeScale = scale
            case 0xD80:
               muxer := string(getNumber(channel, length))
               fmt.Println("[SegmentInfo] Muxer: ", muxer)
            default:
               skipBytes(channel, length)
        }
        
        read += length
    }
    
    return head
}

func processBlock(channel <-chan byte, id uint64, length uint64) {
    skipBytes(channel, length)
}
