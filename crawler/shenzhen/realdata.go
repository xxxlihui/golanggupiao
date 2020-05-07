package shenzhen

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"net/http"
	"nn/spider"
)

var rawListUrl = "http://www.szse.cn/api/report/ShowReport?SHOWTYPE=xlsx&CATALOGID=1110&TABKEY=tab1&random=0.6639499834954136"

//获取股票价格
func GetPrice(code string){

}

//获取股票列表
func GetList() ([]*ListData, error) {
	client := spider.NewClient(spider.RandomUserAgent())
	bys, _, _, err := spider.GetResponseBytes(func() (*http.Response, error) {
		return client.Get(rawListUrl, "", nil, nil)
	})
	if err != nil {
		return nil, err
	}
	f, err := excelize.OpenReader(bytes.NewBuffer(bys))
	if err != nil {
		return nil, err
	}
	rows, err := f.GetRows(f.GetSheetName(1))
	if err != nil {
		return nil, err
	}
	datas:=make([]*ListData,len(rows)-1)
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		data:=&ListData{
			Code: row[5],
			Name: row[6],
			LaunchDate: row[7],
			GeneralCapital: row[8],
			FlowEquity: row[9],
			Area: row[15],
			Pro: row[16],
			City: row[17],
			Industry: row[18],
			HomePage: row[19],
		}
		datas[i-1]=data
	}
	for _, row := range rows {
		fmt.Print(row)
	}
	return nil, nil
}

type ListData struct {
	Code string
	Name string
	//上市日期
	LaunchDate string
	//总股本
	GeneralCapital string
	//流通股本
	FlowEquity string
	//地区
	Area string
	//省份
	Pro string
	//城市
	City string
	//所属行业
	Industry string
	//公司主页
	HomePage string
}
