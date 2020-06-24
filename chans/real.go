package chans

import "nn/data"

var RealChan = make(chan *struct {
	Message string
	Data    []*struct {
		data.PCode
		data.PDayData
		data.PDaySample
	}
}, 10)

var EndChan = make(chan []*struct {
	data.PCode
	data.PDayData
	data.PDaySample
},
)
