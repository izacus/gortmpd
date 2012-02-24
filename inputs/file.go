package file

import (
    "fmt"
    "os"
    "time"
)

type FileInput struct {
    fd  int     // File descriptor
    name string // Filename
}

func readFile(file *os.File, channel chan<- byte) {
    fmt.Println("Starting read...")

    buffer := make([]byte, 1024)    // 512B buffer
    var err error
    var read int
    err = nil
    for err == nil {
        read, err = file.Read(buffer) 
        for i:=0; i < read; i++ {
            channel <- buffer[i]
        }

        time.Sleep(1000)
    }
    
    fmt.Println("File done.")
}

func GetChannel(filename string) (channel <-chan byte) {
    fc := make(chan byte, 51200) 
    fmt.Println("Attempting to open ", filename) 
    file,err := os.Open(filename)
    if err == nil {
        go readFile(file, fc)
    } else {
        fmt.Println("ERROR opening file!")
        return nil
    }

    return fc
}
