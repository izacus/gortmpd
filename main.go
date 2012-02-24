package main

import ( "fmt"
         "time"
         "gortmpd/processor"
         "gortmpd/inputs" )

func main() {
    fmt.Println("goRTMPd starting...")

    channel := file.GetChannel("IrJAwCBbnuc-43.webm")
    go processor.ProcessData(channel)

    for {
        time.Sleep(100)
    }
}
