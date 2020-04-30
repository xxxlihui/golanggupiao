package shanghai

import (
	"net/http"
	"nn/spider"
	"strings"
)

//setcode 0 深圳 sz 1 上海 sh
var rawUrl = "http://www.tdx.com.cn/url/sns.asp?setcode={setcode}&code={code}"

func GetUrl(code string) (string, error) {
	setcode := "0"
	if strings.HasPrefix(code, "sh") {
		setcode = "1"
	}
	code= code[2:]
	targetUrl := strings.Replace(rawUrl, "setcode", setcode, 1)
	targetUrl = strings.Replace(targetUrl, "code", code, 1)

	client := spider.NewClient(spider.RandomUserAgent())
	rspStr, _, _, err := spider.GetResponseString(nil /*func(bys []byte, head http.Header) string {
			by, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(bys)
			return string(by)
		}*/, func() (response *http.Response, e error) {
		return client.Get(targetUrl,
			"", nil, nil)
	})
	if err != nil {
		return "", err
	}
	return rspStr, nil
}
