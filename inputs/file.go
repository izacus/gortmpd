package file

import (
    "fmt"
    "os"
)

type FileInput struct {
    fd  int     // File descriptor
    name string // Filename
}

func readFile(file *os.File, channel chan byte) {
    fmt.Println("Starting read...")

    buffer := make([]byte, 512)    // 512B buffer
    var err error
    var read int
    err = nil
    for err == nil {
        read, err = file.Read(buffer) 
        for i:=0; i < read; i++ {
            channel <- buffer[i]
        }
    }
}

func GetChannel(filename string) (channel chan byte) {
    fc := make(chan byte) 
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
