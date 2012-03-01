package web

import (
	"fmt"
	"net/http"
	"gortmpd/webm"
)

var context *webm.Context

func clientHandler(w http.ResponseWriter, req *http.Request) {
		var packet webm.DispatchPacket

		// Return WebM header
		w.Header().Add("Content-Type", "video/webm")
		w.WriteHeader(http.StatusOK)

		// Write EBML intro
		ebmlHeader := *context.InputStream.GetEBMLHeader()
		packet.Id = 0xA45DFA3
		packet.Length = uint64(len(ebmlHeader))
		packet.Data = ebmlHeader
		w.Write(packet.GetByteRepresentation())

		// Prepare segment header
		packet.Id = 0x8538067
		packet.Length = 0xFFFFFFFFFFFFFFFF
		packet.Data = nil
		w.Write(packet.GetByteRepresentation())
		
		// Send stream info
		segmentInfo := *context.InputStream.GetStreamInfo()
		packet.Id = 0x549A966
		packet.Length = uint64(len(segmentInfo))
		packet.Data = segmentInfo
		w.Write(packet.GetByteRepresentation())

		// Send track info
		trackInfo := *context.InputStream.GetTrackInfo()
		packet.Id = 0x654AE6B
		packet.Length = uint64(len(trackInfo))
		packet.Data = trackInfo
		w.Write(packet.GetByteRepresentation())

		fmt.Println("Client connected, header written...")
}

func StartOutput(ctx *webm.Context) {
	fmt.Println("Output web server starting....")
	context = ctx

	http.HandleFunc("/stream/", clientHandler)
	http.ListenAndServe(":8080", nil)
}