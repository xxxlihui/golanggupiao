package writer

import (
	"errors"
	"net/http"
	"nn/data"
	"nn/spider"
)

type HttpWriter struct {
	Token string
	Url   string
}

func (h *HttpWriter) Write(records []*data.DayRecord) error {
	return PostData(h.Url, h.Token, records)
}

func PostData(url, token string, value interface{}) error {
	client := spider.NewClient(spider.RandomUserAgent())
	rsp, err := client.PostValue(url, "", http.Header{"token": {token}}, value, nil)
	if err != nil {
		return err
	}
	if rsp.StatusCode != http.StatusOK {
		return errors.New("请求失败")
	}
	return nil
}
