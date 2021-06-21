package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"net/url"
	"nn/spider"
	"nn/tool/m38u/cry"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func main() {
	var url, dir, filename string
	var threadCount int
	flags := []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "url",
			Aliases:     []string{"u"},
			Usage:       "地址",
			EnvVars:     []string{"URL"},
			Destination: &url,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "filename",
			Aliases:     []string{"n"},
			Usage:       "文件名",
			Required:    true,
			Destination: &filename,
		}),

		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "dir",
			Aliases:     []string{"d"},
			Usage:       "目录",
			Value:       "./",
			Destination: &dir,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "thread",
			Aliases:     []string{"t"},
			Usage:       "并发下载数",
			Destination: &threadCount,
			Value:       10,
		}),
	}
	app := cli.App{
		Name:    "m38u下载器",
		Version: "1.0",
		Usage:   "m38u下载器,并自动合并",
		Authors: []*cli.Author{{Name: "lhn", Email: "550124023@qq.com"}},
		Flags:   flags,
		Before: altsrc.InitInputSourceWithContext(flags, func(context *cli.Context) (context2 altsrc.InputSourceContext, e error) {
			if _, err := os.Lstat("set.yaml"); err != nil && os.IsNotExist(err) {
				os.OpenFile("set.yaml", os.O_CREATE, os.ModePerm)
			}
			return altsrc.NewYamlSourceFromFile("set.yaml")
		}),
		Action: func(context *cli.Context) error {
			mkdir(dir)
			down(url, dir, filename, threadCount)
			return nil
		},
	}
	app.Run(os.Args)
}
func mkdir(dir string) {
	os.MkdirAll(dir, os.ModePerm)
}
func fileExistsForNewFile(dir, file string) string {
	id := 0
	for {
		f := file
		if id > 0 {
			f = file + fmt.Sprintf("(%d)", id)
		}
		ph := filepath.Join(dir, f)
		_, err := os.Stat(ph)
		if err != nil {
			if os.IsNotExist(err) {
				return f
			}
		}
		id++
	}
}
func down(_url, dir, fileName string, thread int) {
	g := sync.WaitGroup{}
	list, err := getList(_url, dir, fileName)
	if err != nil {
		println("下载错误", err.Error())
		return
	}
	//l := strings.LastIndex(_url, "/")
	name := fileName

	tmpDir := filepath.Join(dir, fileName+"_tmp")
	os.MkdirAll(tmpDir, os.ModePerm)
	total := len(list)
	tr := thread
	chans := make(chan *item, tr)
	okc := make(chan *item)
	g.Add(total)
	//enc := make(chan struct{})
	for i := 0; i < tr; i++ {
		go func() {
			for {
				item := <-chans
				for {
					println("下载", item.name)
					u, _ := url.Parse(_url)
					u2, _ := u.Parse(item.name)
					err := downLoadOne(u2.String(), tmpDir, fmt.Sprintf("%d", item.id), item.decryptFunc)
					if err != nil {
						println(item.name, "id=(", item.id, ")下载失败")
					} else {
						okc <- item
						g.Done()
						break
					}
				}
			}
		}()
	}
	go func() {
		ok := 0
		for {
			i := <-okc
			ok++
			println("完成", i.name, "已完成", ok, "一共", total, "  ", fmt.Sprintf("%.2f", float64(ok)/float64(total)*100))
			if ok >= total {
				break
			}
		}
	}()
	for _, v := range list {
		chans <- v
		//println("还有", total-ok)
	}

	g.Wait()
	//合并文件
	merge(list, dir, tmpDir, name)
	os.RemoveAll(tmpDir)
	println("下载完成")
}

func merge(items []*item, dir, tmpDir, name string) {
	filename := fileExistsForNewFile(dir, name)
	p := filepath.Join(dir, filename+".ts")
	f, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	defer f.Close()
	if err != nil {
		println("创建合并文件失败", err.Error())
		return
	}
	total := len(items)
	ok := 0
	per := total / 10
	for _, v := range items {
		pth := filepath.Join(tmpDir, fmt.Sprintf("%d", v.id))
		bys, err := ioutil.ReadFile(pth)
		/*if decryptFunc != nil {
			bys = decryptFunc(bys)
		}*/
		if err != nil {
			println("读取子文件失败", err.Error())
			return
		}
		f.Write(bys)
		ok++
		if ok == per {
			fmt.Println("合并", fmt.Sprintf("%.2f", float64(ok)/float64(total)))
		}
	}
}

func write(bys []byte, pth string) error {
	f, err := os.OpenFile(pth, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		println("创建文件失败", err.Error())
		return err
	}
	defer f.Close()
	_, err = f.Write(bys)
	return err
}

func downLoadOne(url, dir, name string, decryptFunc func(bys []byte) []byte) error {
	pth := filepath.Join(dir, name)
	info, err := os.Stat(pth)
	if err == nil {
		if info.Size() > 0 {
			return nil
		}
	}

	client := spider.NewClient(spider.RandomUserAgent())
	println("下载", url)
	bys, _, _, err := spider.GetResponseBytes(func() (*http.Response, error) {
		return client.Get(url, "", nil, nil)
	})
	if err != nil {
		//item.status = 0
		return err
	}
	if decryptFunc != nil {
		bys = decryptFunc(bys)
	}
	err = write(bys, pth)

	return err
}
func getTimeoutContext(second int) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(second)*time.Second)
	return ctx
}
func getList(url, dir, name string) ([]*item, error) {
	c := ""
	var strs []string
	cBytes, err := ioutil.ReadFile(filepath.Join(dir, name+".m3u8"))
	if err != nil {
		err = nil
		client := spider.NewClient(spider.RandomUserAgent())
		str, _, _, err1 := spider.GetResponseString(nil, func() (*http.Response, error) {
			return client.Get(url, "", nil, nil)
		})
		if err1 != nil {
			return nil, err1
		}
		strs = strings.Split(str, "\n")
		c = strings.Join([]string{url, str}, "\n")
		ioutil.WriteFile(filepath.Join(dir, name+".m3u8"), []byte(c), os.ModePerm)
	} else {
		c = string(cBytes)
		strs = strings.Split(c, "\n")
		if len(strs) > 1 {
			strs = strs[1:]
		}
	}
	rstrs := make([]*item, 0, len(strs))
	id := 0
	var decryptFunc func(bys []byte) []byte
	for _, v := range strs {
		if v == "" {
			continue
		}
		if !strings.HasPrefix(v, "#") {
			id++
			rstrs = append(rstrs, &item{id: id, name: v, status: 0, decryptFunc: decryptFunc})
		} else {
			if strings.HasPrefix(v, "#EXT-X-KEY") {
				//加密视频
				decryptFunc, _ = cry.GetDecryptFunc(url, v)
			}
		}
	}
	return rstrs, err
}

type item struct {
	id          int
	name        string
	status      int
	decryptFunc func(bys []byte) []byte
	//lock   sync.Mutex
}

func (this *item) Name() string {
	idx := strings.LastIndex(this.name, "/")
	if idx >= 0 {
		return this.name[idx:]
	}
	return this.name
}
func (this *item) start() {
	//this.lock.Lock()
	this.status = 2
}
func (this *item) end(ok int) {
	this.status = ok
	//this.lock.Unlock()
}
func (this *item) rest() {
	this.status = 0
}
