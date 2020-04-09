package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"io/ioutil"
	"net/http"
	"nn/data"
	"nn/log"
	"nn/spider"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var url, token, folder string
	var day int
	flags := []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "server-url",
			Aliases:     []string{"s", "URL"},
			Usage:       "服务端的地址url",
			EnvVars:     []string{"URL"},
			Required:    true,
			Destination: &url,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "token",
			Aliases:     []string{"t"},
			Usage:       "认证token",
			EnvVars:     []string{"TOKEN"},
			Required:    true,
			Destination: &token,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "day",
			Usage:       "要开始导入的时间,int格式：20200318",
			Aliases:     []string{"d"},
			EnvVars:     []string{"DAY"},
			Required:    true,
			Destination: &day,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "folder",
			Aliases:     []string{"f"},
			Usage:       "通达信的目录",
			EnvVars:     []string{"FOLDER"},
			Required:    true,
			Destination: &folder,
		}),
	}
	app := cli.App{
		Name:    "股票数据通达信数据提取客户端",
		Version: "1.0",
		Usage:   "该客户端抓取本地通达信数据,并提交到后台",
		Authors: []*cli.Author{{Name: "lhn", Email: "550124023@qq.com"}},
		Flags:   flags,
		Before: func(context *cli.Context) error {
			log.Debug("-----------------before")
			return nil
		},
		Action: func(context *cli.Context) error {
			fds := []string{filepath.Join(folder, "vipdoc/sh/lday"),
				filepath.Join(folder, "vipdoc/sz/lday")}
			for _, f := range fds {
				if err := ReadFolder(f, url, token, day); err != nil {
					return err
				}
			}
			return nil
		},
	}
	app.Run(os.Args)
}

func ReadFolder(fd, url, token string, day int) error {
	fs, err := ioutil.ReadDir(fd)
	if err != nil {
		return err
	}
	for _, e := range fs {
		if strings.HasPrefix(e.Name(), "sh600") ||
			strings.HasPrefix(e.Name(), "sh601") ||
			strings.HasPrefix(e.Name(), "sh603") ||
			strings.HasPrefix(e.Name(), "sz00") ||
			strings.HasPrefix(e.Name(), "sz30") {
			code := strings.TrimRight(e.Name(), ".day")
			code = strings.ToLower(code)
			records, err := ReadFile(day, code, filepath.Join(fd, e.Name()))
			if err != nil {
				return err
			}
			err = PostData(url, token, &records)
			if err != nil {
				return err
			}
		}
	}
	return nil
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

func ReadFile(day int, code, fp string) ([]*data.DayRecord, error) {
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
		fmt.Printf("day:%d,record:%+v\n", day, record)
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
	record.Open = float64(k) / 100
	err = binary.Read(buf, binary.LittleEndian, &k)
	checkErr(err)
	record.High = float64(k) / 100
	err = binary.Read(buf, binary.LittleEndian, &k)
	checkErr(err)
	record.Low = float64(k) / 100
	err = binary.Read(buf, binary.LittleEndian, &k)
	checkErr(err)
	record.Close = float64(k) / 100
	var f float32
	err = binary.Read(buf, binary.LittleEndian, &f)
	checkErr(err)
	record.Amount = float64(f)
	err = binary.Read(buf, binary.LittleEndian, &k)
	checkErr(err)
	record.Vol = uint64(k)
	return record, nil
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
