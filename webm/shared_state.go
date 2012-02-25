package webm

import (	
	"sync"
)

// This file holds shared state of the transcoders

// Context of an input stream
type InputStreamContext struct {
    ebmlHeader      []byte      // EBML header
    streamInfo      []byte      // Stored incoming stream info
    trackInfo       []byte      // Stored incoming track info

    mu				sync.RWMutex
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

var InputStream InputStreamContext
