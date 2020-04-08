package spider

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"math/rand"
	"net/http"
	_url "net/url"
)

type ResponseFunc func() (*http.Response, error)

func GetAbsoluteUrl(from, u string) (string, error) {
	base, err := _url.Parse(from)
	if err != nil {
		return "", err
	}
	absURL, err := base.Parse(u)
	if err != nil {
		return "", nil
	}
	return absURL.String(), nil
}

func GetResponseBytes(rspFunc ResponseFunc) ([]byte, http.Header, *http.Response, error) {
	rsp, err := rspFunc()
	if err != nil {
		return nil, nil, rsp, err
	}
	defer rsp.Body.Close()
	_body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, rsp.Header, rsp, nil
	}
	return _body, rsp.Header, rsp, nil
}
func GetResponseString(rspFunc ResponseFunc) (string, http.Header, *http.Response, error) {
	_body, _header, rsp, err := GetResponseBytes(rspFunc)
	if err != nil {
		return "", _header, rsp, err
	}
	return string(_body), _header, rsp, nil
}
func GetResponseGoqueryDoc(rspFunc ResponseFunc) (*goquery.Document, http.Header, *http.Response, error) {
	rsp, err := rspFunc()
	if err != nil {
		return nil, nil, nil, err
	}
	defer rsp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	if err != nil {
		return nil, rsp.Header, rsp, err
	}
	return doc, rsp.Header, rsp, nil
}

func RandInt(max, min int) int {
	d := rand.Intn(max)
	if d < min {
		return min
	}
	return d
}
