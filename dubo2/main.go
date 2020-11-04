package main

import (
	"bytes"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"runtime"
	"sort"
	"sync"
)

func main() {
	threadCount := 4 //线程数
	total := 10000   //总投资金额
	xfr := 1.42      //特朗普赔率
	yfr := 3.25      //拜登赔率
	row := 10
	model := 0
	app := &cli.App{
		Name: "总赢算法",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "thread",
				Usage:       "并发数，默认根据cpu核数来做",
				Value:       runtime.NumCPU(),
				Destination: &threadCount,
			},
			&cli.IntFlag{
				Name:        "total",
				Usage:       "一共的本金投入",
				Value:       10000,
				Destination: &total,
			},
			&cli.IntFlag{
				Name:        "row",
				Usage:       "返回的行数",
				Value:       10000,
				Destination: &row,
			},
			&cli.IntFlag{
				Name:        "model",
				Usage:       "模式0 均衡模式,无论谁赢盈利的差距最少 1 特朗普赢 2 拜登赢",
				Value:       0,
				Destination: &model,
			},
			&cli.Float64Flag{
				Name:        "x",
				Usage:       "特朗普赔率",
				Value:       1.28,
				Destination: &xfr,
			},
			&cli.Float64Flag{
				Name:        "y",
				Usage:       "拜登赔率",
				Value:       3.75,
				Destination: &yfr,
			},
		},
		Action: func(context *cli.Context) error {
			//把total分成线程数等份
			max := total / threadCount
			less := total % threadCount

			xf := int(xfr * 100)
			yf := int(yfr * 100)
			mux := sync.Mutex{}
			str := &bytes.Buffer{}
			g := &sync.WaitGroup{}
			totalrst := make([][5]int, 0, total*total/2)
			for c := 0; c < threadCount; c++ {
				start := c * max
				end := start + max
				if c == threadCount-1 {
					end = end + less
				}
				start++
				g.Add(1)
				go func(start, end, total, xf, yf int) {
					fmt.Printf("开始 start:%d,end：%d\n", start, end)
					rst := enums(start, end, total, xf, yf)
					mux.Lock()
					for _, v := range rst {
						totalrst = append(totalrst, v)
					}
					fmt.Printf("完成 start:%d,end：%d\n", start, end)
					mux.Unlock()
					g.Done()
				}(start, end, total, xf, yf)
			}
			g.Wait()
			fmt.Printf("一共%d条记录\n", len(totalrst))
			//各个模式分析
			//特朗普均衡
			sort.Slice(totalrst, func(i, j int) bool {
				return totalrst[i][2] > totalrst[j][2]
			})
			sort.Slice(totalrst, func(i, j int) bool {
				return totalrst[i][4] < totalrst[j][4]
			})
			rst := totalrst[0:row]
			sort.Slice(rst, func(i, j int) bool {
				return rst[i][2] < totalrst[j][2]
			})
			//row=len(totalrst)
			for i := 0; i < row; i++ {
				v := rst[i]
				str.WriteString(fmt.Sprintf("%d,%d,%d,%d,%d\n", v[0], v[1], v[2], v[3], v[4]))
			}
			ioutil.WriteFile("特朗普均衡.csv", str.Bytes(), os.ModePerm)
			fmt.Printf("特朗普均衡 完成\n")
			//拜登均衡
			str.Reset()
			sort.Slice(totalrst, func(i, j int) bool {
				return totalrst[i][3] > totalrst[j][3]
			})
			sort.Slice(totalrst, func(i, j int) bool {
				return totalrst[i][4] < totalrst[j][4]
			})
			rst = totalrst[0:row]
			sort.Slice(rst, func(i, j int) bool {
				return rst[i][3] < totalrst[j][3]
			})
			for i := 0; i < row; i++ {
				v := rst[i]
				str.WriteString(fmt.Sprintf("%d,%d,%d,%d,%d\n", v[0], v[1], v[2], v[3], v[4]))
			}
			ioutil.WriteFile("拜登均衡.csv", str.Bytes(), os.ModePerm)
			fmt.Printf("拜登均衡 完成\n")
			//特朗普模式
			str.Reset()
			sort.Slice(totalrst, func(i, j int) bool {
				return totalrst[i][2] > totalrst[j][2]
			})
			for i := 0; i < row; i++ {
				v := totalrst[i]
				str.WriteString(fmt.Sprintf("%d,%d,%d,%d,%d\n", v[0], v[1], v[2], v[3], v[4]))
			}
			ioutil.WriteFile("特朗普.csv", str.Bytes(), os.ModePerm)
			fmt.Printf("特朗普 完成\n")
			//拜登模式
			str.Reset()
			sort.Slice(totalrst, func(i, j int) bool {
				return totalrst[i][3] > totalrst[j][3]
			})
			for i := 0; i < row; i++ {
				v := totalrst[i]
				str.WriteString(fmt.Sprintf("%d,%d,%d,%d,%d\n", v[0], v[1], v[2], v[3], v[4]))
			}
			ioutil.WriteFile("拜登.csv", str.Bytes(), os.ModePerm)
			fmt.Printf("拜登 完成\n")
			fmt.Printf("投特朗普资金,投拜登资金,特朗普赢的盈利,拜登赢的盈利\n")

			return nil
		},
	}
	app.Run(os.Args)

}

func enums(start, end, total, xf, yf int) [][5]int {
	rst := make([][5]int, 0, (end-start)/2)
	for x := start; x <= end; x++ {
		for y := 1; y <= total && (x+y) <= total; y++ {
			xz := x*xf - y*100
			yz := y*yf - x*100
			if xz > 0 && yz > 0 {
				rst = append(rst, [5]int{x, y, xz / 100, yz / 100, abs(xz - yz)})
			}
		}
	}
	return rst
}
func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -1 * i
}
