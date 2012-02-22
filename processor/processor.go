package processor

import ( 
    "fmt" 
)
func ProcessData(channel <-chan byte) {
    // Presumption: we start on an EMBL, there are no sync bytes in MKV
    fmt.Println("Firing up data processor...")
    packet := getEBMLFromChannel(channel)
    fmt.Println(packet)
}
