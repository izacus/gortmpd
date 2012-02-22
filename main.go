package main

import ( "fmt"
         "time"
         "gortmpd/inputs" )

func receiveData(channel chan byte) {
    buffer := make([]byte, 512)

    for {
        for i := 0; i < 512; i++ {
            buffer[i] = <-channel
        }

        fmt.Println("Read 512 B!")
    }
}

func main() {
    fmt.Println("goRTMPd starting...")

    channel := file.GetChannel("victoria.webm")
    go receiveData(channel)

    for {
        time.Sleep(100)
    }
}
