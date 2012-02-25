package main

import ( "fmt"
         "time"
         "gortmpd/processor"
         "gortmpd/dispatcher"
         "gortmpd/io"
         "gortmpd/webm"
          )

func main() {
    fmt.Println("goRTMPd starting...")
    channel := file.GetInputChannel("IrJAwCBbnuc-43.webm")
    dispatch_channel := make(chan webm.DispatchPacket, 10240)

    var context webm.Context
    context.InputChannel = channel
    context.DispatchChannel = dispatch_channel

    go processor.ProcessData(context)   
    go dispatcher.DispatchPackets(context)

    for {
        time.Sleep(10000000000)
    }
}
