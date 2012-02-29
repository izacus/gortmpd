package dispatcher

import (
	"gortmpd/file"
	"gortmpd/webm"
)

func DispatchPackets(context webm.Context) {
	output_chan := file.GetOutputChannel("output.webm")

	for {
		packet := <- context.DispatchChannel
		bytes := packet.GetByteRepresentation()

		for i := 0; i < len(bytes); i++ {
			output_chan <- bytes[i]
		}
	}
}