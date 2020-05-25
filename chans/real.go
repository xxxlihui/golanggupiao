package chans

import "nn/data"

var RealChan = make(chan *struct {
	Message string
	Data    []*data.RealData
}, 10)

var EndChan = make(chan []*data.RealData)
