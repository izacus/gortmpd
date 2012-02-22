package main

import ( "fmt"
         "time"
         "gortmpd/processor"
         "gortmpd/inputs" )

func main() {
    fmt.Println("goRTMPd starting...")

    channel := file.GetChannel("victoria.webm")
    go processor.ProcessData(channel)

    for {
        time.Sleep(100)
    }
}
