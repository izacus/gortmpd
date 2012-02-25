package webm

import (	
	"sync"
	"fmt"
)

// This file holds shared state of the transcoders
type Context struct {
	InputStream 		InputStreamContext
	InputChannel		<-chan byte
	DispatchChannel 	chan DispatchPacket
}

// Context of an input stream
type InputStreamContext struct {
    ebmlHeader      []byte      // EBML header
    streamInfo      []byte      // Stored incoming stream info
    trackInfo       []byte      // Stored incoming track info

    mu				sync.RWMutex
}

func (isc InputStreamContext) GetEBMLHeader() (header []byte) {
	isc.mu.RLock()
	defer isc.mu.RUnlock()
	return isc.ebmlHeader
}

func (isc InputStreamContext) GetStreamInfo() (info []byte) {
	isc.mu.RLock()
	defer isc.mu.RUnlock()
	return isc.streamInfo
}

func (isc InputStreamContext) GetTrackInfo() (trackInfo []byte) {
	isc.mu.RLock()
	defer isc.mu.RUnlock()
	return isc.trackInfo
}


func (isc InputStreamContext) SetEBMLHeader(header []byte) {
	isc.mu.Lock()
	isc.ebmlHeader = header
	isc.mu.Unlock()
}

func (isc InputStreamContext) SetStreamInfo(info []byte) {
	isc.mu.Lock()
	isc.streamInfo = info
	isc.mu.Unlock()
}

func (isc InputStreamContext) SetTrackInfo(trackInfo []byte) {
	isc.mu.Lock()
	isc.trackInfo = trackInfo
	isc.mu.Unlock()
}

type DispatchPacket struct {
	Id		uint64		// EBML ID of the packet
	Length	uint64		// Length flags for the packet (can be different from data)
	Data    []byte		// Packet data
}

func (dp DispatchPacket) GetByteRepresentation() []byte {

	fmt.Printf("[Packet] ID: %X Len: %d Data: %X", dp.Id, dp.Length, dp.Data)
	panic(nil)

	id_bytes := BuildVintFromNumber(dp.Id)
	length_bytes := BuildVintFromNumber(dp.Length)
	var bytes = make([]byte, len(id_bytes) + len(length_bytes) + len(dp.Data))
	copy(bytes, id_bytes)
	copy(bytes[len(id_bytes):], length_bytes)
	copy(bytes[len(id_bytes)+len(length_bytes):], dp.Data)
	return bytes
}