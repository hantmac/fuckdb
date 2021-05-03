package cmd

import (
	"context"
	"embed"
	"fmt"
	"fuckdb/config"
	"fuckdb/routers"
	"fuckdb/routers/middleware"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start a web server with UI",
	Run: func(cmd *cobra.Command, args []string) {
		// init config
		if err := config.InitConfig(""); err != nil {
			logrus.Fatalln("init config error:%s", err)
		}
		logrus.Infoln("config init success")

		var host = viper.GetString("server.host")
		var port = viper.GetString("server.port")

		go func() {
			time.Sleep(time.Second)
			url := fmt.Sprintf("http://%s:%s/html", host, port)
			openBrowser(url)
		}()

		startHTTPServer(host, port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

//go:embed dist
var dist embed.FS

func startHTTPServer(host, port string) {
	g := gin.Default()
	//Use frontend static
	html, _ := fs.Sub(dist, "dist")
	g.StaticFS("/html", http.FS(html))

	// LoggerWithFormatter middleware
	// By default gin.DefaultWriter = os.Stdout
	g.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		b, _ := ioutil.ReadAll(param.Request.Body)
		// your custom format
		return fmt.Sprintf("StatusCode:%d--%s-ClientIp:%s - TimeStamp:[%s] \"ReqMethod:%s--"+
			"API:%s--ReqProto%s--CostTime:%s \"UserAgent:%s-\" -Error:%s\"\n",
			param.StatusCode,
			string(b),
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	//Add middleware in this slice
	var middlewares = []gin.HandlerFunc{middleware.NoCache, middleware.Secure}

	routers.Load(
		g,
		middlewares...,
	)

	// run server
	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: g,
	}
	processed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			logrus.Errorln("server shutdown failed,err:", err)
		}
		logrus.Infoln("server shutdown gracefully")
		close(processed)
	}()

	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		logrus.Errorln("server not shutdown gracefully,err:", err)
	}

	<-processed
}

func openBrowser(url string) {
	switch runtime.GOOS {
	case "darwin":
		openFreeBSDBrowser(url)

	case "linux":
		openLinuxBrowser(url)

	case "windows":
		openWinBrowser(url)
	default:
		logrus.Errorf("no such os: %v", runtime.GOOS)
	}

}

func openFreeBSDBrowser(url string) {
	openCmd := exec.Command("open", url)
	if err := openCmd.Run(); err != nil {
		logrus.Errorln("open browser error, please open browser manually:", url)
	}
}

func openLinuxBrowser(url string) {
	openCmd := exec.Command("x-www-browser", url)
	if err := openCmd.Run(); err != nil {
		logrus.Errorln("open browser error, please open browser manually:", url)
	}
}

func openWinBrowser(url string) {
	cmd := exec.Command("cmd", "/C", "start", url)

	if err := cmd.Run(); err != nil {
		logrus.Errorln("open browser error, please open browser manually:", url)
	}
}
