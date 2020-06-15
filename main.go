package main

import (
	"github.com/urfave/cli/v2"
	"nn/service"
	"os"
)

func init()  {
	os.Setenv("TZ","Etc/GMT+8")
}
func main() {

	app := cli.App{
		Name:    "股票数据后台",
		Version: "1.0",
		Usage:   "该客户端需要和管理端配合，由管理端调度和管理",
		Authors: []*cli.Author{{Name: "lhn", Email: "550124023@qq.com"}},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "db_host",
				Aliases: []string{"s"},
				Usage:   "数据库ip,默认127.0.0.1",
				EnvVars: []string{"DB_HOST"},
				Value:   "127.0.0.1",
			},
			&cli.IntFlag{
				Name:    "db_port",
				Usage:   "数据库端口默认5432",
				EnvVars: []string{"DB_PORT"},
				Value:   5432,
			},
			&cli.StringFlag{
				Name:    "db_user",
				Usage:   "数据登录账户 默认postgres",
				EnvVars: []string{"DB_USER"},
				Value:   "postgres",
			},
			&cli.StringFlag{
				Name:    "db_pwd",
				Usage:   "数据库密码，默认 123",
				EnvVars: []string{"DB_PWD"},
				Value:   "123",
			},
			&cli.StringFlag{
				Name:    "db_name",
				Usage:   "数据库名称",
				EnvVars: []string{"DB_NAME"},
				Value:   "gupiao",
			},
			&cli.StringFlag{
				Name:    "port",
				Usage:   "服务端口,默认 :80",
				EnvVars: []string{"PORT"},
				Value:   ":80",
			},
		},
		Action: func(context *cli.Context) error {
			defer service.CloseDb()
			service.InitDb(
				context.String("db_host"),
				context.String("db_user"),
				context.String("db_pwd"),
				context.String("db_name"),
				context.Int("db_port"),
			)

			return service.StartServer(context.String("port"))
		},
	}

	app.Run(os.Args)
}
