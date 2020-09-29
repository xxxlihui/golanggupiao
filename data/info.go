package data

//股票信息数据结构

type StockInfo struct {
	PCode
	Name         string `json:"name" gorm:"varchar(20)"` //名称
	LaunchDate   int    `json:"launchDate"`              //上市日期
	DelistedDate int    `json:"delistedDate"`            //退市日期
	//Tags         []string `json:"tags"`                    //标签信息
}

//StockInfos 股票的信息
type StockInfos struct {
	StockInfos []*StockInfo
	//索引
	codeMapper map[string]*StockInfo
}

func NewDefaultStockInfos() *StockInfos {
	return &StockInfos{codeMapper: map[string]*StockInfo{}}
}

func NewStockInfos(stockInfos []*StockInfo) *StockInfos {
	s := &StockInfos{codeMapper: map[string]*StockInfo{}}
	s.Init(stockInfos)
	return s
}

func (receiver *StockInfos) Get(code string) *StockInfo {
	return receiver.codeMapper[code]
}

func (receiver *StockInfos) Init(stockInfos []*StockInfo) {
	receiver.StockInfos = stockInfos
	for _, info := range stockInfos {
		receiver.codeMapper[info.Code] = info
	}
}
