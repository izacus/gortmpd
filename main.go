package main

import ( "fmt"
         "time"
         "gortmpd/processor"
         "gortmpd/dispatcher"
         "gortmpd/inputs" )

func main() {
    fmt.Println("goRTMPd starting...")
    channel := file.GetChannel("IrJAwCBbnuc-43.webm")
//    dispatch_channel := make(chan byte, 10240)

    go processor.ProcessData(channel)   
    go dispatcher.DispatchPackets()

    for {
        time.Sleep(10000000000)
    }
}
