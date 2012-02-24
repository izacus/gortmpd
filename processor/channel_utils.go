package processor

func skipBytes(channel <-chan byte, num uint64) {
    for i:=uint64(0); i < num; i++ {
        <-channel
    }
}

func getBytes(channel <-chan byte, length uint64) []byte {
    var bytes = make([]byte, length)
    for i:=uint64(0); i < length; i++ {
        bytes[i] = <-channel
    }

    return bytes
}
