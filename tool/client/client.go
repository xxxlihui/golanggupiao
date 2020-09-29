package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"io/ioutil"
	"nn/data"
	"nn/log"
	"nn/tool/client/writer"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func main() {
	var url, token, folder string
	var day, concurrency int
	var flags = []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "target-url",
			Aliases:     []string{"target", "URL"},
			Usage:       "目标地址，http地址或者目录，http地址以http开头",
			EnvVars:     []string{"URL"},
			Destination: &url,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "token",
			Aliases:     []string{"t"},
			Usage:       "走http时的认证token",
			EnvVars:     []string{"TOKEN"},
			Destination: &token,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "day",
			Usage:       "要开始导入的时间,int格式：20200318",
			Aliases:     []string{"d"},
			EnvVars:     []string{"DAY"},
			Destination: &day,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "folder",
			Aliases:     []string{"f"},
			Usage:       "通达信的目录",
			EnvVars:     []string{"FOLDER"},
			Destination: &folder,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "concurrency",
			Aliases:     []string{"c"},
			Usage:       "并发数，默认4并发",
			EnvVars:     []string{"C"},
			Destination: &concurrency,
			Value:       4,
		}),
	}
	app := cli.App{
		Name:    "股票数据通达信数据提取客户端",
		Version: "1.0",
		Usage:   "该客户端抓取本地通达信数据,并提交到后台",
		Authors: []*cli.Author{{Name: "lhn", Email: "550124023@qq.com"}},
		Flags:   flags,
		Before: altsrc.InitInputSourceWithContext(flags, func(context *cli.Context) (context2 altsrc.InputSourceContext, e error) {
			return altsrc.NewYamlSourceFromFile("set.yaml")
		}),
		Action: func(context *cli.Context) error {
			fds := []string{filepath.Join(folder, "vipdoc/sh/lday"),
				filepath.Join(folder, "vipdoc/sz/lday")}
			var w writer.Writer
			if strings.HasPrefix(url, "http") {
				w = &writer.HttpWriter{
					Token: token,
					Url:   url,
				}
			} else {
				os.MkdirAll(url, os.ModePerm)
				w = &writer.FileWriter{Folder: url}
			}
			start := time.Now()
			for _, f := range fds {
				if err := ReadFolder(f, url, token, day, concurrency, w); err != nil {
					return err
				}
			}
			diff := time.Now().Sub(start).Milliseconds()
			log.Info("耗时%d", diff)
			return nil
		},
	}
	app.Run(os.Args)
}

type readFile struct {
	day      int
	code     string
	filepath string
}

func check2Err(c chan struct{}, err error) {

}
func ReadFolder(fd, url, token string, day, concurrency int, writer writer.Writer) error {
	fs, err := ioutil.ReadDir(fd)
	if err != nil {
		return err
	}
	chans := make(chan *readFile, 5)
	exitChan := make(chan error)
	wait := sync.WaitGroup{}
	for i := 0; i < concurrency; i++ {
		go func() {
			for {
				select {
				case readFile := <-chans:
					records, err := ReadFile(readFile)
					if err != nil {
						exitChan <- err
						return
					}
					err = writer.Write(records)
					if err != nil {
						exitChan <- err
						return
					}
					wait.Done()
				case <-exitChan:
					return
				}
			}
		}()
	}
	for _, e := range fs {
		if strings.HasPrefix(e.Name(), "sh600") ||
			strings.HasPrefix(e.Name(), "sh601") ||
			strings.HasPrefix(e.Name(), "sh603") ||
			strings.HasPrefix(e.Name(), "sh688") ||
			strings.HasPrefix(e.Name(), "sz00") ||
			strings.HasPrefix(e.Name(), "sz30") {
			code := strings.TrimRight(e.Name(), ".day")
			code = strings.ToLower(code)
			wait.Add(1)
			chans <- &readFile{
				day:      day,
				code:     code,
				filepath: filepath.Join(fd, e.Name()),
			}

		}
	}
	go func() {
		wait.Wait()
		close(exitChan)
	}()
	err, ok := <-exitChan
	if !ok {
		log.Info("完成")
		return nil
	}
	return err
}

func ReadFile(file *readFile) ([]*data.DayRecord, error) {
	fp := file.filepath
	code := file.code
	day := file.day
	log.Debug("读取文件%s", fp)
	bys, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	idx := -1
	l := len(bys)
	records := make([]*data.DayRecord, 0, 20)
	for {
		idx++
		if idx*32 >= l {
			break
		}
		record, err := readDay(bys[idx*32 : idx*32+32])
		if err != nil {
			return nil, err
		}

		if record.Day < day {
			continue
		}
		record.Code = code
		//fmt.Printf("day:%d,record:%+v\n", day, record)
		records = append(records, record)
	}
	return records, nil
}

// 每32个字节为一天数据
//    每4个字节为一个字段，每个字段内低字节在前
//    00 ~ 03 字节：年月日, 整型
//    04 ~ 07 字节：开盘价*100， 整型
//    08 ~ 11 字节：最高价*100,  整型
//    12 ~ 15 字节：最低价*100,  整型
//    16 ~ 19 字节：收盘价*100,  整型
//    20 ~ 23 字节：成交额（元），float型
//    24 ~ 27 字节：成交量（手），整型
//    28 ~ 31 字节：附加信息不知道有什么用, 整型
func readDay(bys []byte) (record *data.DayRecord, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = errors.New(fmt.Sprintf("读取日线数据错误:%v", p))
		}
	}()
	record = &data.DayRecord{}
	buf := bytes.NewBuffer(bys)
	var k uint32
	err = binary.Read(buf, binary.LittleEndian, &k)
	checkErr(err)
	record.Day = int(k)
	err = binary.Read(buf, binary.LittleEndian, &k)
	checkErr(err)
	record.Open = decimal.NewFromInt(int64(k))
	err = binary.Read(buf, binary.LittleEndian, &k)
	checkErr(err)
	record.High = decimal.NewFromInt(int64(k))
	err = binary.Read(buf, binary.LittleEndian, &k)
	checkErr(err)
	record.Low = decimal.NewFromInt(int64(k))
	err = binary.Read(buf, binary.LittleEndian, &k)
	checkErr(err)
	record.Close = decimal.NewFromInt(int64(k))
	var f float32
	err = binary.Read(buf, binary.LittleEndian, &f)
	checkErr(err)
	record.Amount = decimal.RequireFromString(fmt.Sprintf("%d", uint64(f*100)))
	err = binary.Read(buf, binary.LittleEndian, &k)
	checkErr(err)
	record.Volume = decimal.RequireFromString(fmt.Sprintf("%d", uint64(k)))
	return record, nil
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
