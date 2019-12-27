package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mmcdole/gofeed"
	"github.com/oklog/oklog/pkg/group"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	feederPb "github.com/michaljirman/newsapp/newsfeeder-service/pb"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/configs"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/endpoints"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/persistence"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/repositories"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/services"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/transports"
)

func main() {
	logger := logrus.New()
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "debug"
	}
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	logger.SetReportCaller(true)
	logger.SetLevel(ll)

	cfg, err := configs.Get()
	if err != nil {
		logger.Errorf("Failed to load config: %+v", err)
	}

	dbHandle, err := persistence.InitPostgresDbWithMigration(&cfg.Db, logger)
	if err != nil {
		logger.Fatalf("failed to initialise db %s %+v", cfg.Db.GetDSN(), err)
	}
	// Database setup
	repository, err := repositories.NewFeedRepository(&cfg.Db, dbHandle, logger)
	if err != nil {
		logger.Fatalf("failed to create repository for %s %+v", cfg.Db.GetDSN(), err)
	}

	var (
		newsfeederSvc = services.NewFeedService(logger, cfg, repository, gofeed.NewParser())
		endpoints     = endpoints.CreateEndpoints(newsfeederSvc, logger)
		gRPCServer    = transports.NewGRPCServer(endpoints, logger)
	)

	var g group.Group
	{
		// The debug listener mounts the http.DefaultServeMux, and serves up
		// stuff like the Prometheus metrics route, the Go debug and profiling
		// routes, and so on.
		debugListener, err := net.Listen("tcp", cfg.DebugAddr)
		if err != nil {
			logger.Debugf("transport=%s during=%s err=%+v", "debug/HTTP", "Listen", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Debugf("transport=%s addr=%s", "debug/HTTP", cfg.DebugAddr)
			return http.Serve(debugListener, http.DefaultServeMux)
		}, func(error) {
			debugListener.Close()
		})
	}
	{
		// The gRPC listener mounts the Go kit gRPC server we created.
		gRPCListener, err := net.Listen("tcp", cfg.GRPCAddr)
		if err != nil {
			logger.Debugf("transport=%s during=%s err=%+v", "gRPC", "Listen", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Debugf("transport=%s addr=%s", "gRPC", cfg.GRPCAddr)
			baseServer := grpc.NewServer()
			feederPb.RegisterFeederServer(baseServer, gRPCServer)
			return baseServer.Serve(gRPCListener)
		}, func(error) {
			gRPCListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Debugf("group exited %+v", g.Run())
}
