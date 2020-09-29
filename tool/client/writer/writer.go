package writer

import "nn/data"

type Writer interface {
	Write(records []*data.DayRecord) error
}
