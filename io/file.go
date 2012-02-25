package file

import (
    "fmt"
    "os"
)

type FileInput struct {
    fd  int     // File descriptor
    name string // Filename
}

func readFile(file *os.File, channel chan<- byte) {
    fmt.Println("Starting read...")

    for {
        buffer := make([]byte, 1024)    // 512B buffer
        var err error
        var read int
        err = nil
        for err == nil {
            read, err = file.Read(buffer) 
            for i:=0; i < read; i++ {
                channel <- buffer[i]
            }
        }
        
        fmt.Println("File done.")
        file.Seek(0, 0) // Infinite repeat
    }
}

func writeFile(file *os.File, channel <-chan byte) {
    var err error
    err = nil

    buffer := make([]byte, 5024)
    for err == nil {
        for i := 0; i < 5024; i++ {
            buffer[i] = <- channel
        }

        file.Write(buffer)
    }
}

func GetInputChannel(filename string) (channel <-chan byte) {
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

func GetOutputChannel(filename string) (channel chan<- byte) {
    fc := make(chan byte, 51200)
    file,err := os.Create(filename)

    if err == nil {
        go writeFile(file, fc)
    } else {
        fmt.Println("ERROR opening output file!")
        return nil
    }

    return fc
}