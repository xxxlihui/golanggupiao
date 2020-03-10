package service

type DayRecord struct {
	Id     int //day*1000000+code
	Day    int
	Code   int
	High   float32
	Low    float32
	Close  float32
	Vol    uint64
	Amount uint64
	Zt     bool
	Dt     bool
	Zf     float32
	Dm     bool
	Dr     bool
	Pb     bool
	Stop   bool
	A20    bool
	Lb     int
}
type DayStat struct {
	Day    int
	High   float32
	Low    float32
	Close  float32
	Vol    uint64
	Amount uint64
	At     int
	Dt     int
	Zf     int
	Dm     int
	Dr     int
	Pb     int
	a20    int
}
