package web

import (
	"fmt"
	"net/http"
	"gortmpd/webm"
)

var context *webm.Context

func clientHandler(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "video/webm")
		w.WriteHeader(http.StatusOK)
		w.Write(*context.InputStream.GetEBMLHeader())
		w.Write(*context.InputStream.GetStreamInfo())
		w.Write(*context.InputStream.GetTrackInfo())

		fmt.Println("Client connected, header written...")
}

func StartOutput(ctx *webm.Context) {
	fmt.Println("Output web server starting....")
	context = ctx

	http.HandleFunc("/stream/", clientHandler)
	http.ListenAndServe(":8080", nil)
}