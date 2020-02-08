package cmds

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chalvern/apollo/cmds/server"
	"github.com/chalvern/apollo/cmds/sub"
	"github.com/chalvern/apollo/configs/constants"
	"github.com/chalvern/apollo/configs/initializer"
	"github.com/chalvern/sugar"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

// AppInit init
func AppInit() *cli.App {
	app := cli.NewApp()

	// 1. base application info
	// 1. 基础的应用信息
	app.Name = "apollo"
	app.Author = "zhjw43"
	app.Version = "0.0.1"
	app.Copyright = "wheresmile group"
	app.Usage = "backend of some projects"

	// 2. flags
	// 2. 可传入的标识
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "./configs/config.yml",
			Usage: "load configuration from yaml `FILE`",
		},
	}

	app.Before = func(c *cli.Context) error {
		// viper must be config at first
		// viper 必须第一个被配置，因为其他的配置项会依赖它
		configFile := c.String("config")
		sugar.Infof("init viper with config File: %s", configFile)
		initializer.InitViperWithFile(configFile)
		return nil
	}

	// register sub commands
	sub.Init(app)

	app.Action = mainJob
	return app
}

func mainJob(c *cli.Context) error {
	ctx, cancel := context.WithCancel(context.Background())

	// initializer
	initializer.InitJwt()
	initializer.InitSugarWithPath("log/")
	initializer.InitMysql(ctx)
	initializer.InitCaptcha(ctx)

	// the first thread of server
	// 启动服务
	for _, simpleThread := range server.Threads {
		sugar.Infof("start simple thread %s", simpleThread.Name)
		go simpleThread.Thread(ctx)
	}

	// start monitor
	// 启动 prometheus 监控系统
	go func() {
		addr := viper.GetString("core.monitor_addr")
		sugar.Infof("monitor addr: %s", addr)
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(addr, nil)
	}()

	// 4. check system signal for graceful ending.
	// 4 检测系统信号，促使优雅终止
	sg := make(chan os.Signal)
	signal.Notify(sg, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	// stop server
	select {
	case s := <-sg:
		// 结束上下文，通告其他组件结束进程
		cancel()
		sugar.Infof("got signal: %s", s.String())
	}

	if viper.GetString("core.env") == constants.EnvProduction {
		// wait for stopping clear
		time.Sleep(time.Second * 20)
	} else {
		time.Sleep(time.Second * 1)
	}
	sugar.Info("app stoped successfully")
	return nil
}
