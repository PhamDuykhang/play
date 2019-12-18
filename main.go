package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/PhamDuyKhang/userplayboar/internal/app/api"
	"github.com/PhamDuyKhang/userplayboar/internal/app/conf"
	"github.com/PhamDuyKhang/userplayboar/internal/app/errors"
	"github.com/PhamDuyKhang/userplayboar/internal/app/pkg/glog"
)

var (
	log = glog.New().WithPrefix("main")
)

func main() {
	port := flag.Int("p", 8081, "the port of service")
	state := flag.String("s", "dev", "the env of ")
	confPath := flag.String("conf", "config", "where you put your config file")

	flag.Parse()

	cf, err := conf.LoadConfig(state, confPath)
	if err != nil {
		log.WithField("err", err).Errorf("can't load config file")
		panic(err)
	}
	e, err := errors.Init(state, confPath)
	r := api.Init(e, cf)
	if err != nil {
		log.WithField("err", err).Errorf("can't load error message")
		panic(err)
	}

	srv := &http.Server{
		Addr:              fmt.Sprint(":", *port),
		Handler:           r,
		ReadTimeout:       cf.HTTPServer.ReadTimeout,
		WriteTimeout:      cf.HTTPServer.WriteTimeout,
		ReadHeaderTimeout: cf.HTTPServer.ReadHeaderTimeout,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("listen: %s\n", err)
		}
	}()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals

	srvCtx, srvCancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer srvCancel()
	log.Info("shutting down http server...")
	if err := srv.Shutdown(srvCtx); err != nil {
		log.Panic("http server shutdown with error:", err)
	}
}
