package service

//全部浮点数用int表示前三位表示小数位

type DayRecord struct {
	Day    int `gorm:"PRIMARY_KEY"`
	Open   int `gorm:"PRIMARY_KEY"`
	Code   int
	High   int
	Low    int
	Close  int
	Vol    uint64
	Amount uint64
	Zt     bool
	Dt     bool
	Zf     int
	Dm     bool
	Dr     bool
	Pb     bool
	A20    bool
	Lb     int
}
type DayStat struct {
	Day    int `gorm:"PRIMARY_KEY"`
	High   int
	Low    int
	Close  int
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
