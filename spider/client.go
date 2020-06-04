package spider

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

//对httpClient的简单封装使得httpClient简单使用，并支持cookie
type Client struct {
	Client     *http.Client
	UserAgent  string                                         //模拟浏览器的User-Agent
	OnRequest  func(req *http.Request) (*http.Request, error) //请求前统一的处理函数
	OnResponse func(rsp *http.Response, err error)            //返回后统一的处理函数
}

func NewClient(userAgent string) *Client {
	client := &Client{Client: newClient(), UserAgent: userAgent}
	return client
}

//初始化一个带cookie管理的http.Client
func newJarClient() *http.Client {
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	jar, _ := cookiejar.New(nil)
	client.Jar = jar
	return client
}
func newClient() *http.Client {
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	return client
}

//初始化一个爬虫客户端，http.Client
func NewJarClient(userAgent string) *Client {
	client := &Client{Client: newJarClient(), UserAgent: userAgent}
	return client
}

func (c *Client) Get(url, referer string, header http.Header, ctx context.Context) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	copyHeader(header, req.Header)
	return c.Do(referer, req, ctx)
}
func (c *Client) Head(url, referer string, header http.Header, ctx context.Context) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}
	copyHeader(header, req.Header)
	return c.Do(referer, req, ctx)
}
func (c *Client) PostForm(url, referer string, header http.Header, value url.Values, ctx context.Context) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(value.Encode()))
	if err != nil {
		return nil, err
	}
	copyHeader(header, req.Header)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.Do(referer, req, ctx)
}

/*
未实现
func (c *Client) PostMultipartForm(url, referer string, header http.Header, value multipart.Form) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url,nil)
	writer:=multipart.NewWriter(&bytes.Buffer{})
	writer.

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.Do(referer, req)
}*/
func (c *Client) Post(url, referer, contentType string, header http.Header, body io.Reader, ctx context.Context) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	copyHeader(header, req.Header)
	req.Header.Set("Content-Type", contentType)
	return c.Do(referer, req, ctx)
}
func (c *Client) PostValue(url, referer string, header http.Header, value interface{}, ctx context.Context) (*http.Response, error) {
	bys, err := json.Marshal(&value)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(bys)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	copyHeader(header, req.Header)
	req.Header.Set("Content-Type", "application/json")
	return c.Do(referer, req, ctx)
}

//执行请求,from referer
func (c *Client) Do(referer string, req *http.Request, ctx context.Context) (*http.Response, error) {
	c.setHeaderDefault(referer, req.Header)
	if c.OnRequest != nil {
		_req, err := c.OnRequest(req)
		if err != nil {
			return nil, err
		}
		req = _req
	}
	if ctx == nil {

	} else {
		req = req.WithContext(ctx)
	}
	rsp, err := c.Client.Do(req)
	if c.OnResponse != nil {
		c.OnResponse(rsp, err)
	}
	return rsp, err
}

//获取cookies
//u 获取某一个站点的cookie信息的url
//一般情况直接忽略error的判断
func (c *Client) Cookies(u string) []*http.Cookie {
	ur, err := url.Parse(u)
	if err != nil {
		return nil
	}
	return c.Client.Jar.Cookies(ur)
}
func (c *Client) setHeaderDefault(referer string, header http.Header) http.Header {
	if header == nil {
		return http.Header{"User-Agent": []string{c.UserAgent}, "Referer": []string{referer}}
	}
	if header.Get("Referer") == "" {
		header.Set("Referer", referer)
	}
	if header.Get("User-Agent") == "" {
		header.Set("User-Agent", c.UserAgent)
	}
	return header
}

func copyHeader(src, target http.Header) {
	if src == nil {
		return
	}
	for k, v := range src {
		vv := make([]string, len(v))
		copy(vv, v)
		target[k] = vv
	}
}

func randomBoundary() string {
	var buf [30]byte
	_, err := io.ReadFull(rand.Reader, buf[:])
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", buf[:])
}
