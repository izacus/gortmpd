package main

import ( "fmt"
         "time"
         "gortmpd/processor"
         "gortmpd/dispatcher"
         "gortmpd/inputs" )

func main() {
    fmt.Println("goRTMPd starting...")
    channel := file.GetChannel("IrJAwCBbnuc-43.webm")
    dispatch_channel := make(chan dispatcher.DispatchPacket, 10240)

    go processor.ProcessData(channel, dispatch_channel)   
    go dispatcher.DispatchPackets(dispatch_channel)

    for {
        time.Sleep(10000000000)
    }
}
