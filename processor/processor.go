package processor

import ( 
    "fmt" 
)

type ProcessorStates int
const (
    SearchingForHeader = iota
    SearchingForSegment
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
                ebmlId := getVintFromChannel(channel)

                if ebmlId != 0xA45DFA3 {
                    fmt.Println("ERROR - file does not begin with an EBML!")
                    return
                }
                
                ebmlLength := getVintFromChannel(channel) 
                skipBytes(channel, ebmlLength)
                state = SearchingForSegment

            case SearchingForSegment:
                id, length := getEBMLHeaderFromChannel(channel)  
                // Segment header
                if id == 0x8538067 {
                    state = ProcessingBlocks
                    fmt.Printf("[ProcessData] Found segment, length %d bytes.\n", length)
                } else {
                    skipBytes(channel, length)
                }
            case ProcessingBlocks:
                id,length := getEBMLHeaderFromChannel(channel)
                fmt.Printf("Found %X len %d.\n", id, length)
                skipBytes(channel, length)
        }
    }
}
